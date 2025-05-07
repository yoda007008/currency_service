package main

import (
	"currency_service/gateaway/internal/clients/currency"
	"currency_service/gateaway/internal/config"
	"currency_service/gateaway/internal/handler"
	"currency_service/gateaway/internal/middleware"
	"log"
	"net/http"
)

func main() {
	// Загружаем конфигурацию из файла .env
	cfg, err := config.LoadConfig("/home/kirill/GolandProjects/currency_service/crud/config/.env")
	if err != nil {
		log.Fatal("Error loading config:", err)
	}

	// Инициализация клиента для общения с currency сервисом
	currencyClient := currency.NewClient(cfg.CurrencyServiceUrl)

	// Создание маршрутизатора
	mux := http.NewServeMux()

	// Настройка middleware для проверки токенов
	mux.Handle("/currency", middleware.JWTAuth(handler.CurrencyHandler(currencyClient)))

	// Настройка HTTP-сервера
	server := &http.Server{
		Addr:    cfg.GatewayPort,
		Handler: mux,
	}

	log.Printf("API Gateway running on %s", cfg.GatewayPort)
	if err := server.ListenAndServe(); err != nil {
		log.Fatal("Server failed:", err)
	}
}
