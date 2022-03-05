package domain

import (
	"time"
)

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
	ID              string    `gorm:"primarykey" json:"id"`
	TransactionTime time.Time `json:"transaction_time"`
	TransactionType string    `json:"transaction_type"`
	CreditedWallet  string    `json:"credited_wallet"`
	DebitedWallet   string    `json:"debited_wallet,omitempty"`
	Notes           string    `json:"notes"`
	Amount          uint64    `json:"amount"`
}

type WalletUsecase interface {
	GetBalance(walletId string) (uint64, error)
	TopUp(walletId string, amount uint32) (Transaction, error)
	Transfer(senderId string, receiverId string, notes string, amount uint32) (Transaction, error)
}

type WalletRepository interface {
	GetBalance(walletId string) (uint64, error)
	Transaction(transactionId string, transactionTime time.Time, transactionType string, creditedWallet string, debitedWallet string, notes string, amount uint32) (Transaction, error)
}
