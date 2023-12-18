.PHONY: build clean tool lint help

all: build

build:
    # 编译 linux 下的可执行文件
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o ./dist/go-gin-gorm-starter ../main.go

tool:
	go vet ./...; true
	gofmt -w .

help:
	@echo "make: compile packages and dependencies"
	@echo "make tool: run specified go tool"