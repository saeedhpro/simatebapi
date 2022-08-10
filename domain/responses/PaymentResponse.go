package responses

type PaymentTotalResponse struct {
	Total    float64 `json:"total"`
	DueTotal float64 `json:"due_total"`
}
