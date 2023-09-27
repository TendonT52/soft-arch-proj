.PHONY: report-service-postgresql report-service

report-service-postgresql:
	helm install report-service-postgresql ./postgresql \
	--set global.postgresql.auth.postgresPassword=password \
	--set master.persistence.size=10Gi \
	--set master.persistence.path="${DATA_PATH}/postgresql" \
	--set primary.persistence.existingClaim=report-service-postgresql-persistent-claim \
	--set volumePermissions.enabled=true \
	--timeout 10m \
	--wait

report-service:
	helm install report-service ./report-service \
	--set image=${IMAGE} \
	--timeout 10m \
	--wait