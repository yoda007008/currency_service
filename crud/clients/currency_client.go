package clients

import (
	"log"
	"net/http"
	"os"
	"time"
)

type CronAPIClient struct {
	url string
}

func NewClient() *CronAPIClient {
	url := os.Getenv("EXTERNAL_API")
	if url == "" {
		log.Fatal("no environment variable EXTERNAL_API")
	}
	return &CronAPIClient{url: url}
}

func (c *CronAPIClient) RequestToClient() (*http.Response, error) {
	log.Println("Запрос к внешнему API", time.Now())

	resp, err := http.Get(c.url)
	if err != nil {
		return resp, err
	}

	return resp, nil
}
