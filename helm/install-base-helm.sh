cd "$(dirname "$0")"
helm install memphis ./memphis/ -f ./memphis/dev-values.yaml
cd -