package clients

import (
	"context"
	"currency_service/crud/dto"
	"currency_service/crud/repository"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"
)

type CronCurrencyServer struct {
	ApiClient *http.Client
	Repo      *repository.PostgresCurrencyRepository
	BaseURL   string
}

func NewClient(client *http.Client, repo *repository.PostgresCurrencyRepository, baseURL string) (*CronCurrencyServer, error) {
	if client == nil {
		client = &http.Client{Timeout: 10 * time.Second}
	}
	return &CronCurrencyServer{
		ApiClient: client,
		Repo:      repo,
		BaseURL:   baseURL,
	}, nil

}

func (c *CronCurrencyServer) RequestToClient() (*http.Response, error) {
	log.Println("Запрос к внешнему API", time.Now())

	resp, err := c.ApiClient.Get(c.BaseURL)
	if err != nil {
		return resp, fmt.Errorf("Запрос не выпонен")
	}

	if resp.StatusCode != http.StatusOK {
		defer resp.Body.Close()
		body, _ := io.ReadAll(resp.Body)
		return nil, fmt.Errorf("Статус код:", body)
	}
	return resp, nil
}

func (s *CronCurrencyServer) CronUpdateCurrencyRates() error {
	log.Println("Обновление курсов валют", time.Now())

	resp, err := s.RequestToClient()
	if err != nil {
		return fmt.Errorf("Ошибка при запросе: %w", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Ошибка чтения ответа: %w", err)
	}

	var rates dto.CronCurrencyRates
	if err := json.Unmarshal(body, &rates); err != nil {
		return fmt.Errorf("Ошибка парсинга JSON: %w", err)
	}

	for code, value := range rates.Rub {
		if err := s.Repo.CronInsertRate(context.Background(), code, "RUB", value); err != nil {
			log.Printf("Ошибка вставки данных для валюты %s: %v", code, err)
		}
	}
	return nil
}
