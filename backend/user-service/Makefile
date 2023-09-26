create-migrate:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir migrations/ $$name

migrate-up:
	migrate -database "postgres://postgres:password@user-service-postgresql.default:5432/?sslmode=disable" -path migrations up

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

gomock:
	mockgen -source=port/service.go -destination=test/mock_port/mock_service.go -package=mock

dependency:
	mkdir -p ${DATA_PATH}
	@echo "Persistent data will be stored in ${DATA_PATH}"
	parallel ::: \
		"DATA_PATH=${DATA_PATH} make -C ./helm user-service-postgresql" \
		"DATA_PATH=${DATA_PATH} make -C ./helm user-service-redis"

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

k8s-run-dev:
	$(eval DATA_PATH := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))/data)
	DATA_PATH=${DATA_PATH} make -C ../.. k8s-clear

	@echo "Persistent data will be stored in ${DATA_PATH}"
	parallel ::: \
		"telepresence helm install" \
		"DATA_PATH=${DATA_PATH}/user-service make dependency" \
		"DATA_PATH=${DATA_PATH} make -C ../../helm memphis"
	kubectl wait --for=condition=Ready pods --all --all-namespaces

	telepresence connect
	make migrate-up

k8s-run-prod:
	$(eval DATA_PATH := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))/data)
	DATA_PATH=${DATA_PATH} make -C ../.. k8s-clear

	@echo "Persistent data will be stored in ${DATA_PATH}"
	parallel ::: \
		"telepresence helm install" \
		"DATA_PATH=${DATA_PATH}/user-service make dependency" \
		"DATA_PATH=${DATA_PATH} make -C ../../helm memphis"
	MODE=local make deploy
	kubectl wait --for=condition=Ready pods --all --all-namespaces

	telepresence connect
	make migrate-up
