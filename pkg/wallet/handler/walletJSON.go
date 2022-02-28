package handler

type BalanceResponseData struct {
	WalletID string `json:"wallet_id"`
	Balance  int64  `json:"balance"`
}
