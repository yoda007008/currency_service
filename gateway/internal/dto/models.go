package dto

type CurrencyInput struct {
	Code  string  `json:"code"`
	Rate  string  `json:"rate"`
	Value float64 `json:"value"`
}
