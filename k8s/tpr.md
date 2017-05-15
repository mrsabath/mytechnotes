# Using Third Party Resources (TPR)
## update brew, install [helm](https://github.com/kubernetes/helm/blob/master/docs/quickstart.md)

```console
sudo chown -R $(whoami):admin /usr/local
brew update
brew install kubernetes-helm

# setup KUBECONFIG for my k8s cluster:
export KUBECONFIG=/Users/sabath/.fr8r/envs/iris-poc1/shard1/admin/kube-config
# initialize helm, install Tiller
helm init
helm repo update
```

- [Demo Video](https://www.youtube.com/watch?v=qiB4RxCDC8o)
- [Source code](https://github.com/arschles/2017-KubeCon-EU)

```console
git clone https://github.com/arschles/2017-KubeCon-EU.git
helm install --name=kubecon-watcher --namespace=default --set docker.tag=c0ffcab \
 ./charts/watcher
helm install --name=kubecon-backup --namespace=default ./charts/backup
```

## using the `client-go` example

Follow [install](https://github.com/kubernetes/client-go/blob/master/INSTALL.md)
and then [example](https://github.com/kubernetes/client-go/tree/master/examples/third-party-resources)

Steps:
```console
 cd /Users/sabath/workspace/k8s/client-go/examples/third-party-resources
 go get k8s.io/client-go/...
 cd examples/
 cd third-party-resources/
 # this failed the first time, but it registered new tpr
 go run *.go -kubeconfig=/Users/sabath/.fr8r/envs/iris-poc1/shard1/admin/kube-config
 # worked the second time
 go run *.go -kubeconfig=/Users/sabath/.fr8r/envs/iris-poc1/shard1/admin/kube-config

 # setup env.
 export KUBECONFIG=/Users/sabath/.fr8r/envs/iris-poc1/shard1/admin/kube-config
 kubectl get all
 kubectl get thirdpartyresources
 kubectl get thirdpartyresources -o json
 kubectl get example.k8s.io
 kubectl get backup.kubeconeu.deis.com
 kubectl get backup.kubeconeu.deis.com -oyaml
 kubectl get example -o json
 kubectl get example -oyaml > example2.yml
 vi example.yml
 kubectl replace -f example.yml
 kubectl get example -oyaml
```

New `Update` code:
```console
 cd /Users/sabath/workspace/k8s/client-go/examples/ms
 kubectl create -f update_tpr.yaml
 go run *.go -kubeconfig=/Users/sabath/.fr8r/envs/iris-poc1/shard1/admin/kube-config

 # setup env.
 export KUBECONFIG=/Users/sabath/.fr8r/envs/iris-poc1/shard1/admin/kube-config
 kubectl get update -oyaml
 kubectl get update update2 -oyaml > update2.yaml
 mv update2.yaml update3.yaml
 vi update3.yaml
 kubectl create -f update3.yaml
 update "update3" created
 kubectl get update
 NAME      KIND
 update1   Update.v1.k8s.io
 update2   Update.v1.k8s.io
 update3   Update.v1.k8s.io
 kubectl get update update3 -oyaml

```
Newer `Cluster Update Service` using KUBECON Demo:

```Console
/Users/sabath/workspace/go-work/src/github.com/arschles/2017-KubeCon-EU
#/Users/sabath/workspace/github.com/2017-KubeCon-EU
export KUBECONFIG=/Users/sabath/.fr8r/envs/iris-poc1/shard1/admin/kube-config
# register the new TPR
helm install --name=cluster-update --namespace=default ./charts/ClusterUpdateReq
# remove it
helm delete --purge cluster-update

```

Even more newer `Cluster Update Service` using KUBECON Demo:

```Console
# compile, using GOPATH=/Users/sabath/workspace/go-work
cd /Users/sabath/workspace/go-work/src/github.ibm.com/sabath/update-service
make build # to compile
make all   # to push the image

# deploy
export KUBECONFIG=/Users/sabath/.fr8r/envs/iris-poc1/shard1/admin/kube-config

# register the new TPR
helm install --name=cluster-update --namespace=default ./charts/ClusterUpdateReq
kubectl get thirdpartyresources clusterupdatereq.ibm.com -oyaml
# remove it
helm delete --purge cluster-update

# deploy planner
helm install --name=planner --namespace=default ./charts/Planner
helm delete --purge planner

```

Final steps for `Cluster Update Service` using KUBECON Demo:
```Console
# compile, using GOPATH=/Users/sabath/workspace/go-work
cd /Users/sabath/workspace/go-work/src/github.ibm.com/sabath/update-service
make build # to compile
make all   # to push the image

# deploy
export KUBECONFIG=/Users/sabath/.fr8r/envs/iris-poc1/shard1/admin/kube-config

# register the new TPR and planner
helm install --name=update-planner --namespace=default --set docker.tag=cd24d17 ./charts/update-planner
kubectl get thirdpartyresources update.ibm.com -oyaml
kubectl get update -oyaml
# remove it
helm delete --purge update-planner

# request the update
helm install --name=update-req --namespace=default ./charts/update-req
helm delete --purge update-req

helm ls --all
helm delete --purge update-req update-planner


```
