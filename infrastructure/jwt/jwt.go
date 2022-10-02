package jwt

import (
	"time"

	jwtgo "github.com/dgrijalva/jwt-go"
	log "github.com/sirupsen/logrus"
	"github.com/yasszu/go-jwt-auth/domain/entity"
	"github.com/yasszu/go-jwt-auth/domain/service"
	"github.com/yasszu/go-jwt-auth/pkg/conf"
)

var (
	expireTime = time.Hour * 24 * 121
)

type CustomClaims struct {
	AccountID uint `json:"account_id"`
	jwtgo.StandardClaims
}

type Service struct{}

func NewService() *Service {
	return &Service{}
}

var _ service.Jwt = (*Service)(nil)

func (j *Service) Sign(account *entity.Account) (*entity.AccessToken, error) {
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

func (j *Service) Verify(signedToken string) (uint, error) {
	token, err := jwtgo.ParseWithClaims(
		signedToken,
		&CustomClaims{},
		func(token *jwtgo.Token) (interface{}, error) {
			return conf.JWT.SigningKey(), nil
		},
	)
	if err != nil {
		log.Error(err)
		return 0, ErrorParseClaims
	}

	claims, ok := token.Claims.(*CustomClaims)
	if !ok {
		return 0, ErrorParseClaims
	}
	if claims.ExpiresAt < time.Now().Local().Unix() {
		return 0, ErrorTokenExpired
	}

	return claims.AccountID, nil
}
