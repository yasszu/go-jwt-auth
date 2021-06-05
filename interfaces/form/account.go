package form

import (
	"go-jwt-auth/domain/entity"
	"go-jwt-auth/util"

	_validate "github.com/go-playground/validator/v10"
)

type Signup struct {
	Username string `form:"username" validate:"required,max=40"`
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required,min=6,max=64"`
}

func (f Signup) Validate() error {
	validate := _validate.New()
	err := validate.Struct(f)
	return err
}

func (f Signup) Entity() (entity.Account, error) {
	var e entity.Account
	hash, err := util.GenerateBCryptoHash(f.Password)
	if err != nil {
		return entity.Account{}, err
	}
	e.Username = f.Username
	e.Email = f.Email
	e.PasswordHash = hash
	return entity.Account{}, err
}
