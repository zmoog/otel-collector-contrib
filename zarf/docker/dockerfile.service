FROM cgr.dev/chainguard/go AS builder
ENV CGO_ENABLED 0
ARG BUILD_REF
COPY . /app
RUN cd /app && go build -o otelcol-dev ./collector/otelcol-dev

FROM cgr.dev/chainguard/glibc-dynamic
ARG BUILD_DATE
ARG BUILD_REF
COPY --from=builder /app/otelcol-dev /usr/bin/    
COPY collector/config.yaml /app/config.yaml
CMD ["/usr/bin/otelcol-dev", "--config", "/app/config.yaml"]
