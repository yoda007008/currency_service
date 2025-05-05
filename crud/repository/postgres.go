package repository

import (
	"context"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresCurrencyRepository struct {
	db *pgxpool.Pool
}

func NewPostgresCurrencyRepository(connString string) (*PostgresCurrencyRepository, error) {
	//db, err := pgxpool.New(context.Background(), connString)
	//if err != nil {
	//	return nil, fmt.Errorf("fail to connect to database", err)
	//}

	//_, err := db.Exec(context.Background(), ""+
	//	"CREATE TABLE IF NOT EXISTS currency_rates (code VARCHAR PRIMERY KEY);") короче тут создается типо новая таблица, а потом идут запросы
	panic("implement me")
}

func (r *PostgresCurrencyRepository) Create(ctx context.Context, c Currency) error {
	panic("implement me")
}

func (r *PostgresCurrencyRepository) Get(ctx context.Context, code string) (*Currency, error) {
	panic("implement me")
}

func (r *PostgresCurrencyRepository) Update(ctx context.Context, c *Currency) error {
	panic("implement me")
}

func (r *PostgresCurrencyRepository) Delete(ctx context.Context, code string) error {
	panic("implement me")
}
