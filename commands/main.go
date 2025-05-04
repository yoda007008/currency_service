package commands

//
// генерация файлов
//protoc --go_out=proto/gen/go --go_opt=paths=source_relative --go-grpc_out=proto/gen/go --go-grpc_opt=paths=source_relative proto/currency.proto
//запуск grpc
//go run grpc/cmd/main.go --config=grpc/internal/config/local.yaml
