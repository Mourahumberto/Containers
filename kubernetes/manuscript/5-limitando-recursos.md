# Limitando Recursos

Quando criamos um Pod podemos especificar a quantidade de CPU e Memória (RAM) que pode ser consumida em cada contêiner. Quando algum contêiner contém a configuração de limite de recursos o Scheduler fica responsável por alocar esse contêiner no melhor nó possível de acordo com os recursos disponíveis.

Podemos configurar dois tipos de recursos, CPU que é especificada em **unidades de núcleos** e Memória que é especificada em **unidades de bytes**.

Vamos criar nosso primeiro Deployment com limite de recursos, para isso vamos subir a imagem de um ``nginx`` e copiar o ``yaml`` do deployment com o seguinte comando.

Crie o seguinte arquivo:

```
vim deployment-limitado.yaml
```

O conteúdo deve ser o seguinte.

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: nginx
  name: nginx
  namespace: default
spec:
  progressDeadlineSeconds: 600
  replicas: 3
  revisionHistoryLimit: 10
  selector:
    matchLabels:
      app: nginx
  strategy:
    rollingUpdate:
      maxSurge: 25%
      maxUnavailable: 25%
    type: RollingUpdate
  template:
    metadata:
      creationTimestamp: null
      labels:
        app: nginx
    spec:
      containers:
      - image: nginx
        imagePullPolicy: Always
        name: nginx
   # Adicione as seguintes linhas
        resources:
          limits:
            memory: "256Mi"
            cpu: "200m"
          requests:
            memory: "128Mi"
            cpu: "50m"
        terminationMessagePath: /dev/termination-log
        terminationMessagePolicy: File
      dnsPolicy: ClusterFirst
      restartPolicy: Always
      schedulerName: default-scheduler
      securityContext: {}
      terminationGracePeriodSeconds: 30
```

> Atenção! **1 core de CPU** corresponde a ``1000m`` (1000 milicore). Ao especificar ``200m``, estamos querendo reservar 20% de 1 core da CPU. Se fosse informado o valor ``0.2`` teria o mesmo efeito, ou seja, seria reservado 20% de 1 core da CPU.

```
kubectl create -f deployment-limitado.yaml

deployment.apps/nginx created
```

Vamos acessar um contêiner e testar a configuração:

```
kubectl get pod

NAME                    READY   STATUS    RESTARTS   AGE
nginx                   1/1     Running   0          12m
nginx-f89759699-77v8b   1/1     Running   0          117s
nginx-f89759699-ffbgh   1/1     Running   0          117s
nginx-f89759699-vzvlt   1/1     Running   0          2m2s
```

Acessando o shell de um pod:

```
kubectl exec -ti nginx-f89759699-77v8b -- /bin/bash
```

Agora no contêiner, instale e execute o ``stress`` para simular a carga em nossos recursos, no caso CPU e memória.

Instalando o comando stress:

```
apt-get update && apt-get install -y stress
```

Executando o ``stress``:

```
stress --vm 1 --vm-bytes 128M --cpu 1

stress: info: [221] dispatching hogs: 1 cpu, 0 io, 1 vm, 0 hdd
```
Aqui estamos _stressando_ o contêiner, utilizando 128M de RAM e um core de CPU. Brinque de acordo com os limites que você estabeleceu.

Quando ultrapassar o limite configurado, você receberá um erro como mostrado ao executar o seguinte comando, pois ele não conseguirá alocar os recursos de memória:

```
stress --vm 1 --vm-bytes 512M --cpu 1

stress: info: [230] dispatching hogs: 1 cpu, 0 io, 1 vm, 0 hdd
stress: FAIL: [230] (415) <-- worker 232 got signal 9
stress: WARN: [230] (417) now reaping child worker processes
stress: FAIL: [230] (451) failed run completed in 0s
```

Para acompanhar a quantidade de recurso que o pod está utilizando, podemos utilizar o ``kubectl top``. Lembre-se de executar esse comando no node e não no contêiner. :)

```
kubectl top pod --namespace=default nginx-f89759699-77v8b

NAME                     CPU(cores)   MEMORY(bytes)
nginx-85f7fb6b45-b6dsk   201m         226Mi
```

# Namespaces

No kubernetes temos um cara chamado de Namespaces como já vimos anteriormente.

Mas o que é um Namespace? Nada mais é do que um cluster virtual dentro do próprio cluster físico do Kubernetes. Namespaces são uma maneira de dividir recursos de um cluster entre vários ambientes, equipes ou projetos.

Vamos criar nosso primeiro namespaces:

```
kubectl create namespace primeiro-namespace

namespace/primeiro-namespace created
```

Vamos listar todos os namespaces do kubernetes:

```
kubectl get namespaces

NAME                 STATUS   AGE
default              Active   55m
kube-node-lease      Active   56m
kube-public          Active   56m
kube-system          Active   56m
primeiro-namespace   Active   5s
```

Pegando mais informações do nosso namespace:

```
kubectl describe namespace primeiro-namespace

