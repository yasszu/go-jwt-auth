package jwt

import (
	"fmt"
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"github.com/yasszu/go-jwt-auth/domain/entity"
	"github.com/yasszu/go-jwt-auth/util/conf"
)

var (
	expireTime = time.Hour * 24 * 121
)

type CustomClaims struct {
	AccountID uint `json:"account_id"`
	jwtgo.StandardClaims
}

func Sign(account *entity.Account) (*entity.AccessToken, error) {
	now := time.Now()
	expiresAt := now.Add(expireTime)
	claims := &CustomClaims{
		AccountID:      account.ID,
		StandardClaims: jwtgo.StandardClaims{ExpiresAt: expiresAt.Unix()},
	}
	token := jwtgo.NewWithClaims(jwtgo.SigningMethodHS256, claims)

	signedString, err := token.SignedString(conf.JWT.SigningKey())
	if err != nil {
		return nil, err
	}

	accessToken := &entity.AccessToken{
		AccountID: account.ID,
		Token:     signedString,
		ExpiresAt: expiresAt,
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
		log.Error(err)
		return nil, fmt.Errorf("%v: %w", ErrorParseClaims, err)
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return nil, ErrorParseClaims
	}

	if claims.ExpiresAt < time.Now().Local().Unix() {
		return nil, ErrorTokenExpired
	}

	return claims, nil
}
