package wallet

type GetWalletResponse struct {
	Username string  `json:"username"`
	Code     string  `json:"code"`
	Balance  float64 `json:"balance"`
}
