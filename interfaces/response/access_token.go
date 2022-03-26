package response

import (
	"github.com/yasszu/go-jwt-auth/domain/entity"
)

type AccessToken struct {
	AccountID uint   `json:"account_id"`
	Token     string `json:"access_token"`
	ExpiresAt int64  `json:"expired_at"`
}

func NewAccessToken(e *entity.AccessToken) AccessToken {
	return AccessToken{
		AccountID: e.AccountID,
		Token:     e.Token,
		ExpiresAt: e.ExpiresAt.Unix(),
	}
}
