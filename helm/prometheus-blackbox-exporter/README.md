# Prometheus-blackbox_exporter
- É uma aplicação que tem como funcionalidade verificar o uptime da aplicação, verificando se seu dns está apontando pro path certo.
Verificar a latência de resposta, você pode verificar seu endpoint através de http, https, tcp e icmp.
- Este Exemplo aqui, mostrará o blackbox monitorando o uptime das aplicações através de um healthcheck, integrando com o prometheus e alertmanager. Para que seja triggado alertas.

## Install blackbox com helm

- link: https://github.com/helm/charts/blob/master/stable/prometheus-blackbox-exporter/README.md
```
$ helm install --name my-release stable/prometheus-blackbox-exporter
```

## Alteração no chart do prometheus para usar o blackbox e criar alertas

### - Criei um job para fazer scrape do blackbox
- link: https://github.com/prometheus/blackbox_exporter/

```yaml
# scrape_configs:
  - job_name: 'blackbox'
    metrics_path: /probe
    params:
      module: [http_2xx]  # Look for a HTTP 200 response.
    static_configs:
      - targets:
        - http://prometheus.io    # Target to probe with http.
        - https://prometheus.io   # Target to probe with https.
        - http://example.com:8080 # Target to probe with http on port 8080.
    relabel_configs:
      - source_labels: [__address__]
        target_label: __param_target
      - source_labels: [__param_target]
        target_label: instance
      - target_label: __address__
        replacement: blackbox-prometheus-blackbox-exporter:9115  # The blackbox exporter's real hostname:port.
```

obs: adicionei no additionalScrapeConfigs: do própio values.yaml da instalação do prometheus.

### - Criei um alerta para que acionasse caso algum endpoint estivesse off

- link: https://www.digitalocean.com/community/tutorials/how-to-use-alertmanager-and-blackbox-exporter-to-monitor-your-web-server-on-ubuntu-16-04
- Criei uma rule nas próprias roles do prometheus no chart mais especificamente no kubernetes-app.yaml

```yaml
    - alert: HttpEndpointUnreachable-Simcloud
      annotations:
        description: " {{`{{`}} $labels.instance {{`}}`}} unreachable"
        summary: "Http endpoint unreachable"
      expr: probe_success == 0
      for: 30s
      labels:
        severity: warning
{{- if .Values.defaultRules.additionalRuleLabels }}
{{ toYaml .Values.defaultRules.additionalRuleLabels | indent 8 }}
{{- end }}
```

- Outros links interessantes
    - https://devconnected.com/how-to-install-and-configure-blackbox-exporter-for-prometheus/
    - https://github.com/prometheus/blackbox_exporter/blob/master/example.yml
