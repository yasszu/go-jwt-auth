package response

import (
	"github.com/yasszu/go-jwt-auth/domain/entity"
)

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresAt    int64  `json:"expires_at"`
}

func NewTokenResponse(e *entity.TokenPair) *TokenPair {
	return &TokenPair{
		AccessToken:  e.AccessToken.Token,
		RefreshToken: e.RefreshToken.Token,
		ExpiresAt:    e.AccessToken.ExpiresAt.Unix(),
	}
}
