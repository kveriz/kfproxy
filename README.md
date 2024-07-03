### Description

App to publish messages into kafka topics via http proxy implementation.

### Config file

By default app using ssl mode
```
metadata.broker.list={{ .Values.config.bootstrapservers }}
security.protocol=ssl
ssl.ca.location={{ .Values.config.mounts.certs }}/ca.crt
ssl.certificate.location={{ .Values.config.mounts.certs }}/tls.crt
ssl.key.location={{ .Values.config.mounts.certs }}/tls.key
client.id={{ .Values.config.client_id }}
```
You can override config by custom
```
{{ .Values.config.custom }}
```

### Usage
```
curl -X POST http://localhost:8080/publish \
--header "Content-Type:application/json" \
--header 'topicName:TOPIC,TOPICn,TOPICn+1' \
-d '{ "key": "value" }'
```
You can pass comma separated list of topics into topicName header.