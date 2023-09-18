gen-gateway:
	$(eval SERVICE_NAME=user-service)
	mkdir -p "api-gateway/gen/${SERVICE_NAME}"
	protoc \
		--proto_path "backend/${SERVICE_NAME}/proto" \
		--grpc-gateway_out "api-gateway/gen/${SERVICE_NAME}" \
		--grpc-gateway_opt paths=source_relative \
		--go-grpc_out "api-gateway/gen/${SERVICE_NAME}" \
		--go-grpc_opt paths=source_relative \
		--go_out "api-gateway/gen/${SERVICE_NAME}" \
		--go_opt paths=source_relative \
		--openapiv2_out api-gateway/swagger-ui \
		--openapiv2_opt=allow_merge=true \
		--openapiv2_opt=merge_file_name=${SERVICE_NAME} \
		backend/${SERVICE_NAME}/proto/**/*.proto
	$(eval SERVICE_NAME=post-service)
	mkdir -p "api-gateway/gen/${SERVICE_NAME}"
	protoc \
		--proto_path "backend/${SERVICE_NAME}/proto" \
		--grpc-gateway_out "api-gateway/gen/${SERVICE_NAME}" \
		--grpc-gateway_opt paths=source_relative \
		--go-grpc_out "api-gateway/gen/${SERVICE_NAME}" \
		--go-grpc_opt paths=source_relative \
		--go_out "api-gateway/gen/${SERVICE_NAME}" \
		--go_opt paths=source_relative \
		--openapiv2_out api-gateway/swagger-ui \
		--openapiv2_opt=allow_merge=true \
		--openapiv2_opt=merge_file_name=${SERVICE_NAME} \
		backend/${SERVICE_NAME}/proto/**/*.proto

k8s-clear:
	helm list -q | xargs -I {} helm uninstall {} --wait
	kubectl delete secrets --all
	kubectl delete pvc --all
	kubectl delete pv --all
	telepresence helm uninstall || true
	telepresence quit || true
	if [ -d ${DATA_PATH} ]; then \
		rm -rf ${DATA_PATH}/*; \
	fi