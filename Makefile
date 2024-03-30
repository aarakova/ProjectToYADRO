BINARY_NAME=myapp
SRC_NAME=main.go
STOP_WORDS_NAME=stopWords.txt

run:
	go run cmd/task1/${SRC_NAME} -stop cmd/task1/${STOP_WORDS_NAME}

build:
	go build -o bin/${BINARY_NAME} cmd/task1/${SRC_NAME}

clear-bin:
	rm -rf bin/${BINARY_NAME}
