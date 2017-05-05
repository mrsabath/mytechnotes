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
 go get k8s.io/client-go/...
 cd examples/
 cd third-party-resources/
 # this failed the first time, but it registered new tpr
 go run *.go -kubeconfig=/Users/sabath/.fr8r/envs/iris-poc1/shard1/admin/kube-config
 # worked the second time
 go run *.go -kubeconfig=/Users/sabath/.fr8r/envs/iris-poc1/shard1/admin/kube-config
 kubectl get all
 kubectl get thirdpartyresources
 kubectl get thirdpartyresources -o json
 kubectl get example.k8s.io
 kubectl get backup.kubeconeu.deis.com
 kubectl get backup.kubeconeu.deis.com -oyaml
 kubectl get example -o json
 kubectl get example -oyaml > example.yml
 vi example.yml
 kubectl replace -f example.yml
 kubectl get example -oyaml
```
