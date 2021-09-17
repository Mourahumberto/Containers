# Doc para instalação e customização do prometheus
 - Documentação oficial: https://artifacthub.io/packages/helm/prometheus-community/kube-prometheus-stack

# kube-prometheus-stack


## Prerequisites

- Kubernetes 1.16+
- Helm 3+

## Get Repo Info

```console
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update
```
## Install Chart 

### latest/stable
```console
# Helm
$ helm install [RELEASE_NAME] prometheus-community/kube-prometheus-stack
```

### versão específica
```console
# Helm
$ helm install my-kube-prometheus-stack prometheus-community/kube-prometheus-stack --version 18.0.2
```

### valores específicos
```console
# Helm
$ helm install -f values.yaml my-kube-prometheus-stack prometheus-community/kube-prometheus-stack --version 18.0.2 --debug
```

### instalando a partir do repo local
```console
# Helm
$ helm install -f values.yaml my-kube-prometheus-stack ./kube-prometheus-stack/ --debug
```

## Download manifests
```console
# Helm
$ helm pull prometheus-community/kube-prometheus-stack --version 18.0.2
```
## Upgrading Chart

```console
# Helm
$ helm upgrade [RELEASE_NAME] prometheus-community/kube-prometheus-stack
```

## Uninstall Chart

```console
# Helm
$ helm uninstall [RELEASE_NAME]
```

# Alterações feitos no projeto padrão

- alteração no values.yaml parametro config, parâmetro para o alertmanager.yaml
```yaml
  config:
    global:
      slack_api_url: https://discordapp.com/api/webhooks/0000000000000000/X0X0X0X0X0X0X0-X0X0X0X0X0X0X-0X0X0X0X0X-XO/slack
      resolve_timeout: 5m
    route:
      group_by: ['alertname']
      group_wait: 1m
      group_interval: 1m
      repeat_interval: 1h
      receiver: 'teste'
      routes:
      - match:
          severity: critical
        receiver: 'teste'
    receivers:
    - name: 'teste'
      slack_configs:
        - api_url: https://discordapp.com/api/webhooks/0000000000000000/X0X0X0X0X0X0X0-X0X0X0X0X0X0X-0X0X0X0X0X-XO/slack
          send_resolved: true
          # channel: 'testes-de-monitoramento'
          title: "\n({{ .Status }}): Sim-alerts"
          icon_url: https://avatars3.githubusercontent.com/u/3380462
          text: "\nsummary: {{ .CommonAnnotations.summary}}\ndescription: {{ .CommonAnnotations.description }}"
    templates:
    - '/etc/alertmanager/config/*.tmpl'
```

- comentários em algumas rules que não seriam necessárias, no values.yaml, esse parâmetros definem qual rules serão usados para alertas no prometheus.

```yaml
defaultRules:
  create: true
  rules:
    alertmanager: true
    # etcd: true
    # general: true
    # k8s: true
    # kubeApiserver: true
    # kubeApiserverAvailability: true
    # kubeApiserverError: true
    # kubeApiserverSlos: true
    # kubelet: true
    # kubePrometheusGeneral: true
    # kubePrometheusNodeAlerting: true
    # kubePrometheusNodeRecording: true
    # kubernetesAbsent: true
    kubernetesApps: true
    kubernetesResources: true
    # kubernetesStorage: true
    # kubernetesSystem: true
    # kubeScheduler: true
    # kubeStateMetrics: true
    # network: true
    node: true
    # prometheus: true
    # prometheusOperator: true
    # time: true

```

- criação de uma role a mais junto com os parametros de node, para alertar sobre a memória média do cluster no arquivo node-exporter.yaml

```yaml
    - alert: high_memory_load
      annotations:
        description: "Memória do nó com mais de 85%"
        runbook_url: https://github.com/stefanprodan/dockprom/blob/master/prometheus/alert.rules
        summary: "Server memory is almost full"
      expr: |-
        (
          (sum(node_memory_MemTotal_bytes) - sum(node_memory_MemFree_bytes + node_memory_Buffers_bytes + node_memory_Cached_bytes) ) / sum(node_memory_MemTotal_bytes) * 100 > 85
        )
      for: 5m
      labels:
        severity: critical
{{- if .Values.defaultRules.additionalRuleLabels }}
{{ toYaml .Values.defaultRules.additionalRuleLabels | indent 8 }}
{{- end }}
```

- parâmetro para senha que pedirá no grafana

```yaml
  adminPassword: prom-operator
```
## Dependencies

By default this chart installs additional, dependent charts:

- [prometheus-community/kube-state-metrics](https://github.com/prometheus-community/helm-charts/tree/main/charts/kube-state-metrics)
- [prometheus-community/prometheus-node-exporter](https://github.com/prometheus-community/helm-charts/tree/main/charts/prometheus-node-exporter)
- [grafana/grafana](https://github.com/grafana/helm-charts/tree/main/charts/grafana)

# DICA!!!!!!!
o prometheus tem alguns problemas na hora de upgrade or causa de validações.

use os seguintes comando de install e upgrade para não fazer a validação de endpoint.
links do issue :
- https://github.com/helm/charts/issues/19928
- https://lyz-code.github.io/blue-book/devops/prometheus/prometheus_troubleshooting/

```
$ helm upgrade -f values.yaml prometheus ./kube-prometheus-stack/ --debug --set prometheusOperator.admissionWebhooks.enabled=false --set prometheusOperator.admissionWebhooks.patch.enabled=false --set prometheusOperator.tlsProxy.enabled=false

$ helm install -f values.yaml prometheus ./kube-prometheus-stack/ --debug --set prometheusOperator.admissionWebhooks.enabled=false --set prometheusOperator.admissionWebhooks.patch.enabled=false --set prometheusOperator.tlsProxy.enabled=false
```

## Acessando aos serviços

```
kubectl port-forward svc/alertmanager-operated 9093
kubectl port-forward svc/prometheus-operated 90
kubectl port-forward deployment/prometheus-stack-grafana 3000
```