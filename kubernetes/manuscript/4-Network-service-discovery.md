# Services

## Criando um service ClusterIP

Vamos criar um pod a partir de um pod template utilizando os seguintes comandos:

```
kubectl run nginx --image nginx --dry-run=client -o yaml > pod-template.yaml
kubectl create -f pod-template.yaml
pod/nginx created
```

Expondo o pod do Nginx.

```
kubectl expose pod nginx --port=80

service/nginx exposed
```

Obtendo informações do service.

```
kubectl get svc

NAME         TYPE        CLUSTER-IP       EXTERNAL-IP   PORT(S)   AGE
kubernetes   ClusterIP   10.96.0.1        <none>        443/TCP   25m
nginx        ClusterIP   10.104.209.243   <none>        80/TCP    7m15s
```

Execute o seguinte comando para visualizar mais detalhes do service ``nginx``.

```
kubectl describe service nginx

Name:              nginx
Namespace:         default
Labels:            run=nginx
Annotations:       <none>
Selector:          run=nginx
Type:              ClusterIP
IP:                10.104.209.243
Port:              <unset>  80/TCP
TargetPort:        80/TCP
Endpoints:         10.46.0.0:80
Session Affinity:  None
Events:            <none>
```

Acessando o Ningx. Altere o IP do cluster no comando a seguir de acordo com o seu ambiente.

```
curl 10.104.209.243

...
<title>Welcome to nginx!</title>
...
```
Acessando fazendo port-foward no serviço

```
kubectl port-foward svc/nginx 9090:80 -n onamespace
```

Acessando fazendo port-foward no pod
```
kubectl port-foward podname <portlocal>:<port-pod> -n onamespace
```

Acesse o log do Nginx.

```
kubectl logs -f nginx

10.40.0.0 - - [10/May/2020:17:31:56 +0000] "GET / HTTP/1.1" 200 612 "-" "curl/7.58.0" "-"
```

Remova o serviço criado anteriormente.

```
kubectl delete svc nginx

service "nginx" deleted
```

Agora vamos criar nosso service ``ClusterIP``, porém vamos criar um arquivo yaml com suas definições:

```
vim primeiro-service-clusterip.yaml
```

Informe o seguinte conteúdo:

```yaml
apiVersion: v1
kind: Service
metadata:
  labels:
    run: nginx
  name: nginx-clusterip
  namespace: default
spec:
  ports:
  - port: 80
    protocol: TCP
    targetPort: 80
  selector:
    run: nginx
  type: ClusterIP
```

Criando o service:

```
kubectl create -f primeiro-service-clusterip.yaml

service/nginx-clusterip created
```

Obtendo informações do service:

```
kubectl get services

NAME              TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)   AGE
kubernetes        ClusterIP   10.96.0.1       <none>        443/TCP   28m
nginx-clusterip   ClusterIP   10.109.70.243   <none>        80/TCP    71s
```

Visualizando os detalhes do service:

```
kubectl describe service nginx-clusterip

Name:              nginx-clusterip
Namespace:         default
Labels:            run=nginx
Annotations:       <none>
Selector:          run=nginx
Type:              ClusterIP
IP:                10.109.70.243
Port:              <unset>  80/TCP
TargetPort:        80/TCP
Endpoints:         10.46.0.1:80
Session Affinity:  None
Events:            <none>
```

Removendo o service:

```
kubectl delete -f primeiro-service-clusterip.yaml

service "nginx-clusterip" deleted
```

Agora vamos mudar um detalhe em nosso manifesto, vamos brincar com o nosso ``sessionAffinity``:

> **Nota:** Se você quiser ter certeza de que as conexões de um cliente específico sejam passadas para o mesmo pod todas as vezes, você pode selecionar a afinidade da sessão (*session affinity*) com base nos endereços IP do cliente, definindo ``service.spec.sessionAffinity`` como ``ClientIP`` (o padrão é ``None``). Você também pode definir o tempo de permanência máximo da sessão definindo ``service.spec.sessionAffinityConfig.clientIP.timeoutSeconds`` adequadamente (o valor padrão é 10800 segundos, o que resulta em 3 horas).

