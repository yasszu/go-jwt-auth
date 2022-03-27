package jwt

import (
	"errors"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/yasszu/go-jwt-auth/domain/entity"
	"github.com/yasszu/go-jwt-auth/util/conf"
)

const expireHour = 24 * 121

type CustomClaims struct {
	AccountID uint `json:"account_id"`
	jwtgo.StandardClaims
}

func Sign(account *entity.Account) (*entity.AccessToken, error) {
	expiredAt := time.Now().Add(time.Hour * expireHour)
	claims := &CustomClaims{
		AccountID:      account.ID,
		StandardClaims: jwtgo.StandardClaims{ExpiresAt: expiredAt.Unix()},
	}
	token := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, claims)
	signedString, err := token.SignedString(conf.JWT.SigningKey())
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
			return conf.JWT.SigningKey(), nil
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
