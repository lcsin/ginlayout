.PHONY: init swag wire run build gen

init:
	go install github.com/google/wire/cmd/wire@latest
	go install github.com/swaggo/swag/cmd/swag@latest

swag:
	swag init -g cmd/server/main.go -o docs

wire:
	cd cmd/server && wire

run: swag wire
	go run cmd/server/main.go cmd/server/wire_gen.go

build: swag wire
	go build -o bin/server cmd/server/main.go cmd/server/wire_gen.go

gen:
	go run cmd/generate/main.go
