# use specific versions of images
FROM openapitools/openapi-generator-cli:v7.0.1 as api

WORKDIR /local

COPY api api
COPY scripts scripts
COPY .openapi-generator-ignore .openapi-generator-ignore

RUN docker-entrypoint.sh generate -c /local/scripts/generator-cfg.yaml

# not used normally but redefine entrypoint for the case of checking this stage results
ENTRYPOINT ["bash"]

############################################
FROM golang:latest AS build

WORKDIR /app

# download dependencies - low frequency of changes
COPY go.mod go.sum ./
RUN go mod download

# copy sources - higher frequency of changes
COPY internal/ internal/
COPY cmd/ cmd/
COPY --from=api /local/ ./


# create executable - fpjp-webapi-srv
# we want to use scratch image so setting
# the build options in the way that will link all dependencies statically
RUN CGO_ENABLED=0 GOOS=linux \
    go build \
      -ldflags="-w -s" \
      -installsuffix 'static' \
      -o ./fpjp-webapi-srv ./cmd/fpjp-api-service

############################################
FROM scratch

# see https://github.com/opencontainers/image-spec/blob/main/annotations.md for details
LABEL org.opencontainers.image.authors="FPJP"
LABEL org.opencontainers.image.title="Equipment WebAPI Service"
LABEL org.opencontainers.image.description="WEBAPI for managing equipment in departments of a hospital"

# list all variables and their default values for clarity
ENV AMBULANCE_API_ENVIRONMENT=production
ENV AMBULANCE_API_PORT=8080
ENV AMBULANCE_API_MONGODB_HOST=mongo
ENV AMBULANCE_API_MONGODB_PORT=27017
ENV AMBULANCE_API_MONGODB_DATABASE=fpjp-ambulance
ENV AMBULANCE_API_MONGODB_COLLECTION=fpjp-collection
ENV AMBULANCE_API_MONGODB_USERNAME=root
ENV AMBULANCE_API_MONGODB_PASSWORD=
ENV AMBULANCE_API_MONGODB_TIMEOUT_SECONDS=5

COPY --from=build /app/fpjp-webapi-srv ./

# Actual port may be changed during runtime
# Default using for the simple case scenario
EXPOSE 8080
ENTRYPOINT ["./fpjp-webapi-srv"]