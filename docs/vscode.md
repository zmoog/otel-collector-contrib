# Visual Studio Code settings

## Debugging

Here are the launch configurations I use to debug the OpenTelemetry Collector:

```json
{
    "version": "0.2.0",
    "configurations": [
        {
            "name": "Launch otelcol-dev",
            "type": "go",
            "request": "launch",
            "mode": "auto",
            "program": "${workspaceFolder}/collector/otelcol-dev",
            "args": [
                "--config",
                "${workspaceFolder}/collector/config.yaml"
            ],
            "env": {
                "ELASTICSEARCH_ENDPOINTS": "https://<cluster>.eastus2.azure.elastic-cloud.com:443",
                "ELASTICSEARCH_USERNAME": "<username>",
                "ELASTICSEARCH_PASSWORD": "<password>",
                "TOGGL_API_TOKEN": "<toggl-api-token>",
                "TOGGL_LOOKBACK": "2160h",
                "TOGGL_INTERVAL": "1m",
                "WS_USERNAME": "<wavinsentio-username>",
                "WS_PASSWORD": "<wavinsentio-password>",
                "ZCS_CLIENT_ID": "<zcs-client-id>",
                "ZCS_AUTH_KEY": "<zcs-auth-key>",
                "ZCS_THING_KEY": "<zcs-thing-key>",
            }
        }
    ]
}
```

Save this file as `launch.json` in the `.vscode` directory.
