gen-swagger:
	make -C .. gen-gateway
	statik -src=./swagger-ui/

IMAGE=api-gateway
deploy:
	docker build -t ${IMAGE} .
	kubectl create secret generic api-gateway --from-file=app.env
	IMAGE=${IMAGE} make -C ./helm api-gateway

DATA_PATH=$(shell pwd)/data
k8s-run-dev:
	 make -C .. k8s-clear

	@echo "Persistent data will be stored in ${DATA_PATH}"
	parallel ::: \
		"telepresence helm install" \
		"DATA_PATH=${DATA_PATH}/user-service make -C ./backend/user-service user-service-postgresql" \
		"DATA_PATH=${DATA_PATH}/user-service make -C ./backend/user-service user-service-redis" \
		"DATA_PATH=${DATA_PATH} make -C ../helm memphis"
	telepresence connect
	make -C ./backend/user-service migrate-up
	make -C ./backend/user-service seed
	make -C ./backend/user-service mode=main make deploy

	kubectl wait --for=condition=Ready pods --all --all-namespaces

k8s-run-prod:
	make -C .. k8s-clear
	make docker-build

	@echo "Persistent data will be stored in ${DATA_PATH}"
	parallel ::: \
		"telepresence helm install" \
		"DATA_PATH=${DATA_PATH} make -C ./helm user-service-postgresql  " \
		"DATA_PATH=${DATA_PATH} make -C ./helm user-service-redis" \
		"DATA_PATH=${DATA_PATH} make -C ../../helm memphis"
	kubectl create secret generic user-service --from-file=app.env
	IMAGE=${IMAGE} make -C ./helm user-service

	telepresence connect
	kubectl wait --for=condition=Ready pods --all --all-namespaces

	make migrate-up
	make seed