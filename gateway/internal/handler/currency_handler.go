package handler

import (
	"context"
	kirill_sso_v2 "currency_service/crud/proto/gen/go/kirill.sso.v2"
	"currency_service/gateway/internal/dto"
	"currency_service/gateway/internal/middleware"
	"encoding/json"
	"net/http"
)

type CurrencyHandler struct {
	CurrencyService kirill_sso_v2.CrudClient
}

func NewCurrencyHandler(client kirill_sso_v2.CrudClient) *CurrencyHandler {
	return &CurrencyHandler{CurrencyService: client}
}

func (h *CurrencyHandler) GetCurrency(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	if !middleware.ValidateToken(tokenString) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var input dto.CurrencyGet

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	if input.Code == "" {
		http.Error(w, "Currency code is required", http.StatusBadRequest)
		return
	}

	req := &kirill_sso_v2.GetCurrencyRequest{
		Code: input.Code,
	}

	ctx := context.Background()
	resp, err := h.CurrencyService.GetCurrency(ctx, req)
	if err != nil {
		http.Error(w, "Failed to fetch currency: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp.Currency)
}
func (h *CurrencyHandler) CreateCurrency(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	if !middleware.ValidateToken(tokenString) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var input dto.CurrencyInputCreate

	err := json.NewDecoder(r.Body).Decode(&input)
	if err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	if input.Code == "" {
		http.Error(w, "Currency code is required", http.StatusBadRequest)
		return
	}

	if input.Rate == "" {
		http.Error(w, "Currency rate is required", http.StatusBadRequest)
		return
	}

	if input.Value <= 0 {
		http.Error(w, "Currency value must be greater than 0", http.StatusBadRequest)
		return
	}

	req := &kirill_sso_v2.CreateCurrencyRequest{
		Code:  input.Code,
		Rate:  input.Rate,
		Value: input.Value,
	}

	ctx := context.Background()
	_, err = h.CurrencyService.CreateCurrency(ctx, req)
	if err != nil {
		http.Error(w, "Failed to create currency", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	w.Write([]byte(`{"status":"created"}`))
}

func (h *CurrencyHandler) UpdateCurrency(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	if !middleware.ValidateToken(tokenString) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var input dto.CurrencyUpdate
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	if input.Code == "" || input.Rate == "" || input.Value <= 0 {
		http.Error(w, "All fields are required and value must be > 0", http.StatusBadRequest)
		return
	}

	req := &kirill_sso_v2.UpdateCurrencyRequest{
		CurrencyUpdate: &kirill_sso_v2.CurrencyRate{
			Code:  input.Code,
			Rate:  input.Rate,
			Value: input.Value,
		},
	}

	ctx := context.Background()
	_, err := h.CurrencyService.UpdateCurrency(ctx, req)
	if err != nil {
		http.Error(w, "Failed to update currency: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"updated"}`))
}

func (h *CurrencyHandler) DeleteCurrency(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")
	if !middleware.ValidateToken(tokenString) {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	var input dto.CurrencyDelete

	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Invalid JSON body", http.StatusBadRequest)
		return
	}

	if input.Code == "" {
		http.Error(w, "Currency code is required", http.StatusBadRequest)
		return
	}

	req := &kirill_sso_v2.DeleteCurrencyRequest{
		Code: input.Code,
	}

	ctx := context.Background()
	_, err := h.CurrencyService.Delete(ctx, req)
	if err != nil {
		http.Error(w, "Failed to delete currency: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"status":"deleted"}`))
}
