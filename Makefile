run:
	go run cmd/to-do/main.go
build:
	go build -o to-do ./cmd/to-do/

.PHONY: run build
.DEFAULT_GOAL:=run