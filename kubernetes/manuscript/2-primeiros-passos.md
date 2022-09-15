# Primeiros Passos

## Exibindo informações detalhadas sobre os nós

```
kubectl describe node [nome_do_no]
```

Exemplo:

```
kubectl describe node elliot-02

Name:               elliot-02
Roles:              <none>
Labels:             beta.kubernetes.io/arch=amd64
                    beta.kubernetes.io/os=linux
                    kubernetes.io/arch=amd64
                    kubernetes.io/hostname=elliot-02
                    kubernetes.io/os=linux
Annotations:        kubeadm.alpha.kubernetes.io/cri-socket: /var/run/dockershim.sock
                    node.alpha.kubernetes.io/ttl: 0
                    volumes.kubernetes.io/controller-managed-attach-detach: true
```
## gerenciando objetos no kubernetes

### Imperative commands
- comandos criando recurso direto na API sem nem um manifesto(.yaml) kubernetes
```
kubectl create deployment nginx --image nginx
```

### Imperative object configuration
- comandos usado em manifestos para criar objetos no kubernetes
```
kubectl create -f nginx.yaml
kubectl delete -f nginx.yaml -f redis.yaml
kubectl replace -f nginx.yaml
```

### Declarative object configuration
- Faz apply em uma pastas com os manifests
```
kubectl diff -f configs/
kubectl apply -f configs/
Recursive
kubectl diff -R -f configs/
kubectl apply -R -f configs/
```

## Verificando os namespaces e pods

O k8s organiza quase tudo dentro de *namespaces*. Por meio deles, podem ser realizadas limitações de segurança e de recursos dentro do *cluster*, tais como *pods*, *replication controllers* e diversos outros. Para visualizar os *namespaces* disponíveis no *cluster*, digite:

```
kubectl get namespaces
```

Vamos listar os *pods* do *namespace* **kube-system** utilizando o comando a seguir.

```
kubectl get pod -n kube-system
```

Será que há algum *pod* escondido em algum *namespace*? É possível listar todos os *pods* de todos os *namespaces* com o comando a seguir.

```
kubectl get pods --all-namespaces
```

Há a possibilidade ainda, de utilizar o comando com a opção ```-o wide```, que disponibiliza maiores informações sobre o recurso, inclusive em qual nó o *pod* está sendo executado. Exemplo:

```
kubectl get pods --all-namespaces -o wide

NAMESPACE     NAME                                READY   STATUS    RESTARTS   AGE   IP             NODE        NOMINATED NODE   READINESS GATES
default       nginx                               1/1     Running   0          24m   10.44.0.1      docker-02   <none>           <none>
```

## Executando nosso primeiro pod no k8s

Iremos iniciar o nosso primeiro *pod* no k8s. Para isso, executaremos o comando a seguir.

```
kubectl run nginx --image nginx

pod/nginx created
```

Listando os *pods* com ``kubectl get pods``, obteremos a seguinte saída.

```
NAME    READY   STATUS    RESTARTS   AGE
nginx   1/1     Running   0          66s
```

Vamos olhar agora a descrição desse objeto dentro do *cluster*.

```
kubectl describe pod nginx

Name:         nginx
Namespace:    default
Priority:     0
Node:         docker-02/172.16.83.13
Start Time:   Tue, 12 May 2020 02:29:38 -0300
Labels:       run=nginx
Annotations:  <none>
Status:       Running
IP:           10.44.0.1
IPs:
  IP:  10.44.0.1
Containers:
  nginx:
    Container ID:   docker://2719e2bc023944ee8f34db538094c96b24764a637574c703e232908b46b12a9f
    Image:          nginx
    Image ID:       docker-pullable://nginx@sha256:86ae264c3f4acb99b2dee4d0098c40cb8c46dcf9e1148f05d3a51c4df6758c12
    Port:           <none>
    Host Port:      <none>
    State:          Running
      Started:      Tue, 12 May 2020 02:29:42 -0300
```

## Verificar os últimos eventos do cluster

Você pode verificar quais são os últimos eventos do *cluster* com o comando ``kubectl get events``. Serão mostrados eventos como: o *download* de imagens do Docker Hub (ou de outro *registry* configurado), a criação/remoção de *pods*, etc.

A saída a seguir mostra o resultado da criação do nosso contêiner com Nginx.

```
LAST SEEN   TYPE     REASON      OBJECT      MESSAGE
5m34s       Normal   Scheduled   pod/nginx   Successfully assigned default/nginx to docker-02
5m33s       Normal   Pulling     pod/nginx   Pulling image "nginx"
5m31s       Normal   Pulled      pod/nginx   Successfully pulled image "nginx"
5m30s       Normal   Created     pod/nginx   Created container nginx
5m30s       Normal   Started     pod/nginx   Started container nginx
```

No resultado do comando anterior é possível observar que a execução do nginx ocorreu no *namespace* default e que a imagem **nginx** não existia no repositório local e, sendo assim, teve de ser feito download da imagem.

## Efetuar o dump de um objeto em formato YAML

Assim como quando se está trabalhando com *stacks* no Docker Swarm, normalmente recursos no k8s são declarados em arquivos **YAML** ou **JSON** e depois manipulados através do ``kubectl``.

