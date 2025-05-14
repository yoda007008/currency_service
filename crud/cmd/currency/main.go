package main

import (
	cronius "currency_service/crud/cmd/cron"
	migration "currency_service/crud/cmd/migrator"
	"currency_service/crud/handler"
	kirill_sso_v2 "currency_service/crud/proto/gen/go/kirill.sso.v2"
	"currency_service/crud/repository"
	"github.com/robfig/cron/v3"
	"log"
	"net"
	"os"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {
	connStr := os.Getenv("DATABASE_URL")

	migrationsPath := os.Getenv("MIGRATIONS_PATH")
	if err := migration.RunMigrations(connStr, migrationsPath); err != nil {
		log.Println("Migrations is not success", err)
	}
	// миграции
	repo, err := repository.NewPostgresCurrencyRepository(connStr)
	if err != nil {
		log.Fatal("database is not created", err)
	}

	// крон сервис
	c := cron.New()
	exemp := &cronius.Cron{}
	c.AddFunc("@every 1h", func() {
		log.Println("Running scheduled currency update")
		exemp.UpdateCurrencyRates()
	})
	c.Start()

	// Запуск сразу при старте
	exemp.UpdateCurrencyRates()

	grpcPort := os.Getenv("GRPC_PORT")

	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	kirill_sso_v2.RegisterCrudServer(grpcServer, &handler.CurrencyServer{Repo: repo})

	log.Println("gRPC server running on", grpcPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("gRPC serve error: %v", err)
	}
}
