# Instalação Nginx
- Nesse tópico iremos instalar o nginx de maneira simples.

## No minikube
- DOCs
    - [minikube-ingress](https://kubernetes.io/docs/tasks/access-application-cluster/ingress-minikube/)

## Por manifesto
- seguindo essa [Doc](https://kubernetes.github.io/ingress-nginx/deploy/).
vemos como instalar o nginx em vários providers a partir dos manifestos do kubernetes. Usando apenas um "kubernetes apply -f " e mudando algumas variáveis.
- instalando para a [AWS](https://kubernetes.github.io/ingress-nginx/deploy/#aws)

- Edite o arquive e mude o VPC CIDR usado pelo seu cluster k8s:
```
proxy-real-ip-cidr: XXX.XXX.XXX/XX
```
- Mude seu certificado digital também:
```
arn:aws:acm:us-west-2:XXXXXXXX:certificate/XXXXXX-XXXXXXX-XXXXXXX-XXXXXXXX
```

- Deploy the manifest:
```
$ kubectl apply -f deploy-simples.yaml
```

- Deploy de uma aplicação com ingress simples.
```
$ kubectl apply -f app-simples.yaml
```

- Deploy de uma aplicação com ingress com ratelimit.
```
$ kubectl apply -f app-ratelimit.yaml
```