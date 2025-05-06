package repository

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresCurrencyRepository struct {
	db *pgxpool.Pool
}

func NewPostgresCurrencyRepository(connString string) (*PostgresCurrencyRepository, error) {
	db, err := pgxpool.New(context.Background(), connString)
	if err != nil {
		return nil, fmt.Errorf("fail to connect to database", err)
	}

	_, err = db.Exec(context.Background(), `
			CREATE TABLE IF NOT EXISTS currency_rates (
    		code VARCHAR PRIMERY KEY
            rate VARCHAR NOT NULL
            value DOUBLE PRECISION NOT NULL);`)
	if err != nil {
		return nil, fmt.Errorf("fail to create table", err)
	}
	return &PostgresCurrencyRepository{db: db}, nil
}

func (r *PostgresCurrencyRepository) Create(ctx context.Context, c Currency) error {
	_, err := r.db.Exec(ctx, `INSERT INTO currency_rates(code, rate, value) VALUES ($1, $2, $3)`, c.Code, c.Rate, c.Value)
	return err
}

func (r *PostgresCurrencyRepository) Get(ctx context.Context, code string) (Currency, error) {
	row := r.db.QueryRow(ctx, `SELECT code, rate, value FROM currency_rates WHERE code = $1`, code)
	var c Currency
	err := row.Scan(&c.Code, &c.Rate, &c.Value)
	return c, err
}

func (r *PostgresCurrencyRepository) Update(ctx context.Context, c Currency) error {
	_, err := r.db.Exec(ctx, `UPDATE currency_rates SET rate=$2, value=$3 WHERE code=$1`, c.Code, c.Rate, c.Value)
	return err
}

func (r *PostgresCurrencyRepository) Delete(ctx context.Context, code string) error {
	_, err := r.db.Exec(ctx, `DELETE FROM currency_rates WHERE code=$1`, code)
	return err
}
