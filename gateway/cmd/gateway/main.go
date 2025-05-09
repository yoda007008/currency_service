package main

import (
	"context"
	"currency_service/crud/handler"
	"currency_service/gateway/internal/config"
	gatewayHandler "currency_service/gateway/internal/handler"
	"currency_service/gateway/internal/middleware"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// загрузка конфигурации
	cfg, err := config.LoadConfig("gateway/internal/config/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// инициализация middleware для валидации токенов
	authMiddleware := middleware.NewAuthMiddleware(cfg.Auth.URL)

	// инициализация обработчиков
	currencyHandler := gatewayHandler.NewCurrencyHandler(handler.CurrencyServer{})

	// настройка маршрутов
	mux := http.NewServeMux()
	mux.HandleFunc("/currency/get", authMiddleware.Validate(currencyHandler.GetCurrency))
	mux.HandleFunc("/currency/create", authMiddleware.Validate(currencyHandler.CreateCurrency))
	mux.HandleFunc("/currency/update", authMiddleware.Validate(currencyHandler.UpdateCurrency))
	mux.HandleFunc("/currency/delete", authMiddleware.Validate(currencyHandler.DeleteCurrency))

	// создание HTTP сервера
	server := &http.Server{
		Addr:    fmt.Sprintf("%s:%d", cfg.Server.Host, cfg.Server.Port),
		Handler: mux,
	}

	// канал для получения сигналов завершения
	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	// запуск сервера в отдельной горутине
	go func() {
		log.Printf("Starting server on %s:%d", cfg.Server.Host, cfg.Server.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start server: %v", err)
		}
	}()

	// ожидание сигнала завершения
	<-stop
	log.Println("Shutting down server...")

	// cоздание контекста с таймаутом для graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// graceful shutdown
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited properly")
}
