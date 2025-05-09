package handler

import (
	"context"
	"currency_service/crud/handler"
	kirill_sso_v2 "currency_service/crud/proto/gen/go/kirill.sso.v2"
	"currency_service/gateway/internal/middleware"
	"net/http"
	"strconv"
)

type CurrencyHandler struct {
	CurrencyService handler.CurrencyServer
}

func NewCurrencyHandler(cs handler.CurrencyServer) *CurrencyHandler {
	return &CurrencyHandler{CurrencyService: cs}
}

func (h *CurrencyHandler) GetCurrency(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	if !middleware.ValidateToken(tokenString) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Currency code is required", http.StatusBadRequest)
		return
	}

	req := &kirill_sso_v2.GetCurrencyRequest{
		Code: code,
	}

	ctx := context.Background()
	_, err := h.CurrencyService.GetCurrency(ctx, req)
	if err != nil {
		http.Error(w, "Fail to fetch currency code", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
}

func (h *CurrencyHandler) CreateCurrency(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	if !middleware.ValidateToken(tokenString) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Currency code is required", http.StatusBadRequest)
		return
	}

	rate := r.URL.Query().Get("code")
	if rate == "" {
		http.Error(w, "Currency rate is required", http.StatusBadRequest)
		return
	}

	var value string

	valueFloat, err := strconv.ParseFloat(value, 64)
	if err != nil {
		http.Error(w, "Invalid currency value format", http.StatusBadRequest)
		return
	}

	if valueFloat <= 0 {
		http.Error(w, "Currency value must be greater than 0", http.StatusBadRequest)
		return
	}

	req := &kirill_sso_v2.CreateCurrencyRequest{
		Code:  code,
		Rate:  rate,
		Value: valueFloat,
	}

	ctx := context.Background()
	_, err = h.CurrencyService.CreateCurrency(ctx, req)
	if err != nil {
		http.Error(w, "Failed to create currency", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
func (h *CurrencyHandler) UpdateCurrency(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	if !middleware.ValidateToken(tokenString) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Currency code is required", http.StatusBadRequest)
		return
	}

	rate := r.URL.Query().Get("code")
	if rate == "" {
		http.Error(w, "Currency rate is required", http.StatusBadRequest)
		return
	}

	var value string

	valueFloat, err := strconv.ParseFloat(value, 64)
	if err != nil {
		http.Error(w, "Invalid currency value format", http.StatusBadRequest)
		return
	}

	if valueFloat <= 0 {
		http.Error(w, "Currency value must be greater than 0", http.StatusBadRequest)
		return
	}

	req := &kirill_sso_v2.UpdateCurrencyRequest{
		CurrencyUpdate: &kirill_sso_v2.CurrencyRate{
			Code:  code,
			Rate:  rate,
			Value: valueFloat,
		},
	}

	ctx := context.Background()
	_, err = h.CurrencyService.UpdateCurrency(ctx, req)
	if err != nil {
		http.Error(w, "Failed to update currency", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}

func (h *CurrencyHandler) DeleteCurrency(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	if !middleware.ValidateToken(tokenString) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	code := r.URL.Query().Get("code")
	if code == "" {
		http.Error(w, "Currency code is required", http.StatusBadRequest)
		return
	}

	req := &kirill_sso_v2.DeleteCurrencyRequest{
		Code: code,
	}

	ctx := context.Background()
	_, err := h.CurrencyService.Delete(ctx, req)
	if err != nil {
		http.Error(w, "Failed to delete currency", http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
}
