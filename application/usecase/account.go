package usecase

import (
	"context"
	"go-jwt-auth/domain/entity"
	"go-jwt-auth/domain/repository"
	"go-jwt-auth/infrastructure/jwt"
	"go-jwt-auth/util"
	"log"
)

type AccountUsecase interface {
	SignUp(c context.Context, account entity.Account) (*entity.AccessToken, error)
	Login(c context.Context, email string, password string) (*entity.AccessToken, error)
	Me(c context.Context, accountID uint) (*entity.Account, error)
}

type accountUsecase struct {
	accountRepository repository.AccountRepository
}

func NewAccountUsecase(accountRepository repository.AccountRepository) AccountUsecase {
	return &accountUsecase{
		accountRepository: accountRepository,
	}
}

func (u *accountUsecase) SignUp(c context.Context, account entity.Account) (*entity.AccessToken, error) {
	if err := u.accountRepository.CreateAccount(&account); err != nil {
		log.Println(err.Error())
		return nil, err
	}

	token, err := jwt.Sign(&account)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return token, nil
}

func (u *accountUsecase) Login(c context.Context, email string, password string) (*entity.AccessToken, error) {
	account, err := u.accountRepository.GetAccountByEmail(email)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	if err = util.ComparePassword(account.PasswordHash, password); err != nil {
		log.Println(err.Error())
		return nil, err
	}

	token, err := jwt.Sign(account)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return token, nil
}

func (u *accountUsecase) Me(c context.Context, accountID uint) (*entity.Account, error) {
	account, err := u.accountRepository.GetAccountById(accountID)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}

	return account, nil
}
