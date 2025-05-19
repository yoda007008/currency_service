package main

import (
	"currency_service/crud/clients"
	"currency_service/crud/repository"
	"currency_service/crud/worker"
	_ "github.com/lib/pq"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	connStr := os.Getenv("DATABASE_URL")
	url := os.Getenv("EXTERNAL_API")

	if url == "" {
		log.Fatal("no environment variable EXTERNAL_API")
	}

	repo, err := repository.NewPostgresCurrencyRepository(connStr)
	if err != nil {
		log.Fatalf("not connect to database", err)
	}

	apiClient := &http.Client{Timeout: time.Second * 10}

	svc := clients.CronCurrencyServer{
		Repo:      repo,
		ApiClient: apiClient,
		BaseURL:   url,
	}

	cronJob := worker.NewCron(&svc)
	cronJob.Start()

	stop := make(chan os.Signal, 1)
	signal.Notify(stop, os.Interrupt, syscall.SIGTERM)

	<-stop

	log.Println("Stop cron")

	cronJob.Stop()

	time.Sleep(1 * time.Second)
}
