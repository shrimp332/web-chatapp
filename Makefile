OS=$(shell uname -s)


all: build

dev:
ifeq ($(OS), Linux)
	@sh -c "sleep 1; xdg-open http://localhost:8000" &
else ifeq ($(OS), Darwin)
	@sh -c "sleep 1; open http://localhost:8000" &
endif
	@air \
		-build.bin "./main" \
		-build.cmd "make build"

build:
	@go build -o main ./cmd/main.go

run: build
	@./main

clean:
	@rm main

.PHONY: all build run test clean dev
