.PHONY: build run clean # Declare targets that are not actual files, but rather commands

build:
	@go build -o bin/api ./cmd/api/main.go

run:build
	@./bin/api

# if we use @ before any command, it will not print the command itself

# run: build # this means run will depend on build, first execute build then run