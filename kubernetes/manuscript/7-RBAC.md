# RBAC
- O Kubernetes RBAC é um controle de segurança fundamental para garantir que os usuários e cargas de trabalho do cluster tenham acesso apenas aos recursos necessários para executar suas funções. É importante garantir que, ao projetar permissões para usuários do cluster, o administrador do cluster entenda as áreas em que o escalonamento de privilégios pode ocorrer, para reduzir o risco de acesso excessivo que leva a incidentes de segurança.
- RBAC pode ser aplicado a pessoas e a services accounts
- Sempre que possivel user RoleBindings no lugar de clusterrolebindings, para dar permissões a namespaces e não ao cluster todo.

## Criando um usuário no Kubernetes
- habilitando o CSR https://kubernetes.io/docs/reference/access-authn-authz/certificate-signing-requests/
- criando um usuário no minikube: https://medium.com/@HoussemDellai/rbac-with-kubernetes-in-minikube-4deed658ea7b

Para criar um usuário no Kubernetes, vamos precisar gerar um CSR (*Certificate Signing Request*) para o usuário. O usuário que vamos utilizar como exemplo é o ``linuxtips``.

Comando para gerar o CSR:

```
openssl req -new -newkey rsa:4096 -nodes -keyout linuxtips.key -out linuxtips.csr -subj "/CN=linuxtips"
```

Agora vamos fazer o request do CSR no cluster:

```
cat <<EOF | kubectl apply -f -
apiVersion: certificates.k8s.io/v1
kind: CertificateSigningRequest
metadata:
  name: linuxtips-csr
  namespace: default
spec:
  groups:
  request: $(cat linuxtips.csr | base64 | tr -d '\n')
  signerName: kubernetes.io/kube-apiserver-client
  usages:
  - client auth
EOF
```

Para ver os CSR criados, utilize o seguinte comando:

```
kubectl get csr
```

O CSR deverá estar com o status ``Pending``, vamos aprová-lo

```
kubectl certificate approve linuxtips-csr
```

Agora o certificado foi assinado pela CA (*Certificate Authority*) do cluster, para pegar o certificado assinado vamos usar o seguinte comando:

```
kubectl get csr linuxtips-csr -o jsonpath='{.status.certificate}' | base64 --decode > linuxtips.crt
```

Será necessário para a configuração do ``kubeconfig`` o arquivo referente a CA do cluster, para obtê-lá vamos extrai-lá do ``kubeconf`` atual que estamos utilizando:

```
kubectl config view -o jsonpath='{.clusters[0].cluster.certificate-authority-data}' --raw | base64 --decode - > ca.crt
```

Feito isso vamos montar nosso ``kubeconfig`` para o novo usuário:

Vamos pegar as informações de IP cluster:

```
kubectl config set-cluster $(kubectl config view -o jsonpath='{.clusters[0].name}') --server=$(kubectl config view -o jsonpath='{.clusters[0].cluster.server}') --certificate-authority=ca.crt --kubeconfig=linuxtips-config --embed-certs
```

Agora setando as confs de ``user`` e ``key``:

```
kubectl config set-credentials linuxtips --client-certificate=linuxtips.crt --client-key=linuxtips.key --embed-certs --kubeconfig=linuxtips-config
```

Agora vamos definir o ``context linuxtips`` e, em seguida, vamos utilizá-lo:

```
kubectl config set-context linuxtips --cluster=$(kubectl config view -o jsonpath='{.clusters[0].name}')  --user=linuxtips --kubeconfig=linuxtips-config
```

```
kubectl config use-context linuxtips --kubeconfig=linuxtips-config
```

Vamos ver um teste:

```
kubectl version --kubeconfig=linuxtips-config
```

Pronto! Agora só associar um ``role`` com as permissões desejadas para o usuário.

# RBAC

O controle de acesso baseado em funções (*Role-based Access Control* - RBAC) é um método para fazer o controle de acesso aos recursos do Kubernetes com base nas funções dos administradores individuais em sua organização.

A autorização RBAC usa o ``rbac.authorization.k8s.io`` para conduzir as decisões de autorização, permitindo que você configure políticas dinamicamente por meio da API Kubernetes.

## Role e ClusterRole

Um RBAC ``Role`` ou ``ClusterRole`` contém regras que representam um conjunto de permissões. As permissões são puramente aditivas (não há regras de "negação").

- Uma ``Role`` sempre define permissões em um determinado namespace, ao criar uma função, você deve especificar o namespace ao qual ela pertence.
- O ``ClusterRole``, por outro lado, é um recurso sem namespaces.

Você pode usar um ``ClusterRole`` para:

* Definir permissões em recursos com namespace e ser concedido dentro de namespaces individuais;
* Definir permissões em recursos com namespaces e ser concedido em todos os namespaces;
* Definir permissões em recursos com escopo de cluster.

