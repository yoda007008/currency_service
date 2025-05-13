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
	"time"
)

type Cron struct {
	Rub map[string]float64 `json:"rub"`
}

func UpdateCurrencyRates(db *pgxpool.Pool) {
	fmt.Println("Обновление курсов валют", time.Now())

	resp, err := http.Get("https://www.floatrates.com/daily/rub.json")
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

	var rates Cron
	if err := json.Unmarshal(body, &rates); err != nil {
		log.Println("Ошибка парсина JSON")
	}

	for code, value := range rates.Rub {
		_, err = db.Exec(context.Background(), `
	INSERT INTO currency_rates (code, rate, value)
	VALUES ($1, $2, $3)
	ON CONFLICT (code) DO UPDATE SET rate=$2, value=$3
`, code, "RUB", value)
		if err != nil {
			log.Printf("Ошибка вставки данных для валюты", code, err)
		}
	}
}
