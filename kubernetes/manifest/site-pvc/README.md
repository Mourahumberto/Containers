DOC: https://docs.aws.amazon.com/eks/latest/userguide/efs-csi.html

Usando o rook para nfs e sharedstorage
DOC: https://rook.io/docs/rook/v1.8/quickstart.html

or documentation on running Rook in your Kubernetes cluster see the [Kubernetes Quickstart Guide](/Documentation/quickstart.md)

step-by-step
1- kubectl create -f crds.yaml -f common.yaml -f operator.yaml

2- verify the rook-ceph-operator is in the `Running` state before proceeding
kubectl -n rook-ceph get pod

3- 
3.1- change cluster-on-pvc.yaml in OVH
storageClassName: csi-cinder-high-speed

3.2 change cluster-on-pvc.yaml in AWS
storageClassName: gp2

4- create cluster 
kubectl create -f cluster-on-pvc.yaml
kubectl -n rook-ceph get pod
```
NAME                                                 READY   STATUS      RESTARTS   AGE
csi-cephfsplugin-provisioner-d77bb49c6-n5tgs         5/5     Running     0          140s
csi-cephfsplugin-provisioner-d77bb49c6-v9rvn         5/5     Running     0          140s
csi-cephfsplugin-rthrp                               3/3     Running     0          140s
csi-rbdplugin-hbsm7                                  3/3     Running     0          140s
csi-rbdplugin-provisioner-5b5cd64fd-nvk6c            6/6     Running     0          140s
csi-rbdplugin-provisioner-5b5cd64fd-q7bxl            6/6     Running     0          140s
rook-ceph-crashcollector-minikube-5b57b7c5d4-hfldl   1/1     Running     0          105s
rook-ceph-mgr-a-64cd7cdf54-j8b5p                     1/1     Running     0          77s
rook-ceph-mon-a-694bb7987d-fp9w7                     1/1     Running     0          105s
rook-ceph-mon-b-856fdd5cb9-5h2qk                     1/1     Running     0          94s
rook-ceph-mon-c-57545897fc-j576h                     1/1     Running     0          85s
rook-ceph-operator-85f5b946bd-s8grz                  1/1     Running     0          92m
rook-ceph-osd-0-6bb747b6c5-lnvb6                     1/1     Running     0          23s
rook-ceph-osd-1-7f67f9646d-44p7v                     1/1     Running     0          24s
rook-ceph-osd-2-6cd4b776ff-v4d68                     1/1     Running     0          25s
rook-ceph-osd-prepare-node1-vx2rz                    0/2     Completed   0          60s
rook-ceph-osd-prepare-node2-ab3fd                    0/2     Completed   0          60s
rook-ceph-osd-prepare-node3-w4xyz                    0/2     Completed   0          60s
```

5 - choosing de storage
```
For a walkthrough of the three types of storage exposed by Rook, see the guides for:

Block: Create block storage to be consumed by a pod (RWO)
Shared Filesystem: Create a filesystem to be shared across multiple pods (RWX)
Object: Create an object store that is accessible inside or outside the Kubernetes cluster
```
5.1- Create Shared Filesystem
kubectl create -f myfilesystem.yaml

```
# To confirm the filesystem is configured, wait for the mds pods to start
kubectl -n rook-ceph get pod -l app=rook-ceph-mds
```

6- Provision Storage
kubectl create -f deploy/examples/csi/cephfs/storageclass.yaml

7- Consume the Shared Filesystem: hmoura Sample
kubectl create -f hmoura/teste.yaml

X- Teardown 
kubectl delete -f hmoura/teste.yaml
kubectl -n rook-ceph delete cephfilesystem myfs
kubectl delete -n rook-ceph cephblockpool replicapool
kubectl delete storageclass rook-ceph-block
kubectl delete -f csi/cephfs/kube-registry.yaml
kubectl delete storageclass csi-cephfs


CEPH DASH https://rook.io/docs/rook/v1.8/ceph-dashboard.html
kubectl -n rook-ceph get service
kubectl port-forward svc/rook-ceph-mgr-dashboard 7000 -n rook-ceph
Login Credentials
kubectl -n rook-ceph get secret rook-ceph-dashboard-password -o jsonpath="{['data']['password']}" | base64 --decode && echo

Arquitetura