.PHONY: build
CURRENT_DIR := $(shell pwd)
ENTRY_POINT := $(CURRENT_DIR)/cmd/analizer/main.go
ENV_FILE := $(CURRENT_DIR)/.env
CONFIG_PATH := $(CURRENT_DIR)/config/config.yaml

export CONFIG_PATH

run: 
	go run $(ENTRY_POINT)

build: ./main
	go build $(ENTRY_POINT)

up_container:
	cd ./build && docker-compose --env-file $(ENV_FILE) up -d

stop_container:
	cd ./build && docker-compose --env-file $(ENV_FILE) stop

down_container:
	cd ./build && docker-compose --env-file $(ENV_FILE) down
