proto:
	protoc --go_out=. --go-grpc_out=. ./pkg/pb/product.proto

	protoc --go_out=. --go-grpc_out=. ./pkg/pb/order.proto

server:
	go run cmd/main.go