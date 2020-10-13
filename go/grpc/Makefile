codegen:
	protoc -I. --go_out=plugins=grpc:. *.proto

test: codegen
	go test -v ./...
