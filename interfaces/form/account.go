package form

import (
	"github.com/yasszu/go-jwt-auth/domain/entity"
	"github.com/yasszu/go-jwt-auth/util/crypt"

	_validate "github.com/go-playground/validator/v10"
)

type Signup struct {
	Username string `form:"username" validate:"required,max=40"`
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required,min=6,max=64"`
}

func (f *Signup) Validate() error {
	validate := _validate.New()
	if err := validate.Struct(f); err != nil {
		return err
	}
	return nil
}

func (f *Signup) Entity() (entity.Account, error) {
	hash, err := crypt.GenerateBCryptoHash(f.Password)
	if err != nil {
		return entity.Account{}, err
	}

	return entity.Account{
		Username:     f.Username,
		Email:        f.Email,
		PasswordHash: hash,
	}, nil
}

type Login struct {
	Email    string `form:"email" validate:"required,email"`
	Password string `form:"password" validate:"required,min=6,max=64"`
}

func (f *Login) Validate() error {
	validate := _validate.New()
	if err := validate.Struct(f); err != nil {
		return err
	}
	return nil
}
