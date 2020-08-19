package repository

import (
	"database/sql"

	"github.com/go-playground/validator/v10"

	"go-jwt-auth/model"
	"go-jwt-auth/util"
)

type (
	Account = model.Account
	AccountForm = model.AccountForm
)

type IAccountRepository interface {
	GetAccountByEmail(email string) (Account, error)
	GetAccountById(id int64) (Account, error)
	CreateAccount(form AccountForm) (int64, error)
}

type AccountRepository struct {
	DB *sql.DB
}

var validate *validator.Validate

func NewAccountRepository(db *sql.DB) *AccountRepository {
	return &AccountRepository{db}
}

func (a AccountRepository) GetAccountByEmail(email string) (Account, error) {
	var account Account
	row := a.DB.QueryRow(`SELECT account_id, email, password FROM Accounts WHERE email = $1`, email)
	err := row.Scan(&account.AccountID, &account.Email, &account.Password)
	return account, err
}

func (a AccountRepository) GetAccountById(id int64) (Account, error) {
	var account Account
	row := a.DB.QueryRow(`SELECT account_id, email, password FROM Accounts WHERE account_id = $1`, id)
	err := row.Scan(&account.AccountID, &account.Email, &account.Password)
	return account, err
}

func (a AccountRepository) CreateAccount(form AccountForm) (int64, error) {
	var accountID int64
	hash := util.Password(form.Password).SHA256()
	row := a.DB.QueryRow(`INSERT INTO Accounts (username, email, password) VALUES ($1, $2, $3) RETURNING account_id`, form.Username, form.Email, hash)
	err := row.Scan(&accountID)
	return accountID, err
}
