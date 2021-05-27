package usecase

import (
	"go-jwt-auth/domain/entity"
	"go-jwt-auth/domain/repository"
	"go-jwt-auth/jwt"
	"go-jwt-auth/util"

	"github.com/labstack/echo/v4"
)

type AccountUsecase interface {
	SignUp(c echo.Context, account entity.Account) (*entity.AccessToken, error)
	Login(c echo.Context, email string, password string) (*entity.AccessToken, error)
	Me(c echo.Context, accountID uint) (*entity.Account, error)
}

type accountUsecase struct {
	accountRepository repository.AccountRepository
}

func NewAccountUsecase(accountRepository repository.AccountRepository) AccountUsecase {
	return &accountUsecase{
		accountRepository: accountRepository,
	}
}

func (u *accountUsecase) SignUp(c echo.Context, account entity.Account) (*entity.AccessToken, error) {
	if err := u.accountRepository.CreateAccount(&account); err != nil {
		c.Logger().Error(err)
		return nil, err
	}

	token, err := jwt.Sign(&account)
	if err != nil {
		c.Logger().Error(err)
		return nil, err
	}

	return token, nil
}

func (u *accountUsecase) Login(c echo.Context, email string, password string) (*entity.AccessToken, error) {
	account, err := u.accountRepository.GetAccountByEmail(email)
	if err != nil {
		c.Logger().Error(err)
		return nil, err
	}

	if err = util.ComparePassword(account.PasswordHash, password); err != nil {
		c.Logger().Error(err)
		return nil, err
	}

	token, err := jwt.Sign(account)
	if err != nil {
		c.Logger().Error(err)
		return nil, err
	}

	return token, nil
}

func (u *accountUsecase) Me(c echo.Context, accountID uint) (*entity.Account, error) {
	account, err := u.accountRepository.GetAccountById(accountID)
	if err != nil {
		c.Logger().Error(err)
		return nil, err
	}

	return account, nil
}
