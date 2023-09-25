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

	$(eval real_dir := $(shell realpath ${DATA_PATH}))
	$(eval parent_dir := $(shell dirname $(realpath $(firstword $(MAKEFILE_LIST)))))
	@echo "Removing all files in ${real_dir} that should be in ${parent_dir} and end with /data"
	@if [[ ${real_dir} == ${parent_dir}/* ]] && [[ ${real_dir} == */data ]]; then \
		rm -rf ${real_dir}/* ; \
	else \
		echo "DATA_PATH=${real_dir} is not set to a directory that is a child of the ${parent_dir} directory and ends with /data"; \
		exit 1; \
	fi