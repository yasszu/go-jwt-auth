package response

import "github.com/yasszu/go-jwt-auth/domain/entity"

type TokenPair struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
}

func NewTokenResponse(e *entity.TokenPair) *TokenPair {
	return &TokenPair{
		AccessToken:  e.AccessToken.Token,
		RefreshToken: e.RefreshToken.Token,
	}
}
