.PHONY: post-service-postgresql post-service

post-service-postgresql:
	helm install post-service-postgresql ./postgresql \
	--set global.postgresql.auth.postgresPassword=password \
	--set master.persistence.size=10Gi \
	--set master.persistence.path="${DATA_PATH}/postgresql" \
	--set primary.persistence.existingClaim=post-service-postgresql-persistent-claim \
	--set volumePermissions.enabled=true \
	--timeout 10m \
	--wait

post-service:
	helm install post-service ./post-service \
	--set image=${IMAGE} \
	--timeout 10m \
	--wait