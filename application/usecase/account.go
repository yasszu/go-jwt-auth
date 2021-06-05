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
	SignUp(ctx context.Context, account *entity.Account) (*entity.AccessToken, error)
	Login(ctx context.Context, email, password string) (*entity.AccessToken, error)
	Me(ctx context.Context, accountID uint) (*entity.Account, error)
}

type accountUsecase struct {
	accountRepository repository.AccountRepository
}

func NewAccountUsecase(accountRepository repository.AccountRepository) AccountUsecase {
	return &accountUsecase{
		accountRepository: accountRepository,
	}
}

func (u *accountUsecase) SignUp(_ context.Context, account *entity.Account) (*entity.AccessToken, error) {
	if err := u.accountRepository.CreateAccount(account); err != nil {
		log.Println(err.Error())
		return nil, &entity.UnexpectedError{Err: err}
	}

	token, err := jwt.Sign(account)
	if err != nil {
		log.Println(err.Error())
		return nil, &entity.UnexpectedError{Err: err}
	}

	return token, nil
}

func (u *accountUsecase) Login(_ context.Context, email, password string) (*entity.AccessToken, error) {
	account, err := u.accountRepository.GetAccountByEmail(email)
	if err != nil {
		log.Println(err.Error())
		return nil, &entity.UnexpectedError{Err: err}
	}

	if err = util.ComparePassword(account.PasswordHash, password); err != nil {
		log.Println(err.Error())
		return nil, &entity.UnauthorizedError{
			Massage: "invalid password",
		}
	}

	token, err := jwt.Sign(account)
	if err != nil {
		log.Println(err.Error())
		return nil, &entity.UnexpectedError{Err: err}
	}

	return token, nil
}

func (u *accountUsecase) Me(_ context.Context, accountID uint) (*entity.Account, error) {
	account, err := u.accountRepository.GetAccountByID(accountID)
	if err != nil {
		log.Println(err.Error())
		return nil, &entity.UnexpectedError{Err: err}
	}

	return account, nil
}
