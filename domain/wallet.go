package domain

type Wallet struct {
	ID      string `gorm:"primarykey"`
	UserID  uint
	User    User
	Balance uint64 `gorm:"not null;default:0"`
}
}

type WalletUsecase interface {
	GetBalance(walletId string) (int64, error)
}

type WalletRepository interface {
	GetBalance(walletId string) (int64, error)
}
