package clients

import (
	"log"
	"net/http"
	"time"
)

type CronAPIClient struct {
	url string
}

func NewClient(url string) *CronAPIClient {
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
