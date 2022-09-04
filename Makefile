run:
	go run mainFile/main.go
build:
	go build -o todos ./todos/todo.go

.PHONY: run build
.DEFAULT_GOAL:=run