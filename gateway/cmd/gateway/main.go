package main

import (
	"context"
	kirill_sso_v2 "currency_service/crud/proto/gen/go/kirill.sso.v2"
	"currency_service/gateway/internal/config"
	gatewayHandler "currency_service/gateway/internal/handler"
	"currency_service/gateway/internal/middleware"
	"encoding/json"
	"fmt"
	"google.golang.org/grpc"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	// запуск auth сервиса
	startAuthService()

	// загрузка конфигурации
	cfg, err := config.LoadConfig("gateway/internal/config/config.yaml")
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	// инициализация middleware для валидации токенов
	authMiddleware := middleware.NewAuthMiddleware(cfg.Auth.URL)

	// инициализация обработчиков
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure()) // временно без TLS
	if err != nil {
		log.Fatalf("Failed to connect to gRPC service: %v", err)
	}
	defer conn.Close()

	// Создание gRPC клиента
	grpcClient := kirill_sso_v2.NewCrudClient(conn)

	// Передаём gRPC клиента в HTTP handler
	currencyHandler := gatewayHandler.NewCurrencyHandler(grpcClient)
	//repo := repository.PostgresCurrencyRepository{} // <-- передайте нужные зависимости
	//currencyHandler := gatewayHandler.NewCurrencyHandler(handler.CurrencyServer{})
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
		log.Printf("Starting service on %s:%d", cfg.Server.Host, cfg.Server.Port)
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Failed to start service: %v", err)
		}
	}()

	// ожидание сигнала завершения
	<-stop
	log.Println("Shutting down service...")

	// cоздание контекста с таймаутом для graceful shutdown
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// graceful shutdown
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server forced to shutdown: %v", err)
	}

	log.Println("Server exited properly")
}

func startAuthService() {
	http.HandleFunc("/validate", func(w http.ResponseWriter, r *http.Request) {
		type TokenRequest struct {
			Token string `json:"token"`
		}
		var req TokenRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		if !middleware.ValidateToken(req.Token) {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		w.WriteHeader(http.StatusOK)
	})

	go func() {
		log.Println("Starting auth service on :8081")
		if err := http.ListenAndServe(":8081", nil); err != nil {
			log.Fatalf("Auth service failed: %v", err)
		}
	}()
}
