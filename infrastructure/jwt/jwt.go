package jwt

import (
	"context"
	"errors"
	"os"
	"time"

	"go-jwt-auth/domain/entity"

	jwtgo "github.com/dgrijalva/jwt-go"
)

const (
	expireHour   = 24 * 121
	jwtSecretKey = "JWT_SECRET"
)

var (
	signingKey []byte
)

type CustomClaims struct {
	AccountID uint `json:"account_id"`
	jwtgo.StandardClaims
}

func init() {
	_ = os.Setenv(jwtSecretKey, "b5a636fc-bd01-41b1-9780-7bbd906fa4c0")
	secret := os.Getenv(jwtSecretKey)
	signingKey = []byte(secret)
}

func Sign(account *entity.Account) (*entity.AccessToken, error) {
	expiredAt := time.Now().Add(time.Hour * expireHour)
	claims := &CustomClaims{
		AccountID:      account.ID,
		StandardClaims: jwtgo.StandardClaims{ExpiresAt: expiredAt.Unix()},
	}
	token := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, claims)
	signedString, err := token.SignedString(signingKey)
	if err != nil {
		return nil, err
	}
	accessToken := &entity.AccessToken{
		AccountID: account.ID,
		Token:     signedString,
		ExpiresAt: expiredAt,
	}
	return accessToken, nil
}

func ValidateToken(signedToken string) (*CustomClaims, error) {
	token, err := jwtgo.ParseWithClaims(
		signedToken,
		&CustomClaims{},
		func(token *jwtgo.Token) (interface{}, error) {
			return signingKey, nil
		},
	)
	if err != nil {
		return nil, err
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, errors.New("couldn't parse claims")
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, errors.New("JWT is expired")
	}

	return claims, nil
}

func GetAccountID(ctx context.Context) (uint, bool) {
	accountID, ok := ctx.Value(entity.ContextKeyAccountID).(uint)
	return accountID, ok
}
