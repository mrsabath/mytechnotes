# Testing Update Commands

## setup creds on silesia using `iris-poc1` cluster:
```
# IRIS POC: KUBECONFIG=/Users/sabath/.fr8r/envs/iris-poc1/shard1/admin/kube-config
export KUBECONFIG=/Users/sabath/.fr8r/envs/fyre-01/radiant01/admin/kube-config
# iris armada
export KUBECONFIG=/Users/sabath/.bluemix/plugins/container-service/clusters/ksquad-dal12-01/kube-config-dal12-ksquad-dal12-01.yml
```

## HELM deploy
```console
# list
helm list --all
# install
helm install --name=update-planner --namespace=default --set docker.tag=devel ./charts/update-planner
helm install --name=update-req --namespace=default ./charts/update-req
helm install --name=update-executor --namespace=default --set docker.tag=devel  --set executor.podCount=2 ./charts/update-executor
helm delete --purge update-executor update-req update-planner
# tests:
helm install --name=synthetic-load --namespace=default --set webService.podCount=15 --set webDeplGroup.podCount=10 ./charts/synthetic-load
```

Testing...
```
# list pods for on all nodes:
kubectl describe nodes | grep Non-terminated
# cleanup after leftovers:
for m in `kubectl get pods | grep web-ms- | awk {'print $1'}`;do kubectl delete pod $m; done
# list inventory status:
kubectl get inventories  -o json | jq '.items[] | ([.metadata.name, .status ])'
```

Testing DNS problems
```
kubectl run curl --image=radial/busyboxplus:curl -i --tty
nslookup kubernetes.default
#re-attach to this container
kubectl attach curl2-449093818-0m2tf -c curl2 -i -t
# delete kube-dns and vpn pods. This should fix the problem
```

# review results of inventories
```
kubectl get inventories  -o json | jq '.items[] | ([.metadata.name, .status ])'
kubectl get inventories  -ojson | jq ".items[].status"
kubectl get inventories  -ojson | jq ".items[].metadata.name"
```

## uncordon all nodes:

```
for N in $(kubectl get node | grep SchedulingDisabled | awk '{ print $1 }');do kubectl uncordon $N; done
```

## run tests:
```
cd /Users/sabath/workspace/alchemy-containers/fr8r-ansible/examples/apps/k8s
cd "/Users/sabath/Box Sync/projects/mytechnotes/examples/k8s"
```

## Handy commands:
```
# resize replicas:
kubectl scale deployment $name --replicas 10
# drain node:
kubectl drain --ignore-daemonsets --force --delete-local-data $nodes
# uncordon_all:
for m in `kubectl get nodes -l update=true | grep SchedulingDisabled | awk {'print $1'}`;do kubectl uncordon $m; done
```

## web server testing:
```
kubectl create -f web-ms-deployment.yaml
deployment "web-ms-deployment" created
service "web-ms-service" created
kubectl scale deployment web-ms-deployment --replicas 10

# list current containers:
command="kubectl get pods -o wide --all-namespaces"
while true; do NOW=$(date +%s); $command; NOW2=$(date +%s);echo $(($NOW2-$NOW)); sleep 2; done

# list current state of the nodes:
command="kubectl get nodes -owide"
while true; do NOW=$(date +%s); $command; NOW2=$(date +%s);echo $(($NOW2-$NOW)); sleep 2; done

# probing the web server:
export command="curl --max-time 1 9.12.238.14:30000"
while true; do NOW=$(date +%s); $command; NOW2=$(date +%s);echo $(($NOW2-$NOW)); sleep 1; done

# evict containers:
kubectl drain --ignore-daemonsets --force --delete-local-data 9.12.238.152
```
## label nodes with update-zones:
```
kubectl get nodes -l failure-domain.beta.kubernetes.io/zone=update-zone-1
kubectl label nodes 10.115.252.14 failure-domain.beta.kubernetes.io/zone=update-zone-1
```

