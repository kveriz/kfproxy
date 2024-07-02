**Description**
App to publish messages into kafka topics via http proxy implementation

**Config file**
By default app using ssl mode
```
metadata.broker.list={{ .Values.config.bootstrapservers }}
security.protocol=ssl
ssl.ca.location={{ .Values.config.paths.ca }}
ssl.certificate.location={{ .Values.config.paths.crt }}
ssl.key.location={{ .Values.config.paths.key }}
client.id={{ .Values.config.client_id }}
```
You can override config by custom
```
{{ .Values.config.custom }}
```
**Usage**
```
curl -X POST http://localhost:8080/publish \
--header "Content-Type:application/json" \
--header 'topicName:TOPIC' \
-d '{ "key": "value" }'
```