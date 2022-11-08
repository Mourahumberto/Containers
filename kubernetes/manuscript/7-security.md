# Security
- Isso mostra um overview sobre segurança em um contexto cloud native.

### Os 4 C's do cloud native security
- Cloud, Cluster, Container e Code

## Cloud
- Cada cloud tem suas regras e dicas de segurança, verifique isso na cloud que está usando.

## Cluster
- sugestões de segurança

  ### Infraestrutura

  - **Network acesso para o Control plane**: O Control plane não pode ser aberta a internet, Ela deve ter uma lista restrita dos IP's que podem acessá-la
  - **Network Access para os Nós**: Se possível não deixe esses nós públicos na internet, e que eles recebam conexão apenas do controlplane em portas específicas, ou aceitem requisições a serviços NodePort e load balances.
  - **Acesso ao Kubernetes API**: Cada cloud provider tem suas formas diferentes de autentificação a API a documentação do Kops explica sobre IAM roles e polices. https://github.com/kubernetes/kops/blob/master/docs/iam_roles.md#iam-roles

  ### Componentes do cluster
  - RBAC Authorization
  - Authentication
  - Aplication secrets
  - Garantir que os pods atendam aos padrões de segurança
  - Network Polices
  - TLS kubernetes Ingress

## Container
- Vulnerabilidades de OS ou pacotes nos containers, um container deve ter o mínimo possível que precisa pra rodar.
- Containers de imagens oficiais
- Desativar privilégio de usuário, criar usuários dentro do container que tenham permissão de rodar o processo e não de root.

## Code
- Acesso apenas por TLS, se sua aplicação se comunica com TCP use TLS ou mTLS
- Limite o número de portas expostas por sua aplicação expondo só as que forem necessária como a que o serviço usa e a de coleta de métricas.
- Faça o scan das vulnerabilidades das bibliotecas usadas.
- analize estática dos códigos
- algumas ferramentas fazem o teste dinâmica do código a mais conhecida é a OWASP https://www.zaproxy.org/


## Pod Security Admission

- O Kubernetes oferece um controlador de admissão Pod Security integrado para aplicar os padrões de segurança do Pod. As restrições de segurança do pod são aplicadas no nível do namespace quando os pods são criados.

# Multi tanancy
- Aqui irei falar sobre as melhores práticas quando se tem multiplos times em um mesmo cluster kubernetes.

## Namespaces
- è interessante que cada usuário só possa acessar recursos que estão no namespace destinado para ele, Desta forma com o RBAC você cria usuários e service acounts que só podem criar e ler recurso dentro do namespace estipulado.

## Quotas
- Os workload consomem recurso computacional, desta forma é interessante colocar quotas em cada namespace, para evitar "vizinhos barulhentos". desta forma, você pode definir quantos pods, configmaps e etc cada namespace pode ter.
- doc quotas https://kubernetes.io/pt-br/docs/concepts/policy/resource-quotas/
- doc limit ranges https://kubernetes.io/pt-br/docs/concepts/policy/limit-range/

## Network Isolation
- Por default todos os pods dentro do cluster podem se comunicar e não há encripitação na comunicação entre os pods.
- a comunicação de pod-to-pod pode ser controlada por network polices. Que podem ser isoladas usando namespaces labels ou range de ips.

## storage Isolation
- o provisionamento de volume dinâmico é recomendado e os tipos de volume que usam o disco do nó devem ser evitados.
- StorageClasses permite você descobrir classes de storage oferecidos no seu cluster. os Pods pode pedir esse storage através do recurso PersistentStorageClain, que habilita o isolamento de partes do storage.
- Para isolar você pode criar um StorageClass pra cada cliente ou marcar o reclaim policy como delete, para ter certeza que o PV não vai ser reutilazado por outro recurso.

## Nodes Isolation
- Você pode criar nós isolados que apenas rodam os pods de um cliente.Então todos os clientes terão nós dedicados.
  
## Quality-of-Service(QoS)
- Caso você tenha vários clientes com diferentes tiers, Existem várias features que te ajudam, como requests/limits, network Qos, storage class e pods priority e preemption

# Implementações
- Existem duas formas de implementar, por namespace isolation ou por virtualização de clusters.

## Namespace per tenant(client)
[- doc namespace per tenant](https://kubernetes.io/docs/concepts/security/multi-tenancy/#namespace-per-tenant)

## Virtual control plane
- nesse você consegue virtualizar vários control planes e clusters dentro de um só cluste.
[- doc vcluster per tenat](https://kubernetes.io/docs/concepts/security/multi-tenancy/#virtual-control-plane-per-tenant)

## Checklist de segurança
- https://kubernetes.io/docs/concepts/security/security-checklist/