# Makefile for Protobuf generation

.PHONY: all clean

# Directories
PROTO_DIR_CURRENCY=currency_service/proto
GEN_DIR_CURRENCY=currency_service/pkg/grpc/proto
PROTO_DIR_CRUD=crud/proto
GEN_DIR_CRUD=crud/gen/go

# Tools
PROTOC=protoc
PROTOC_GEN_GO_PLUGIN=protoc-gen-go
PROTOC_GEN_GO_GRPC_PLUGIN=protoc-gen-go-grpc

all: currency crud

# Currency service protobuf generation
currency:
	@mkdir -p $(GEN_DIR_CURRENCY)
	$(PROTOC) --go_out=$(GEN_DIR_CURRENCY) --go-grpc_out=$(GEN_DIR_CURRENCY) \
	-I$(PROTO_DIR_CURRENCY) $(PROTO_DIR_CURRENCY)/currency.proto

# CRUD service protobuf generation
crud:
	@mkdir -p $(GEN_DIR_CRUD)
	$(PROTOC) --go_out=$(GEN_DIR_CRUD) --go-grpc_out=$(GEN_DIR_CRUD) \
	-I$(PROTO_DIR_CRUD) $(PROTO_DIR_CRUD)/crud.proto

clean:
	rm -rf $(GEN_DIR_CURRENCY)/*
	rm -rf $(GEN_DIR_CRUD)/*