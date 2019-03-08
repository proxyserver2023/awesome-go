.PHONY: build

build:
	protoc -I . --go_out=plugins=grpc:. ./pb/*.proto
