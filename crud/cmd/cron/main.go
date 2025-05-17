package main

import (
	"currency_service/crud/clients"
	"currency_service/crud/repository"
	"currency_service/crud/service"
	"currency_service/crud/worker"
	_ "github.com/lib/pq"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	connStr := os.Getenv("DATABASE_URL")

	repo, err := repository.NewPostgresCurrencyRepository(connStr)
	if err != nil {
		log.Fatalf("not connect to database", err)
	}

	apiClient := clients.NewClient()

	svc := service.CronCurrencyServer{Repo: repo, APIClient: apiClient}

	cronJob := worker.NewCron(&svc)
	cronJob.Start()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	log.Println("Stop cron")

	cronJob.Stop()

	time.Sleep(1 * time.Second)
}
