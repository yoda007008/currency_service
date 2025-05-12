package dto

type CurrencyInputCreate struct {
	Code  string  `json:"code"`
	Rate  string  `json:"rate"`
	Value float64 `json:"value"`
}

type CurrencyInputUpdate struct {
}

type CurrencyInputDelete struct {
}
