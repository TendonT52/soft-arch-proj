service_name="user-service"
input_dir="backend/$service_name/proto"
output_dir="api-gateway/gen/$service_name"

mkdir -p "$output_dir"
mkdir -p "api-gateway/swagger"
protoc \
    --proto_path "$input_dir" \
    --grpc-gateway_out "$output_dir" \
    --grpc-gateway_opt paths=source_relative \
    --go-grpc_out "$output_dir" \
    --go-grpc_opt paths=source_relative \
    --go_out "$output_dir" \
    --go_opt paths=source_relative \
    --openapiv2_out api-gateway/swagger-ui \
    --openapiv2_opt=allow_merge=true \
    --openapiv2_opt=merge_file_name=sofe-arch-prog \
    "$input_dir"/**/*.proto
