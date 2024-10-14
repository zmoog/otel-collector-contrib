# toggltrackreceiver

To avoid duplicates, I am adding an ingest pipeline to set the time entry `id` field as document `_id` in Elasticsearch.

```text
GET _index_template/logs-toggl.track
{
  "index_templates": [
    {
      "name": "logs-toggl.track",
      "index_template": {
        "index_patterns": [
          "logs-toggl.track-*"
        ],
        "template": {
          "settings": {
            "index": {
              "final_pipeline": "toggl-track-pipeline"
            }
          },
          "mappings": {
            "_routing": {
              "required": false
            },
            "numeric_detection": false,
            "dynamic_date_formats": [
              "strict_date_optional_time",
              "yyyy/MM/dd HH:mm:ss Z||yyyy/MM/dd Z"
            ],
            "_meta": {
              "package": {
                "name": "azure"
              },
              "managed_by": "fleet",
              "managed": true
            },
            "dynamic": true,
            "_source": {
              "excludes": [],
              "includes": [],
              "enabled": true
            },
            "dynamic_templates": [],
            "date_detection": true
          }
        },
        "composed_of": [
          "logs@settings",
          "ecs@mappings"
        ],
        "priority": 200,
        "_meta": {
          "package": {
            "name": "toggl"
          }
        },
        "data_stream": {
          "hidden": false,
          "allow_custom_routing": false
        }
      }
    }
  ]
}
```

```text
{
  "toggl-track-pipeline": {
    "processors": [
      {
        "remove": {
          "field": "_id",
          "ignore_missing": true
        }
      },
      {
        "set": {
          "field": "_id",
          "copy_from": "Attributes.id"
        }
      },
      {
        "set": {
          "field": "Attributes.project_name",
          "value": "Elastic",
          "if": "ctx.Attributes?.project_id == '178435728'"
        }
      },
      {
        "set": {
          "field": "Attributes.project_name",
          "value": "Maintenance",
          "if": "ctx.Attributes?.project_id == '28041930'"
        }
      },
      {
        "set": {
          "field": "Attributes.project_name",
          "value": "Professional",
          "if": "ctx.Attributes?.project_id == '95029662'"
        }
      }
    ]
  }
}
```
