package handler

import (
	"context"
	"currency_service/crud/dto"
	kirill_sso_v2 "currency_service/crud/proto/gen/go/kirill.sso.v2"
	"currency_service/crud/repository"
)

type CurrencyServer struct {
	kirill_sso_v2.UnimplementedCrudServer
	Repo *repository.PostgresCurrencyRepository
}

func (s *CurrencyServer) Delete(ctx context.Context, req *kirill_sso_v2.DeleteCurrencyRequest) (*kirill_sso_v2.DeleteCurrencyResponse, error) {
	err := s.Repo.Delete(ctx, req.Code)
	return &kirill_sso_v2.DeleteCurrencyResponse{}, err
}

func (s *CurrencyServer) CreateCurrency(ctx context.Context, req *kirill_sso_v2.CreateCurrencyRequest) (*kirill_sso_v2.CreateCurrencyResponse, error) {
	err := s.Repo.Create(ctx, dto.Currency{
		Code:  req.Code,
		Rate:  req.Rate,
		Value: req.Value,
	})
	return &kirill_sso_v2.CreateCurrencyResponse{}, err
}

func (s *CurrencyServer) UpdateCurrency(ctx context.Context, req *kirill_sso_v2.UpdateCurrencyRequest) (*kirill_sso_v2.UpdateCurrencyResponse, error) {
	err := s.Repo.Update(ctx, dto.Currency{
		Code:  req.CurrencyUpdate.Code,
		Rate:  req.CurrencyUpdate.Rate,
		Value: req.CurrencyUpdate.Value,
	})
	return &kirill_sso_v2.UpdateCurrencyResponse{}, err
}

func (s *CurrencyServer) GetCurrency(ctx context.Context, req *kirill_sso_v2.GetCurrencyRequest) (*kirill_sso_v2.GetCurrencyResponse, error) {
	c, err := s.Repo.Get(ctx, req.Code)
	if err != nil {
		return nil, err
	}

	return &kirill_sso_v2.GetCurrencyResponse{
		Currency: &kirill_sso_v2.CurrencyRate{
			Code:  c.Code,
			Rate:  c.Rate,
			Value: c.Value,
		},
	}, nil
}
