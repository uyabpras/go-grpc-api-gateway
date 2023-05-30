proto:
	protoc -Ipkg/auth/proto  --go_out=. --go-grpc_out=require_unimplemented_servers=false:. pkg/auth/proto/*.proto
	protoc -Ipkg/order/proto  --go_out=. --go-grpc_out=require_unimplemented_servers=false:. pkg/order/proto/*.proto
	protoc -Ipkg/product/proto  --go_out=. --go-grpc_out=require_unimplemented_servers=false:. pkg/product/proto/*.proto

server:
	go run cmd/main.go