```
vim primeiro-service-clusterip.yaml
```

O conteúdo deve ser o seguinte:

```yaml
apiVersion: v1
kind: Service
metadata:
  labels:
    run: nginx
  name: nginx-clusterip
  namespace: default
spec:
  ports:
  - port: 80
    protocol: TCP
    targetPort: 80
  selector:
    run: nginx
  sessionAffinity: ClientIP
  type: ClusterIP
```

Criando o service novamente:

```
kubectl create -f primeiro-service-clusterip.yaml

service/nginx-clusterip created
```

Obtendo informações do service:

```
kubectl get services

NAME              TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)   AGE
kubernetes        ClusterIP   10.96.0.1      <none>        443/TCP   29m
nginx-clusterip   ClusterIP   10.96.44.114   <none>        80/TCP    7s
```

Visualizando os detalhes do service:

```
kubectl describe service nginx

Name:              nginx-clusterip
Namespace:         default
Labels:            run=nginx
Annotations:       <none>
Selector:          run=nginx
Type:              ClusterIP
IP:                10.96.44.114
Port:              <unset>  80/TCP
TargetPort:        80/TCP
Endpoints:         10.46.0.1:80
Session Affinity:  ClientIP
Events:            <none>
```

Com isso, agora temos como manter a sessão, ou seja, ele irá manter a conexão com o mesmo pod, respeitando o IP de origem do cliente.

Caso precise, é possível alterar o valor do timeout para o ``sessionAffinity`` (O valor padrão é de 10800 segundos, ou seja 3 horas), apenas adicionando a configuração abaixo.

```
  sessionAffinityConfig:
    clientIP:
      timeoutSeconds: 10
```

Agora podemos remover o service:

```
kubectl delete -f primeiro-service-clusterip.yaml

service "nginx-clusterip" deleted
```

## Criando um service NodePort

Execute o comando a seguir para exportar o pod usando o service NodePort. Lembrando que o range de portas internas é entre 30000/TCP a 32767/TCP.

```
kubectl expose pods nginx --type=NodePort --port=80

service/nginx exposed
```

Obtendo informações do service:

```
kubectl get svc

NAME         TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)        AGE
kubernetes   ClusterIP   10.96.0.1       <none>        443/TCP        29m
nginx        NodePort    10.101.42.230   <none>        80:31858/TCP   5s
```

Removendo o service:

```
kubectl delete svc nginx

service "nginx" deleted
```

Agora vamos criar um service NodePort, porém vamos criar um manifesto yaml com suas definições.

```
vim primeiro-service-nodeport.yaml
```

O conteúdo deve ser o seguinte.

```yaml
apiVersion: v1
kind: Service
metadata:
  labels:
    run: nginx
  name: nginx-nodeport
  namespace: default
spec:
  externalTrafficPolicy: Cluster
  ports:
  - nodePort: 31111
    port: 80
    protocol: TCP
    targetPort: 80
  selector:
    run: nginx
  sessionAffinity: None
  type: NodePort
```

Criando o service:

```
kubectl create -f primeiro-service-nodeport.yaml

service/nginx-nodeport created
```

Obtendo informações do service:

```
kubectl get services

NAME             TYPE        CLUSTER-IP     EXTERNAL-IP   PORT(S)        AGE
kubernetes       ClusterIP   10.96.0.1      <none>        443/TCP        30m
nginx-nodeport   NodePort    10.102.91.81   <none>        80:31111/TCP   7s
```

Visualizando os detalhes do service:

