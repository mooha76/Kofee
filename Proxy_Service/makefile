proto: 
	protoc --go_out=. --go-grpc_out=. ./proto/*.proto

mock:
	mockgen -source=./pkg/pb/auth_grpc.pb.go -destination=./pkg/mock/client_mock.go -package=mock

wire:
	cd di && wire

run: 
	go run cmd/main.go

