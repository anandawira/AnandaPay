package domain

type Wallet struct {
	ID      string `gorm:"unique;not null"`
	UserID  uint
	User    User
	Balance int64 `gorm:"not null;default:0"`
}

type WalletUsecase interface {
	GetBalance(walletId string) (int64, error)
}

type WalletRepository interface {
	GetBalance(walletId string) (int64, error)
}
