
# Services
- É uma maneira de de expor um conjunto de app, em um endpoint único, desta forma é dado um dns interno e um ip único que balancea a carga entre os pods.
- O service usa selectores para encontar os pods que serão usados para balancear a carga, desta forma ele trabalha como service discovery.

## Multi-Port Services
- Um serviço pode expor multiplas portas.

```yaml
apiVersion: v1
kind: Service
metadata:
  name: my-service
spec:
  selector:
    app.kubernetes.io/name: MyApp
  ports:
    - name: http
      protocol: TCP
      port: 80
      targetPort: 9376
    - name: https
      protocol: TCP
      port: 443
      targetPort: 9377
```
## Discovering services
- existem duas formas de o Pod encontrar services dentro do cluster. Um é através de variáveis de ambientes todo container tem as variáveis de ambientes com o ip do serviço
e a porta desta forma:"{SVCNAME}_SERVICE_HOST and {SVCNAME}_SERVICE_PORT". è importante saber que isso acontece quando o serviço foi criado antes do pod. se o service
for criado depois do pod, pode ser que não tenha a vari´qavel do service dentro do container.
- Outra forma de encontrar os services é a través do DNS, todo service tem um dns ***myservice.myns***
- O Kubernetes também oferece suporte a registros DNS SRV (Serviço) para portas nomeadas. Se o serviço my-service.my-ns tiver uma porta chamada http com o protocolo definido como TCP, você poderá fazer uma consulta DNS SRV para _http._tcp.my-service.my-ns para descobrir o número da porta para http, conforme bem como o endereço IP.


## Tipo ClusterIp
- Expõe o service com um ip interno visto dentro do cluster.

Vamos criar um pod a partir de um pod template utilizando os seguintes comandos:

```
kubectl run nginx --image nginx --dry-run=client -o yaml > pod-template.yaml
kubectl create -f pod-template.yaml
pod/nginx created
```
- Criando o service:

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

Com isso, agora temos como manter a sessão, ou seja, ele irá manter a conexão com o mesmo pod, respeitando o IP de origem do cliente.

Caso precise, é possível alterar o valor do timeout para o ``sessionAffinity`` (O valor padrão é de 10800 segundos, ou seja 3 horas), apenas adicionando a configuração abaixo.

```
  sessionAffinityConfig:
    clientIP:
      timeoutSeconds: 10
```
## Criando um service NodePort

Execute o comando a seguir para exportar o pod usando o service NodePort. Lembrando que o range de portas internas é entre 30000/TCP a 32767/TCP.

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

## Criando um service LoadBalancer

Execute o comando a seguir para exportar o pod usando o service LoadBalancer.

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

# Network Policy
- Você consegue definir políticas de rede, de ingress e egress para os pods e namespaces. 
- Quando definimos uma política de rede baseada em pod ou namespace, utiliza-se um selector para especificar qual tráfego é permitido de e para o(s) Pod(s) que correspondem ao seletor.
- Quando uma política de redes baseada em IP é criada, nós definimos a política baseada em blocos de IP (faixas CIDR).
- por padrão todos os pods são não isolados.
Ex: polices
https://kubernetes.io/pt-br/docs/concepts/services-networking/network-policies/

# DNS para Services e Pods
- O kubernetes cria um nome no dns para cada pod e service.
- o DNS do kubernetes pode ser visto em todos os pods, o kubelet cria o arquivo em todos os pods. /etc/resolv.conf
- Ele herda também o dns do próprio host (pelo que eu vi)

- Services records: my-svc.my-namespace.svc.cluster-domain.example e _my-port-name._my-port-protocol.my-svc.my-namespace.svc.cluster-domain.example

## Pod's DNS Policy
- você pode criar policys para que os pods não usem o dns padrão do kubernetes e use um dns server criado por você https://kubernetes.io/docs/concepts/services-networking/dns-pod-service/#pod-dns-config

## Routing em uma mesma zona
https://kubernetes.io/docs/concepts/services-networking/topology-aware-hints/

## routing no mesmo nó
https://kubernetes.io/docs/concepts/services-networking/service-traffic-policy/