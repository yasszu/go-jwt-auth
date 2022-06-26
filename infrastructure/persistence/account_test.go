package persistence

import (
	"context"
	"reflect"
	"testing"

	"github.com/yasszu/go-jwt-auth/domain/entity"
)

func TestAccountRepository_GetAccountByEmail(t *testing.T) {
	account := &entity.Account{
		ID:           1,
		Username:     "user1",
		Email:        "user1@example.com",
		PasswordHash: "password123",
	}

	prepare(t, func() {
		db.Create(account)
	})

	type args struct {
		email string
	}
	tests := []struct {
		name string
		args struct {
			email string
		}
		want    *entity.Account
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				email: "user1@example.com",
			},
			want:    account,
			wantErr: false,
		},
		{
			name: "not_found_user",
			args: args{
				email: "user2@example.com",
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &AccountRepository{db: db}
			got, err := r.GetAccountByEmail(context.Background(), tt.args.email)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAccountByEmail() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAccountByEmail() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAccountRepository_GetAccountByID(t *testing.T) {
	account := &entity.Account{
		ID:           2,
		Username:     "user2",
		Email:        "user2@example.com",
		PasswordHash: "password123",
	}

	prepare(t, func() {
		db.Create(account)
	})

	type args struct {
		id uint
	}
	tests := []struct {
		name string
		args struct {
			id uint
		}
		want    *entity.Account
		wantErr bool
	}{
		{
			name: "success",
			args: args{
				id: account.ID,
			},
			want:    account,
			wantErr: false,
		},
		{
			name: "not_found_user",
			args: args{
				id: 9999,
			},
			want:    nil,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &AccountRepository{db: db}
			got, err := r.GetAccountByID(context.Background(), tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetAccountByID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetAccountByID() got = %v, want %v", got, tt.want)
			}
		})
	}
}
