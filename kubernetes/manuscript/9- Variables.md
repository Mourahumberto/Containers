## Sumário

- [Deployments](#deployments)
  - [Filtrando por Labels](#filtrando-por-labels)
  - [Node Selector](#node-selector)
  - [Kubectl Edit](#kubectl-edit)
- [ReplicaSet](#replicaset)
- [DaemonSet](#daemonset)
- [Rollouts e Rollbacks](#rollouts-e-rollbacks)
- [Cron Jobs](#cron-jobs)

# Deployments

O **Deployment** é um recurso com a responsabilidade de instruir o Kubernetes a criar, atualizar e monitorar a saúde das instâncias de suas aplicações.

Um Deployment é o responsável por gerenciar o seu **ReplicaSet** (que iremos falar logo menos), ou seja, o Deployment é quem vai determinar a configuração de sua aplicação e como ela será implementada. O Deployment é o **controller** que irá cuidar, por exemplo, uma instância de sua aplicação por algum motivo for interrompida. O **Deployment controller** irá identificar o problema com a instância e irá criar uma nova em seu lugar.

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