Se você quiser definir uma função em um namespace, use uma ``Role``, caso queria definir uma função em todo o cluster, use um ``ClusterRole``. :)

## RoleBinding e ClusterRoleBinding

Uma ``RoleBinding`` concede as permissões definidas em uma função a um usuário ou conjunto de usuários. Ele contém uma lista de assuntos (usuários, grupos ou contas de serviço) e uma referência à função que está sendo concedida. Um ``RoleBinding`` concede permissões dentro de um namespace específico, enquanto um ``ClusterRoleBinding`` concede esse acesso a todo o cluster.

Um RoleBinding pode fazer referência a qualquer papel no mesmo namespace. Como alternativa, um RoleBinding pode fazer referência a um ClusterRole e vincular esse ClusterRole ao namespace do RoleBinding. Se você deseja vincular um ClusterRole a todos os namespaces em seu cluster, use um ClusterRoleBinding.

Com o ``ClusterRoleBinding`` você pode associar uma conta de serviço a uma determinada ClusterRole.

Primeiro vamos exibir as ClusterRoles:

```
kubectl get clusterrole

NAME                                                                   CREATED AT
admin                                                                  2020-09-20T04:35:27Z
calico-kube-controllers                                                2020-09-20T04:35:34Z
calico-node                                                            2020-09-20T04:35:34Z
cluster-admin                                                          2020-09-20T04:35:27Z
edit                                                                   2020-09-20T04:35:27Z
kubeadm:get-nodes                                                      2020-09-20T04:35:30Z
system:aggregate-to-admin                                              2020-09-20T04:35:28Z
system:aggregate-to-edit                                               2020-09-20T04:35:28Z
system:aggregate-to-view                                               2020-09-20T04:35:28Z
system:auth-delegator                                                  2020-09-20T04:35:28Z
system:basic-user                                                      2020-09-20T04:35:27Z
system:certificates.k8s.io:certificatesigningrequests:nodeclient       2020-09-20T04:35:28Z
system:certificates.k8s.io:certificatesigningrequests:selfnodeclient   2020-09-20T04:35:28Z
system:certificates.k8s.io:kube-apiserver-client-approver              2020-09-20T04:35:28Z
system:certificates.k8s.io:kube-apiserver-client-kubelet-approver      2020-09-20T04:35:28Z
system:certificates.k8s.io:kubelet-serving-approver                    2020-09-20T04:35:28Z
system:certificates.k8s.io:legacy-unknown-approver                     2020-09-20T04:35:28Z
system:controller:attachdetach-controller                              2020-09-20T04:35:28Z
system:controller:certificate-controller                               2020-09-20T04:35:28Z
system:controller:clusterrole-aggregation-controller                   2020-09-20T04:35:28Z
system:controller:cronjob-controller                                   2020-09-20T04:35:28Z
system:controller:daemon-set-controller                                2020-09-20T04:35:28Z
system:controller:deployment-controller                                2020-09-20T04:35:28Z
system:controller:disruption-controller                                2020-09-20T04:35:28Z
system:controller:endpoint-controller                                  2020-09-20T04:35:28Z
system:controller:endpointslice-controller                             2020-09-20T04:35:28Z
system:controller:endpointslicemirroring-controller                    2020-09-20T04:35:28Z
system:controller:expand-controller                                    2020-09-20T04:35:28Z
system:controller:generic-garbage-collector                            2020-09-20T04:35:28Z
system:controller:horizontal-pod-autoscaler                            2020-09-20T04:35:28Z
system:controller:job-controller                                       2020-09-20T04:35:28Z
system:controller:namespace-controller                                 2020-09-20T04:35:28Z
system:controller:node-controller                                      2020-09-20T04:35:28Z
system:controller:persistent-volume-binder                             2020-09-20T04:35:28Z
system:controller:pod-garbage-collector                                2020-09-20T04:35:28Z
system:controller:pv-protection-controller                             2020-09-20T04:35:28Z
system:controller:pvc-protection-controller                            2020-09-20T04:35:28Z
system:controller:replicaset-controller                                2020-09-20T04:35:28Z
system:controller:replication-controller                               2020-09-20T04:35:28Z
system:controller:resourcequota-controller                             2020-09-20T04:35:28Z
system:controller:route-controller                                     2020-09-20T04:35:28Z
system:controller:service-account-controller                           2020-09-20T04:35:28Z
system:controller:service-controller                                   2020-09-20T04:35:28Z
system:controller:statefulset-controller                               2020-09-20T04:35:28Z
system:controller:ttl-controller                                       2020-09-20T04:35:28Z
system:coredns                                                         2020-09-20T04:35:30Z
system:discovery                                                       2020-09-20T04:35:27Z
system:heapster                                                        2020-09-20T04:35:28Z
system:kube-aggregator                                                 2020-09-20T04:35:28Z
system:kube-controller-manager                                         2020-09-20T04:35:28Z
system:kube-dns                                                        2020-09-20T04:35:28Z
system:kube-scheduler                                                  2020-09-20T04:35:28Z
system:kubelet-api-admin                                               2020-09-20T04:35:28Z
system:node                                                            2020-09-20T04:35:28Z
system:node-bootstrapper                                               2020-09-20T04:35:28Z
system:node-problem-detector                                           2020-09-20T04:35:28Z
system:node-proxier                                                    2020-09-20T04:35:28Z
system:persistent-volume-provisioner                                   2020-09-20T04:35:28Z
system:public-info-viewer                                              2020-09-20T04:35:27Z
system:volume-scheduler                                                2020-09-20T04:35:28Z
view                                                                   2020-09-20T04:35:28Z
```

