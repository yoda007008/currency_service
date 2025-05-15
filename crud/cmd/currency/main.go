package main

import (
	migration "currency_service/crud/cmd/migrator"
	"currency_service/crud/handler"
	kirill_sso_v2 "currency_service/crud/proto/gen/go/kirill.sso.v2"
	"currency_service/crud/repository"
	"log"
	"net"
	"os"

	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

func main() {
	connStr := os.Getenv("DATABASE_URL")

	// миграции
	migrationsPath := os.Getenv("MIGRATIONS_PATH")
	if err := migration.RunMigrations(connStr, migrationsPath); err != nil {
		log.Println("Migrations is not success", err)
	}
	repo, err := repository.NewPostgresCurrencyRepository(connStr)
	if err != nil {
		log.Fatal("not connect to database", err)
	}

	grpcPort := os.Getenv("GRPC_PORT")

	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	kirill_sso_v2.RegisterCrudServer(grpcServer, &handler.CurrencyServer{Repo: repo})

	log.Println("gRPC service running on", grpcPort)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("gRPC serve error: %v", err)
	}
}
