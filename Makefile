proto:
	protoc -Ipkg/auth/proto  --go_out=. --go-grpc_out=. pkg/auth/proto/*.proto
	protoc -Ipkg/order/proto  --go_out=. --go-grpc_out=. pkg/order/proto/*.proto
	protoc -Ipkg/product/proto  --go_out=. --go-grpc_out=. pkg/product/proto/*.proto

server:
	go run cmd/main.go