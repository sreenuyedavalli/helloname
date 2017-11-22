{{- define "hello_env_vars" -}}
- name: APP_DB_USERNAME
  valueFrom:
    secretKeyRef:
      name: cloudsql-db-credentials
      key: username
- name: APP_DB_PASSWORD
  valueFrom:
    secretKeyRef:
      name: cloudsql-db-credentials
      key: password
- name: APP_DB_NAME
  valueFrom:
    secretKeyRef:
      name: hello-name-secret
      key: dbname
- name: APP_DB_HOST
  valueFrom:
    secretKeyRef:
      name: hello-name-secret
      key: host
{{- end -}}
