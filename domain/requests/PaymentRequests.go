package requests

type CreatePaymentRequest struct {
	UserID    *uint64 `json:"user_id"`
	Created   string  `json:"created"`
	Amount    float64 `json:"amount"`
	Paytype   uint    `json:"paytype"`
	PaidFor   uint    `json:"paid_for"`
	TraceCode string  `json:"trace_code"`
	PaidTo    string  `json:"paid_to"`
	CheckNum  string  `json:"check_num"`
	CheckBank string  `json:"check_bank"`
	CheckDate string  `json:"check_date"`
	Info      string  `json:"info"`
}
