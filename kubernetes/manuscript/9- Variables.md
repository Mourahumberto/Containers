# Variables
- Quando você cria um pod você pode colocar variáveis de ambiente nos containers
```
apiVersion: v1
kind: Pod
metadata:
  name: envar-demo
  labels:
    purpose: demonstrate-envars
spec:
  containers:
  - name: envar-demo-container
    image: gcr.io/google-samples/node-hello:1.0
    env:
    - name: DEMO_GREETING
      value: "Hello from the environment"
    - name: DEMO_FAREWELL
      value: "Such a sweet sorrow
```

# Secrets

Objetos do tipo **Secret** são normalmente utilizados para armazenar informações confidenciais, como por exemplo tokens e chaves SSH. Deixar senhas e informações confidenciais em arquivo texto não é um bom comportamento visto do olhar de segurança. Colocar essas informações em um objeto ``Secret`` permite que o administrador tenha mais controle sobre eles reduzindo assim o risco de exposição acidental.

Vamos criar nosso primeiro objeto ``Secret`` utilizando o arquivo ``secret.txt`` que vamos criar logo a seguir.

```
echo -n "giropops strigus girus" > secret.txt
```

Agora que já temos nosso arquivo ``secret.txt`` com o conteúdo ``descomplicando-k8s`` vamos criar nosso objeto ``Secret``.

```
kubectl create secret generic my-secret --from-file=secret.txt

secret/my-secret created
```

Vamos ver os detalhes desse objeto para ver o que realmente aconteceu.

```
kubectl describe secret my-secret

Name:         my-secret
Namespace:    default
Labels:       <none>
Annotations:  <none>

Type:  Opaque

Data
====
secret.txt:  18 bytes
```

Observe que não é possível ver o conteúdo do arquivo utilizando o ``describe``, isso é para proteger a chave de ser exposta acidentalmente.

Para verificar o conteúdo de um ``Secret`` precisamos decodificar o arquivo gerado, para fazer isso temos que verificar o manifesto do do mesmo.

```
kubectl get secret

NAME              TYPE             DATA      AGE
my-secret         Opaque           1         13m
```

```
kubectl get secret my-secret -o yaml

apiVersion: v1
data:
  secret.txt: Z2lyb3BvcHMgc3RyaWd1cyBnaXJ1cw==
kind: Secret
metadata:
  creationTimestamp: 2018-08-26T17:10:14Z
  name: my-secret
  namespace: default
  resourceVersion: "3296864"
  selfLink: /api/v1/namespaces/default/secrets/my-secret
  uid: e61d124a-a952-11e8-8723-42010a8a0002
type: Opaque
```

Agora que já temos a chave codificada basta decodificar usando ``Base64``.

```
echo 'Z2lyb3BvcHMgc3RyaWd1cyBnaXJ1cw==' | base64 --decode

giropops strigus girus
```

Tudo certo com nosso ``Secret``, agora vamos utilizar ele dentro de um Pod, para isso vamos precisar referenciar o ``Secret`` dentro do Pod utilizando volumes, vamos criar nosso manifesto.

```
vim pod-secret.yaml
```

Informe o seguinte conteúdo:

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: test-secret
  namespace: default
spec:
  containers:
  - image: busybox
    name: busy
    command:
      - sleep
      - "3600"
    volumeMounts:
    - mountPath: /tmp/giropops
      name: my-volume-secret
  volumes:
  - name: my-volume-secret
    secret:
      secretName: my-secret
```

Nesse manifesto vamos utilizar o volume ``my-volume-secret`` para montar dentro do contêiner a Secret ``my-secret`` no diretório ``/tmp/giropos``.

```
kubectl create -f pod-secret.yaml

pod/test-secret created
```

Vamos verificar se o ``Secret`` foi criado corretamente:

```
kubectl exec -ti test-secret -- ls /tmp/giropops

secret.txt
```

```
kubectl exec -ti test-secret -- cat /tmp/giropops/secret.txt

giropops strigus girus
```

Sucesso! Esse é um dos modos de colocar informações ou senha dentro de nossos Pods, mas existe um jeito ainda mais bacana utilizando os Secrets como variável de ambiente.

Vamos dar uma olhada nesse cara, primeiro vamos criar um novo objeto ``Secret`` usando chave literal com chave e valor.

```
kubectl create secret generic my-literal-secret --from-literal user=linuxtips --from-literal password=catota

secret/my-literal-secret created
```

Vamos ver os detalhes do objeto ``Secret`` ``my-literal-secret``:

```
kubectl describe secret my-literal-secret

Name:         my-literal-secret
Namespace:    default
Labels:       <none>
Annotations:  <none>

Type:  Opaque

Data
====
password:  6 bytes
user:      9 bytes
```

Acabamos de criar um objeto ``Secret`` com duas chaves, um ``user`` e outra ``password``, agora vamos referenciar essa chave dentro do nosso Pod utilizando variável de ambiente, para isso vamos criar nosso novo manifesto.

```
vim pod-secret-env.yaml
```

Informe o seguinte conteúdo:

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: teste-secret-env
  namespace: default
spec:
  containers:
  - image: busybox
    name: busy-secret-env
    command:
      - sleep
      - "3600"
    env:
    - name: MEU_USERNAME
      valueFrom:
        secretKeyRef:
          name: my-literal-secret
          key: user
    - name: MEU_PASSWORD
      valueFrom:
        secretKeyRef:
          name: my-literal-secret
          key: password
```

Vamos criar nosso pod.

