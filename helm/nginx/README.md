# Instalação Nginx
- Nesse tópico iremos instalar o nginx de maneira simples.

# Pelo helm
- instalando com values default


```bash
$ helm repo add ingress-nginx https://kubernetes.github.io/ingress-nginx
$ helm repo update
$ helm install -f values-nginx.yaml uniq-ingress-nginx ingress-nginx/ingress-nginx -n nginx --create-namespace
$ helm upgrade -f values-nginx.yaml uniq-ingress-nginx ingress-nginx/ingress-nginx -n nginx

- neste values, já está com anotations para ssl com acm e um nlb. Caso não precise baixe o values default do nginx
```
outros comandos
```bash
$ helm list --all-namespaces
$ helm show values ingress-nginx/ingress-nginx > values-nginx.yaml
$ helm install -f prod/values-nginx.yaml uniq-ingress-nginx ingress-nginx/ingress-nginx -n nginx --create-namespace
$ helm upgrade -f prod/values-nginx.yaml uniq-ingress-nginx ingress-nginx/ingress-nginx -n nginx

```

[repo helm](https://artifacthub.io/packages/helm/ingress-nginx/ingress-nginx)

