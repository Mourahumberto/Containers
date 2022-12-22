# Kubectl taint
doc: https://kubernetes.io/pt-br/docs/concepts/scheduling-eviction/taint-and-toleration/

O **Taint** nada mais é do que adicionar propriedades ao nó do cluster para impedir que os pods sejam alocados em nós inapropriados.

Por exemplo, todo nó ``master`` do cluster é marcado para não receber pods que não sejam de gerenciamento do cluster.

O nó ``master`` está marcado com o taint ``NoSchedule``, assim o scheduler do Kubernetes não aloca pods no nó master, e procura outros nós no cluster sem essa marca.

Visualizando os nodes do cluster:

```
kubectl get nodes

NAME           STATUS   ROLES    AGE     VERSION
elliot-01   Ready    master   7d14h   v1.18.2
elliot-02   Ready    <none>   7d14h   v1.18.2
elliot-03   Ready    <none>   7d14h   v1.18.2
```

Visualizando as labels Taints do node ``master``:

```
kubectl describe node elliot-01 | grep -i taint

Taints:             node-role.kubernetes.io/master:NoSchedule
```

**Vamos testar algumas coisas e permitir que o nó master rode outros pods.**

Primeiro vamos rodar 3 réplicas de ``nginx``:

```
kubectl create deployment nginx --image=nginx

deployment.apps/nginx created
```

Visualizando os deployments:

```
kubectl get deployments.apps

NAME    READY   UP-TO-DATE   AVAILABLE   AGE
nginx   1/1     1            1           5s
```

Escalando o deployment do nginx para 3 réplicas:

```
kubectl scale deployment nginx --replicas=3

deployment.apps/nginx scaled
```

Visualizando novamente os deployments:

```
kubectl get deployments.apps

NAME    READY   UP-TO-DATE   AVAILABLE   AGE
nginx   3/3     3            3           1m5s
```

Visualizando os detalhes dos pods:

```
kubectl get pods -o wide

NAME                     READY   STATUS    RESTARTS   AGE     IP          NODE               NOMINATED NODE   READINESS GATES
limit-pod                1/1     Running   0          3m44s   10.32.0.4   elliot-02   <none>           <none>
nginx                    1/1     Running   0          25m     10.46.0.1   elliot-03    <none>           <none>
nginx-85f7fb6b45-9bzwc   1/1     Running   0          6m7s    10.32.0.3   elliot-02   <none>           <none>
nginx-85f7fb6b45-cbmtr   1/1     Running   0          6m7s    10.46.0.2   elliot-03    <none>           <none>
nginx-85f7fb6b45-rprz5   1/1     Running   0          6m7s    10.32.0.2   elliot-02   <none>           <none>
```

Vamos adicionar a marca ``NoSchedule`` aos nós worker também para ver como eles se comportam.

Node worker 1:

```
kubectl taint node elliot-02 key1=value1:NoSchedule

node/elliot-02 tainted
```

Node worker 2:

```
kubectl taint node elliot-03 key1=value1:NoSchedule

node/elliot-03 tainted
```

Visualizando a label Taint no node worker 1:

```
kubectl describe node elliot-02 | grep -i taint

Taints:             key1=value1:NoSchedule
```

Visualizando a label Taint no node worker 2:

```
kubectl describe node elliot-03 | grep -i taint

Taints:             key1=value1:NoSchedule
```

Agora vamos aumentar a quantidade de réplicas:

```
kubectl scale deployment nginx --replicas=5

deployment.apps/nginx scaled
```

Visualizando os detalhes dos pods:

```
kubectl get pods  -o wide

NAME                     READY   STATUS    RESTARTS   AGE     IP          NODE               NOMINATED NODE   READINESS GATES
limit-pod                1/1     Running   0          5m23s   10.32.0.4   elliot-02   <none>           <none>
nginx                    1/1     Running   0          27m     10.46.0.1   elliot-03    <none>           <none>
nginx-85f7fb6b45-9bzwc   1/1     Running   0          7m46s   10.32.0.3   elliot-02   <none>           <none>
nginx-85f7fb6b45-cbmtr   1/1     Running   0          7m46s   10.46.0.2   elliot-03    <none>           <none>
nginx-85f7fb6b45-qnhtl   0/1     Pending   0          18s     <none>      <none>             <none>           <none>
nginx-85f7fb6b45-qsvpp   0/1     Pending   0          18s     <none>      <none>             <none>           <none>
nginx-85f7fb6b45-rprz5   1/1     Running   0          7m46s   10.32.0.2   elliot-02   <none>           <none>
```

