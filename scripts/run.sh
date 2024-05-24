#!/bin/bash

command="$1"

if [ -z "$command" ]; then
    command="start"
fi

ProjectRoot="$(dirname "$(realpath "$0")")/.."

export FPJP_API_ENVIRONMENT="Development"
export FPJP_API_PORT="8080"
export FPJP_API_MONGODB_USERNAME="root"
export FPJP_API_MONGODB_PASSWORD="neUhaDnes"

function mongo {
    docker compose --file "${ProjectRoot}/deployments/docker-compose/compose.yaml" "$@"
}

case "$command" in
  "openapi")
    docker run --rm -ti -v "${ProjectRoot}:/local" openapitools/openapi-generator-cli generate -c /local/scripts/generator-cfg.yaml
    ;;
  "start")
    {
      mongo up --detach
      go run "${ProjectRoot}/cmd/fpjp-api-service"
    } || {
      mongo down
    }
    ;;
  "mongo")
    mongo up
    ;;
  *)
    echo "Unknown command: $command"
    exit 1
    ;;
esac
