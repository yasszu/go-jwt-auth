package middleware

import "github.com/yasszu/go-jwt-auth/domain/service"

type Middleware struct {
	jwtService service.Jwt
}

func NewHandler(jwtService service.Jwt) *Middleware {
	return &Middleware{
		jwtService: jwtService,
	}
}
