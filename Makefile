BINARY=server.bin

all: clean deps build run

build: src/main.go
	@echo "\nBuilding the web server:"
	go build -o ./bin/$(BINARY) ./src

run:
	@echo "\nStarting the web server on http://localhost:8080 :"
	./bin/$(BINARY)

test:
	@echo "\nRunning the tests:"
	go test -v src/main.go

clean:
	@echo "\nCleaning up:"
	go clean
	rm -f ./bin/$(BINARY)

deps:
	@echo "\nDownloading dependencies:"
	go mod download
