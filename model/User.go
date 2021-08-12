package model

type User struct {
	UserId        string `gorm:"primary_key not_null"`
	AccountType   string
	StudentNumber float64
	Password      string
}
