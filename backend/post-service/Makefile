create-migrate:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir migrations/ $$name

migrate-up:
	migrate -database "postgres://postgres:password@localhost:5432/?sslmode=disable" -path migrations up

migrate-down:
	migrate -database "postgres://postgres:password@localhost:5432/?sslmode=disable" -path migrations down

clean-db:
	psql postgres://postgres:password@localhost:5432/?sslmode=disable -a -f ./cleanup.sql

gen-proto:
	protoc \
    --proto_path=proto \
    proto/**/*.proto \
    --go_out=gen \
    --go_opt=paths=source_relative \
    --go-grpc_out=gen \
    --go-grpc_opt=paths=source_relative

test-all:
	go clean -testcache && go test ./...