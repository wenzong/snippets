.PHONY: codegen proto wire
codegen: proto wire

buf:
	buf --config api/pb/buf.yaml check lint

proto: buf
	protoc -I api/pb/ --go_out=plugins=grpc:api/pb api/pb/*.proto

wire:
	wire ./...

.PHONY: e2etest
e2etest: codegen
	go test -v ./test/

.PHONY: serve
serve: codegen
	go run main.go serve

.PHONY: docker
docker:
	docker build -f build/Dockerfile -t demo:latest .
	docker run -it --rm -v `pwd`/config.toml:/config.toml demo:latest app serve
