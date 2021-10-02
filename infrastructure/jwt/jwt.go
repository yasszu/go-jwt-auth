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
	accessTokenExpireTime  = 1 * time.Hour
	refreshTokenExpireTime = 168 * time.Hour
)

type CustomClaims struct {
	AccountID uint `json:"account_id"`
	jwtgo.StandardClaims
}

func GenerateAccessToken(account *entity.Account) (*entity.AccessToken, error) {
	expiredAt := time.Now().Add(accessTokenExpireTime)
	claims := &CustomClaims{
		AccountID:      account.ID,
		StandardClaims: jwtgo.StandardClaims{ExpiresAt: expiredAt.Unix()},
	}
	token := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, claims)

	signedString, err := token.SignedString(conf.JWT.AccessTokenSigningKey())
	if err != nil {
		return nil, err
	}

	return &entity.AccessToken{
		AccountID: account.ID,
		Token:     signedString,
		ExpiresAt: expiredAt,
	}, nil
}

func GenerateRefreshToken(account *entity.Account) (*entity.AccessToken, error) {
	expiredAt := time.Now().Add(refreshTokenExpireTime)
	claims := &CustomClaims{
		AccountID:      account.ID,
		StandardClaims: jwtgo.StandardClaims{ExpiresAt: expiredAt.Unix()},
	}
	token := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, claims)

	signedString, err := token.SignedString(conf.JWT.RefreshTokenSigningKey())
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
	token, err := jwtgo.ParseWithClaims(
		signedToken,
		&CustomClaims{},
		func(token *jwtgo.Token) (interface{}, error) {
			return conf.JWT.AccessTokenSigningKey(), nil
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

func ValidateRefreshToken(signedToken string) (*CustomClaims, error) {
	token, err := jwtgo.ParseWithClaims(
		signedToken,
		&CustomClaims{},
		func(token *jwtgo.Token) (interface{}, error) {
			return conf.JWT.RefreshTokenSigningKey(), nil
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
