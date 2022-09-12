# Deployments

O **Deployment** é um recurso com a responsabilidade de instruir o Kubernetes a criar, atualizar e monitorar a saúde das instâncias de suas aplicações.

Um Deployment é o responsável por gerenciar o seu **ReplicaSet** , ou seja, o Deployment é quem vai determinar a configuração de sua aplicação e como ela será implementada. O Deployment é o **controller** que irá cuidar, por exemplo, uma instância de sua aplicação por algum motivo for interrompida. O **Deployment controller** irá identificar o problema com a instância e irá criar uma nova em seu lugar.

Quando você utiliza o ``kubectl create deployment``, você está realizando o deploy de um objeto chamado **Deployment**. Como outros objetos, o Deployment também pode ser criado através de um arquivo [YAML](https://en.wikipedia.org/wiki/YAML) ou de um [JSON](https://www.json.org/json-en.html), conhecidos por **manifestos**.

Se você deseja alterar alguma configuração de seus objetos, como o pod, você pode utilizar o ``kubectl apply``, através de um manifesto, ou ainda através do ``kubectl edit``. Normalmente, quando você faz uma alteração em seu Deployment, é criado uma nova versão do ReplicaSet, esse se tornando o ativo e fazendo com que seu antecessor seja desativado. As versões anteriores dos ReplicaSets são mantidas, possibilitando o _rollback_ em caso de falhas.

As **labels** são importantes para o gerenciamento do cluster, pois com elas é possível buscar ou selecionar recursos em seu cluster, fazendo com que você consiga organizar em pequenas categorias, facilitando assim a sua busca e organizando seus pods e seus recursos do cluster. As labels não são recursos do API server, elas são armazenadas no metadata em formato chave-valor.

Antes nos tínhamos somente o RC, _Replication Controller_, que era um controle sobre o número de réplicas que determinado pod estava executando, o problema é que todo esse gerenciamento era feito do lado do *client*. Para solucionar esse problema, foi adicionado o objeto Deployment, que permite a atualização pelo lado do *server*. **Deployments** geram **ReplicaSets**, que oferecerem melhores opções do que o **ReplicationController**, e por esse motivo está sendo substituído.

Podemos criar nossos deployments a partir do template:

```
kubectl create deployment --dry-run=client -o yaml --image=nginx nginx-template > primeiro-deployment-template.yaml
ou
kubectl-example deployment > primeiro-deployment-template.yaml
```
Vamos criar os nossos primeiros Deployments:
```
vim primeiro-deployment.yaml
```

O conteúdo deve ser o seguinte:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  creationTimestamp: null
  labels:
    app: nginx-template
  name: nginx-template
spec:
  replicas: 1
  selector:
    matchLabels:
      app: nginx-template
  strategy: {}
  template:
    metadata:
      creationTimestamp: null
      labels:
        app: nginx-template
    spec:
      containers:
      - image: nginx
        name: nginx
        resources: {}
status: {}

```

Vamos criar o deployment a partir do manifesto:

```
kubectl create -f primeiro-deployment.yaml

deployment.extensions/primeiro-deployment created
```

Crie um segundo deployment:

```
vim segundo-deployment.yaml
```

O conteúdo deve ser o seguinte:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    run: nginx
  name: segundo-deployment
  namespace: default
spec:
  replicas: 1
  selector:
    matchLabels:
      run: nginx
  template:
    metadata:
      labels:
        run: nginx
        dc: Netherlands
    spec:
      containers:
      - image: nginx
        imagePullPolicy: Always
        name: nginx2
        ports:
        - containerPort: 80
          protocol: TCP
        resources: {}
        terminationMessagePath: /dev/termination-log
        terminationMessagePolicy: File
      dnsPolicy: ClusterFirst
      restartPolicy: Always
      schedulerName: default-scheduler
      securityContext: {}
      terminationGracePeriodSeconds: 30
```

Vamos criar o deployment a partir do manifesto:

```
kubectl create -f segundo-deployment.yaml

deployment.extensions/segundo-deployment created
```

Visualizando os deployments:

```
kubectl get deployment

NAME                DESIRED   CURRENT   UP-TO-DATE   AVAILABLE   AGE
primeiro-deployment  1         1         1            1           6m
segundo-deployment   1         1         1            1           1m
```

Visualizando os pods:

```
kubectl get pods

NAME                                 READY  STATUS    RESTARTS   AGE
primeiro-deployment-68c9dbf8b8-kjqpt 1/1    Running   0          19s
segundo-deployment-59db86c584-cf9pp  1/1    Running   0          15s
```

Visualizando os detalhes do ``pod`` criado a partir do **primeiro deployment**:

```
kubectl describe pod primeiro-deployment-68c9dbf8b8-kjqpt

Name:               primeiro-deployment-68c9dbf8b8-kjqpt
Namespace:          default
Priority:           0
PriorityClassName:  <none>
Node:               elliot-02/10.138.0.3
Start Time:         Sat, 04 Aug 2018 00:45:29 +0000
Labels:             dc=UK
                    pod-template-hash=2475869464
                    run=nginx
Annotations:        <none>
Status:             Running
IP:                 10.46.0.1
Controlled By:      ReplicaSet/primeiro-deployment-68c9dbf8b8
Containers:
  nginx2:
    Container ID:   docker://963ec997a0aa4aa3cecabdb3c59f67d80e7010c51eac23735524899f7f2dd4f9
    Image:          nginx
    Image ID:       docker-pullable://nginx@sha256:d85914d547a6c92faa39ce7058bd7529baacab7e0cd4255442b04577c4d1f424
    Port:           80/TCP
    Host Port:      0/TCP
    State:          Running
      Started:      Sat, 04 Aug 2018 00:45:36 +0000
    Ready:          True
    Restart Count:  0
    Environment:    <none>
    Mounts:
      /var/run/secrets/kubernetes.io/serviceaccount from default-token-np77m (ro)
Conditions:
  Type              Status
  Initialized       True
  Ready             True
  ContainersReady   True
  PodScheduled      True
Volumes:
  default-token-np77m:
    Type:        Secret (a volume populated by a Secret)
    SecretName:  default-token-np77m
    Optional:    false
QoS Class:       BestEffort
Node-Selectors:  <none>
Tolerations:     node.kubernetes.io/not-ready:NoExecute for 300s
                 node.kubernetes.io/unreachable:NoExecute for 300s
Events:
  Type    Reason     Age   From                Message
  ----    ------     ----  ----                -------
  Normal  Scheduled  51s   default-scheduler   Successfully assigned default/primeiro-deployment-68c9dbf8b8-kjqpt to elliot-02
  Normal  Pulling    50s   kubelet, elliot-02  pulling image "nginx"
  Normal  Pulled     44s   kubelet, elliot-02  Successfully pulled image "nginx"
  Normal  Created    44s   kubelet, elliot-02  Created container
  Normal  Started    44s   kubelet, elliot-02  Started container
```

Visualizando os detalhes do ``pod`` criado a partir do **segundo deployment**:

```
kubectl describe pod segundo-deployment-59db86c584-cf9pp

Name:               segundo-deployment-59db86c584-cf9pp
Namespace:          default
Priority:           0
PriorityClassName:  <none>
Node:               elliot-02/10.138.0.3
Start Time:         Sat, 04 Aug 2018 00:45:49 +0000
Labels:             dc=Netherlands
                    pod-template-hash=1586427140
                    run=nginx
Annotations:        <none>
Status:             Running
IP:                 10.46.0.2
Controlled By:      ReplicaSet/segundo-deployment-59db86c584
Containers:
  nginx2:
    Container ID:   docker://a9e6b5463341e62eff9e45c8c0aace14195f35e41be088ca386949500a1f2bb0
    Image:          nginx
    Image ID:       docker-pullable://nginx@sha256:d85914d547a6c92faa39ce7058bd7529baacab7e0cd4255442b04577c4d1f424
    Port:           80/TCP
    Host Port:      0/TCP
    State:          Running
      Started:      Sat, 04 Aug 2018 00:45:51 +0000
    Ready:          True
    Restart Count:  0
    Environment:    <none>
    Mounts:
      /var/run/secrets/kubernetes.io/serviceaccount from default-token-np77m (ro)
Conditions:
  Type              Status
  Initialized       True
  Ready             True
  ContainersReady   True
  PodScheduled      True
Volumes:
  default-token-np77m:
    Type:        Secret (a volume populated by a Secret)
    SecretName:  default-token-np77m
    Optional:    false
QoS Class:       BestEffort
Node-Selectors:  <none>
Tolerations:     node.kubernetes.io/not-ready:NoExecute for 300s
                 node.kubernetes.io/unreachable:NoExecute for 300s
Events:
  Type    Reason     Age   From                Message
  ----    ------     ----  ----                -------
  Normal  Scheduled  2m    default-scheduler   Successfully assigned default/segundo-deployment-59db86c584-cf9pp to elliot-02
  Normal  Pulling    2m    kubelet, elliot-02  pulling image "nginx"
  Normal  Pulled     2m    kubelet, elliot-02  Successfully pulled image "nginx"
  Normal  Created    2m    kubelet, elliot-02  Created container
  Normal  Started    2m    kubelet, elliot-02  Started container
```

Visualizando os detalhes do **primeiro deployment**:

```
kubectl describe deployment primeiro-deployment

Name:                   primeiro-deployment
Namespace:              default
CreationTimestamp:      Sat, 04 Aug 2018 00:45:29 +0000
Labels:                 app=giropops
                        run=nginx
Annotations:            deployment.kubernetes.io/revision=1
Selector:               run=nginx
Replicas:               1 desired | 1 updated | 1 total | 1 available | 0 unavailable
StrategyType:           RollingUpdate
MinReadySeconds:        0
RollingUpdateStrategy:  1 max unavailable, 1 max surge
Pod Template:
  Labels:  dc=UK
           run=nginx
  Containers:
   nginx2:
    Image:        nginx
    Port:         80/TCP
    Host Port:    0/TCP
    Environment:  <none>
    Mounts:       <none>
  Volumes:        <none>
Conditions:
  Type           Status  Reason
  ----           ------  ------
  Available      True    MinimumReplicasAvailable
  Progressing    True    NewReplicaSetAvailable
OldReplicaSets:  <none>
NewReplicaSet:   primeiro-deployment-68c9dbf8b8 (1/1 replicas created)
Events:
  Type    Reason             Age   From                   Message
  ----    ------             ----  ----                   -------
  Normal  ScalingReplicaSet  3m    deployment-controller  Scaled up replica set primeiro-deployment-68c9dbf8b8 to 1
```

Visualizando os detalhes do **segundo deployment**:

```
kubectl describe deployment segundo-deployment

Name:                   segundo-deployment
Namespace:              default
CreationTimestamp:      Sat, 04 Aug 2018 00:45:49 +0000
Labels:                 run=nginx
Annotations:            deployment.kubernetes.io/revision=1
Selector:               run=nginx
Replicas:               1 desired | 1 updated | 1 total | 1 available | 0 unavailable
StrategyType:           RollingUpdate
MinReadySeconds:        0
RollingUpdateStrategy:  1 max unavailable, 1 max surge
Pod Template:
  Labels:  dc=Netherlands
           run=nginx
  Containers:
   nginx2:
    Image:        nginx
    Port:         80/TCP
    Host Port:    0/TCP
    Environment:  <none>
    Mounts:       <none>
  Volumes:        <none>
Conditions:
  Type           Status  Reason
  ----           ------  ------
  Available      True    MinimumReplicasAvailable
  Progressing    True    NewReplicaSetAvailable
OldReplicaSets:  <none>
NewReplicaSet:   segundo-deployment-59db86c584 (1/1 replicas created)
Events:
  Type    Reason             Age   From                   Message
  ----    ------             ----  ----                   -------
  Normal  ScalingReplicaSet  3m    deployment-controller  Scaled up replica set segundo-deployment-59db86c584 to 1
```

## Filtrando por Labels

Quando criamos nossos Deployments adicionamos as seguintes labels:

```yaml
  labels:
    run: nginx
    dc: UK
---
  labels:
    run: nginx
    dc: Netherlands
```

As Labels são utilizadas para a organização do cluster, vamos listar nossos pods procurando pelas Labels.

Primeiro vamos realizar uma pesquisa utilizando as labels ``dc=UK`` e ``dc=Netherlands``:

Pesquisando pela label ``UK``:

```
kubectl get pods -l dc=UK

NAME                                 READY  STATUS   RESTARTS   AGE
primeiro-deployment-68c9dbf8b8-kjqpt 1/1    Running  0          3m
```

Pesquisando pela label ``Netherlands``:

```
kubectl get pods -l dc=Netherlands

NAME                                READY STATUS    RESTARTS   AGE
segundo-deployment-59db86c584-cf9pp 1/1   Running   0          4m
```

Caso queira uma saída mais personalizada podemos listar da seguinte forma, veja:

```
kubectl get pod -L dc

NAME                         READY STATUS   RESTARTS AGE DC
primeiro-deployment-68c9...  1/1   Running  0        5m  UK
segundo-deployment-59db ...  1/1   Running  0        5m  Netherlands
```

## Node Selector

O **Node Selector** é uma forma de classificar nossos nodes como por exemplo nosso node ``elliot-02`` que possui disco **SSD** e está localizado no DataCenter ``UK``, e o node ``elliot-03`` que possui disco **HDD** e está localizado no DataCenter ``Netherlands``.

Agora que temos essas informações vamos criar essas labels em nossos nodes, para utilizar o ``nodeSelector``.

Criando a label ``disk`` com o valor ``SSD`` no worker 1:

```
kubectl label node elliot-02 disk=SSD

node/elliot-02 labeled
```

Criando a label ``dc`` com o valor ``UK`` no worker 1:

```
kubectl label node elliot-02 dc=UK

node/elliot-02 labeled
```

Criando a label ``dc`` com o valor ``Netherlands`` no worker 2:

```
kubectl label node elliot-03 dc=Netherlands

node/elliot-03 labeled
```

Criando a label ``disk`` com o valor ``hdd`` no worker 2:

```
kubectl label nodes elliot-03 disk=hdd

node/elliot-03 labeled
```

Opa! Acabamos declarando o ``disk=hdd`` em letra minúscula, como arrumamos isso? Subscrevendo a label como no comando a seguir.

```
kubectl label nodes elliot-03 disk=HDD --overwrite

node/elliot-03 labeled
```

Para saber as labels configuradas em cada node basta executar o seguinte comando:

No worker 1:

```
kubectl label nodes elliot-02 --list

dc=UK
disk=SSD
kubernetes.io/hostname=elliot-02
beta.kubernetes.io/arch=amd64
beta.kubernetes.io/os=linux
```

No worker 2:

```
kubectl label nodes elliot-03 --list

beta.kubernetes.io/os=linux
dc=Netherlands
disk=HDD
kubernetes.io/hostname=elliot-03
beta.kubernetes.io/arch=amd64
```

Agora, basta realizar o deploy novamente, porém antes vamos adicionar duas novas opções ao YAML e vamos ver a mágica acontecer. O nosso pod irá ser criado no node ``elliot-02``, onde possui a label ``disk=SSD``.

Crie o arquivo ``terceiro-deployment.yaml``:

```
vim terceiro-deployment.yaml
```

Informe o seguinte conteúdo:

```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    run: nginx
  name: terceiro-deployment
  namespace: default
spec:
  replicas: 1
  selector:
    matchLabels:
      run: nginx
  template:
    metadata:
      creationTimestamp: null
      labels:
        run: nginx
        dc: Netherlands
    spec:
      containers:
      - image: nginx
        imagePullPolicy: Always
        name: nginx2
        ports:
        - containerPort: 80
          protocol: TCP
        resources: {}
        terminationMessagePath: /dev/termination-log
        terminationMessagePolicy: File
      dnsPolicy: ClusterFirst
      restartPolicy: Always
      schedulerName: default-scheduler
      securityContext: {}
      terminationGracePeriodSeconds: 30
      nodeSelector:
        disk: SSD
```

Crie o deployment a partir do manifesto:

```
kubectl create -f terceiro-deployment.yaml

deployment.extensions/terceiro-deployment created
```

Visualizando os detalhes dos pods:

```
kubectl get pods -o wide

NAME                        READY STATUS  RESTARTS  AGE  IP           NODE
primeiro-deployment-56d9... 1/1   Running  0      14m  172.17.0.4 elliot-03
segundo-deployment-869f...  1/1   Running  0      14m  172.17.0.5 elliot-03
terceiro-deployment-59cd... 1/1   Running  0      22s  172.17.0.6 elliot-02
```

Removendo a label ``dc`` de um node worker:

```
kubectl label nodes elliot-02 dc-
```

Removendo uma determinada label de todos os nodes:

```
kubectl label nodes --all dc-
```

Agora imagine as infinitas possibilidades que isso poderá lhe proporcionar… Já estou pensando em várias, como por exemplo se é produção ou não, se consome muita CPU ou muita RAM, se precisa estar em determinado rack e por aí vai. 😃

Simples como voar, não?

## Kubectl Edit

Agora vamos fazer o seguinte, vamos utilizar o comando ``kubectl edit`` para editar nosso primeiro deployment, digamos que a "quente" com o pod ainda em execução.

```
kubectl edit deployment primeiro-deployment
```

Abriu um editor, correto? Vamos alterar a label ``DC``. Vamos imaginar que esse Deployment agora rodará no ``DC`` de ``Netherlands``. Precisamos adicionar a ``Label`` e o ``nodeSelector``.

O conteúdo deve ser o seguinte:

```yaml
spec:
  replicas: 1
  selector:
    matchLabels:
      run: nginx
  strategy:
    rollingUpdate:
      maxSurge: 1
      maxUnavailable: 1
    type: RollingUpdate
  template:
    metadata:
      creationTimestamp: null
      labels:
        dc: Netherlands
        app: giropops
        run: nginx
spec:
      containers:
      - image: nginx
        imagePullPolicy: Always
        name: nginx2
        ports:
        - containerPort: 80
          protocol: TCP
        resources: {}
        terminationMessagePath: /dev/termination-log
        terminationMessagePolicy: File
      dnsPolicy: ClusterFirst
      nodeSelector:
        dc: Netherlands
...

deployment.extensions/primeiro-deployment edited
```

Como podemos ver mudamos o valor da label ``dc`` e também modificamos o ``nodeSelector``, onde ele agora subirá no node que tiver a label ``dc`` com o valor ``Netherlands``, fácil! 😀

Veja se o resultado foi conforme esperado:

```
kubectl get pods -l dc=Netherlands -o wide

NAME                     READY  STATUS    RESTARTS  AGE ..NODE
primeiro-deployment-7..  1/1    Running   0         3m    elliot-03
segundo-deployment-5..   1/1    Running   0         49m   elliot-02
terceiro-deployment-5..  1/1    Running   0         14m   elliot-02
```

Com certeza, esse pod foi criado no node ``elliot-03``, pois havíamos dito que ele possuía essa label anteriormente.

# ReplicaSet

O **ReplicaSet** garante a quantidade solicitada de pods e os recursos necessários para um Deployment. Uma vez que o Deployment é criado, é o ReplicaSet que controla a quantidade de pods em execução, caso algum pod seja finalizado, ele que irá detectar e solicitar que outro pod seja executado em seu lugar, garantindo assim a quantidade de réplicas solicitadas.

Caso você delete um Pod ou o worker que ele esteja fique down, o replicaset irá se incarregar de subir outro pod.

# DaemonSet

**DaemonSet** é basicamente a mesma coisa do que o ReplicaSet, com a diferença que quando você utiliza o DaemonSet você não especifica o número de réplicas, ele subirá um pod por node em seu cluster.

É sempre interessante quando criar usar e abusar das labels, assim você conseguirá ter melhor flexibilidade na distribuição mais adequada de sua aplicação.

Ele é bem interessante para serviços que necessitem rodar em todos os nodes do cluster, como por exemplo, coletores de logs e agentes de monitoração.

# Rollouts e Rollbacks

**CRIAR**

# Cron Jobs

Um serviço **CronJob** nada mais é do que uma linha de um arquivo crontab o mesmo arquivo de uma tabela ``cron``. Ele agenda e executa tarefas periodicamente em um determinado cronograma.

Mas para que podemos usar os **Cron Jobs****? As "Cron" são úteis para criar tarefas periódicas e recorrentes, como executar backups ou enviar e-mails.

Vamos criar um exemplo para ver como funciona, bora criar nosso manifesto:

```
vim primeiro-cron.yaml
```

Informe o seguinte conteúdo.

```yaml
apiVersion: batch/v1beta1
kind: CronJob
metadata:
  name: giropops-cron
spec:
  schedule: "*/1 * * * *"
  jobTemplate:
    spec:
      template:
        spec:
          containers:
          - name: giropops-cron
            image: busybox
            args:
            - /bin/sh
            - -c
            - date; echo Bem Vindo ao Descomplicando Kubernetes - LinuxTips VAIIII ;sleep 30
          restartPolicy: OnFailure
```

Nosso exemplo de ``CronJobs`` anterior imprime a hora atual e uma mensagem de de saudação a cada minuto.

Vamos criar o ``CronJob`` a partir do manifesto.

```
kubectl create -f primeiro-cron.yaml

cronjob.batch/giropops-cron created
```

Agora vamos listar e detalhar melhor nosso ``Cronjob``.

```
kubectl get cronjobs

NAME            SCHEDULE      SUSPEND   ACTIVE   LAST SCHEDULE   AGE
giropops-cron   */1 * * * *   False     1        13s             2m
```

Vamos visualizar os detalhes do ``Cronjob`` ``giropops-cron``.

```
kubectl describe cronjobs.batch giropops-cron

Name:                       giropops-cron
Namespace:                  default
Labels:                     <none>
Annotations:                <none>
Schedule:                   */1 * * * *
Concurrency Policy:         Allow
Suspend:                    False
Starting Deadline Seconds:  <unset>
Selector:                   <unset>
Parallelism:                <unset>
Completions:                <unset>
Pod Template:
  Labels:  <none>
  Containers:
   giropops-cron:
    Image:      busybox
    Port:       <none>
    Host Port:  <none>
    Args:
      /bin/sh
      -c
      date; echo LinuxTips VAIIII ;sleep 30
    Environment:     <none>
    Mounts:          <none>
  Volumes:           <none>
Last Schedule Time:  Wed, 22 Aug 2018 22:33:00 +0000
Active Jobs:         <none>
Events:
  Type    Reason            Age   From                Message
  ----    ------            ----  ----                -------
  Normal  SuccessfulCreate  1m    cronjob-controller  Created job giropops-cron-1534977120
  Normal  SawCompletedJob   1m    cronjob-controller  Saw completed job: giropops-cron-1534977120
  Normal  SuccessfulCreate  41s   cronjob-controller  Created job giropops-cron-1534977180
  Normal  SawCompletedJob   1s    cronjob-controller  Saw completed job: giropops-cron-1534977180
  Normal  SuccessfulDelete  1s    cronjob-controller  Deleted job giropops-cron-1534977000
```

Olha que bacana, se observar no ``Events`` do cluster o ``cron`` já está agendando e executando as tarefas.

Agora vamos ver esse ``cron`` funcionando através do comando ``kubectl get`` junto do parâmetro ``--watch`` para verificar a saída das tarefas, preste atenção que a tarefa vai ser criada em cerca de um minuto após a criação do ``CronJob``.

```
kubectl get jobs --watch

NAME                       DESIRED  SUCCESSFUL   AGE
giropops-cron-1534979640   1         1            2m
giropops-cron-1534979700   1         1            1m
```

Vamos visualizar o CronJob:

```
kubectl get cronjob giropops-cron

NAME           SCHEDULE      SUSPEND   ACTIVE    LAST SCHEDULE   AGE
giropops-cron  */1 * * * *   False     1         26s             48m
```

Como podemos observar que nosso ``cron`` está funcionando corretamente. Para visualizar a saída dos comandos executados pela tarefa vamos utilizar o comando ``logs`` do ``kubectl``.

Para isso vamos listar os pods em execução e, em seguida, pegar os logs do mesmo.

```
kubectl get pods

NAME                            READY     STATUS      RESTARTS   AGE
giropops-cron-1534979940-vcwdg  1/1       Running     0          25s
```

Vamos visualizar os logs:

```
kubectl logs giropops-cron-1534979940-vcwdg

Wed Aug 22 23:19:06 UTC 2018
LinuxTips VAIIII
```

O ``cron`` está executando corretamente as tarefas de imprimir a data e a frase que criamos no manifesto.

Se executarmos um ``kubectl get pods`` poderemos ver os Pods criados e utilizados para executar as tarefas a todo minuto.

```
kubectl get pods

NAME                             READY    STATUS      RESTARTS   AGE
giropops-cron-1534980360-cc9ng   0/1      Completed   0          2m
giropops-cron-1534980420-6czgg   0/1      Completed   0          1m
giropops-cron-1534980480-4bwcc   1/1      Running     0          4s
```

---

> **Atenção!!!** Por padrão, o Kubernetes mantém o histórico dos últimos 3 ``cron`` executados, concluídos ou com falhas.
Fonte: https://kubernetes.io/docs/tasks/job/automated-tasks-with-cron-jobs/#jobs-history-limits

---

Agora vamos deletar nosso CronJob:

```
kubectl delete cronjob giropops-cron

cronjob.batch "giropops-cron" deleted
```