## label nodes for update
```
kubectl get nodes -owide
kubectl label nodes $nodename update=true
# select the node without updates
kubectl label nodes 9.12.238.14 update=false
# get nodes with label:
kubectl get nodes -l update=true
kubectl get nodes -l update=false
```

## redis test:
```
# make sure the node with update=false is ready for deployment, deploy master:
kubectl create -f redis-master.yaml
# cordon
kubectl create -f redis-slave.yaml
kubectl create -f frontend.yaml

# using browser, submit few messages:
curl http://9.12.238.14:30003/

# using redis client connect to slave:
redis-cli -h 9.12.238.14 -p 30002
keys *
get messages

# looped client request:
command="redis-cli -h 9.12.238.14 -p 30001 get messages"
while true; do NOW=$(date +%s); $command; NOW2=$(date +%s);echo $(($NOW2-$NOW)); sleep 1; done

# using redis client connect to redis master:
redis-cli -h 9.12.238.14 -p 30001

# using redis client connect to redis slave:
redis-cli -h 9.12.238.14 -p 30002

# using front end to set data:
# set:
curl "http://9.12.238.14:30003/guestbook.php?cmd=get&key=messages"
{"data": "1,2,3"}
# get:
curl "http://9.12.238.14:30003/guestbook.php?cmd=set&key=messages1&value=1,2,3"

# looping:
# set:
while true; do NOW=$(date +%s); curl "http://9.12.238.14:30003/guestbook.php?cmd=set&key=messages&value=$NOW"; NOW2=$(date +%s);echo $NOW $(($NOW2-$NOW)); sleep 1; done

# get:
command="redis-cli -h 9.12.238.14 -p 30002 get messages"
while true; do NOW=$(date +%s); $command; sleep 1; done

# all-in-one:
 while true; do NOW=$(date +%s); curl "http://9.12.238.14:30003/guestbook.php?cmd=set&key=messages&value=$NOW"; NOW2=$(date +%s);result=`/usr/local/bin/redis-cli -h 9.12.238.14 -p 30002 get messages`;echo $result=$NOW    $(($NOW2-$NOW)); sleep 1; done

# get:
command="redis-cli -h 9.12.238.14 -p 30002 get messages"
while true; do NOW=$(date +%s); $command; sleep 1; done
```

# service example:
https://github.ibm.com/dettori/interconnect-demo/blob/master/nodeApp/kube-manifests/twitter2redis-service.yaml
https://github.ibm.com/dettori/interconnect-demo/blob/master/nodeApp/kube-manifests/twitter2redis-deployment.yaml


# list status of all the NotReady nodes:
for m in `kubectl get nodes | grep NotReady | awk {'print $1'}`; do kubectl get node $m -owide; done

# remove all the NotReady nodes:
for m in `kubectl get nodes | grep NotReady | awk {'print $1'}`; do kubectl delete node $m; done

## testing ETCD

Setting up NFS:
https://www.digitalocean.com/community/tutorials/how-to-set-up-an-nfs-mount-on-ubuntu-16-04

```
cd "/Users/sabath/Box Sync/projects/mytechnotes/examples/k8s"
create -f pvol.yaml
create -f etcd_ss.yaml

# testing:
cd /Users/sabath/workspace/alchemy-containers/fr8r-ansible/examples/apps/k8s
./test_etcd.sh

curl -l http://$HOST:32379/v2/keys
curl -l -X PUT http://$HOST:32379/v2/keys/test -d value="11"
```

## Multiple Update-Zones
[k8s doc](https://kubernetes.io/docs/concepts/configuration/assign-pod-node/)

Good [example](https://github.com/kubernetes/kubernetes/issues/39310)
```
kubectl label nodes 9.12.238.116 failure-domain.beta.kubernetes.io/zone=update-zone-1
kubectl get nodes -l failure-domain.beta.kubernetes.io/zone=update-zone-1
```
