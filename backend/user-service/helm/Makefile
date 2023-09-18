.PHONY: user-service-postgresql user-service-redis user-service

user-service-postgresql:
	helm install user-service-postgresql ./postgresql \
	--set global.postgresql.auth.postgresPassword=password \
	--set master.persistence.size=10Gi \
	--set master.persistence.path="${DATA_PATH}/postgresql" \
	--set primary.persistence.existingClaim=user-service-postgresql-persistent-claim \
	--set volumePermissions.enabled=true \
	--wait

user-service-redis:
	helm install user-service-redis ./redis \
	--set auth.enabled=true \
	--set auth.password=password \
	--set architecture=standalone \
	--set master.persistence.size=10Gi \
	--set master.persistence.path="${DATA_PATH}/redis" \
	--set master.persistence.existingClaim=user-service-redis-persistent-claim \
	--set volumePermissions.enabled=true \
	--wait

user-service:
	helm install user-service ./user-service \
	--set image=${IMAGE} \
	--wait