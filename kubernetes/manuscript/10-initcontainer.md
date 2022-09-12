# InitContainers

O objeto do tipo **InitContainers** são um ou mais contêineres que são executados antes do contêiner de um aplicativo em um Pod. Os contêineres de inicialização podem conter utilitários ou scripts de configuração não presentes em uma imagem de aplicativo.

- Os contêineres de inicialização sempre são executados até a conclusão.
- Cada contêiner init deve ser concluído com sucesso antes que o próximo comece.

Se o contêiner init de um Pod falhar, o Kubernetes reiniciará repetidamente o Pod até que o contêiner init tenha êxito. No entanto, se o Pod tiver o ``restartPolicy`` como ``Never``, o Kubernetes não reiniciará o Pod, e o contêiner principal não irá ser executado.

Crie o Pod a partir do manifesto:

```
vim nginx-initcontainer.yaml
```

Informe o seguinte conteúdo:

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: init-demo
spec:
  containers:
  - name: nginx
    image: nginx
    ports:
    - containerPort: 80
    volumeMounts:
    - name: workdir
      mountPath: /usr/share/nginx/html
  initContainers:
  - name: install
    image: busybox
    command: ['wget','-O','/work-dir/index.html','http://linuxtips.io']
    volumeMounts:
    - name: workdir
      mountPath: "/work-dir"
  dnsPolicy: Default
  volumes:
  - name: workdir
    emptyDir: {}
```

Crie o pod a partir do manifesto.

```
kubectl create -f nginx-initcontainer.yaml

pod/init-demo created
```

Visualize o conteúdo de um arquivo dentro de um contêiner do Pod.

```
kubectl exec -ti init-demo -- cat /usr/share/nginx/html/index.html
```

Vamos ver os detalhes do Pod:

```
kubectl describe pod init-demo

Name:               init-demo
Namespace:          default
Priority:           0
PriorityClassName:  <none>
Node:               k8s3/172.31.28.13
Start Time:         Sun, 11 Nov 2018 13:10:02 +0000
Labels:             <none>
Annotations:        <none>
Status:             Running
IP:                 10.44.0.13
Init Containers:
  install:
    Container ID:  docker://d86ec7ce801819a2073b96098055407dec5564a678c2548dd445a613314bac8e
    Image:         busybox
    Image ID:      docker-pullable://busybox@sha256:2a03a6059f21e150ae84b0973863609494aad70f0a80eaeb64bddd8d92465812
    Port:          <none>
    Host Port:     <none>
    Command:
      wget
      -O
      /work-dir/index.html
      http://kubernetes.io
    State:          Terminated
      Reason:       Completed
      Exit Code:    0
      Started:      Sun, 11 Nov 2018 13:10:03 +0000
      Finished:     Sun, 11 Nov 2018 13:10:03 +0000
    Ready:          True
    Restart Count:  0
    Environment:    <none>
    Mounts:
      /var/run/secrets/kubernetes.io/serviceaccount from default-token-x5g8c (ro)
      /work-dir from workdir (rw)
Containers:
  nginx:
    Container ID:   docker://1ec37249f98ececdfb5f80c910142c5e6edbc9373e34cd194fbd3955725da334
    Image:          nginx
    Image ID:       docker-pullable://nginx@sha256:d59a1aa7866258751a261bae525a1842c7ff0662d4f34a355d5f36826abc0341
    Port:           80/TCP
    Host Port:      0/TCP
    State:          Running
      Started:      Sun, 11 Nov 2018 13:10:04 +0000
    Ready:          True
    Restart Count:  0
    Environment:    <none>
    Mounts:
      /usr/share/nginx/html from workdir (rw)
      /var/run/secrets/kubernetes.io/serviceaccount from default-token-x5g8c (ro)
Conditions:
  Type              Status
  Initialized       True
  Ready             True
  ContainersReady   True
  PodScheduled      True
Volumes:
  workdir:
    Type:    EmptyDir (a temporary directory that shares a pod's lifetime)
    Medium:
  default-token-x5g8c:
    Type:        Secret (a volume populated by a Secret)
    SecretName:  default-token-x5g8c
    Optional:    false
QoS Class:       BestEffort
Node-Selectors:  <none>
Tolerations:     node.kubernetes.io/not-ready:NoExecute for 300s
                 node.kubernetes.io/unreachable:NoExecute for 300s
Events:
  Type    Reason     Age    From               Message
  ----    ------     ----   ----               -------
  Normal  Scheduled  3m1s   default-scheduler  Successfully assigned default/init-demo to k8s3
  Normal  Pulling    3m     kubelet, k8s3      pulling image "busybox"
  Normal  Pulled     3m     kubelet, k8s3      Successfully pulled image "busybox"
  Normal  Created    3m     kubelet, k8s3      Created container
  Normal  Started    3m     kubelet, k8s3      Started container
  Normal  Pulling    2m59s  kubelet, k8s3      pulling image "nginx"
  Normal  Pulled     2m59s  kubelet, k8s3      Successfully pulled image "nginx"
  Normal  Created    2m59s  kubelet, k8s3      Created container
  Normal  Started    2m59s  kubelet, k8s3      Started container
```

Coletando os logs do contêiner init:

```
kubectl logs init-demo -c install

Connecting to linuxtips.io (23.236.62.147:80)
Connecting to www.linuxtips.io (35.247.254.172:443)
wget: note: TLS certificate validation not implemented
saving to '/work-dir/index.html'
index.html           100% |********************************|  765k  0:00:00 ETA
'/work-dir/index.html' saved
```

E por último vamos remover o Pod a partir do manifesto:

```
kubectl delete -f nginx-initcontainer.yaml

pod/init-demo deleted
```

