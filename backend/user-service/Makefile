create-migrate:
	@read -p "Enter migration name: " name; \
	migrate create -ext sql -dir migrations/ $$name

migrate-up:
	migrate -database "postgres://postgres:password@user-service-postgresql.default:5432/?sslmode=disable" -path migrations up

migrate-down:
	migrate -database "postgres://postgres:password@user-service-postgresql.default:5432/?sslmode=disable" -path migrations down

migrate-force:
	migrate -database "postgres://postgres:password@user-service-postgresql.default:5432/?sslmode=disable" -path migrations force 1

migrate-drop:
	migrate -database "postgres://postgres:password@user-service-postgresql.default:5432/?sslmode=disable" -path migrations drop

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

k8s-get:
	kubectl get all -A

gomock:
	mockgen -source=port/service.go -destination=test/mock_port/mock_service.go -package=mock

DATA_PATH=$(shell pwd)/data
k8s-init :
	@echo "Persistent data will be stored in ${DATA_PATH}"

	helm install user-service-postgresql ./helm/postgresql  \
	--set global.postgresql.auth.postgresPassword=password \
	--set master.persistence.size=10Gi \
	--set master.persistence.path="${DATA_PATH}/postgresql" \
	--set primary.persistence.existingClaim=user-service-postgresql-persistent-claim \
	--set volumePermissions.enabled=true

	helm install user-service-redis ./helm/redis \
	--set auth.enabled=true \
	--set auth.password=password \
	--set architecture=standalone \
	--set master.persistence.size=10Gi \
	--set master.persistence.path="${DATA_PATH}/redis" \
	--set master.persistence.existingClaim=user-service-redis-persistent-claim \
	--set volumePermissions.enabled=true

	helm install memphis ../../helm/memphis \
	--set memphis.creds.rootPwd=password \
	--set cluster.replicas=1 \
	--set master.persistence.size=10Gi \
	--set master.persistence.path="${DATA_PATH}/memphis" \
	--set memphis.storageEngine.fileStorage.existingClaim=memphis-persistent-claim \
	--set master.metadata.persistence.size=10Gi \
	--set master.metadata.persistence.path="${DATA_PATH}/memphis-meta" \
	--set metadata.persistence.existingClaim=memphis-metadata-persistent-claim \
	--wait

k8s-run-dev:
	telepresence helm install
	telepresence connect
	make k8s-init
	make migrate-up
	make seed

k8s-run-service:
	kubectl create secret generic user-service --from-file=app.env
	helm install user-service ./helm/user-service
