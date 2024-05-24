## Instalando Loki.
- Doc loki: https://grafana.com/docs/loki/latest/get-started/

### install Loki
- helm loki: https://artifacthub.io/packages/helm/grafana/loki
- helm install --values values.yaml loki grafana/loki -n monitoring

- o que alterei no values.yaml

```yaml
    type: s3
    bucketNames:
      chunks: lokiuniqlogs
      ruler: lokiuniqlogs
      admin: lokiuniqlogs
    type: s3
    s3:
      endpoint: s3.amazonaws.com
      region: us-east-1
      secretAccessKey: 
      accessKeyId: 
      s3ForcePathStyle: false
      insecure: false  
      http_config: {}
      # -- Check https://grafana.com/docs/loki/latest/configure/#s3_storage_config for more info on how to provide a backoff_config
      backoff_config: {}
```

- tive que alterar também, a questão de auth pois sempre dava erro
```yaml
  # Should authentication be enabled
  auth_enabled: false
```
- doc que ajudou: https://github.com/grafana/helm-charts/issues/921#issuecomment-1183597102
### install promtail
- helm doc: https://artifacthub.io/packages/helm/grafana/promtail
- helm install --values values-promtail.yaml --install promtail grafana/promtail -n monitoring

## DATASET
- adicionando o dataset no grafana : https://grafana.com/blog/2023/04/12/how-to-collect-and-query-kubernetes-logs-with-grafana-loki-grafana-and-grafana-agent/