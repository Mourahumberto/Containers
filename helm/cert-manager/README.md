# Doc para instalação e customização do prometheus
 - Documentação oficial: https://artifacthub.io/packages/helm/cert-manager/cert-manager

# kube-prometheus-stack


## Prerequisites

- Kubernetes 1.18+
- Helm 3+

## Installing the Chart 
- instala os CRD's, esse recurso é independente do chart, o chart pode ser instalado e feito upgrade independente dos crd's
```
kubectl apply -f https://github.com/jetstack/cert-manager/releases/download/v1.7.2/cert-manager.crds.yaml
```

## Get Repo Info

```console
helm repo add jetstack https://charts.jetstack.io
helm repo update
```
## Install Chart 

### versão específica
```console
# Helm
$ helm install [my-release] --namespace cert-manager --version v1.7.2 jetstack/cert-manager
```
### valores específicos
```console
# Helm
$ helm install -f values.yaml [my-release] --namespace cert-manager --version v1.7.2 jetstack/cert-manager
```

### instalando a partir do repo local
```console
# Helm
$ helm install -f values.yaml [my-release] --namespace cert-manager ./certmanager/ --debug
```

## Download manifests
```console
# Helm
$ helm pull jetstack/cert-manager --version 18.0.2
```
## Upgrading Chart

```console
# Helm
$ helm upgrade [RELEASE_NAME] prometheus-community/kube-prometheus-stack --version v1.7.2
```

## Uninstall Chart

```console
# Helm
$ helm uninstall [RELEASE_NAME]
```

# Próximos passos

- Em sequencia precisará criar uma emissão de certificado, precisará de um ClusterIssuer
tipos de emissões suportadas https://cert-manager.io/docs/configuration/
- será usado o ACME e criaremos um basic ACME Issuer
- Primeiro será criado um ```staging_issuer.yaml```. Para que seja testado a emissão do certificado.


```yaml
apiVersion: cert-manager.io/v1alpha2
kind: ClusterIssuer
metadata:
 name: letsencrypt-staging
 namespace: cert-manager
spec:
 acme:
   # The ACME server URL
   server: https://acme-staging-v02.api.letsencrypt.org/directory
   # Email address used for ACME registration
   email: your_email_address_here
   # Name of a secret used to store the ACME account private key
   privateKeySecretRef:
     name: letsencrypt-staging
   # Enable the HTTP-01 challenge provider
   solvers:
   - http01:
       ingress:
         class:  nginx
```
```$ kubectl create -f staging_issuer.yaml```

- agora criaremos uma aplicação que use o certificado. ```site.yaml```

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: nginx
  labels:
    app: nginx
spec:
  replicas: 1
  selector:
    matchLabels:
      app: nginx
  template:
    metadata:
      labels:
        app: nginx
    spec:
      containers:
      - name: nginx
        image: nginx:1.14.2
        ports:
        - containerPort: 80

---

apiVersion: v1
kind: Service
metadata:
  name: nginx
  labels:
    app: nginx
spec:
  ports:
    - port: 8080 # the port that this service should serve on
      targetPort: 80 #Container Port
      protocol: TCP
  selector:
    app: nginx

---

apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: example
  namespace: default
  annotations:
    cert-manager.io/cluster-issuer: "letsencrypt-staging"
spec:
  tls:
  - hosts:
    - <seu.dominio>.com.br
    secretName: site-nginx-teste
  ingressClassName: nginx
  rules:
    - host: <seu.dominio>.com.br
      http:
        paths:
          - backend:
              service:
                name: nginx
                port:
                  number: 8080
            path: /
            pathType: Prefix

```
```kubectl apply -f site.yaml```
- para verificar se o certificado foi criado
```kubectl describe ingress```

- será emitido o certificado. Porém não estará funcionando o https ainda. Para isso ccriaremos um issuer com o servidor do letsencrypt de prod.

```yaml
apiVersion: cert-manager.io/v1alpha2
kind: ClusterIssuer
metadata:
  name: letsencrypt-prod
  namespace: cert-manager
spec:
  acme:
    # The ACME server URL
    server: https://acme-v02.api.letsencrypt.org/directory
    # Email address used for ACME registration
    email: your_email_address_here
    # Name of a secret used to store the ACME account private key
    privateKeySecretRef:
      name: letsencrypt-prod
    # Enable the HTTP-01 challenge provider
    solvers:
    - http01:
        ingress:
          class: nginx
```
```kubectl create -f prod_issuer.yaml```
- e altera no ingress o anotation

```yaml
  annotations:
    cert-manager.io/cluster-issuer: "letsencrypt-prod"
```
- verifica o certificado
```kubectl describe certificate <secretname>```

DOC complementar: https://www.digitalocean.com/community/tutorials/how-to-set-up-an-nginx-ingress-with-cert-manager-on-digitalocean-kubernetes-pt