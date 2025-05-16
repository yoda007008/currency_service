package service

import (
	"context"
	"currency_service/crud/dto"
	"currency_service/crud/repository"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type CronCurrencyServer struct {
	Repo *repository.PostgresCurrencyRepository
}

func (s *CronCurrencyServer) CronUpdateCurrencyRates() error {
	log.Println("Обновление курсов валют", time.Now())

	api := os.Getenv("EXTERNAL_API")

	resp, err := http.Get(api)
	if err != nil {
		return fmt.Errorf("Ошибка при запросе курсов", err)
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("Ошибка чтения ответа:", err)
	}

	var rates dto.CronCurrencyRates
	if err := json.Unmarshal(body, &rates); err != nil {
		return fmt.Errorf("Ошибка парсинга JSON:", err)
	}

	for code, value := range rates.Rub {
		if err := s.Repo.CronInsertRate(context.Background(), code, "RUB", value); err != nil {
			fmt.Errorf("Ошибка вставки данных для валюты %s: %v", code, err)
		}
	}
	return nil
}
