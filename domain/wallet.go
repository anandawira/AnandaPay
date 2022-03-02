package domain

import "time"

type Wallet struct {
	ID      string `gorm:"primarykey"`
	UserID  uint
	User    User
	Balance uint64 `gorm:"not null;default:0"`
}

const (
	TYPE_TRANSFER = "transfer"
	TYPE_TOPUP    = "topup"
)

type Transaction struct {
	ID              string `gorm:"primarykey"`
	TransactionTime time.Time
	CreditedWallet  string
	DebitedWallet   string
	Notes           string
	Amount          uint64
}

type WalletUsecase interface {
	GetBalance(walletId string) (uint64, error)
	TopUp(walletId string, amount uint32) error
}

type WalletRepository interface {
	GetBalance(walletId string) (uint64, error)
	TopUp(transactionId string, transactionTime time.Time, creditedWallet string, notes string, amount uint32) error
}
