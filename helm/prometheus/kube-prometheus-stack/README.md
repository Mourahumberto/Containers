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

- parâmetro para senha que pedirá no grafana

```yaml
  adminPassword: prom-operator
```
## Dependencies

By default this chart installs additional, dependent charts:

- [prometheus-community/kube-state-metrics](https://github.com/prometheus-community/helm-charts/tree/main/charts/kube-state-metrics)
- [prometheus-community/prometheus-node-exporter](https://github.com/prometheus-community/helm-charts/tree/main/charts/prometheus-node-exporter)
- [grafana/grafana](https://github.com/grafana/helm-charts/tree/main/charts/grafana)

## Acessando aos serviços

```
kubectl port-forward svc/alertmanager-operated 9093
kubectl port-forward svc/prometheus-operated 90
kubectl port-forward deployment/prometheus-stack-grafana 3000
```
