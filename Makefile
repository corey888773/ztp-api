
OUT := bin/app
SRC := main.go

.PHONY: build
build:
	go build -o $(OUT) $(SRC)

dev:
	docker compose -f docker-compose-dev.yml up -d && go run main.go

.PHONY: compose\:down\:orphans
compose\:down\:orphans:
	docker compose down --remove-orphans

.PHONY: compose\:down\:all
compose\:down\:all:
	docker compose down --volumes --remove-orphans --rmi all