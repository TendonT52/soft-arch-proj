gen-swagger:
	make -C .. gen-gateway
	statik -f -src=./swagger-ui/

deploy:
	@if [ "${MODE}" = "remote" ]; then \
		kubectl create secret generic api-gateway --from-file=app.env; \
		IMAGE=ghcr.io/tendont52/soft-arch-proj-backend-api-gateway:latest make -C ./helm api-gateway; \
	elif [ "${MODE}" = "local" ]; then \
		docker build -t api-gateway .; \
		kubectl create secret generic api-gateway --from-file=app.env; \
		IMAGE=api-gateway make -C ./helm api-gateway; \
	else \
		echo "MODE is not set to main or local"; \
		exit 1; \
	fi

k8s-run-dev:
	make gen-swagger
	$(eval DATA_PATH := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))/data)
	DATA_PATH=${DATA_PATH} make -C .. k8s-clear

	@echo "Persistent data will be stored in ${DATA_PATH}"
	parallel ::: \
		"telepresence helm install" \
		"DATA_PATH=${DATA_PATH} make -C ../helm memphis" \
		"DATA_PATH=${DATA_PATH}/user-service make -C ../backend/user-service dependency" \
		"DATA_PATH=${DATA_PATH}/post-service make -C ../backend/post-service dependency" \
		"DATA_PATH=${DATA_PATH}/report-service make -C ../backend/report-service dependency" \
		"MODE=remote make -C ../backend/user-service deploy" \
		"MODE=remote make -C ../backend/post-service deploy" \
		"MODE=remote make -C ../backend/report-service deploy"
	telepresence connect
	kubectl wait --for=condition=Ready pods --all --all-namespaces

k8s-run-prod:
	make gen-swagger
	$(eval DATA_PATH := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST))))/data)
	DATA_PATH=${DATA_PATH} make -C .. k8s-clear

	@echo "Persistent data will be stored in ${DATA_PATH}"
	parallel ::: \
		"telepresence helm install" \
		"DATA_PATH=${DATA_PATH} make -C ../helm memphis" \
		"DATA_PATH=${DATA_PATH}/user-service make -C ../backend/user-service dependency" \
		"DATA_PATH=${DATA_PATH}/post-service make -C ../backend/post-service dependency" \
		"DATA_PATH=${DATA_PATH}/report-service make -C ../backend/report-service dependency" \
		"MODE=remote make -C ../backend/user-service deploy" \
		"MODE=remote make -C ../backend/post-service deploy" \
		"MODE=remote make -C ../backend/report-service deploy" \
		"MODE=local make deploy"
	telepresence connect
	kubectl wait --for=condition=Ready pods --all --all-namespaces
