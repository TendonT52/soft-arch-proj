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

deploy:
	@if [ "${MODE}" = "remote" ]; then \
		kubectl create secret generic mail-service --from-file=app.env; \
		IMAGE=ghcr.io/tendont52/soft-arch-proj-backend-mail-service:latest make -C ./helm mail-service; \
	elif [ "${MODE}" = "local" ]; then \
		docker build -t mail-service .; \
		kubectl create secret generic mail-service --from-file=app.env; \
		IMAGE=mail-service make -C ./helm mail-service; \
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
		"DATA_PATH=${DATA_PATH} make -C ../../helm memphis"
	kubectl wait --for=condition=Ready pods --all --all-namespaces

	telepresence connect

k8s-run-prod:
	$(eval DATA_PATH := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))/data)
	DATA_PATH=${DATA_PATH} make -C ../.. k8s-clear

	@echo "Persistent data will be stored in ${DATA_PATH}"
	parallel ::: \
		"telepresence helm install" \
		"DATA_PATH=${DATA_PATH} make -C ../../helm memphis"
	MODE=local make deploy
	kubectl wait --for=condition=Ready pods --all --all-namespaces

	telepresence connect
