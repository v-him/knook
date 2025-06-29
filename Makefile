BINARY_NAME=bin/knook

build:
	go build -o ${BINARY_NAME} main.go

clean:
	go clean
	rm ${BINARY_NAME}

.PHONY: build clean
