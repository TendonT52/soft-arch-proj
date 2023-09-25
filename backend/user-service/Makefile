create-migrate:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir migrations/ $$name

migrate-up:
	migrate -database "postgres://postgres:password@user-service-postgresql.default:5432/?sslmode=disable" -path migrations up

seed:
	psql postgres://postgres:password@user-service-postgresql.default:5432/?sslmode=disable -a -f ./seed.sql

clean-db:
	psql postgres://postgres:password@user-service-postgresql.default:5432/?sslmode=disable -a -f ./cleanup.sql

gen-proto:
	protoc \
    --proto_path=proto \
    proto/**/*.proto \
    --go_out=gen \
    --go_opt=paths=source_relative \
    --go-grpc_out=gen \
    --go-grpc_opt=paths=source_relative

server:
	go run .

test-all:
	go test ./...

gomock:
	mockgen -source=port/service.go -destination=test/mock_port/mock_service.go -package=mock

deploy:
	@if [ "${MODE}" = "remote" ]; then \
		kubectl create secret generic user-service --from-file=app.env; \
		IMAGE=ghcr.io/tendont52/soft-arch-proj-backend-user-service:latest make -C ./helm user-service; \
	elif [ "${MODE}" = "local" ]; then \
		docker build -t user-service .; \
		kubectl create secret generic user-service --from-file=app.env; \
		IMAGE=user-service make -C ./helm user-service; \
	else \
		echo "MODE is not set to main or local"; \
		exit 1; \
	fi

DATA_PATH=$(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))/data
k8s-run-dev:
	DATA_PATH=${DATA_PATH} make -C ../.. k8s-clear

	@echo "Persistent data will be stored in ${DATA_PATH}"
	parallel ::: \
		"telepresence helm install" \
		"DATA_PATH=${DATA_PATH} make -C ./helm user-service-postgresql" \
		"DATA_PATH=${DATA_PATH} make -C ./helm user-service-redis" \
		"DATA_PATH=${DATA_PATH} make -C ../../helm memphis"
	telepresence connect
	kubectl wait --for=condition=Ready pods --all --all-namespaces
	make migrate-up
	make seed

k8s-run-prod:
	DATA_PATH=${DATA_PATH} make -C ../.. k8s-clear

	@echo "Persistent data will be stored in ${DATA_PATH}"
	parallel ::: \
		"telepresence helm install" \
		"DATA_PATH=${DATA_PATH} make -C ./helm user-service-postgresql  " \
		"DATA_PATH=${DATA_PATH} make -C ./helm user-service-redis" \
		"DATA_PATH=${DATA_PATH} make -C ../../helm memphis"

	telepresence connect
	kubectl wait --for=condition=Ready pods --all --all-namespaces
	make migrate-up
	make seed
	MODE=local make deploy