Lembra que falamos anteriormente que para associar uma ``ClusterRole``, precisamos criar uma ``ClusterRoleBinding``?

Para ver a função de cada uma delas, você pode executar o ``describe``, como fizemos para o **cluster-admin**, repare em ``Resources`` que tem o ``*.*``, ou seja, a ``ServiceAccount`` que pertence a este grupo, pode fazer tudo dentro do cluster.

```
kubectl describe clusterrole cluster-admin

Name:         cluster-admin
Labels:       kubernetes.io/bootstrapping=rbac-defaults
Annotations:  rbac.authorization.kubernetes.io/autoupdate: true
PolicyRule:
  Resources  Non-Resource URLs  Resource Names  Verbs
  ---------  -----------------  --------------  -----
  *.*        []                 []              [*]
             [*]                []              [*]
```

Vamos criar uma ``ServiceAccount`` na unha para ser administrador do nosso cluster:

```
kubectl create serviceaccount jeferson
```

Conferindo se a ``ServiceAccount`` foi criada:

```
kubectl get serviceaccounts

NAME      SECRETS   AGE
default   1         10d
jeferson  1         10s
```

Visualizando detalhes da ``ServiceAccount``:

```
kubectl describe serviceaccounts jeferson

Name:                jeferson
Namespace:           default
Labels:              <none>
Annotations:         <none>
Image pull secrets:  <none>
Mountable secrets:   jeferson-token-h8dz6
Tokens:              jeferson-token-h8dz6
Events:              <none>
```

Crie uma ``ClusterRoleBinding`` e associe a ``ServiceAccount`` **jeferson** e com a função ``ClusterRole`` **cluster-admin**:

```
kubectl create clusterrolebinding toskeria --serviceaccount=default:jeferson --clusterrole=cluster-admin
```

Exibindo a ``ClusterRoleBinding`` que acabamos de criar:

```
kubectl get clusterrolebindings.rbac.authorization.k8s.io toskeria

NAME       ROLE                        AGE
toskeria   ClusterRole/cluster-admin   118m
```

Vamos mandar um ``describe`` do ``ClusterRoleBinding`` **toskeira**.

```
kubectl describe clusterrolebindings.rbac.authorization.k8s.io toskeria

Name:         toskeria
Labels:       <none>
Annotations:  <none>
Role:
  Kind:  ClusterRole
  Name:  cluster-admin
Subjects:
  Kind            Name      Namespace
  ----            ----      ---------
  ServiceAccount  jeferson  default
```

Agora iremos criar um ``ServiceAccount`` **admin-user** a partir do manifesto ``admin-user.yaml``.

Crie o arquivo:

```
vim admin-user.yaml
```

Informe o seguinte conteúdo:

```yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  name: admin-user
  namespace: kube-system
```

Crie a ``ServiceAccount`` partir do manifesto.

```
kubectl create -f admin-user.yaml
```

Checando se a ``ServiceAccount`` foi criada:

```
kubectl get serviceaccounts --namespace=kube-system admin-user

NAME         SECRETS   AGE
admin-user   1         10s
```

Agora iremos associar o **admin-user** ao **cluster-admin** também a partir do manifesto ``admin-cluster-role-binding.yaml``.

Crie o arquivo:

```
vim admin-cluster-role-binding.yaml
```

Informe o seguinte conteúdo:

```yaml
apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
  name: admin-user
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: cluster-admin
subjects:
- kind: ServiceAccount
  name: admin-user
  namespace: kube-system
```

Crie o ``ClusterRoleBinding`` partir do manifesto.

```
kubectl create -f admin-cluster-role-binding.yaml

clusterrolebinding.rbac.authorization.k8s.io/admin-user created
```

Cheque se o ``ClusterRoleBinding`` foi criado:

```
kubectl get clusterrolebindings.rbac.authorization.k8s.io admin-user

NAME         ROLE                        AGE
admin-user   ClusterRole/cluster-admin   21m
```

Pronto! Agora o usuário **admin-user** foi associado ao ``ClusterRoleBinding`` **admin-user** com a função **cluster-admin**.

Lembrando que toda vez que nos referirmos ao ``ServiceAccount``, estamos referindo à uma conta de usuário.

