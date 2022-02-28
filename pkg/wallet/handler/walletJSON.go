package handler

type BalanceResponseData struct {
	WalletID string `json:"wallet_id,omitempty"`
	Balance  int64  `json:"balance,omitempty"`
}
