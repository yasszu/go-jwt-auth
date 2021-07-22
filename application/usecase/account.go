package usecase

import (
	"context"

	log "github.com/sirupsen/logrus"
	"github.com/yasszu/go-jwt-auth/domain/entity"
	"github.com/yasszu/go-jwt-auth/domain/repository"
	"github.com/yasszu/go-jwt-auth/infrastructure/jwt"
	"github.com/yasszu/go-jwt-auth/util/crypt"
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
		log.Error(err)
		return nil, &entity.UnexpectedError{Err: err}
	}

	token, err := jwt.Sign(account)
	if err != nil {
		log.Error(err)
		return nil, &entity.UnexpectedError{Err: err}
	}

	return token, nil
}

func (u *accountUsecase) Login(_ context.Context, email, password string) (*entity.AccessToken, error) {
	account, err := u.accountRepository.GetAccountByEmail(email)
	if err != nil {
		log.Error(err)
		return nil, &entity.UnexpectedError{Err: err}
	}

	if err = crypt.ComparePassword(account.PasswordHash, password); err != nil {
		log.Error(err)
		return nil, &entity.UnauthorizedError{
			Massage: "invalid password",
		}
	}

	token, err := jwt.Sign(account)
	if err != nil {
		log.Error(err)
		return nil, &entity.UnexpectedError{Err: err}
	}

	return token, nil
}

func (u *accountUsecase) Me(_ context.Context, accountID uint) (*entity.Account, error) {
	account, err := u.accountRepository.GetAccountByID(accountID)
	if err != nil {
		log.Error(err)
		return nil, &entity.UnexpectedError{Err: err}
	}

	return account, nil
}
