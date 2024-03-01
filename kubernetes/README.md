# Kubernetes Documentos e exemplos

Documentação com exemplos e conhecimentos sobre k8s.

## Conteúdo

<details>
<summary>0-indrotução</summary>

- [0-Introdução](manuscript/0-introdution.md)
  - [Alguns sites que devemos visitar](manuscript/0-introdution.md##Alguns-sites-que-devemos-visitar)
  - [Arquitetura do k8s](manuscript/0-introdution.md##Arquitetura-do-k8s)
  - [Conceitos-chave do k8s](manuscript/0-introdution.md##Conceitos-chave-do-k8s)
 
</details>

<details>
<summary>1-Instalação</summary>

- [1-Instalação](manuscript/1-install.md)
  - [Instalação do kubectl](manuscript/1-install.md##Instalação-do-kubectl)
  - [Instalação do Minikube Linux](manuscript/1-install.md##Instalação-do-Minikube-Linux)
  - [Iniciando, parando e excluindo o Minikube](manuscript/1-install.md##Iniciando-parando-e-excluindo-o-Minikube)
  - [Descobrindo o endereço do Minikube](manuscript/1-install.md##Descobrindo-o-endereço-do-Minikube)
  - [Dashboard](manuscript/1-install.md##Dashboard)

</details>

<details>
<summary>2-primeiros-passos</summary>

- [2-primeiros-passos](manuscript/2-primeiros-passos)
  - [Exibindo informações detalhadas sobre os nós](manuscript/2-primeiros-passos##Exibindo-informações-detalhadas-sobre-os-nós)
  - [Gerenciando objetos no kubernetes](manuscript/2-primeiros-passos##gerenciando-objetos-no-kubernetes)
  - [Verificando os namespaces e pods](manuscript/2-primeiros-passos##-Verificando-os-namespaces-e-pods)
  - [Executando nosso primeiro pod no k8s](manuscript/2-primeiros-passos##Executando-nosso-primeiro-pod-no-k8s)
  - [Verificar os últimos eventos do cluster](manuscript/2-primeiros-passos##Verificar-os-últimos-eventos-do-cluster)
  - [Efetuar o dump de um objeto em formato YAML](manuscript/2-primeiros-passos##-Efetuar-o-dump-de-um-objeto-em-formato-YAML)
  - [Expondo o pod](manuscript/2-primeiros-passos##Expondo-o-pod)

</details>

<details>
<summary>3-workloads</summary>

- [3-workloads](manuscript/3-workloads.md)
- [Namespace](manuscript/3-workloads.md#Namespace)
- [Cron Jobs](manuscript/3-workloads.md#Cron-Jobs)
- [3.1-pods](manuscript/3.1-pods.md)
    - [O Deployment e o ReplicaSet](manuscript/3.1-pods.md)
    - [Criando um ReplicaSet](manuscript/3.1-pods.md##Usando-Labels-e-selectors)
    - [Container Probes](manuscript/3.1-pods.md##Container-Probes)
    - [Termination Probe](manuscript/3.1-pods.md##Termination-Probe)
    - [Init container](manuscript/3.1-pods.md##init-container)
    - [Pod disruption budgets](manuscript/3.1-pods.md##Pod-disruption-budgets)  

</details>

<details>
<summary>DAY-5</summary>

- [DAY-5](day-5/README.md#day-5)
- [Conteúdo do Day-5](day-5/README.md#conteúdo-do-day-5)
- [Inicio da aula do Day-5](day-5/README.md#inicio-da-aula-do-day-5)
  - [O que iremos ver hoje?](day-5/README.md#o-que-iremos-ver-hoje)
  - [Instalação de um cluster Kubernetes](day-5/README.md#instalação-de-um-cluster-kubernetes)
    - [O que é um cluster Kubernetes?](day-5/README.md#o-que-é-um-cluster-kubernetes)
    - [Formas de instalar o Kubernetes](day-5/README.md#formas-de-instalar-o-kubernetes)
    - [Criando um cluster Kubernetes com o kubeadm](day-5/README.md#criando-um-cluster-kubernetes-com-o-kubeadm)
      - [Instalando o kubeadm](day-5/README.md#instalando-o-kubeadm)
      - [Desativando o uso do swap no sistema](day-5/README.md#desativando-o-uso-do-swap-no-sistema)
      - [Carregando os módulos do kernel](day-5/README.md#carregando-os-módulos-do-kernel)
      - [Configurando parâmetros do sistema](day-5/README.md#configurando-parâmetros-do-sistema)
      - [Instalando os pacotes do Kubernetes](day-5/README.md#instalando-os-pacotes-do-kubernetes)
      - [Instalando o Docker e o containerd](day-5/README.md#instalando-o-docker-e-o-containerd)
      - [Configurando o containerd](day-5/README.md#configurando-o-containerd)
      - [Habilitando o serviço do kubelet](day-5/README.md#habilitando-o-serviço-do-kubelet)
      - [Configurando as portas](day-5/README.md#configurando-as-portas)
      - [Iniciando o cluster](day-5/README.md#iniciando-o-cluster)
      - [Entendendo o arquivo admin.conf](day-5/README.md#entendendo-o-arquivo-adminconf)
      - [Instalando o Weave Net](day-5/README.md#instalando-o-weave-net)
      - [O que é o CNI?](day-5/README.md#o-que-é-o-cni)
    - [Visualizando detalhes dos nodes](day-5/README.md#visualizando-detalhes-dos-nodes)
  - [A sua lição de casa](day-5/README.md#a-sua-lição-de-casa)
- [Final do Day-5](day-5/README.md#final-do-day-5)

</details>

<details>
<summary>DAY-6</summary>

- [DAY-6](day-6/README.md#day-6)
  - [Conteúdo do Day-6](day-6/README.md#conteúdo-do-day-6)
  - [Inicio da aula do Day-6](day-6/README.md#inicio-da-aula-do-day-6)
    - [O que iremos ver hoje?](day-6/README.md#o-que-iremos-ver-hoje)
      - [O que são volumes?](day-6/README.md#o-que-são-volumes)
        - [EmpytDir](day-6/README.md#empytdir)
        - [Storage Class](day-6/README.md#storage-class)
        - [PV - Persistent Volume](day-6/README.md#pv---persistent-volume)
        - [PVC - Persistent Volume Claim](day-6/README.md#pvc---persistent-volume-claim)
    - [A sua lição de casa](day-6/README.md#a-sua-lição-de-casa)
  - [Final do Day-6](day-6/README.md#final-do-day-6)

</details>

<details>
<summary>DAY-7</summary>

- [DAY-7](day-7/README.md#day-7)
- [Conteúdo do Day-7](day-7/README.md#conteúdo-do-day-7)
  - [O que iremos ver hoje?](day-7/README.md#o-que-iremos-ver-hoje)
    - [O que é um StatefulSet?](day-7/README.md#o-que-é-um-statefulset)
      - [Quando usar StatefulSets?](day-7/README.md#quando-usar-statefulsets)
      - [E como ele funciona?](day-7/README.md#e-como-ele-funciona)
      - [O StatefulSet e os volumes persistentes](day-7/README.md#o-statefulset-e-os-volumes-persistentes)
      - [O StatefulSet e o Headless Service](day-7/README.md#o-statefulset-e-o-headless-service)
      - [Criando um StatefulSet](day-7/README.md#criando-um-statefulset)
      - [Excluindo um StatefulSet](day-7/README.md#excluindo-um-statefulset)
      - [Excluindo um Headless Service](day-7/README.md#excluindo-um-headless-service)
      - [Excluindo um PVC](day-7/README.md#excluindo-um-pvc)
    - [Services](day-7/README.md#services)
      - [Tipos de Services](day-7/README.md#tipos-de-services)
      - [Como os Services funcionam](day-7/README.md#como-os-services-funcionam)
      - [Os Services e os Endpoints](day-7/README.md#os-services-e-os-endpoints)
      - [Criando um Service](day-7/README.md#criando-um-service)
        - [ClusterIP](day-7/README.md#clusterip)
        - [ClusterIP](day-7/README.md#clusterip-1)
        - [LoadBalancer](day-7/README.md#loadbalancer)
        - [ExternalName](day-7/README.md#externalname)
      - [Verificando os Services](day-7/README.md#verificando-os-services)
      - [Verificando os Endpoints](day-7/README.md#verificando-os-endpoints)
      - [Removendo um Service](day-7/README.md#removendo-um-service)
  - [A sua lição de casa](day-7/README.md#a-sua-lição-de-casa)
- [Final do Day-7](day-7/README.md#final-do-day-7)

&nbsp;

## O treinamento Descomplicando o Kubernetes - Expert Mode

Pensamos em fazer um treinamento realmente prático onde a pessoa consiga aprender os conceitos e teoria com excelente didática, utilizando exemplos e desafios práticos para que você consiga executar todo o conhecimento adquirido. Isso é muito importante para que você consiga fixar e explorar ainda mais o conteúdo do treinamento.
E por fim, vamos simular algumas conversas para que fique um pouco mais parecido com o dia-a-dia no ambiente de trabalho.

Durante o treinamento vamos passar por todos os tópicos importantes do Kubernetes, para que no final do treinamento você possua todo conhecimento e também toda a segurança para implementar e administrar o Kubernetes em ambientes críticos e complexos.

Estamos prontos para iniciar a nossa viagem?
&nbsp;

### O conteúdo programático

O conteúdo ainda será ajustado, e no final do treinamento teremos o conteúdo completo.

&nbsp;

### Como adquirir o treinamento?

Para adquirir o treinamento [Descomplicando o Kubernetes](https://www.linuxtips.io/) você deverá ir até a loja da [LINUXtips](https://www.linuxtips.io/).

&nbsp;

## A ideia do formato do treinamento

Ensinar Kubernetes de uma forma mais real, passando todo o conteúdo de forma prática e trazendo uma conexão com o ambiente real de trabalho.

Esse é o primeiro treinamento sobre Kubernetes de forma realmente prática, da vida real. Pois entendemos que prática é o conjunto de entendimento sobre um assunto, seguido de exemplos reais que possam ser reproduzidos e conectando tudo isso com a forma como trabalhamos.

Assim a definição de prática passa a ser um focada em o conhecimento da ferramenta e adicionando a realidade de um profissional no seu dia-a-dia aprendendo uma nova tecnologia, uma nova ferramenta.

Prepare-se para um novo tipo de treinamento, e o melhor, prepare-se para um novo conceito sobre treinamento prático e de aprendizado de tecnologia.
&nbsp;

### As pessoas (personagens) no treinamento

Temos algumas pessoas que vão nos ajudar durante o treinamento, simulando uma dinâmica um pouco maior e ajudando na imersão que gostaríamos.

Ainda estamos desenvolvendo e aprimorando os personagens e o enredo, portanto ainda teremos muitas novidades.

&nbsp;

#### A Pessoa_X

A Pessoa_X é uma das pessoas responsáveis pela loja de meias Strigus Socket, que está no meio da modernização de seu infra e das ferramentas que são utilizadas.

Segundo uma pessoa que já trabalhou com a Pessoa_X, ela é a pessoa que está sempre procurando aprender para inovar em seu ambiente. Normalmente é através dela que surgem as novas ferramentas, bem como a resolução de um monte de problemas.

O nível de conhecimento dela é sempre iniciante quando ela entra em um novo projeto, porém ao final dele, ela se torna uma especialista e com uma boa experiência prática, pois ela foi exposta a diversas situações, que a fizeram conhecer a nova tecnologia muito bem e se sentindo muito confortável em trabalhar no projeto.

Pessoa_X, foi um prazer fazer essa pequena descrição sobre você!

Seja bem-vinda nesse novo projeto e espero que você se divirta como sempre!

Lembre-se sempre que eu, Jeferson, estarei aqui para apoiar você em cada etapa dessa jornada! Eu sou o seu parceiro nesse projeto e tudo o que você precisar nessa jornada! Bora!

&nbsp;

#### A Pessoa_Lider_X

Iremos criando a personalidade dessa pessoa durante o treinamento.
O que já sabemos é que ela é a pessoa líder imediata da Pessoa_X, e que irá demandar a maioria das tarefas. E tem como o esteriótipo um líder meio tosco.
&nbsp;

#### A Pessoa_Diretora_X

Líder imediato da Pessoa_Lider_X e que tem um sobrinho 'jênio' e que está ali, dando os seus pitacos no setor de tecnologia, por que ele 'mereceu', entendeu?

&nbsp;

#### A Pessoa_RH_X

A pessoa responsável pelo RH da empresa, no decorrer do treinamento vamos faz
endo a história e características dela.
&nbsp;

## Vamos começar?

Agora que você já conhece mais detalhes sobre o treinamento, acredito que já podemos começar, certo?

Lembrando que o treinamento está disponível na plataforma da escola da LINUXtips, que não é o mesmo endereço da [loja](https://www.linuxtips.io/), para acessar a escola [CLIQUE AQUI](https://www.linuxtips.io).
&nbsp;

