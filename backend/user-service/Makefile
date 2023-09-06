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
	./k8s/install-helm.sh
	./../../helm/install-base-helm.sh
	telepresence helm install
	telepresence connect