package domain

type Wallet struct {
	ID      string `gorm:"unique;not null"`
	UserID  uint
	User    User
	Balance int64 `gorm:"not null;default:0"`
}
