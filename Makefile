# ==============================================================================
# Define dependencies
BASE_IMAGE_NAME := zmoog
SERVICE_NAME    := otel-collector
SERVICE_VERSION := 0.5-$(shell git rev-parse --short HEAD)
SERVICE_IMAGE   := $(BASE_IMAGE_NAME)/$(SERVICE_NAME):$(SERVICE_VERSION)

KO_DOCKER_REPO  := ghcr.io/zmoog/otel-collector-contrib

# ==============================================================================
# Define targets

# BUILD_DIR ?= _build
# export GOBIN = $(shell realpath $(BUILD_DIR))/_bin

# $(BUILD_DIR):
# 	@mkdir -p $(BUILD_DIR)

# $(GOBIN): tools/go.mod
# 	cd tools && go install go.opentelemetry.io/collector/cmd/mdatagen
# 	cd tools && go install go.opentelemetry.io/collector/cmd/builder
# 	cd tools && go install golang.org/x/tools/cmd/goimports
# 	cd tools && go install honnef.co/go/tools/cmd/staticcheck


.PHONY: generate
generate:
	# look inside the receiver directory
	# and run mdatagen against the metadata.yaml
	# found there
	find receiver -name go.mod -execdir go tool mdatagen metadata.yaml \;

.PHONY: staticcheck
staticcheck:
	# run staticcheck for all go
	# directories that have a go.mod
	# file present
	find . -name go.mod -execdir go tool staticcheck ./... \;

.PHONY: fmt
fmt:
	go tool goimports -local github.com/zmoog/ -w .

.PHONY: generate-otelcol
generate-otelcol:
	cd collector && go tool builder --config builder-config.yaml

.PHONY: run
run:
	cd collector/otelcol && ./otelcol --config ../../config.yaml

.PHONY: service
service:
	cd collector/otelcol && KO_DOCKER_REPO=$(KO_DOCKER_REPO) go tool ko build . \
		--platform=linux/amd64,linux/arm64 \
		--bare
