package usecase

import (
	"context"

	log "github.com/sirupsen/logrus"
	"github.com/yasszu/go-jwt-auth/domain/entity"
	"github.com/yasszu/go-jwt-auth/domain/repository"
	"github.com/yasszu/go-jwt-auth/domain/service"
	"github.com/yasszu/go-jwt-auth/util/crypt"
)

//go:generate mockgen -source=./account.go -destination=./mock/account.go -package=mock
type AccountUsecase interface {
	SignUp(ctx context.Context, account *entity.Account) (*entity.AccessToken, error)
	Login(ctx context.Context, email, password string) (*entity.AccessToken, error)
	Me(ctx context.Context, accountID uint) (*entity.Account, error)
}

type accountUsecase struct {
	accountRepository repository.Account
	jwtService        service.Jwt
}

func NewAccountUsecase(accountRepository repository.Account, jwtService service.Jwt) AccountUsecase {
	return &accountUsecase{
		accountRepository: accountRepository,
		jwtService:        jwtService,
	}
}

func (u *accountUsecase) SignUp(ctx context.Context, account *entity.Account) (*entity.AccessToken, error) {
	if err := u.accountRepository.CreateAccount(ctx, account); err != nil {
		log.Error(err)
		return nil, newUnexpectedError()
	}

	token, err := u.jwtService.Sign(account)
	if err != nil {
		log.Error(err)
		return nil, newUnexpectedError()
	}

	return token, nil
}

func (u *accountUsecase) Login(ctx context.Context, email, password string) (*entity.AccessToken, error) {
	account, err := u.accountRepository.GetAccountByEmail(ctx, email)
	if err != nil {
		log.Error(err)
		return nil, newUnexpectedError()
	}
	if account == nil {
		return nil, newNotFoundError()
	}

	if err = crypt.ComparePassword(account.PasswordHash, password); err != nil {
		log.Error(err)
		return nil, newErrorUnauthorized()
	}

	token, err := u.jwtService.Sign(account)
	if err != nil {
		log.Error(err)
		return nil, newUnexpectedError()
	}

	return token, nil
}

func (u *accountUsecase) Me(ctx context.Context, accountID uint) (*entity.Account, error) {
	account, err := u.accountRepository.GetAccountByID(ctx, accountID)
	if err != nil {
		log.Error(err)
		return nil, newUnexpectedError()
	}

	return account, nil
}
