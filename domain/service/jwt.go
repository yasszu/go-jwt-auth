package service

import "github.com/yasszu/go-jwt-auth/domain/entity"

//go:generate mockgen -source=./jwt.go -destination=./mock/jwt.go -package=mock
type Jwt interface {
	Sign(account *entity.Account) (*entity.AccessToken, error)
	Verify(signedToken string) (uint, error)
}
