.PHONY: memphis
memphis:
	helm install memphis ./memphis \
	--set memphis.creds.rootPwd=password \
	--set cluster.replicas=1 \
	--set master.persistence.size=10Gi \
	--set master.persistence.path="${DATA_PATH}/memphis" \
	--set memphis.storageEngine.fileStorage.existingClaim=memphis-persistent-claim \
	--set master.metadata.persistence.size=10Gi \
	--set master.metadata.persistence.path="${DATA_PATH}/memphis-meta" \
	--set metadata.persistence.existingClaim=memphis-metadata-persistent-claim \
	--timeout 10m \
	--wait