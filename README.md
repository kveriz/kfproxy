Config example

metadata.broker.list={{ .Values.config.bootstrapservers }}
security.protocol=ssl
ssl.ca.location={{ .Values.config.paths.ca }}
ssl.certificate.location={{ .Values.config.paths.crt }}
ssl.key.location={{ .Values.config.paths.key }}
client.id={{ .Values.config.client_id }}