Name:         primeiro-namespace
Labels:       <none>
Annotations:  <none>
Status:       Active

No resource quota.

No LimitRange resource.
```

Como podemos ver, nosso namespace ainda está sem configurações, então iremos utilizar o ``LimitRange`` para adicionar limites de recursos.

Vamos criar o manifesto do ``LimitRange``:

```
vim limitando-recursos.yaml
```

O conteúdo deve ser o seguinte.

```yaml
apiVersion: v1
kind: LimitRange
metadata:
  name: limitando-recursos
spec:
  limits:
  - default:
      cpu: 1
      memory: 100Mi
    defaultRequest:
      cpu: 0.5
      memory: 80Mi
    type: Container
```

Agora vamos adicionar esse ``LimitRange`` ao Namespace:

```
kubectl create -f limitando-recursos.yaml -n primeiro-namespace

limitrange/limitando-recursos created
```

Listando o ``LimitRange``:

```
kubectl get limitranges

No resources found in default namespace.
```

Opa! Não encontramos não é mesmo? Mas claro, esquecemos de passar nosso namespace na hora de listar:

```
kubectl get limitrange -n primeiro-namespace

NAME                 CREATED AT
limitando-recursos   2020-05-10T18:02:51Z
```

Ou:

```
kubectl get limitrange --all-namespaces

NAMESPACE            NAME                 CREATED AT
primeiro-namespace   limitando-recursos   2020-05-10T18:02:51Z
```

Vamos dar um describe no ``LimitRange``:

```
kubectl describe limitrange -n primeiro-namespace

Name:       limitando-recursos
Namespace:  primeiro-namespace
Type        Resource  Min  Max  Default Request  Default Limit  Max Limit/Request Ratio
----        --------  ---  ---  ---------------  -------------  -----------------------
Container   cpu       -    -    500m             1              -
Container   memory    -    -    80Mi             100Mi          -
```

Como podemos observar, adicionamos limites de memória e cpu para cada contêiner que subir nesse ``Namespace``, se algum contêiner for criado dentro do ``Namespace`` sem as configurações de ``Limitrange``, o contêiner irá herdar as configurações de limites de recursos do Namespace.

Vamos criar um pod para verificar se o limite se aplicará:

```
vim pod-limitrange.yaml
```

O conteúdo deve ser o seguinte.

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: limit-pod
spec:
  containers:
  - name: meu-container
    image: nginx
```

Agora vamos criar um pod fora do namespace default e outro dentro do namespace limitado (primeiro-namespace), e vamos observar os limites de recursos de cada contêiner e como foram aplicados:

Criando o pod no namespace ``default``:

```
kubectl create -f pod-limitrange.yaml

pod/limit-pod created
```

Criando o pod no namespace ``primeiro-namespace``:

```
kubectl create -f pod-limitrange.yaml -n primeiro-namespace

pod/limit-pod created
```

Vamos listar esses pods e na sequência ver mais detalhes :

```
kubectl get pods --all-namespaces

NAMESPACE            NAME                                      READY   STATUS    RESTARTS   AGE
default              limit-pod                                 1/1     Running   0          10s
primeiro-namespace   limit-pod                                 1/1     Running   0          5s
```

Vamos ver mais detalhes do pod no namespace ``default``.

```
kubectl describe pod limit-pod

Name:         limit-pod
Namespace:    default
Priority:     0
Node:         elliot-02/172.31.19.123
Start Time:   Sun, 10 May 2020 18:03:52 +0000
Labels:       <none>
Annotations:  <none>
Status:       Running
IP:           10.32.0.4
IPs:
  IP:  10.32.0.4
Containers:
  meu-container:
    Container ID:   docker://19850dc935ffa41b1754cb58ab60ec5bb3616bbbe6a958abe1b2575bd26ee73d
    Image:          nginx
...
```

Vamos ver mais detalhes do pod no namespace ``primeiro-namespace``.

```
kubectl describe pod limit-pod -n primeiro-namespace

Name:         limit-pod
Namespace:    primeiro-namespace
Priority:     0
Node:         elliot-03/172.31.24.60
Start Time:   Sun, 10 May 2020 18:03:57 +0000
Labels:       <none>
Annotations:  kubernetes.io/limit-ranger:
                LimitRanger plugin set: cpu, memory request for container meu-container; cpu, memory limit for container meu-container
Status:       Running
IP:           10.46.0.3
IPs:
  IP:  10.46.0.3
Containers:
  meu-container:
    Container ID:   docker://ad3760837d71955df47263a2031d4238da2e94002ce4a0631475d301faf1ddef
    Image:          nginx
...
    Limits:
      cpu:     1
      memory:  100Mi
    Requests:
      cpu:        500m
      memory:     80Mi
```

Como podemos ver o **Pod** no namespace ``primeiro-namespace`` está com limit de recursos configurados.