Como podemos ver, as nova réplicas ficaram órfãs esperando aparecer um nó com as prioridades adequadas para o Scheduler.

Vamos remover esse Taint dos nossos nós worker:

Removendo o taint do worker 1:

```
kubectl taint node elliot-02 key1:NoSchedule-

node/elliot-02 untainted
```

Removendo o taint do worker 2:

```
kubectl taint node elliot-03 key1:NoSchedule-

node/elliot-03 untainted
```

Visualizando os detalhes dos pods:

```
kubectl get pods  -o wide

NAME                     READY   STATUS    RESTARTS   AGE     IP          NODE               NOMINATED NODE   READINESS GATES
limit-pod                1/1     Running   0          6m17s   10.32.0.4   elliot-02          <none>           <none>
nginx                    1/1     Running   0          27m     10.46.0.1   elliot-03          <none>           <none>
nginx-85f7fb6b45-9bzwc   1/1     Running   0          8m40s   10.32.0.3   elliot-02          <none>           <none>
nginx-85f7fb6b45-cbmtr   1/1     Running   0          8m40s   10.46.0.2   elliot-03          <none>           <none>
nginx-85f7fb6b45-qnhtl   1/1     Running   0          72s     10.46.0.5   elliot-03          <none>           <none>
nginx-85f7fb6b45-qsvpp   1/1     Running   0          72s     10.46.0.4   elliot-03          <none>           <none>
nginx-85f7fb6b45-rprz5   1/1     Running   0          8m40s   10.32.0.2   elliot-02          <none>           <none>
```

Existem vários tipos de marcas que podemos usar para classificar os nós, vamos testar uma outra chamada ``NoExecute``, que impede o Scheduler de agendar Pods nesses nós.

Adicionando a marca ``NoExecute`` no worker 1:

```
kubectl taint node elliot-02 key1=value1:NoExecute

node/elliot-02 tainted
```

Adicionando a marca ``NoExecute`` no worker 2:

```
kubectl taint node elliot-03 key1=value1:NoExecute

node/elliot-03 tainted
```

Visualizando os detalhes dos pods:

```
kubectl get pods

NAME                     READY   STATUS    RESTARTS   AGE
nginx-85f7fb6b45-87sq5   0/1     Pending   0          20s
nginx-85f7fb6b45-8q99g   0/1     Pending   0          20s
nginx-85f7fb6b45-drmzz   0/1     Pending   0          20s
nginx-85f7fb6b45-hb4dp   0/1     Pending   0          20s
nginx-85f7fb6b45-l6zln   0/1     Pending   0          20s
```

Como podemos ver todos os Pods estão órfãs. Porque o nó ``master`` tem a marca taint ``NoScheduler`` default do kubernetes e os nós worker tem a marca ``NoExecute``.

Vamos diminuir a quantidade de réplicas para ver o que acontece.

Reduzindo a quantidade de réplicas no worker 1:

```
kubectl scale deployment nginx --replicas=1

deployment.apps/nginx scaled
```

Reduzindo a quantidade de réplicas no worker 2:

```
kubectl get pods

nginx-85f7fb6b45-drmzz   0/1     Pending   0          43s
```

Vamos remover o taint ``NoExecute`` do nós workers.

Removendo o taint no worker 1:

```
kubectl taint node elliot-02 key1:NoExecute-

node/elliot-02 untainted
```

Removendo o taint no worker 2:

```
kubectl taint node elliot-03 key1:NoExecute-

node/elliot-03 untainted
```

Visualizando os detalhes dos pods:

```
kubectl get pods

NAME                     READY   STATUS    RESTARTS   AGE
nginx-85f7fb6b45-drmzz   1/1     Running   0          76s
```

Agora temos um nó operando normalmente.

Mas e se nossos workers ficarem indisponíveis, podemos rodar Pods no nó master?

Claro que podemos, vamos configurar nosso nó master para que o Scheduler consiga agenda Pods nele.

```
kubectl taint nodes --all node-role.kubernetes.io/master-

node/elliot-01 untainted
```

Visualizando a marca taint no nó master.

