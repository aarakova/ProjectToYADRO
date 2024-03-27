BINARY_NAME=myapp

build:
	go build -o bin/myapp cmd/task1/main.go

clear-bin:
	rm -rf bin/*
