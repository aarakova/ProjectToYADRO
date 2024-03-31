BINARY_NAME=myapp
SRC_NAME=main.go
STOP_WORDS_NAME=stopWords.txt

build:
	go build -o bin/${BINARY_NAME} cmd/task1/${SRC_NAME}

clean:
	rm -rf bin/${BINARY_NAME}
