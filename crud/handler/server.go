package handler

import (
	"context"
	kirill_sso_v2 "currency_service/crud/proto/gen/go/kirill.sso.v2"
	"fmt"
)

type Currency struct {
	Code  string
	Rate  string
	Value float64
}
type CurrencyService struct {
	kirill_sso_v2.UnimplementedCrudServer
	Currency []Currency
}

func (s *CurrencyService) Delete(ctx context.Context, request *kirill_sso_v2.DeleteCurrencyRequest) (*kirill_sso_v2.DeleteCurrencyResponse, error) {
	for _, i := range s.Currency {
		if i.Code == request.Code {
			s.Currency = s.Currency[:len(s.Currency)-1]
			return &kirill_sso_v2.DeleteCurrencyResponse{}, nil
		}
	}
	return &kirill_sso_v2.DeleteCurrencyResponse{}, nil
}

func (s *CurrencyService) CreateCurrency(ctx context.Context, req *kirill_sso_v2.CreateCurrencyRequest) (*kirill_sso_v2.CreateCurrencyResponse, error) {
	newCurrency := Currency{
		Code:  req.Code,
		Value: req.Value,
		Rate:  req.Rate,
	}
	s.Currency = append(s.Currency, newCurrency)

	return &kirill_sso_v2.CreateCurrencyResponse{}, nil
}

func (s *CurrencyService) UpdateCurrency(ctx context.Context, req *kirill_sso_v2.UpdateCurrencyRequest) (*kirill_sso_v2.UpdateCurrencyResponse, error) {
	for index, currency := range s.Currency {
		if currency.Code == req.CurrencyUpdate.Code {
			s.Currency[index].Rate = req.CurrencyUpdate.Code
			s.Currency[index].Value = req.CurrencyUpdate.Rate
		}
	}
	return &kirill_sso_v2.UpdateCurrencyResponse{}, nil
}

func (s *CurrencyService) AddCurrency(ctx context.Context, req *kirill_sso_v2.AddCurrencyRequest) (*kirill_sso_v2.AddCurrencyResponse, error) {
	for _, i := range s.Currency {
		if i.Code == req.Code {
			return nil, fmt.Errorf("such an element already exists")
		}
	}
	newCurrency := Currency{
		Code:  req.Code,
		Value: req.Rate,
		Rate:  req.Date,
	}

	s.Currency = append(s.Currency, newCurrency)

	return &kirill_sso_v2.AddCurrencyResponse{}, nil
}

func (s *CurrencyService) GetCurrency(ctx context.Context, req *kirill_sso_v2.GetCurrencyRequest) (*kirill_sso_v2.GetCurrencyResponse, error) {
	for _, i := range s.Currency {
		if i.Code == req.Code {
			return &kirill_sso_v2.GetCurrencyResponse{
				Currency: &kirill_sso_v2.CurrencyRate{},
			}, nil
		}
	}
	return &kirill_sso_v2.GetCurrencyResponse{}, nil
}
