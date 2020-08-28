package model

type Account struct {
	AccountID int64
	Username  string
	Email     string
	Password  string
}

type AccountForm struct {
	Username string `validate:"required,max=40"`
	Email string `validate:"required,email"`
	Password string `validate:"required,min=6,max=64"`
}