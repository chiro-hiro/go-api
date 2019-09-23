CC=@go build
TARGET_AUTH=go-api
BIN_DIR=/usr/bin/

CUR_DIR=$(shell dirname $(realpath $(lastword $(MAKEFILE_LIST))))
BLD_DIR=$(CUR_DIR)/build
GOPATH=$(shell go env GOPATH)
BASE_DIR=$(GOPATH)/src/github.com/chiro-hiro
PROJECT_DIR=$(BASE_DIR)/go-api

default:
	mkdir -p $(BLD_DIR)
	mkdir -p $(BASE_DIR)
	ln -sf $(CUR_DIR) $(PROJECT_DIR)
	cd $(PROJECT_DIR)/users; go get
	cd $(PROJECT_DIR)/exports;go get
	cd $(PROJECT_DIR)/utilities;go get
	go get github.com/go-sql-driver/mysql
auth:
	go build -o $(BLD_DIR)/$(TARGET_AUTH) $(CUR_DIR)/auththentication/main.go
all:
	make clean
	make
	make auth
clean:
	rm -rf $(BASE_DIR)
	rm -rf $(BLD_DIR)
