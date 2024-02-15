https://artifacthub.io/packages/helm/ingress-nginx/ingress-nginx

https://kubernetes.github.io/ingress-nginx/deploy/#aws
1)helm upgrade --install ingress-nginx ingress-nginx   --repo https://kubernetes.github.io/ingress-nginx   --namespace ingress-nginx --create-namespace

2) helm list --all-namespaces
3)helm show values ingress-nginx/ingress-nginx > values-nginx.yaml
