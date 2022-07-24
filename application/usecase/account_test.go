package usecase_test

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/yasszu/go-jwt-auth/application/usecase"
	"github.com/yasszu/go-jwt-auth/domain/entity"
	repository "github.com/yasszu/go-jwt-auth/domain/repository/mock"
	service "github.com/yasszu/go-jwt-auth/domain/service/mock"
)

func Test_accountUsecase_Me(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name      string
		prepare   func(ctx context.Context, ctrl *gomock.Controller) usecase.AccountUsecase
		accountID uint
		want      *entity.Account
		wantErr   bool
	}{
		{
			name: "success",
			prepare: func(ctx context.Context, ctrl *gomock.Controller) usecase.AccountUsecase {
				accountRepository := repository.NewMockAccount(ctrl)
				accountRepository.EXPECT().GetAccountByID(gomock.Any(), uint(1)).Return(
					&entity.Account{
						ID:           1,
						Username:     "test1",
						Email:        "test1@exapmle.com",
						PasswordHash: "password",
					}, nil)
				jwtService := service.NewMockJwt(ctrl)
				uc := usecase.NewAccountUsecase(accountRepository, jwtService)
				return uc
			},
			accountID: 1,
			want: &entity.Account{
				ID:           1,
				Username:     "test1",
				Email:        "test1@exapmle.com",
				PasswordHash: "password",
			},
			wantErr: false,
		},
		{
			name: "error_when_GetAccountByID()_returns_error",
			prepare: func(ctx context.Context, ctrl *gomock.Controller) usecase.AccountUsecase {
				accountRepository := repository.NewMockAccount(ctrl)
				accountRepository.EXPECT().GetAccountByID(gomock.Any(), uint(1)).Return(nil, sql.ErrNoRows)
				jwtService := service.NewMockJwt(ctrl)
				uc := usecase.NewAccountUsecase(accountRepository, jwtService)
				return uc
			},
			accountID: 1,
			want:      nil,
			wantErr:   true,
		},
		{
			name: "error_when_account_is_nil",
			prepare: func(ctx context.Context, ctrl *gomock.Controller) usecase.AccountUsecase {
				accountRepository := repository.NewMockAccount(ctrl)
				accountRepository.EXPECT().GetAccountByID(gomock.Any(), uint(1)).Return(nil, nil)
				jwtService := service.NewMockJwt(ctrl)
				uc := usecase.NewAccountUsecase(accountRepository, jwtService)
				return uc
			},
			accountID: 1,
			want:      nil,
			wantErr:   true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			ctrl := gomock.NewController(t)
			ctx := context.Background()
			u := tt.prepare(ctx, ctrl)

			got, err := u.Me(ctx, tt.accountID)
			if (err != nil) != tt.wantErr {
				t.Errorf("Me() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Me() got = %v, want %v", got, tt.want)
			}
		})
	}
}
