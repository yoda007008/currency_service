package main

import (
	"currency_service/crud/handler"
	kirill_sso_v2 "currency_service/crud/proto/gen/go/kirill.sso.v2"
	"fmt"
	"google.golang.org/grpc"
	"net"
)

func main() {
	lis, err := net.Listen("tcp", ":50051")
	fmt.Println(err)
	grpcServer := grpc.NewServer()
	kirill_sso_v2.RegisterCrudServer(
		grpcServer,
		&handler.CurrencyService{
			UnimplementedCrudServer: kirill_sso_v2.UnimplementedCrudServer{},
			Currency:                []handler.Currency{},
		},
	)
	err = grpcServer.Serve(lis)
	fmt.Println(err)
}
