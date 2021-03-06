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
## Adding new TPR
Find the proper URL using curl:
```console
export DIR=/Users/sabath/.fr8r/envs/iris-poc1/shard1/admin
export HOST=9.59.149.16:443
alias curl='/usr/local/Cellar/curl/7.46.0/bin/curl'

# here:
curl --key ${DIR}/admin-key.pem --cert ${DIR}/admin.pem  --cacert ${DIR}/ca.pem -v -XGET  -H "Accept: application/json" -H "User-Agent: kubectl/v1.5.1 (linux/amd64) kubernetes/82450d0" https://${HOST}/apis/ibm.com/v1alpha1 | jq

# get more:
curl --key ${DIR}/admin-key.pem --cert ${DIR}/admin.pem  --cacert ${DIR}/ca.pem -v -XGET  -H "Accept: application/json" -H "User-Agent: kubectl/v1.5.1 (linux/amd64) kubernetes/82450d0" https://${HOST}/apis/ibm.com/v1alpha1/inventories | jq

curl --key ${DIR}/admin-key.pem --cert ${DIR}/admin.pem  --cacert ${DIR}/ca.pem -v -XGET  -H "Accept: application/json" -H "User-Agent: kubectl/v1.5.1 (linux/amd64) kubernetes/82450d0" https://${HOST}/apis/extensions/v1beta1/thirdpartyresources | jq

curl --key ${DIR}/admin-key.pem --cert ${DIR}/admin.pem  --cacert ${DIR}/ca.pem -v -XGET  -H "Accept: application/json" -H "User-Agent: kubectl/v1.5.1 (linux/amd64) kubernetes/82450d0" https://${HOST}/apis/extensions/v1beta1/thirdpartyresourcesupdate.ibm.com
```

## testing the planner

### cordoning the node
```
# request body:
"spec":{"externalID":"9.12.238.116","unschedulable":true}
 curl -k -v -XPUT  -H "Accept: application/json, */*" -H "Content-Type: application/json" -H "User-Agent: kubectl/v1.6.2 (darwin/amd64) kubernetes/477efc3" https://9.59.149.16:443/api/v1/nodes/9.12.238.116
 ```

