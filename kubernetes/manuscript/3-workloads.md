# Namespace
No kubernetes, namespaces fornece a possibilidade de isolar grupos de recursos dentro de um mesmo cluster. Nomes de recursos devem ser únicos em um dado namespace,
mas não entre os namespaces. Existem recursos que devem estar dentro de um namespace(deployments...) e outos que estão a nível do cluster (PV, nodes,StorageClass...).
```
# Em um namespace
kubectl api-resources --namespaced=true

# Sem namespace
kubectl api-resources --namespaced=false
```
## Namespace DNS
Quando você cria um service ele cria um dns de entrada, se você quer pesquisar um serviço fora do namespace, necessita desse formato.

```
<service-name>.<namespace-name>.svc.cluster.local
```

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