```
kubectl describe service nginx

Name:                     nginx-nodeport
Namespace:                default
Labels:                   run=nginx
Annotations:              <none>
Selector:                 run=nginx
Type:                     NodePort
IP:                       10.102.91.81
Port:                     <unset>  80/TCP
TargetPort:               80/TCP
NodePort:                 <unset>  31111/TCP
Endpoints:                10.46.0.1:80
Session Affinity:         None
External Traffic Policy:  Cluster
Events:                   <none>
```

Removendo o service:

```
kubectl delete -f primeiro-service-nodeport.yaml

service "nginx-nodeport" deleted
```

## Criando um service LoadBalancer

Execute o comando a seguir para exportar o pod usando o service LoadBalancer.

```
kubectl expose pod nginx --type=LoadBalancer --port=80

service/nginx exposed
```

Obtendo informações do service:

```
kubectl get svc

NAME         TYPE           CLUSTER-IP      EXTERNAL-IP   PORT(S)        AGE
kubernetes   ClusterIP      10.96.0.1       <none>        443/TCP        32m
nginx        LoadBalancer   10.110.198.89   <pending>     80:30728/TCP   4s
```

Removendo o service:

```
kubectl delete svc nginx

service "nginx" deleted
```

Agora vamos criar service LoadBalancer, porém vamos criar um yaml com suas definições.

```
vim primeiro-service-loadbalancer.yaml
```

O conteúdo deve ser o seguinte.

```yaml
apiVersion: v1
kind: Service
metadata:
  labels:
    run: nginx
  name: nginx-loadbalancer
  namespace: default
spec:
  externalTrafficPolicy: Cluster
  ports:
  - nodePort: 31222
    port: 80
    protocol: TCP
    targetPort: 80
  selector:
    run: nginx
  sessionAffinity: None
  type: LoadBalancer
```

Criando o service:

```
kubectl create -f primeiro-service-loadbalancer.yaml

service/nginx-loadbalancer created
```

Obtendo informações do service:

```
kubectl get services

NAME                 TYPE           CLUSTER-IP     EXTERNAL-IP   PORT(S)        AGE
kubernetes           ClusterIP      10.96.0.1      <none>        443/TCP        33m
nginx-loadbalancer   LoadBalancer   10.96.67.165   <pending>     80:31222/TCP   4s
```

Visualizando informações do service:

```
kubectl describe service nginx

Name:                     nginx-loadbalancer
Namespace:                default
Labels:                   run=nginx
Annotations:              <none>
Selector:                 run=nginx
Type:                     LoadBalancer
IP:                       10.96.67.165
Port:                     <unset>  80/TCP
TargetPort:               80/TCP
NodePort:                 <unset>  31222/TCP
Endpoints:                10.46.0.1:80
Session Affinity:         None
External Traffic Policy:  Cluster
Events:                   <none>
```

Removendo o service:

```
kubectl delete -f primeiro-service-loadbalancer.yaml

service "nginx-loadbalancer" deleted
```

## EndPoint

Sempre que criamos um service, automaticamente é criado um endpoint. O endpoint nada mais é do que o IP do pod que o service irá utilizar, por exemplo, quando criamos um service do tipo ClusterIP temos o seu IP, correto?

Agora, quando batemos nesse IP ele redireciona a conexão para o **Pod** através desse IP, o **EndPoint**.

Para listar os EndPoints criados, execute o comando:

```
kubectl get endpoints

NAME         ENDPOINTS         AGE
kubernetes   10.142.0.5:6443   4d
```

Vamos verificar esse endpoint com mais detalhes.

```
kubectl describe endpoints kubernetes

Name:         kubernetes
Namespace:    default
Labels:       <none>
Annotations:  <none>
Subsets:
  Addresses:          172.31.17.67
  NotReadyAddresses:  <none>
  Ports:
    Name   Port  Protocol
    ----   ----  --------
    https  6443  TCP

Events:  <none>
```

Removendo o deployment ``nginx``:

```
kubectl delete deployment nginx

deployment.apps "nginx" deleted
```

Removendo o service:

```
kubectl delete service nginx

service "nginx" deleted
```

