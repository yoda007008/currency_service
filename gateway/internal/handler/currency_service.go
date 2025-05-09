package handler

import (
	"context"
	_ "currency_service/crud/handler"
	kirill_sso_v2 "currency_service/crud/proto/gen/go/kirill.sso.v2"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewCurrencyServiceClient(url string) (*CurrencyServiceClient, error) {
	conn, err := grpc.Dial(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	client := kirill_sso_v2.NewCrudClient(conn)
	return &CurrencyServiceClient{client: client}, nil
}

type CurrencyServiceClient struct {
	client kirill_sso_v2.CrudClient
}

func (c *CurrencyServiceClient) CreateCurrency(ctx context.Context, req *kirill_sso_v2.CreateCurrencyRequest) (*kirill_sso_v2.CreateCurrencyResponse, error) {
	return c.client.CreateCurrency(ctx, req)
}

func (c *CurrencyServiceClient) UpdateCurrency(ctx context.Context, req *kirill_sso_v2.UpdateCurrencyRequest) (*kirill_sso_v2.UpdateCurrencyResponse, error) {
	return c.client.UpdateCurrency(ctx, req)
}

func (c *CurrencyServiceClient) Delete(ctx context.Context, req *kirill_sso_v2.DeleteCurrencyRequest) (*kirill_sso_v2.DeleteCurrencyResponse, error) {
	return c.client.Delete(ctx, req)
}

func (c *CurrencyServiceClient) GetCurrency(ctx context.Context, req *kirill_sso_v2.GetCurrencyRequest) (*kirill_sso_v2.GetCurrencyResponse, error) {
	return c.client.GetCurrency(ctx, req)
}
