# vi: ft=make

GOPATH:=$(shell go env GOPATH)

.PHONY: proto test

proto:
	go get github.com/golang/protobuf/protoc-gen-go
	protoc -I . ./idl/content.proto --lile-server_out=. --go_out=plugins=grpc:${GOPATH}/src

build: proto
	go build -o build/sundogrd/content sundogrd/content/main.go
    
test:
	@go get github.com/rakyll/gotest
	gotest -p 1 -v ./...
