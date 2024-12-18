.PHONY: build clean tool lint help

all: build

build:
	@go build -v .

vet :
	go vet ./...; true

lint:
	golint ./...

clean:
	rm -rf go-gin-example
	go clean -i .

model:
	/root/go/bin/gentool -dsn "root:root@galaxy1024@tcp(192.168.23.53:3306)/bingcool?charset=utf8mb4&parseTime=True&loc=Local" -outPath=./domain/query

help:
	@echo "make: compile packages and dependencies"
	@echo "make tool: run specified go tool"
	@echo "make lint: golint ./..."
	@echo "make clean: remove object files and cached files"