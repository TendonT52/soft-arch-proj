create-migrate:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir migrations/ $$name

server:
	go run .

k8s-get:
	kubectl get all -A

gomock:
	mockgen -source=port/service.go -destination=test/mock_port/mock_service.go -package=mock

init-k8s:

	helm install todo-postgresql ./k8s/postgresql  \
	--set global.postgresql.auth.postgresPassword=password \
	--set master.persistence.size=10Gi \
	--set master.persistence.path="${PWD}/data/postgresql" \
	--set primary.persistence.existingClaim=todo-postgresql-persistent-claim \
	--set volumePermissions.enabled=true

	helm install todo-redis ./k8s/redis \
	--set auth.enabled=true \
	--set auth.password=password \
	--set architecture=standalone \
	--set master.persistence.size=10Gi \
	--set master.persistence.path="${PWD}/data/redis" \
	--set master.persistence.existingClaim=todo-redis-persistent-claim \
	--set volumePermissions.enabled=true

	helm install todo-memphis ../../helm/memphis \
	--set memphis.creds.rootPwd=password \
	--set cluster.replicas=1 \
	--set master.persistence.size=10Gi \
	--set master.persistence.path="${PWD}/data/redis" \
	--set memphis.storageEngine.fileStorage.existingClaim=memphis-persistent-claim \
	--wait

	telepresence helm install
	telepresence connect

# ./k8s/install-helm.sh
# ./../../helm/install-base-helm.sh