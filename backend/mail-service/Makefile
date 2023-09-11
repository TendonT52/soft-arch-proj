server:
	go run .

gen-proto:
	protoc \
    --proto_path=proto \
    proto/**/*.proto \
    --go_out=gen \
    --go_opt=paths=source_relative \
    --go-grpc_out=gen \
    --go-grpc_opt=paths=source_relative

test-all:
	 go test ./...