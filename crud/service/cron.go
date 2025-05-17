package service

import (
	"context"
	"currency_service/crud/clients"
	"currency_service/crud/dto"
	"currency_service/crud/repository"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"time"
)

type CronCurrencyServer struct {
	Repo      *repository.PostgresCurrencyRepository
	APIClient *clients.CronAPIClient
}

func (s *CronCurrencyServer) CronUpdateCurrencyRates() error {
	log.Println("Обновление курсов валют", time.Now())

	resp, err := s.APIClient.RequestToClient()
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
