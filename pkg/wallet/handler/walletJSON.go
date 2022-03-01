package handler

type BalanceResponseData struct {
	WalletID string `json:"wallet_id"`
	Balance  uint64  `json:"balance"`
}
