package transaction

type WithdrawRequest struct {
	Code   string  `json:"code"`
	Amount float64 `json:"amount"`
}
