.PHONY: client server

client:
	@echo "Running client"
	@go run client/main.go

server:
	@echo "Running server"
	@go run server/main.go
