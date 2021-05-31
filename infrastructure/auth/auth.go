package auth

import (
	"errors"
	"go-jwt-auth/domain/entity"
	"net/http"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	expireHour   = 24 * 121
	jwtSecretKey = "JWT_SECRET"
	AccountIdKey = "AccountId"
)

var (
	signingKey []byte
)

type CustomClaims struct {
	AccountID uint `json:"account_id"`
	jwt.StandardClaims
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
		StandardClaims: jwt.StandardClaims{ExpiresAt: expiredAt.Unix()},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
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
	token, err := jwt.ParseWithClaims(
		signedToken,
		&CustomClaims{},
		func(token *jwt.Token) (interface{}, error) {
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

func GetAccountID(r *http.Request) (uint, bool) {
	ctx := r.Context()
	accountID, ok := ctx.Value(AccountIdKey).(uint)
	return accountID, ok
}
