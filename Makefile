BINARY_NAME=myapp
SRC_NAME=main.go

run:
	go run cmd/${SRC_NAME}

build:
	go build -o bin/${BINARY_NAME} cmd/task1/${SRC_NAME}

clear-bin:
	rm -rf bin/${BINARY_NAME}
