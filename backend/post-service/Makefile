create-migrate:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir migrations/ $$name

migrate-up:
	migrate -database "postgres://postgres:password@localhost:5432/?sslmode=disable" -path migrations up

clean-db:
	psql postgres://postgres:password@localhost:5432/?sslmode=disable -a -f ./cleanup.sql

gen-proto:
	protoc \
    --proto_path=proto \
    proto/**/*.proto \
    --go_out=gen \
    --go_opt=paths=source_relative \
    --go-grpc_out=gen \
    --go-grpc_opt=paths=source_relative & \
	protoc \
    --proto_path=../user-service/proto \
    ../user-service/proto/**/*.proto \
    --go_out=gen \
    --go_opt=paths=source_relative \
    --go-grpc_out=gen \
    --go-grpc_opt=paths=source_relative

test-all:
	go test ./...

dependency:
	mkdir -p ${DATA_PATH}
	@echo "Persistent data will be stored in ${DATA_PATH}"
	parallel ::: \
		"DATA_PATH=${DATA_PATH} make -C ./helm post-service-postgresql"

deploy:
	@if [ "${MODE}" = "remote" ]; then \
		kubectl create secret generic post-service --from-file=app.env; \
		IMAGE=ghcr.io/tendont52/soft-arch-proj-backend-post-service:latest make -C ./helm post-service; \
	elif [ "${MODE}" = "local" ]; then \
		docker build -t post-service .; \
		kubectl create secret generic post-service --from-file=app.env; \
		IMAGE=post-service make -C ./helm post-service; \
	else \
		echo "MODE is not set to main or local"; \
		exit 1; \
	fi

k8s-run-dev:
	$(eval DATA_PATH := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))/data)
	DATA_PATH=${DATA_PATH} make -C ../.. k8s-clear

	@echo "Persistent data will be stored in ${DATA_PATH}"
	parallel ::: \
		"telepresence helm install" \
		"DATA_PATH=${DATA_PATH} make -C ../../helm memphis" \
		"DATA_PATH=${DATA_PATH}/user-service make -C ../user-service dependency" \
		"DATA_PATH=${DATA_PATH} make dependency"
	MODE=remote make -C ../user-service deploy
	kubectl wait --for=condition=Ready pods --all --all-namespaces

	telepresence connect
	make -C ../user-service migrate-up

k8s-run-prod:
	$(eval DATA_PATH := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))/data)
	DATA_PATH=${DATA_PATH} make -C ../.. k8s-clear

	@echo "Persistent data will be stored in ${DATA_PATH}"
	parallel ::: \
		"telepresence helm install" \
		"DATA_PATH=${DATA_PATH} make -C ../../helm memphis" \
		"DATA_PATH=${DATA_PATH}/user-service make -C ../user-service dependency" \
		"DATA_PATH=${DATA_PATH} make dependency"
	MODE=remote make -C ../user-service deploy
	kubectl wait --for=condition=Ready pods --all --all-namespaces

	telepresence connect
	make -C ../user-service migrate-up

	MODE=local make deploy
