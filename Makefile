.PHONY: build
ENTRY_POINT := ./cmd/analizer/main.go
CURRENT_DIR := $(shell pwd)
ENV_FILE := $(CURRENT_DIR)/.env
CONFIG_PATH := $(CURRENT_DIR)/config/config.yaml

export CONFIG_PATH

run: build
	./main

build: ./main
	go build $(ENTRY_POINT)

start_env:
	cd ./build && docker-compose --env-file $(ENV_FILE) up -d

stop_env:
	cd ./build && docker-compose --env-file $(ENV_FILE) stop