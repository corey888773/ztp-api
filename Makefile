
OUT := bin/app
SRC := main.go

.PHONY: build
build:
	go build -o $(OUT) $(SRC)
