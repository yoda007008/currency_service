package dto

import "net/http"

type CurrencyInputCreate struct {
	Code  string  `json:"code"`
	Rate  string  `json:"rate"`
	Value float64 `json:"value"`
}

type GatewayMethods interface {
	GetCurrency(w http.ResponseWriter, r *http.Request)
	CreateCurrency(w http.ResponseWriter, r *http.Request)
	UpdateCurrency(w http.ResponseWriter, r *http.Request)
	DeleteCurrency(w http.ResponseWriter, r *http.Request)
}

type CurrencyInputUpdate struct {
}

type CurrencyInputDelete struct {
}