### draining a nodes
```

I0518 13:24:26.650482   51317 request.go:991] Request Body: {"kind":"Eviction","apiVersion":"policy/v1beta1","metadata":{"name":"redis-master-2683564498-vhhpv","namespace":"straore-default","creationTimestamp":null},"deleteOptions":{}}
I0518 13:24:26.650491   51317 request.go:991] Request Body: {"kind":"Eviction","apiVersion":"policy/v1beta1","metadata":{"name":"frontend-1078725745-w9rnf","namespace":"straore-default","creationTimestamp":null},"deleteOptions":{}}
I0518 13:24:26.650496   51317 request.go:991] Request Body: {"kind":"Eviction","apiVersion":"policy/v1beta1","metadata":{"name":"db-1391787668-zm8xj","namespace":"straore-default","creationTimestamp":null},"deleteOptions":{}}
I0518 13:24:26.650503   51317 request.go:991] Request Body: {"kind":"Eviction","apiVersion":"policy/v1beta1","metadata":{"name":"kube-dns-v20-qf0fz","namespace":"kube-system","creationTimestamp":null},"deleteOptions":{}}
I0518 13:24:26.650498   51317 request.go:991] Request Body: {"kind":"Eviction","apiVersion":"policy/v1beta1","metadata":{"name":"frontend-2415135947-s7bj5","namespace":"sbwolfe-default","creationTimestamp":null},"deleteOptions":{}}
I0518 13:24:26.650534   51317 round_trippers.go:398] curl -k -v -XPOST  -H "Accept: application/json, */*" -H "Content-Type: application/json" -H "User-Agent: kubectl/v1.6.2 (darwin/amd64) kubernetes/477efc3" https://9.59.149.16:443/api/v1/namespaces/kube-system/pods/kube-dns-v20-qf0fz/eviction
I0518 13:24:26.650533   51317 round_trippers.go:398] curl -k -v -XPOST  -H "User-Agent: kubectl/v1.6.2 (darwin/amd64) kubernetes/477efc3" -H "Content-Type: application/json" -H "Accept: application/json, */*" https://9.59.149.16:443/api/v1/namespaces/straore-default/pods/redis-master-2683564498-vhhpv/eviction
I0518 13:24:26.650551   51317 round_trippers.go:398] curl -k -v -XPOST  -H "Accept: application/json, */*" -H "Content-Type: application/json" -H "User-Agent: kubectl/v1.6.2 (darwin/amd64) kubernetes/477efc3" https://9.59.149.16:443/api/v1/namespaces/straore-default/pods/frontend-1078725745-w9rnf/eviction
I0518 13:24:26.650572   51317 round_trippers.go:398] curl -k -v -XPOST  -H "Accept: application/json, */*" -H "Content-Type: application/json" -H "User-Agent: kubectl/v1.6.2 (darwin/amd64) kubernetes/477efc3" https://9.59.149.16:443/api/v1/namespaces/sbwolfe-default/pods/frontend-2415135947-s7bj5/eviction
I0518 13:24:26.650919   51317 round_trippers.go:398] curl -k -v -XPOST  -H "User-Agent: kubectl/v1.6.2 (darwin/amd64) kubernetes/477efc3" -H "Content-Type: application/json" -H "Accept: application/json, */*" https://9.59.149.16:443/api/v1/namespaces/straore-default/pods/db-1391787668-zm8xj/eviction
I0518 13:24:26.804901   51317 round_trippers.go:417] POST https://9.59.149.16:443/api/v1/namespaces/straore-default/pods/redis-master-2683564498-vhhpv/eviction 201 Created in 154 milliseconds
I0518 13:24:26.804922   51317 round_trippers.go:423] Response Headers:
I0518 13:24:26.804929   51317 round_trippers.go:426]     Date: Thu, 18 May 2017 17:24:16 GMT
I0518 13:24:26.804931   51317 round_trippers.go:426]     Content-Type: application/json
I0518 13:24:26.804934   51317 round_trippers.go:426]     Content-Length: 80
I0518 13:24:26.805007   51317 request.go:991] Response Body: {"kind":"Status","apiVersion":"v1","metadata":{},"status":"Success","code":201}
I0518 13:24:26.805437   51317 round_trippers.go:398] curl -k -v -XGET  -H "Accept: application/json, */*" -H "User-Agent: kubectl/v1.6.2 (darwin/amd64) kubernetes/477efc3" https://9.59.149.16:443/api/v1/namespaces/straore-default/pods/redis-master-2683564498-vhhpv


I0518 13:24:26.142366   51317 round_trippers.go:417] GET https://9.59.149.16:443/api/v1/nodes/9.12.238.118 200 OK in 156 milliseconds
I0518 13:24:26.142389   51317 round_trippers.go:423] Response Headers:
I0518 13:24:26.142393   51317 round_trippers.go:426]     Content-Type: application/json
I0518 13:24:26.142397   51317 round_trippers.go:426]     Date: Thu, 18 May 2017 17:24:16 GMT
I0518 13:24:26.144119   51317 request.go:991] Response Body: {"kind":"Node","apiVersion":"v1","metadata":{"name":"9.12.238.118","selfLink":"/api/v1/nodes9.12.238.118","uid":"e1338fcb-2463-11e7-a53b-06



I0518 13:24:26.148221   51317 round_trippers.go:398] curl -k -v -XGET  -H "Accept: application/json, */*" -H "User-Agent: kubectl/v1.6.2 (darwin/amd64) kubernetes/477efc3" https://9.59.149.16:443/api/v1/pods?fieldSelector=spec.nodeName%3D9.12.238.118
I0518 13:24:26.179808   51317 round_trippers.go:417] GET https://9.59.149.16:443/api/v1/pods?fieldSelector=spec.nodeName%3D9.12.238.118 200 OK in 31 milliseconds
I0518 13:24:26.179829   51317 round_trippers.go:423] Response Headers:
I0518 13:24:26.179833   51317 round_trippers.go:426]     Content-Type: application/json
I0518 13:24:26.179837   51317 round_trippers.go:426]     Date: Thu, 18 May 2017 17:24:16 GMT
I0518 13:24:26.217886   51317 request.go:991] Response Body: {"kind":"PodList","apiVersion":"v1","metadata":{"selfLink":"/api/v1/pods","resourceVersion":"8828233"},"items":[{"metadata":{"name":"crawle
```
