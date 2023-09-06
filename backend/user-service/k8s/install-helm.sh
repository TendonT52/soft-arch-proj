cd "$(dirname "$0")"
helm install todo-redis ./redis/ -f ./redis/dev-values.yaml
helm install todo-postgresql ./postgresql/ -f ./postgresql/dev-values.yaml
cd -