package entity

type Account struct {
	ID           uint `gorm:"primaryKey"`
	Username     string
	Email        string
	PasswordHash string
}
