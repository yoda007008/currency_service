package service

import (
	"context"
	"currency_service/crud/dto"
	"currency_service/crud/repository"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"os"
	"time"
)

type CronCurrencyServer struct {
	Repo *repository.PostgresCurrencyRepository
}

func (s *CronCurrencyServer) CronUpdateCurrencyRates() {
	log.Println("Обновление курсов валют", time.Now())

	api := os.Getenv("EXTERNAL_API")

	resp, err := http.Get(api)
	if err != nil {
		log.Println("Ошибка при запросе курсов:", err)
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("Ошибка чтения ответа:", err)
		return
	}

	var rates dto.CronCurrencyRates
	if err := json.Unmarshal(body, &rates); err != nil {
		log.Println("Ошибка парсинга JSON:", err)
		return
	}

	for code, value := range rates.Rub {
		if err := s.Repo.CronInsertRate(context.Background(), code, "RUB", value); err != nil {
			log.Printf("Ошибка вставки данных для валюты %s: %v", code, err)
		}
	}
}
