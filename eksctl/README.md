site oficial: https://eksctl.io/
github: https://github.com/weaveworks/eksctl

 1- Criação de role do EKS, que posteriormente será usado no EKS.

 2-criação de usuários e chaves ssh

 3-instalação de aws-cli, eksctl e kubectl


CLUSTER 
# Para criar o cluster
$ eksctl create cluster -f hom-kubernetes.yml

# Para deletar o cluster
$ eksctl delete cluster -f hom-kubernetes.yml

descobrir os clusters
eksctl get cluster

NODEGROUP 
link: https://eksctl.io/usage/managing-nodegroups/

# Para scalar o nodegroup por linha de comando
$ eksctl scale nodegroup --cluster=EKS-course-cluster --nodes=5 --name=ng-1 --nodes-min=3  --nodes-max=5

# Para escalar para cima/baixo um recurso usando o yaml.
$ eksctl scale nodegroup --name=perf-1 --config-file=perf-kubernetes.yml

# Para adicionar um novo nodegroup
$ eksctl create nodegroup --config-file=eks-hom.yml --include='ng-teste'

# Para deletar um nodegrroup apenas
$ eksctl delete nodegroup --config-file=eks-hom.yml --include=ng-testando --approve

# descobrir os nodes groups
$ eksctl get nodegroup --cluster basic-cluster

VPC
exemplo com cidr: https://eksctl.io/examples/reusing-iam-and-vpc/

## Adicionando o contexto do nosso cluster ao kubectl

```bash
aws eks --region us-east-1 update-kubeconfig --name nome-do-cluster
aws eks --region us-east-1 update-kubeconfig --name k8s-demo
aws eks --region us-east-1 update-kubeconfig --name simcloud-eks --profile sim

```


