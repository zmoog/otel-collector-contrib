# ==============================================================================
# Define dependencies
BASE_IMAGE_NAME := zmoog
# BASE_IMAGE_NAME := ghcr.io/zmoog/benderr
SERVICE_NAME    := otel-collector
SERVICE_VERSION := 0.1.0
SERVICE_IMAGE   := $(BASE_IMAGE_NAME)/$(SERVICE_NAME):$(SERVICE_VERSION)

OTELCOL_VERSION := 0.111.0

# ==============================================================================
# Define targets

ocb:
	cd collector && \
	curl --proto '=https' --tlsv1.2 -fL -o ocb https://github.com/open-telemetry/opentelemetry-collector-releases/releases/download/cmd%2Fbuilder%2Fv${OTELCOL_VERSION}/ocb_${OTELCOL_VERSION}_darwin_arm64 && \
	chmod +x ocb

build-collector: ocb
	cd collector && \
	./ocb --config builder-config.yaml

run-collector:
	go run ./collector/otelcol-dev --config collector/config.yaml

# =============================================================================
# Building containers

service:
	docker build \
		-f Dockerfile \
		-t ${SERVICE_IMAGE} \
		--build-arg BUILD_REF=$(SERVICE_VERSION) \
		--build-arg BUILD_DATE=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ') \
		.	