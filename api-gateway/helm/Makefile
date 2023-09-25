.PHONY: api-gateway

api-gateway:
	helm install api-gateway ./api-gateway \
	--set image=${IMAGE} \
	--timeout 10m \
	--wait