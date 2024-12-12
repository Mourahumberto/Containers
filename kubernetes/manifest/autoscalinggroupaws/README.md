# Criando o autoscaling automático na AWS.

## Pontos importantes
- Antes de rodar o '''autoscaling.yaml''' precisamos que o pod tenha uma role associada a ele, atraves do OIDC, role direto no nó ou atraves de secret e access key.
- A policy necessária é a seguinte.
```yaml
{
  "Version": "2012-10-17",
  "Statement": [
    {
      "Effect": "Allow",
      "Action": [
        "autoscaling:DescribeAutoScalingGroups",
        "autoscaling:DescribeAutoScalingInstances",
        "autoscaling:DescribeLaunchConfigurations",
        "autoscaling:DescribeScalingActivities",
        "ec2:DescribeImages",
        "ec2:DescribeInstanceTypes",
        "ec2:DescribeLaunchTemplateVersions",
        "ec2:GetInstanceTypesFromInstanceRequirements",
        "eks:DescribeNodegroup"
      ],
      "Resource": ["*"]
    },
    {
      "Effect": "Allow",
      "Action": [
        "autoscaling:SetDesiredCapacity",
        "autoscaling:TerminateInstanceInAutoScalingGroup"
      ],
      "Resource": ["*"]
    }
  ]
}
```
## Aplicação do manifesto
- Irei aplicar esse manifesto usando uma role no pod já que o OIDC já está configurado e criei uma role usando esse identity provider.
- vídeo que mostra como configurar o oidc parao pod ter uma role atachada(https://www.youtube.com/watch?v=bu0M2y2g1m8)

## Alteração no manifesto
- adição do metadata com a role criada
```yaml
apiVersion: v1
kind: ServiceAccount
metadata:
  annotations: #
    eks.amazonaws.com/role-arn: arn:aws:iam::654654356287:role/testerole #
    eks.amazonaws.com/sts-regional-endpoints: "true" #
  labels:
    k8s-addon: cluster-autoscaler.addons.k8s.io
    k8s-app: cluster-autoscaler
  name: cluster-autoscaler
  namespace: kube-system
```

- altere para seu cluster a linha 165.
```
--node-group-auto-discovery=asg:tag=k8s.io/cluster-autoscaler/enabled,k8s.io/cluster-autoscaler/<YOUR CLUSTER NAME>
```
- e depois rode um 
```
$kubectl apply -f autoscaling.yaml
```

# DOC
- https://github.com/kubernetes/autoscaler/tree/master/cluster-autoscaler/cloudprovider/aws#Auto-discovery-setup
- https://aws.github.io/aws-eks-best-practices/cluster-autoscaling/