Para nos poupar o trabalho de escrever o arquivo inteiro, pode-se utilizar como *template* o *dump* de um objeto já existente no k8s, como mostrado a seguir.

```
kubectl get pod nginx -o yaml > meu-primeiro.yaml
```

Será criado um novo arquivo chamado ``meu-primeiro.yaml``, resultante do redirecionamento da saída do comando ``kubectl get pod nginx -o yaml``.

Abrindo o arquivo com ``vim meu-primeiro.yaml`` (você pode utilizar o editor que você preferir), teremos o seguinte conteúdo.

Vamos agora remover o nosso *pod* com o seguinte comando.

```
kubectl delete pod nginx
```

A saída deve ser algo como:

```
pod "nginx" deleted
```

Vamos recriá-lo, agora a partir do nosso arquivo YAML.

```
kubectl create -f meu-primeiro.yaml

pod/nginx created
```

Observe que não foi necessário informar ao ``kubectl`` qual tipo de recurso seria criado, pois isso já está contido dentro do arquivo.

Listando os *pods* disponíveis com o seguinte comando.

```
kubectl get pods
```

Deve-se obter uma saída similar à esta:

```
NAME    READY   STATUS    RESTARTS   AGE
nginx   1/1     Running   0          109s
```

Uma outra forma de criar um arquivo de *template* é através da opção ``--dry-run`` do ``kubectl``, com o funcionamento ligeiramente diferente dependendo do tipo de recurso que será criado. Exemplos:

Para a criação do template de um *pod*:

```
kubectl run meu-nginx --image nginx --dry-run=client -o yaml > pod-template.yaml
```

Para a criação do *template* de um *deployment*:

```
kubectl create deployment meu-nginx --image=nginx --dry-run=client -o yaml > deployment-template.yaml
```

A vantagem deste método é que não há a necessidade de limpar o arquivo, além de serem apresentadas apenas as opções necessárias do recurso.

## Socorro, são muitas opções!

Calma, nós sabemos. Mas o ``kubectl`` pode lhe auxiliar um pouco em relação a isso. Ele contém a opção ``explain``, que você pode utilizar caso precise de ajuda com alguma opção em específico dos arquivos de recurso. A seguir alguns exemplos de sintaxe.

```
kubectl explain [recurso]

kubectl explain [recurso.caminho.para.spec]

kubectl explain [recurso.caminho.para.spec] --recursive
```

Exemplos:

```
kubectl explain deployment

kubectl explain pod --recursive

kubectl explain deployment.spec.template.spec
```

## Expondo o pod

Dispositivos fora do *cluster*, por padrão, não conseguem acessar os *pods* criados, como é comum em outros sistemas de contêineres. Para expor um *pod*, execute o comando a seguir.

```
kubectl expose pod nginx
```

Será apresentada a seguinte mensagem de erro:

```
error: couldn't find port via --port flag or introspection
See 'kubectl expose -h' for help and examples
```

O erro ocorre devido ao fato do k8s não saber qual é a porta de destino do contêiner que deve ser exposta (no caso, a 80/TCP). Para configurá-la, vamos primeiramente remover o nosso *pod* antigo:

```
kubectl delete -f meu-primeiro.yaml
```

Abra agora o arquivo ``meu-primeiro.yaml`` e adicione o bloco a seguir.

```yaml
...
spec:
       containers:
       - image: nginx
         imagePullPolicy: Always
         ports:
         - containerPort: 80
         name: nginx
         resources: {}
...
```

> Atenção!!! Arquivos YAML utilizam para sua tabulação dois espaços e não *tab*.

Feita a modificação no arquivo, salve-o e crie novamente o *pod* com o comando a seguir.

```
kubectl create -f meu-primeiro.yaml

pod/nginx created
```

Liste o pod.

```
kubectl get pod nginx

NAME    READY   STATUS    RESTARTS   AGE
nginx   1/1     Running   0          32s
```

O comando a seguir cria um objeto do k8s chamado de *Service*, que é utilizado justamente para expor *pods* para acesso externo.

```
kubectl expose pod nginx
```

Podemos listar todos os *services* com o comando a seguir.

```
kubectl get services

NAME         TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)   AGE
kubernetes   ClusterIP   10.96.0.1       <none>        443/TCP   8d
nginx        ClusterIP   10.105.41.192   <none>        80/TCP    2m30s
```

Como é possível observar, há dois *services* no nosso *cluster*: o primeiro é para uso do próprio k8s, enquanto o segundo foi o quê acabamos de criar. Utilizando o ``curl`` contra o endereço IP mostrado na coluna *CLUSTER-IP*, deve nos ser apresentada a tela principal do Nginx.

```
curl 10.105.41.192

```

Este *pod* está disponível para acesso a partir de qualquer nó do *cluster*.

## Limpando tudo e indo para casa

Para mostrar todos os recursos recém criados, pode-se utilizar uma das seguintes opções a seguir.

```
kubectl get all

kubectl get pod,service

kubectl get pod,svc
```

Note que o k8s nos disponibiliza algumas abreviações de seus recursos. Com o tempo você irá se familiar com elas. Para apagar os recursos criados, você pode executar os seguintes comandos.

```
kubectl delete -f meu-primeiro.yaml

kubectl delete service nginx
```

Liste novamente os recursos para ver se os mesmos ainda estão presentes.
