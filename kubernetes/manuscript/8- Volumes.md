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

## Persistent Volumes

Um **PersistentVolume (PV)** é uma parte do armazenamento no cluster que foi provisionado por um administrador ou provisionado dinamicamente usando **storage class**. O PV é um recurso assim como um nó e tem um ciclo de vida independente de um pod. 
Um **PersistentVolumeClaim (PVC)** é um requerimento de recurso do storage feito pelo usuário. é similar a um pod "pedindo" recurso para um nó. os PVC's consomem o PV. eles podem ser montados com ReadWriteOnce, ReadOnlyMany or ReadWriteMany.
