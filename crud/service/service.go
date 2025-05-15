package service

import (
	"context"
	"currency_service/crud/dto"
	kirill_sso_v2 "currency_service/crud/proto/gen/go/kirill.sso.v2"
	"currency_service/crud/repository"
)

type CurrencyService struct {
	kirill_sso_v2.UnimplementedCrudServer
	Repo *repository.PostgresCurrencyRepository
}

func (s *CurrencyService) Delete(ctx context.Context, req *kirill_sso_v2.DeleteCurrencyRequest) (*kirill_sso_v2.DeleteCurrencyResponse, error) {
	err := s.Repo.Delete(ctx, req.Code)
	return &kirill_sso_v2.DeleteCurrencyResponse{}, err
}

func (s *CurrencyService) CreateCurrency(ctx context.Context, req *kirill_sso_v2.CreateCurrencyRequest) (*kirill_sso_v2.CreateCurrencyResponse, error) {
	err := s.Repo.Create(ctx, dto.Currency{
		Code:  req.Code,
		Rate:  req.Rate,
		Value: req.Value,
	})
	return &kirill_sso_v2.CreateCurrencyResponse{}, err
}

func (s *CurrencyService) UpdateCurrency(ctx context.Context, req *kirill_sso_v2.UpdateCurrencyRequest) (*kirill_sso_v2.UpdateCurrencyResponse, error) {
	err := s.Repo.Update(ctx, dto.Currency{
		Code:  req.CurrencyUpdate.Code,
		Rate:  req.CurrencyUpdate.Rate,
		Value: req.CurrencyUpdate.Value,
	})
	return &kirill_sso_v2.UpdateCurrencyResponse{}, err
}

func (s *CurrencyService) GetCurrency(ctx context.Context, req *kirill_sso_v2.GetCurrencyRequest) (*kirill_sso_v2.GetCurrencyResponse, error) {
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
