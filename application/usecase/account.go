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
	SignUp(ctx context.Context, account *entity.Account) (*entity.TokenPair, error)
	Login(ctx context.Context, email, password string) (*entity.TokenPair, error)
	RefreshToken(ctx context.Context, refreshToken string) (*entity.TokenPair, error)
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

func (u *accountUsecase) SignUp(_ context.Context, account *entity.Account) (*entity.TokenPair, error) {
	if err := u.accountRepository.CreateAccount(account); err != nil {
		log.Error(err)
		return nil, &entity.UnexpectedError{Err: err}
	}

	accessToken, err := jwt.GenerateAccessToken(account)
	if err != nil {
		log.Error(err)
		return nil, &entity.UnexpectedError{Err: err}
	}

	refreshToken, err := jwt.GenerateRefreshToken(account)
	if err != nil {
		log.Error(err)
		return nil, &entity.UnexpectedError{Err: err}
	}

	return &entity.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (u *accountUsecase) Login(_ context.Context, email, password string) (*entity.TokenPair, error) {
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

	accessToken, err := jwt.GenerateAccessToken(account)
	if err != nil {
		log.Error(err)
		return nil, &entity.UnexpectedError{Err: err}
	}

	refreshToken, err := jwt.GenerateRefreshToken(account)
	if err != nil {
		log.Error(err)
		return nil, &entity.UnexpectedError{Err: err}
	}

	return &entity.TokenPair{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}, nil
}

func (u *accountUsecase) RefreshToken(ctx context.Context, refreshToken string) (*entity.TokenPair, error) {
	claims, err := jwt.ValidateRefreshToken(refreshToken)
	if err != nil {
		return nil, err
	}

	account, err := u.accountRepository.GetAccountByID(claims.AccountID)
	if err != nil {
		log.Error(err)
		return nil, &entity.UnexpectedError{Err: err}
	}

	newAccessToken, err := jwt.GenerateAccessToken(account)
	if err != nil {
		log.Error(err)
		return nil, &entity.UnexpectedError{Err: err}
	}

	newRefreshToken, err := jwt.GenerateRefreshToken(account)
	if err != nil {
		log.Error(err)
		return nil, &entity.UnexpectedError{Err: err}
	}

	return &entity.TokenPair{
		AccessToken:  newAccessToken,
		RefreshToken: newRefreshToken,
	}, nil
}

func (u *accountUsecase) Me(_ context.Context, accountID uint) (*entity.Account, error) {
	account, err := u.accountRepository.GetAccountByID(accountID)
	if err != nil {
		log.Error(err)
		return nil, &entity.UnexpectedError{Err: err}
	}

	return account, nil
}
