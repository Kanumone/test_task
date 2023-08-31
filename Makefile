.PHONY: build run up_container stop_container down_container env
CURRENT_DIR := $(shell pwd)
ENTRY_POINT := $(CURRENT_DIR)/cmd/analizer/main.go
ENV_FILE := $(CURRENT_DIR)/.env
CONFIG_PATH := $(CURRENT_DIR)/config/config.yaml

export CONFIG_PATH

run:
	go run $(ENTRY_POINT)

build: $(CURRENT_DIR)/main
	go build $(ENTRY_POINT)
	./main

up_container:
	cd ./build && docker-compose --env-file $(ENV_FILE) up -d

stop_container:
	cd ./build && docker-compose --env-file $(ENV_FILE) stop

down_container:
	cd ./build && docker-compose --env-file $(ENV_FILE) down

env: .env
	cp $(CURRENT_DIR)/.env.example $(CURRENT_DIR)/.env
