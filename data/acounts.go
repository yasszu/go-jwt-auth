package data

import (
	"database/sql"
	"go-jwt-auth/model"
	"go-jwt-auth/util"

	"github.com/labstack/echo"
)

type Account = model.Account

type Accounts struct {
	DB *sql.DB
}

func NewAccounts(c echo.Context) *Accounts {
	db := c.Get("db").(*sql.DB)
	return &Accounts{DB: db}
}

func (a Accounts) GetAccountByEmail(email string) (Account, error) {
	var account Account
	row := a.DB.QueryRow(`SELECT account_id, email, password FROM accounts WHERE email = $1`, email)
	err := row.Scan(&account.AccountID, &account.Email, &account.Password)
	return account, err
}

func (a Accounts) GetAccountById(id int64) (Account, error) {
	var account Account
	row := a.DB.QueryRow(`SELECT account_id, email, password FROM accounts WHERE account_id = $1`, id)
	err := row.Scan(&account.AccountID, &account.Email, &account.Password)
	return account, err
}

func (a Accounts) CreateAccount(email string, password string) (int64, error) {
	var accountID int64
	hash := util.Password(password).SHA256()
	row := a.DB.QueryRow(`INSERT INTO accounts (email, password) VALUES ($1, $2) RETURNING account_id`, email, hash)
	err := row.Scan(&accountID)
	return accountID, err
}