```
kubectl describe node elliot-01 | grep -i taint

Taints:             <none>
```

Agora vamos aumentar a quantidade de réplicas do nosso pod ``nginx``.

```
kubectl scale deployment nginx --replicas=4

deployment.apps/nginx scaled
```

Visualizando os detalhes dos pods:

```
kubectl get pods -o wide

NAME                     READY   STATUS    RESTARTS   AGE    IP          NODE               NOMINATED NODE   READINESS GATES
nginx-85f7fb6b45-2c6dm   1/1     Running   0          9s     10.32.0.2   elliot-02          <none>           <none>
nginx-85f7fb6b45-4jzcn   1/1     Running   0          9s     10.32.0.3   elliot-02          <none>           <none>
nginx-85f7fb6b45-drmzz   1/1     Running   0          114s   10.46.0.1   elliot-03          <none>           <none>
nginx-85f7fb6b45-rstvq   1/1     Running   0          9s     10.46.0.2   elliot-03          <none>           <none>
```

Vamos adicionar o Taint ``NoExecute`` nos nós worker para ver o que acontece.

Adicionando o taint no worker 1:

```
kubectl taint node elliot-02 key1=value1:NoExecute

node/elliot-02 tainted
```

Adicionando o taint no worker 2:

```
kubectl taint node elliot-03 key1=value1:NoExecute

node/elliot-03 tainted
```

Visualizando os detalhes do pods:

```
kubectl get pods -o wide

NAME                     READY   STATUS    RESTARTS   AGE   IP          NODE              NOMINATED NODE   READINESS GATES
nginx-85f7fb6b45-49knz   1/1     Running   0          14s   10.40.0.5   elliot-01         <none>           <none>
nginx-85f7fb6b45-4cm9x   1/1     Running   0          14s   10.40.0.4   elliot-01         <none>           <none>
nginx-85f7fb6b45-kppnd   1/1     Running   0          14s   10.40.0.6   elliot-01         <none>           <none>
nginx-85f7fb6b45-rjlmj   1/1     Running   0          14s   10.40.0.3   elliot-01         <none>           <none>
```

Removendo o deployment ``nginx``:

```
kubectl delete deployment nginx

deployment.extensions "nginx" deleted
```

O Scheduler alocou tudo no nó ``master``, como podemos ver o Taint pode ser usado para ajustar configurações de qual Pod deve ser alocado em qual nó.

Vamos permitir que nosso Scheduler aloque e execute os Pods em todos os nós:

Removendo o taint ``NoSchedule`` em todos os nós do cluster:

```
kubectl taint node --all key1:NoSchedule-

node/elliot-01 untainted
node/elliot-02 untainted
node/elliot-03 untainted
```

Removendo o taint ``NoExecute`` em todos os nós do cluster:

```
kubectl taint node --all key1:NoExecute-

node/kube-worker1 untainted
node/kube-worker2 untainted
error: taint "key1:NoExecute" not found
```

Visualizando os taints dos nodes:

```
kubectl describe node elliot-01 | grep -i taint

Taints:             <none>
```

# Colocando o nó em modo de manutenção

Para colocar o nó em manutenção iremos utilizar o ``cordon``.

```
kubectl cordon elliot-02

node/elliot-02 cordoned
```

Visualizando o node em manutenção.

```
kubectl get nodes

NAME        STATUS                      ROLES    AGE     VERSION
elliot-01   Ready                       master   7d14h   v1.18.2
elliot-02   Ready,SchedulingDisabled    <none>   7d14h   v1.18.2
elliot-03   Ready                       <none>   7d14h   v1.18.2
```

Repare que o nó ``elliot-02`` ficou com o status ``Ready,SchedulingDisabled``, agora você pode fazer a manutenção no seu node tranquilamente.
Para retirar nó de modo de manutenção, iremos utilizar o ``uncordon``.

```
kubectl uncordon elliot-02

node/elliot-02 uncordoned
```

Visualizando novamente os nós.

```
kubectl get nodes

NAME           STATUS   ROLES    AGE     VERSION
elliot-01   Ready    master   7d14h   v1.18.2
elliot-02   Ready    <none>   7d14h   v1.18.2
elliot-03   Ready    <none>   7d14h   v1.18.2
```

Pronto, agora seu nó não está mais em modo de manutenção.
