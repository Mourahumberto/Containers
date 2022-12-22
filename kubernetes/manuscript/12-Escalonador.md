# Scheduling, preemption e eviction
- kubernetes scheduller se refere se um pod está dando match em algum nó. Premption é o processo de encerrar pods com menos prioridades, para dar lugar a pods com mais prioridade e eviction é o propcesso de encerrar pods em determinados nós.

### Kube-scheduler
- É o escalonador padrão do kubernetes. Esse escalonador é executado como parte do control plane. o escalonador leva alguns fatores na hora de schedular um pod.
como restrições de hadware, software, políticas, especificações, afinidade e anti-afinidadee assim por diante. O node seleciona o nó paraq um pod levando em conta duas Etapas:
  - Filtragem
  - Pontuação

## Assinando Pods aos nós
- Algumas vezes é importante restringir pods a rodar em nós específicos por N fatores, então você pode fazer isso com as seguintes estratégias
1. NodeSelector
   - Você usa essa estratégia, quando quer que determinados pods rodem em nodes com determinadas labels.
2. affinity e anti-affinity
   - É uma forma mais flexivel de escalonar um Pod, você pode colocar que um nó é preferêncial e caso não der pra colocar nesse nó o pod irá para outro. Node affinity
   - Também pode ser colocado afinidades em pods, para que um pod prefira ou evita ficar em junto a outros pods. Inter-pod affinity/anti-affinity 
   1. Node Afinnity é similar ao node selector com dois tipos de afinidade.
     - exemplo: | [afinidades de nós](../manifest/pods-affinity/node-affinity.yaml) |
     - requiredDuringSchedulingIgnoredDuringExecution:(hard) Desta forma assim como o node selector, ele só pode schedular quando o pod tem met com o nó. Porém, de forma mais expressiva.
     - preferredDuringSchedulingIgnoredDuringExecution:(soft) o scheduler tenta encontrar o nó para o pod com as labels desejadas, caso não encontre ele coloca em um nó qualquer.
   - Os operadores usados são: In, NotIn, Exists, DoesNotExist, Gt and Lt.
   2. Node affinity weight
     - você pode colocar uma medida de 1 a 100, se tiverem várias regras e elas tiverem dois grupos de nós que podem receber os pods, o scheduler irá ver qual tem a pontuação maior, caso você coloque um com 99 e outro com 1 ele irá escolher o nó com número maior
     - exemplo: | [afinidades de nós com pontuação](../manifest/pods-affinity/node-affinity-weight.yaml) |
   3. Node affinity per scheduling profile
     - Quando são configurado multiplos scheduling profiles 
   4. Inter pod affinitty e anti-affinity
     - Essa ferramenta te permite squedular o pod baseado em quais pod tem ou não tem naquele nó. Desta forma você pode preferir que pods do mesmo microsserviço tenham uma anbti-affinidade, para que prefiram ficar em nós separados para auta-disponibilidade.
     - Lembrando que quanto mais regras para squedular mais processamento do squeduler é necessário, não recomendado para cluster com centenas de nós.
     - exemplo: | [afinidades de pods](../manifest/pods-affinity/pod-with-pod-affinity.yaml) |
     - outro exemplo do redis onde os pods devem ficar em nós separados para ter auta disponibilidade | [redis](../manifest/pods-affinity/redis-antiaffinity.yamll) |
     - exemplo de aplicação que quer ficar próximos ao pods de redis mas separados de suas cópias. [app e redis](../manifest/pods-affinity/app-redis.yaml) |
  - nodeName
  - Pod topology spread constrains
      1. Como motivação, caso você queira separ replicas do mesmo deployment, porém quer deixar as réplicas próximos de outros deployments ou zonas.
      2. https://kubernetes.io/docs/concepts/scheduling-eviction/topology-spread-constraints/

## Pod Priority and Preemption
- Os pods do seu sistema pode ter prioridade a cima de outros pods. indicando a importância dele em relação a outros pods.
- doc: https://kubernetes.io/docs/concepts/scheduling-eviction/pod-priority-preemption/