.PHONY: mail-service

mail-service:
	helm install mail-service ./mail-service \
	--set image=${IMAGE} \
	--timeout 10m \
	--wait