package handler

import (
	"currency_service/gateaway/internal/clients/currency"
	"encoding/json"
	"net/http"
)

func CurrencyHandler(client *currency.Client) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := r.URL.Path[len("/currency/"):]
		result, err := client.GetCurrencyRate(code)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(result)
	}
}
