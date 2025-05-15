package dto

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Currency struct {
	Code  string
	Rate  string
	Value float64
}

type CurrencyRepository interface {
	Create(ctx context.Context, c Currency) error
	Delete(ctx context.Context, code string) error
	Update(ctx context.Context, c Currency) error
	Get(ctx context.Context, code string) (Currency, error)
	GetDB() *pgxpool.Pool
	CronInsertRate(ctx context.Context, code, base string, value float64) error //типо метод, чтобы вынести cron
}

type CronCurrencyRates struct {
	Rub map[string]float64 `json:"rub"`
}
