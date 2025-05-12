package dto

import "net/http"

type GatewayMethods interface {
	GetCurrency(w http.ResponseWriter, r *http.Request)
	CreateCurrency(w http.ResponseWriter, r *http.Request)
	UpdateCurrency(w http.ResponseWriter, r *http.Request)
	DeleteCurrency(w http.ResponseWriter, r *http.Request)
}

type CurrencyInputCreate struct {
	Code  string  `json:"code"`
	Rate  string  `json:"rate"`
	Value float64 `json:"value"`
}

type CurrencyUpdate struct {
	Code  string  `json:"code"`
	Rate  string  `json:"rate"`
	Value float64 `json:"value"`
}

type CurrencyDelete struct {
	Code string `json:"code"`
}

type Currency struct {
	Code string `json:"code"`
}
