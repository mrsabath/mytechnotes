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
