BINARY=server.bin

all: clean build run

build: main.go
	go build -o ./bin/$(BINARY) .

run:
	./bin/$(BINARY)

test:
	go test -v main.go

clean:
	go clean
	rm -f ./bin/$(BINARY)

deps:
	go mod download