```
kubectl create -f pod-secret-env.yaml

pod/teste-secret-env created
```

Agora vamos listar as variáveis de ambiente dentro do contêiner para verificar se nosso Secret realmente foi criado.

```
kubectl exec teste-secret-env -c busy-secret-env -it -- printenv

PATH=/usr/local/sbin:/usr/local/bin:/usr/sbin:/usr/bin:/sbin:/bin
HOSTNAME=teste-secret-env
TERM=xterm
MEU_USERNAME=linuxtips
MEU_PASSWORD=catota
KUBERNETES_PORT_443_TCP_PROTO=tcp
KUBERNETES_PORT_443_TCP_PORT=443
KUBERNETES_PORT_443_TCP_ADDR=10.96.0.1
KUBERNETES_SERVICE_HOST=10.96.0.1
KUBERNETES_SERVICE_PORT=443
KUBERNETES_SERVICE_PORT_HTTPS=443
KUBERNETES_PORT=tcp://10.96.0.1:443
KUBERNETES_PORT_443_TCP=tcp://10.96.0.1:443
HOME=/root
```

Viram? Agora podemos utilizar essa chave dentro do contêiner como variável de ambiente, caso alguma aplicação dentro do contêiner precise se conectar ao um banco de dados por exemplo utilizando usuário e senha, basta criar um ``secret`` com essas informações e referenciar dentro de um Pod depois é só consumir dentro do Pod como variável de ambiente ou um arquivo texto criando volumes.

# ConfigMaps

Os Objetos do tipo **ConfigMaps** são utilizados para separar arquivos de configuração do conteúdo da imagem de um contêiner, assim podemos adicionar e alterar arquivos de configuração dentro dos Pods sem buildar uma nova imagem de contêiner.

Para nosso exemplo vamos utilizar um ``ConfigMaps`` configurado com dois arquivos e  um valor literal.

Vamos criar um diretório chamado ``frutas`` e nele vamos adicionar frutas e suas características.

```
mkdir frutas

echo -n amarela > frutas/banana

echo -n vermelho > frutas/morango

echo -n verde > frutas/limao

echo -n "verde e vermelha" > frutas/melancia

echo -n kiwi > predileta
```

Crie o ``Configmap``.

```
kubectl create configmap cores-frutas --from-literal=uva=roxa --from-file=predileta --from-file=frutas/
```

Visualize o Configmap.

```
kubectl get configmap
```

Vamos criar um pod para usar o Configmap:

```
vim pod-configmap.yaml
```

Informe o seguinte conteúdo:

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: busybox-configmap
  namespace: default
spec:
  containers:
  - image: busybox
    name: busy-configmap
    command:
      - sleep
      - "3600"
    env:
    - name: frutas
      valueFrom:
        configMapKeyRef:
          name: cores-frutas
          key: predileta
```

Crie o pod a partir do manifesto.

```
kubectl create -f pod-configmap.yaml
```

Após a criação, execute o comando ``set`` dentro do contêiner, para listar as variáveis de ambiente e conferir se foi criada a variável de acordo com a ``key=predileta`` que definimos em nosso arquivo yaml.

Repare no final da saída do comando ``set`` a env ``frutas='kiwi'``.

```
kubectl exec -ti busybox-configmap -- sh 
/ # set
...
frutas='kiwi'
```

Vamos criar um pod utilizando utilizando mais de uma variável:

```
vim pod-configmap-env.yaml
```

Informe o seguinte conteúdo:

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: busybox-configmap-env
  namespace: default
spec:
  containers:
  - image: busybox
    name: busy-configmap
    command:
      - sleep
      - "3600"
    envFrom:
    - configMapRef:
        name: cores-frutas
```

Crie o pod a partir do manifesto:

```
kubectl create -f pod-configmap-env.yaml
```

Vamos entrar no contêiner e executar o comando ``set`` novamente para listar as variáveis, repare que foi criada todas as variáveis.

```
kubectl exec -ti busybox-configmap-env -- sh

/ # set
...
banana='amarela'
limao='verde'
melancia='verde e vermelha'
morango='vermelho'
predileta='kiwi'
uva='roxa'
```

Agora vamos criar um pod para usar outro Configmap, só que dessa vez utilizando volume:

```
vim pod-configmap-file.yaml
```

Informe o seguinte conteúdo:

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: busybox-configmap-file
  namespace: default
spec:
  containers:
  - image: busybox
    name: busy-configmap
    command:
      - sleep
      - "3600"
    volumeMounts:
    - name: meu-configmap-vol
      mountPath: /etc/frutas
  volumes:
  - name: meu-configmap-vol
    configMap:
      name: cores-frutas
```

Crie o pod a partir do manifesto.

```
kubectl create -f pod-configmap-file.yaml
```

Após a criação do pod, vamos conferir o nosso configmap como arquivos.

```
kubectl exec -ti busybox-configmap-file -- sh
/ # ls -lh /etc/frutas/
total 0      
lrwxrwxrwx    1 root     root          13 Sep 23 04:56 banana -> ..data/banana
lrwxrwxrwx    1 root     root          12 Sep 23 04:56 limao -> ..data/limao
lrwxrwxrwx    1 root     root          15 Sep 23 04:56 melancia -> ..data/melancia
lrwxrwxrwx    1 root     root          14 Sep 23 04:56 morango -> ..data/morango
lrwxrwxrwx    1 root     root          16 Sep 23 04:56 predileta -> ..data/predileta
lrwxrwxrwx    1 root     root          10 Sep 23 04:56 uva -> ..data/uva
```
