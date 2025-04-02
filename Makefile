.PHONY: clean swag install build run dev

APP_NAME = middleware-with-gofiber
BUILD_DIR = ./build
TEMP_DIR = ./tmp

clean:
	rm -rf $(BUILD_DIR)/*
	rm -rf *.out
	rm -rf $(TEMP_DIR)/*

swag:
	swag init

install:
	go mod tidy

build:
	go build -o $(BUILD_DIR)/$(APP_NAME) main.go

run:
	go run main.go

dev:
	air

all-dev:
	make clean && make swag && make install && make dev