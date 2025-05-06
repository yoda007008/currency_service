package repository

import "context"

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
}
