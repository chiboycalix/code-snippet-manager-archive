build:
		go build -o bin/main main.go

run:
		go run main.go

test:
		go test -v ./...

clean:
		rm -rf bin

all: build run