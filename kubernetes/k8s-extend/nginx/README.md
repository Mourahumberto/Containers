# Instalação Nginx
- Nesse tópico iremos instalar o nginx de maneira simples.

## No minikube
- DOCs
    - [minikube-ingress](https://kubernetes.io/docs/tasks/access-application-cluster/ingress-minikube/)

## Por manifesto
- seguindo essa [Doc](https://kubernetes.github.io/ingress-nginx/deploy/).
vemos como instalar o nginx em vários providers a partir dos manifestos do kubernetes. Usando apenas um "kubernetes apply -f " e mudando algumas variáveis.
- instalando para a [AWS](https://kubernetes.github.io/ingress-nginx/deploy/#aws)

- Edite o arquive e mude o VPC CIDR usado pelo seu cluster k8s e mude seu certificado digital também.

```yaml
proxy-real-ip-cidr: XXX.XXX.XXX/XX
...
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

# Nginx com features de firewall
- É importante colocar algumas regras de firewall no seu nginx, [annotations](https://github.com/kubernetes/ingress-nginx/blob/main/docs/user-guide/nginx-configuration/annotations.md) suportados no nginx.

## Rate Limit
- ferramenta importante contra atacks ddos. Pois ela coloca um limite máximo de reqs por minutos de um determinado IP.
- no anotation na criação do ingress colocar a seguinte anotação.
- Docs legais:
  -[rate limit e teste](https://www.nginx.com/blog/microservices-march-protect-kubernetes-apis-with-rate-limiting/)

```yaml
apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: minimal-ingress
  annotations:
    nginx.ingress.kubernetes.io/rewrite-target: /$1
    nginx.ingress.kubernetes.io/limit-rpm: "10"

spec:
  ingressClassName: nginx
  rules:
  - http:
      paths:
      - path: /teste
        pathType: Prefix
        backend:
          service:
            name: nginx
            port:
              number: 80
      - path: /teste2
        pathType: Prefix
        backend:
          service:
            name: nginx2
            port:
              number: 80
```

- Deploy de uma aplicação com ingress simples.
```
$ kubectl apply -f app-ratelimit.yaml
```

## modsecurity no nginx [WAF](https://github.com/kubernetes/ingress-nginx/blob/main/docs/user-guide/nginx-configuration/annotations.md#modsecurity)

## Habilitando o modsecurity o modsecurity, com os manifestos.

- Exemplo que eu me basiei [DOC](https://thelinuxnotes.com/index.php/how-to-install-and-configure-modsecurity-waf-in-kubernetes/#google_vignette)

- Primeiro você altera o configmap na parte data: e adiciona algumas labels.
```
data:
  allow-snippet-annotations: "true"
  enable-modsecurity: "true"
  enable-owasp-modsecurity-crs: "true"
```

- Segundo você cria annotations em seu ingress
```
  annotations:
    nginx.ingress.kubernetes.io/limit-rpm: "40"
    nginx.ingress.kubernetes.io/proxy-body-size: 8m
    nginx.ingress.kubernetes.io/enable-modsecurity: "true"
    nginx.ingress.kubernetes.io/enable-owasp-core-rules: "true"   
    nginx.ingress.kubernetes.io/modsecurity-snippet: |
     SecDebugLog /tmp/modsec_debug.log 
     SecRuleEngine On
     SecRequestBodyAccess On
```
- Deploy the manifest:
```
$ kubectl apply -f deploy-modsecurity.yaml
```

- Deploy de uma aplicação com ingress simples.
```
$ kubectl apply -f app-modsecurity.yaml
```

- exemplo de curl teste para testar o modsecurity.
```bash
curl 'https://seudominio.com.br/seupath/?param="><script>alert(1);</script>'
curl -X POST https://seudominio.com.br/seupath  -F "user='<script><alert>Hello></alert></script>'"
```

### DOCS interessantes
- Exemplo de como criar sua própria regra [página-web](https://thelinuxnotes.com/index.php/how-to-install-and-configure-modsecurity-waf-in-kubernetes/#google_vignette)
- configs mod-security [página web](https://github.com/owasp-modsecurity/ModSecurity/blob/v3/master/modsecurity.conf-recommended)
- Anottations [página web](https://kubernetes.github.io/ingress-nginx/user-guide/nginx-configuration/annotations/#modsecurity)
