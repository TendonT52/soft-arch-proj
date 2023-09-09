create-migrate:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir migrations/ $$name

server:
	go run .

k8s-get:
	kubectl get all -A

gomock:
	mockgen -source=port/service.go -destination=test/mock_port/mock_service.go -package=mock

k8s-run-dev:
	helm install user-service-postgresql ./helm/postgresql  \
	--set global.postgresql.auth.postgresPassword=password \
	--set master.persistence.size=10Gi \
	--set master.persistence.path="${PWD}/data/postgresql" \
	--set primary.persistence.existingClaim=user-service-postgresql-persistent-claim \
	--set volumePermissions.enabled=true

	helm install user-service-redis ./helm/redis \
	--set auth.enabled=true \
	--set auth.password=password \
	--set architecture=standalone \
	--set master.persistence.size=10Gi \
	--set master.persistence.path="${PWD}/data/redis" \
	--set master.persistence.existingClaim=user-service-redis-persistent-claim \
	--set volumePermissions.enabled=true

	helm install user-service-memphis ../../helm/memphis \
	--set memphis.creds.rootPwd=password \
	--set cluster.replicas=1 \
	--set master.persistence.size=10Gi \
	--set master.persistence.path="${PWD}/data/memphis" \
	--set memphis.storageEngine.fileStorage.existingClaim=memphis-persistent-claim \
	--set master.metadata.persistence.size=10Gi \
	--set master.metadata.persistence.path="${PWD}/data/memphis-meta" \
	--set metadata.persistence.existingClaim=memphis-metadata-persistent-claim \
	--wait

	kubectl create secret generic user-service --from-file=app.env
	helm install user-service ./helm/user-service

	telepresence helm install
	telepresence connect


# ./k8s/install-helm.sh
# ./../../helm/install-base-helm.sh