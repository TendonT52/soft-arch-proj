gen-swagger:
	statik -src=./swagger-ui/


k8s-run-dev:
	telepresence helm install
	telepresence connect

	make -C ../backend/user-service/ k8s-init DATA_PATH=$(shell pwd)/data
	make -C ../backend/user-service/ migrate-up
	make -C ../backend/user-service/ seed
	make -C ../backend/user-service/ k8s-run-service
