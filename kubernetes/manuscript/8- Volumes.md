## Sumário

- [Volumes](#volumes)
  - [Persistent Volumes](#persistent-volumes)
  - [Types of Volumes](#types-of-volumes)
  - [Kubectl Edit](#kubectl-edit)
- [ReplicaSet](#replicaset)
- [DaemonSet](#daemonset)
- [Rollouts e Rollbacks](#rollouts-e-rollbacks)
- [Cron Jobs](#cron-jobs)

# Volumes

Os Discos no container é ephemero. No caso quando um pod é restartado o pod volta ao estado inicial do seu storage.
Porém os Persistents volumes existem para que o volume existe independente do tempo de vida de um pod.
Tipos de volumes: ``https://kubernetes.io/docs/concepts/storage/volumes/``

## Types of Volumes
Kubernetes supports several types of volumes.
awsElasticBlockStore
azureDisk
azureFile
cephfs
cinder
configMap
downwardAPI
emptyDir
fc (fibre channel)
flocker (deprecated)
gcePersistentDisk
gitRepo (deprecated)
glusterfs
hostPath
iscsi
local
nfs
persistentVolumeClaim
portworxVolume
projected
quobyte (deprecated)
rbd
secret
storageOS (deprecated)
vsphereVolume

## Access Modes
 - ReadWriteOnce(RWO) — volume can be mounted as read-write by a single node.
 - ReadOnlyMany(ROX) — volume can be mounted read-only by many nodes.
 - ReadWriteMany(RWX) — volume can be mounted as read-write by many nodes.
 - ReadWriteOncePod(RWOP) — volume can be mounted as read-write by a single Pod.

## Persistent Volumes
- Um **PersistentVolume (PV)** é uma parte do armazenamento no cluster que foi provisionado por um administrador ou provisionado dinamicamente usando **storage class**. O PV é um recurso assim como um nó e tem um ciclo de vida independente de um pod.
- PV não precisa de namespace, logo são acessiveis em todo o cluster e em todos os namespaces.
- Ao contrário do Volumes, o ciclo de vida dos PVs é gerenciado pelo Kubernetes.

## Persistent Volume Clain
- Um **PersistentVolumeClaim (PVC)** é um requerimento de recurso do storage feito pelo usuário. é similar a um pod "pedindo" recurso para um nó. os PVC's consomem o PV.
- Um PVC é a ligação do pod ao PV
- Um PVC descreve a capacidade e as caracteristicas do disco necessárias para o pod, e o cluster tenta ligar esse e provisionar esse PV.
- O PVC deve estar no mesmo namespace do pod, e os clains podem especificar o tamanho e o access mode.
### Access Modes
 - ReadWriteOnce(RWO) — volume can be mounted as read-write by a single node.
 - ReadOnlyMany(ROX) — volume can be mounted read-only by many nodes.
 - ReadWriteMany(RWX) — volume can be mounted as read-write by many nodes.
 - ReadWriteOncePod(RWOP) — volume can be mounted as read-write by a single Pod.



## Storage Class
- Um **Storage Class** é uma forma dinâmica de provisionar volumes persistentes. Ele é usado em conjunto com um pvc para que pods façam requisição de novos storages.
- Storage class usa provisioners que são específicos de cada cloud provide, que dão ao kubernetes acesso ao storage físico.

- forma para ver qual storage class é criado no seu cloud provider.
```$ kubectl get storageclasses.storage.k8s.io```


A figura a seguir mostra a arquitetura de PVC, PV e storage class.

| ![Arquitetura Kubernetes](../images/pvc_pv_sc.png) |
|:---------------------------------------------------------------------------------------------:|
| *Arquitetura Kubernetes Volumes  [Ref: Ashish Patel KB article](https://medium.com/devops-mojo/kubernetes-storage-options-overview-persistent-volumes-pv-claims-pvc-and-storageclass-sc-k8s-storage-df71ca0fccc3)*

### Exemplos de PVC usando storage class RWO
