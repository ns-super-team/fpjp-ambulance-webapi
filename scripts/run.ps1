param (
    $command
)

if (-not $command)  {
    $command = "start"
}

$ProjectRoot = "${PSScriptRoot}/.."

$env:FPJP_API_ENVIRONMENT="Development"
$env:FPJP_API_PORT="8080"
$env:FPJP_API_MONGODB_USERNAME="root"
$env:FPJP_API_MONGODB_PASSWORD="neUhaDnes"

function mongo {
    docker compose --file ${ProjectRoot}/deployments/docker-compose/compose.yaml $args
}

switch ($command) {
  "openapi" {
    docker run --rm -ti -v ${ProjectRoot}:/local openapitools/openapi-generator-cli generate -c /local/scripts/generator-cfg.yaml
  }
  "start" {
    try {
      mongo up --detach
      go run ${ProjectRoot}/cmd/fpjp-api-service
    } finally {
      mongo down
    }
  }
  "mongo" {
    mongo up
  }
  default {
    throw "Unknown command: $command"
  }
}