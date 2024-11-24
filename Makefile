# ==============================================================================
# Define dependencies
KIND            := kindest/node:v1.29.1@sha256:a0cc28af37cf39b019e2b448c54d1a3f789de32536cb5a5db61a49623e527144
KIND_CLUSTER    := otel-collector
NAMESPACE       := otel-collector-system
APP             := otel-collector
BASE_IMAGE_NAME := zmoog
SERVICE_NAME    := otel-collector
SERVICE_VERSION := 0.2.0
SERVICE_IMAGE   := $(BASE_IMAGE_NAME)/$(SERVICE_NAME):$(SERVICE_VERSION)

OTELCOL_VERSION := 0.111.0

# ==============================================================================
# Define targets

ocb:
	cd collector && \
	curl --proto '=https' --tlsv1.2 -fL -o ocb https://github.com/open-telemetry/opentelemetry-collector-releases/releases/download/cmd%2Fbuilder%2Fv${OTELCOL_VERSION}/ocb_${OTELCOL_VERSION}_darwin_arm64 && \
	chmod +x ocb

generate-collector: ocb
	cd collector && \
	./ocb --config builder-config.yaml

run-collector:
	go run ./collector/otelcol-dev --config collector/config.yaml

# =============================================================================
# Building containers

service:
	docker build \
		-f zarf/docker/dockerfile.service \
		-t ${SERVICE_IMAGE} \
		--build-arg BUILD_REF=$(SERVICE_VERSION) \
		--build-arg BUILD_DATE=$(shell date -u +'%Y-%m-%dT%H:%M:%SZ') \
		.	
# =============================================================================
# Running from within k8s/kind

dev-up:
	kind create cluster \
		--image=${KIND} \
		--name=${KIND_CLUSTER} \
		--config=zarf/k8s/dev/kind-config.yaml

dev-down:
	kind delete cluster --name ${KIND_CLUSTER}

dev-load:
	kind load docker-image ${SERVICE_IMAGE} --name ${KIND_CLUSTER}

dev-apply:
	kustomize build zarf/k8s/dev/otel-collector | kubectl apply -f -
	kubectl wait pods --namespace=$(NAMESPACE) --selector app=$(APP) --for=condition=Ready --timeout=60s

# =============================================================================

dev-status:
	kubectl get nodes -o wide
	kubectl get svc -o wide
	kubectl get pods -o wide --watch --all-namespaces

dev-restart:
	kubectl rollout restart deployment $(APP) --namespace=$(NAMESPACE)

dev-update: service dev-load dev-restart

dev-update-apply: service dev-load dev-apply

# =============================================================================

dev-logs:
	# @kubectl logs --namespace=$(NAMESPACE) -l app=$(APP) --all-containers=true --tail=100 -f | go run app/tooling/logfmt/main.go
	@kubectl logs --namespace=$(NAMESPACE) -l app=$(APP) --all-containers=true --tail=100 -f

dev-describe-deployment:
	@kubectl describe deployment --namespace=$(NAMESPACE) $(APP)

dev-describe-otel-collector:
	@kubectl describe pods --namespace=$(NAMESPACE) -l app=$(APP)

# =============================================================================