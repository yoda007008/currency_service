package cron

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type Cron struct {
	Rub map[string]float64 `json:"rub"`
	db  *pgxpool.Pool
}

func (c *Cron) UpdateCurrencyRates() {
	fmt.Println("Обновление курсов валют", time.Now())

	api := os.Getenv("EXTERNAL_API")

	resp, err := http.Get(api)
	if err != nil {
		log.Println("Ошибка при запросе курсов, пожалуйста попробуйте еще раз")
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println("Ошибка чтения ответа")
		return
	}

	var rates Cron // это надо убрать как-то
	if err := json.Unmarshal(body, &rates); err != nil {
		log.Println("Ошибка парсина JSON")
	}

	for code, value := range rates.Rub {
		_, err = c.db.Exec(context.Background(), `
	INSERT INTO currency_rates (code, rate, value)
	VALUES ($1, $2, $3)
`, code, "RUB", value)
		if err != nil {
			log.Printf("Ошибка вставки данных для валюты", code, err)
		}
	}
}
