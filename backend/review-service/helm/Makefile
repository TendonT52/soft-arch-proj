.PHONY: review-service-postgresql review-service

review-service-postgresql:
	helm install review-service-postgresql ./postgresql \
	--set global.postgresql.auth.postgresPassword=password \
	--set master.persistence.size=10Gi \
	--set master.persistence.path="${DATA_PATH}/postgresql" \
	--set primary.persistence.existingClaim=review-service-postgresql-persistent-claim \
	--set volumePermissions.enabled=true \
	--timeout 10m \
	--wait

review-service:
	helm install review-service ./review-service \
	--set image=${IMAGE} \
	--timeout 10m \
	--wait