package response

import "github.com/yasszu/go-jwt-auth/domain/entity"

type Account struct {
	AccountID uint   `json:"account_id"`
	Username  string `json:"username"`
	Email     string `json:"email"`
}

func NewAccount(e *entity.Account) Account {
	return Account{
		AccountID: e.ID,
		Username:  e.Username,
		Email:     e.Email,
	}
}
