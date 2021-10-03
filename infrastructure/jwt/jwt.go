package jwt

import (
	"context"
	"errors"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/yasszu/go-jwt-auth/domain/entity"
	"github.com/yasszu/go-jwt-auth/util/conf"
)

const (
	accessTokenExpireTime  = 15 * time.Minute
	refreshTokenExpireTime = 168 * time.Hour
)

type CustomClaims struct {
	AccountID uint `json:"account_id"`
	jwtgo.StandardClaims
}

func GenerateAccessToken(account *entity.Account) (*entity.AccessToken, error) {
	return generateToken(account, accessTokenExpireTime, conf.JWT.AccessTokenSigningKey())
}

func GenerateRefreshToken(account *entity.Account) (*entity.AccessToken, error) {
	return generateToken(account, refreshTokenExpireTime, conf.JWT.RefreshTokenSigningKey())
}

func generateToken(account *entity.Account, expiresTime time.Duration, secret []byte) (*entity.AccessToken, error) {
	expiredAt := time.Now().Add(expiresTime)
	claims := &CustomClaims{
		AccountID:      account.ID,
		StandardClaims: jwtgo.StandardClaims{ExpiresAt: expiredAt.Unix()},
	}
	token := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, claims)

	signedString, err := token.SignedString(secret)
	if err != nil {
		return nil, err
	}

	return &entity.AccessToken{
		AccountID: account.ID,
		Token:     signedString,
		ExpiresAt: expiredAt,
	}, nil
}

func ValidateAccessToken(signedToken string) (*CustomClaims, error) {
	return validateToken(signedToken, conf.JWT.AccessTokenSigningKey())
}

func ValidateRefreshToken(signedToken string) (*CustomClaims, error) {
	return validateToken(signedToken, conf.JWT.RefreshTokenSigningKey())
}

func validateToken(signedToken string, secret []byte) (*CustomClaims, error) {
	token, err := jwtgo.ParseWithClaims(
		signedToken,
		&CustomClaims{},
		func(token *jwtgo.Token) (interface{}, error) {
			return secret, nil
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
		return nil, errors.New("jWT is expired")
	}

	return claims, nil
}

func GetAccountID(ctx context.Context) (uint, bool) {
	accountID, ok := ctx.Value(entity.ContextKeyAccountID).(uint)
	return accountID, ok
}
