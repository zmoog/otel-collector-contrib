# Collector

How to run the custom collector:

```shell
# use the ";EntityPath=hubName" format
export EVENTHUB_CONNECTION_STRING=""
export ELASTICSEARCH_ENDPOINTS=""
export ELASTICSEARCH_USERNAME=""
export ELASTICSEARCH_PASSWORD=""
```

```shell
$ ./otelcol-dev/otelcol-dev --config config.yaml

2024-10-05T08:52:06.129+0200    info    service@v0.110.0/service.go:137 Setting up own telemetry...
2024-10-05T08:52:06.132+0200    info    service@v0.110.0/service.go:186 Skipped telemetry setup.
2024-10-05T08:52:06.132+0200    debug   builders/builders.go:24 Beta component. May change in the future.       {"kind": "exporter", "data_type": "logs", "name": "elasticsearch"}
2024-10-05T08:52:06.133+0200    warn    elasticsearchexporter@v0.110.0/config.go:359    dedot has been deprecated: in the future, dedotting will always be performed in ECS mode only   {"kind": "exporter", "data_type": "logs", "name": "elasticsearch"}
2024-10-05T08:52:06.134+0200    info    builders/builders.go:26 Development component. May change in the future.        {"kind": "exporter", "data_type": "logs", "name": "debug"}
2024-10-05T08:52:06.134+0200    debug   builders/builders.go:24 Beta component. May change in the future.       {"kind": "processor", "name": "batch", "pipeline": "logs"}
2024-10-05T08:52:06.134+0200    debug   builders/builders.go:24 Alpha component. May change in the future.      {"kind": "receiver", "name": "azureeventhub", "data_type": "logs"}
2024-10-05T08:52:06.134+0200    debug   builders/extension.go:48        Beta component. May change in the future.       {"kind": "extension", "name": "basicauth"}
2024-10-05T08:52:06.135+0200    info    service@v0.110.0/service.go:208 Starting otelcol-dev... {"Version": "1.0.0", "NumCPU": 10}
2024-10-05T08:52:06.135+0200    info    extensions/extensions.go:39     Starting extensions...
2024-10-05T08:52:06.135+0200    info    extensions/extensions.go:42     Extension is starting...        {"kind": "extension", "name": "basicauth"}
2024-10-05T08:52:06.135+0200    info    extensions/extensions.go:59     Extension started.      {"kind": "extension", "name": "basicauth"}
2024-10-05T08:52:15.466+0200    info    service@v0.110.0/service.go:234 Everything is ready. Begin running and processing data.
```
