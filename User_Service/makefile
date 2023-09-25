proto:
	protoc --go_out=. --go-grpc_out=. ./proto/*.proto

wire:
	cd di && wire

run:
	go run cmd/main.go

mock:
	mockgen -source=./pkg/usecase/interfaces/user.go -destination=./pkg/mock/usecase_mock.go -package=mock

docker-build:
	docker build -t nikhil382/ecommerce-microservice-clean-arch-user-service .