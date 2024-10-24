FROM cgr.dev/chainguard/go AS builder
ENV CGO_ENABLED 0
ARG BUILD_REF

# RUN mkdir /service
COPY . /app
# WORKDIR /app
# RUN go mod download
RUN cd /app && go build -o otelcol-dev ./collector/otelcol-dev

# Copy the sourcecode into the container.
# COPY . /service

# WORKDIR /service/app/services/otel-collector
# RUN go build -ldflags "-X main.build=${BUILD_REF}" ./collector/otelcol-dev

FROM cgr.dev/chainguard/glibc-dynamic
ARG BUILD_DATE
ARG BUILD_REF
# RUN addgroup -g 1000 -S otel-collector && \
    # adduser -u 1000 -h /service -G otel-collector -S otel-collector

COPY --from=builder /app/otelcol-dev /usr/bin/    
# COPY --from=builder /service/app/services/otel-collector/otel-collector /service/otel-collector
COPY collector/config.yaml /app/config.yaml
# WORKDIR /service
# USER bender
CMD ["/usr/bin/otelcol-dev", "--config", "/app/config.yaml"]

# LABEL org.opencontainers.image.created="${BUILD_DATE}" \
#       org.opencontainers.image.title="bender-bot" \
#       org.opencontainers.image.authors="Maurizio Branca <maurizio.branca@gmail.com>" \
#       org.opencontainers.image.source="https://github.com/zmoog/service/app/bender-bot" \
#       org.opencontainers.image.revision="${BUILD_REF}" \
#       org.opencontainers.image.zmoog="zmoog labs" 
