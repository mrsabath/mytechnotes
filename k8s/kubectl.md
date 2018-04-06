# Various K8s and kubectl Commands

## install kubectl

install kubectl on host using Docker:
```
docker run -e LICENSE=accept --net=host -v /usr/local/bin:/data ibmcom/kubernetes:v1.7.3 cp /kubectl /data
```

[other instructions](https://kubernetes.io/docs/user-guide/prereqs/)

```
brew install kubectl
# or
curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/darwin/amd64/kubectl

chmod +x kubectl
ln -s kubectl /usr/local/bin/kubectl
/usr/local/bin/kubectl -> /Users/sabath/workspace/kubectl/kubectl

# set kubeconfig:
export KUBECONFIG="/home/vagrant/.fr8r/envs/dev-vbox/shard1/user1/kube-config"

# or on every call:
kubectl --kubeconfig=/home/vagrant/.fr8r/envs/dev-vbox/shard1/user1/kube-config create -f pod-web.yaml  --validate=false
pod "kube-web-ms" created
vagrant@client:~$ kubectl --kubeconfig=/home/vagrant/.fr8r/envs/dev-vbox/shard1/user1/kube-config get pods
NAME          READY     STATUS    RESTARTS   AGE
kube-web-ms   1/1       Running   0          <invalid>
vagrant@client:~$ kubectl --kubeconfig=/home/vagrant/first-user/kube-config get pods
No resources found.
```
## Simple kubectl test:
```console
# '--' separates the command attributes from kubectl attributes
kubectl run hello --image=busybox  -- /bin/sh -c "while true; do sleep 10; echo test; done;"
kubectl logs -f $(kubectl get po --selector=run=hello --output=jsonpath={.items..metadata.name})
```

## various useful combinations:
```console
# show logs for specific pod:
kubectl logs $(kubectl get pods  --show-all --selector=job-name=pi --output=jsonpath={.items..metadata.name})
# show logs for first pod with this name:
kubectl logs $(kubectl get pods --selector=app=att-client --output=jsonpath={.items..metadata.name} | awk -F ' ' '{print $1}' | sed -n 1p)
# or
kubectl logs -f $(kubectl get pods | grep att-client | awk -F ' ' '{print $1}' | sed -n 1p)
```

UI:
```
kubectl proxy
Starting to serve on 127.0.0.1:8001
```

## create alias
```console
alias k='kubectl --kubeconfig="/Users/sabath/.fr8r/envs/iris-poc1/shard1/admin/kube-config"'
k get po
```

## port forward
Forward the local port on the client 9000 (e.g. mac) to the service port on pod (9001)
```
$k port-forward minio-cc6d4ffdf-4qqww  -n heptio-ark-server 9000:9001 &
```

## testing access to k8s API:
```
curl --key /etc/kubernetes/cert/admin-key.pem  --cert /etc/kubernetes/cert/admin.pem  --cacert /etc/kubernetes/cert/ca.pem -v -XGET  -H "Accept: application/json" -H "User-Agent: kubectl/v1.5.1 (linux/amd64) kubernetes/82450d0" https://9.12.235.3:8443/api
```

## using curl on mac with admin
```console
export DIR=/Users/sabath/.fr8r/envs/iris-poc1/shard1/admin
export HOST=9.59.149.16:443
alias curl='/usr/local/Cellar/curl/7.46.0/bin/curl'

# api
curl --key ${DIR}/admin-key.pem --cert ${DIR}/admin.pem  --cacert ${DIR}/ca.pem \
-v -XGET  -H "Accept: application/json" -H "User-Agent: kubectl/v1.5.1 (linux/amd64) kubernetes/82450d0" https://${HOST}/api

# apis
curl --key ${DIR}/admin-key.pem --cert ${DIR}/admin.pem  --cacert ${DIR}/ca.pem \
-v -XGET  -H "Accept: application/json" -H "User-Agent: kubectl/v1.5.1 (linux/amd64) kubernetes/82450d0" https://${HOST}/apis

# swagger
curl --key ${DIR}/admin-key.pem --cert ${DIR}/admin.pem  --cacert ${DIR}/ca.pem \
-v -XGET  -H "Accept: application/json" -H "User-Agent: kubectl/v1.5.1 (linux/amd64) kubernetes/82450d0" https://${HOST}/swagger
```

## debug:
```
kubectl --v=10 get pods
watch kubectl get pods
```

## simple run command
```
kubectl run update-deployer --image fr8r/update-deployer:devel --command -- sleep 50000
```

## uncordon all nodes:
```
for I in $(kubectl get node | grep SchedulingDisabled | awk '{ print $1 }');do kubectl uncordon $I; done
```

## setup kubectl:
```
./kubectl.sh config set-cluster local --server=http://127.0.0.1:8080 --insecure-skip-tls-verify=true
./kubectl.sh config set-context local --cluster=local
./kubectl.sh config use-context local
./kubectl.sh --help
```

## execute kubectl
```
./kubectl.sh run cont1 --image=busybox --command -- httpd -f -p 3000
./kubectl.sh get pods
./kubectl.sh get rc
get nodes
```

# create pod at specific namespace
```
cat my-namespace.yaml:
kind: Namespace
apiVersion: v1
metadata:
  name: my-namespace
  labels:
    name: my-namespace

./kubectl.sh create -f my-namespace.yaml
./kubectl.sh create namespace <name>`
./kubectl.sh get namespaces
./kubectl.sh run cont6 --image=mrsabath/web-ms --port=80 --namespace=default --command -- httpd -f -p 3000
```

## deployments (rc):
```
kubectl run k2 --image=busybox sleep 864000
kubectl scale  deployment --replicas=3 k2
kubectl get deployment
kubectl get pods
kubectl delete pods --all
```

## custom columns
```
watch kubectl get in,up -o=custom-columns=NAME:.metadata.name,STATUS:.status.state
```

## replicationcontroller
```
kubectl get replicationcontrollers
kubectl get rc
kubectl delete rc <name>

./kubectl.sh run group1 --image=mrsabath/web-ms --port=80 --namespace=ms-namespace --replicas=1 --env="test1=1" --command --
./kubectl.sh run group1 --image=mrsabath/web-ms --port=80 --namespace=default --replicas=1 --command --
./kubectl.sh run group1 --image=mrsabath/web-ms --port=80 --namespace=default --replicas=1 --env="test1=1" --env="test2=2" --command --
./kubectl.sh get  rc --namespace=my-namespace
./kubectl.sh scale rc --replicas=3 --namespace=my-namespace group1
./kubectl.sh delete rc group1 --namespace=my-namespace
./kubectl.sh delete rc --all --namespace=my-namespace

./kubectl run group37 --image=10.140.132.215:5001/mrsabath/web-ms --port=80 --namespace=ms-namespace --replicas=1 --requests="memory=128Mi" --env="test1=1" --env="test2=2" --command --

# list repl. controllers
root@default:/go/src/github.com/kubernetes/kubernetes# ./kubectl.sh get  rc
CONTROLLER   CONTAINER(S)   IMAGE(S)   SELECTOR    REPLICAS   AGE
cont1        cont1          busybox    run=cont1   1          23m
cont2        cont2          busybox    run=cont2   1          23m

# list a specific one:
root@default:/go/src/github.com/kubernetes/kubernetes# ./kubectl.sh get  rc -l run=cont1
CONTROLLER   CONTAINER(S)   IMAGE(S)   SELECTOR    REPLICAS   AGE
cont1        cont1          busybox    run=cont1   1          24m

# scale up
./kubectl.sh scale --replicas=3 rc cont1
./kubectl.sh scale --namespace=my-namespace --replicas=3 rc cont1
replicationcontroller "cont1" scaled
root@default:/go/src/github.com/kubernetes/kubernetes# ./kubectl.sh get  rc -l run=cont1
CONTROLLER   CONTAINER(S)   IMAGE(S)   SELECTOR    REPLICAS   AGE
cont1        cont1          busybox    run=cont1   3          25m
# scale down
root@default:/go/src/github.com/kubernetes/kubernetes# ./kubectl.sh scale --replicas=0 rc cont1
replicationcontroller "cont1" scaled
root@default:/go/src/github.com/kubernetes/kubernetes# ./kubectl.sh get  rc -l run=cont1
CONTROLLER   CONTAINER(S)   IMAGE(S)   SELECTOR    REPLICAS   AGE
cont1        cont1          busybox    run=cont1   0          26m
root@default:/go/src/github.com/kubernetes/kubernetes#

# resize
root@default:/go/src/github.com/kubernetes/kubernetes# ./kubectl.sh scale --replicas=2 rc cont1
replicationcontroller "cont1" scaled

# delete
./kubectl.sh delete rc group1
./kubectl.sh delete rc --all
./kubectl.sh delete rc --all --namespace=default
replicationcontroller "cont1" deleted

./kubectl.sh  describe pod
./kubectl exec -it group1-dgbf9 /bin/bash
```

## HTTP representation of the commands (using proxy)
```
curl -XPOST -H "X-Tls-Client-Dn: /CN=$user" -H "Content-Type: application/json"   localhost:8087/kubeinit
curl -XGET -H "X-Tls-Client-Dn: /CN=$user" -H "Content-Type: application/json"   localhost:8087/api
# does not work
curl -XGET -H "X-Tls-Client-Dn: /CN=$user" -H "Content-Type: application/json"   localhost:8087/apis
curl -XGET -H "X-Tls-Client-Dn: /CN=$user" -H "Content-Type: application/json"   localhost:8087/api/v1/namespaces/default/pods
curl -XPOST -H "X-Tls-Client-Dn: /CN=$user" -H "Content-Type: application/json" -H "Accept: application/json, */*" -H "User-Agent: kubectl/v1.2.0 (linux/amd64) kubernetes/d800dca" -d '{"kind":"Deployment","apiVersion":"extensions/v1beta1","metadata":{"name":"test2","creationTimestamp":null,"labels":{"run":"test2"}},"spec":{"replicas":2,"selector":{"matchLabels":{"run":"test2"}},"template":{"metadata":{"creationTimestamp":null,"labels":{"run":"test2"}},"spec":{"containers":[{"name":"test2","image":"mrsabath/web-ms","resources":{"requests":{"memory":"128Mi"}}}]}},"strategy":{}},"status":{}}' localhost:8087/apis/extensions/v1beta1/namespaces/default/deployments
```
# HTTP representation of the commands (directly)
```
export DOCKER_CERT_PATH=/Users/sabath/.ice/certs/containers-api-dev.stage1.ng.bluemix.net/7ce53472-e001-40df-b06b-03cc5e5967a0
export KUBEHOST=https://10.140.146.5:443    # "dev-mon-kube02"
export KUBEHOST=https://10.140.171.205:443  # "dev-mon01-radiant11"
```
## problem with curl on mac, needs cert conversion:
```
openssl pkcs12 -export -inkey $DOCKER_CERT_PATH/key.pem -in $DOCKER_CERT_PATH/cert.pem -name test-curl-client-side -out test-curl.p12 -password pass:mysecret
curl -k --cert test-curl.p12:mysecret $KUBEHOST/api
curl -k --cert test-curl.p12:mysecret $KUBEHOST/apis
# create a pod:
Request Body: {"kind":"Pod","apiVersion":"v1","metadata":{"name":"redis-django","namespace":"default","creationTimestamp":null,"labels":{"app":"web"}},"spec":{"containers":[{"name":"key-value-store","image":"redis","ports":[{"containerPort":6379,"protocol":"TCP"}],"resources":{},"terminationMessagePath":"/dev/termination-log","imagePullPolicy":"Always"},{"name":"frontend","image":"django","ports":[{"containerPort":8000,"protocol":"TCP"}],"resources":{},"terminationMessagePath":"/dev/termination-log","imagePullPolicy":"Always"}],"restartPolicy":"Always","terminationGracePeriodSeconds":30,"dnsPolicy":"ClusterFirst","securityContext":{}},"status":{}}
I0610 02:57:06.678099   31723 round_trippers.go:299] curl -k -v -XPOST  -H "Accept: application/json, */*" -H "Content-Type: application/json" -H "User-Agent: kubectl/v0.0.0 (linux/amd64) kubernetes/fbc5424" https://10.140.146.5:443/api/v1/namespaces/default/pods
I0610 02:57:06.688726   31723 round_trippers.go:318] POST https://10.140.146.5:443/api/v1/namespaces/default/pods 201 Created in 10 milliseconds

## this works in non-mac:
curl -XGET -H "Accept: application/json, */*" -H "User-Agent: kubectl/v1.2.0 (linux/amd64) kubernetes/d800dca" -key $DOCKER_CERT_PATH/key.pem -cacert $DOCKER_CERT_PATH/cert.pem  $KUBEHOST/api

curl -k -vvv -XGET  -H "Accept: application/json, */*" -H "User-Agent: kubectl/v1.2.0 (linux/amd64) kubernetes/d800dca"  -key /Users/sabath/.ice/certs/containers-api-dev.stage1.ng.bluemix.net/7ce53472-e001-40df-b06b-03cc5e5967a0/key.pem -cacert /Users/sabath/.ice/certs/containers-api-dev.stage1.ng.bluemix.net/7ce53472-e001-40df-b06b-03cc5e5967a0/ca.pem  -cert /Users/sabath/.ice/certs/containers-api-dev.stage1.ng.bluemix.net/7ce53472-e001-40df-b06b-03cc5e5967a0/cert.pem  $KUBEHOST/api

curl -k -vvv -XGET  -H "Accept: application/json, */*" -H "User-Agent: kubectl/v1.2.0 (linux/amd64) kubernetes/d800dca"  -key key.pem -cacert ca.pem  -cert cert.pem   $KUBEHOST/api

```

## ussing curl
```
curl --version
curl 7.43.0 (x86_64-apple-darwin15.0) libcurl/7.43.0 SecureTransport zlib/1.2.5

curl --cert cert.pem --key key.pem --cacert ca.pem -XGET -H "X-Tls-Client-Dn: /CN=test1" https://localhost:8087/api

# works on ubuntu:
curl  --key /etc/kubernetes/cert/admin-key.pem --cert /etc/kubernetes/cert/admin.pem --cacert /etc/kubernetes/cert/ca.pem   -k -v -XGET  -H "Accept: application/json, */*" -H "User-Agent: kubectl/v0.0.0 (linux/amd64) kubernetes/fbc5424" https://192.168.10.2:443/api

# works on mac:
/usr/local/Cellar/curl/7.46.0/bin/curl --key key.pem --cert cert.pem --cacert ca.pem -k -v -XGET  -H "Accept: application/json, */*" -H "User-Agent: kubectl/v1.3.0 (darwin/amd64) kubernetes/df66b90" https://192.168.10.2:443/apis

/usr/local/Cellar/curl/7.46.0/bin/curl --key key.pem  --cert cert.pem  --cacert ca.pem  -k -v -XGET  -H "Accept: application/json, */*" -H "User-Agent: kubectl/v1.3.0 (darwin/amd64) kubernetes/df66b90" https://localhost:8087/apis

## fixed by:
brew rm curl && brew install curl --with-openssl
https://coderwall.com/p/h3zzrw/using-client-ssl-certificates-for-php-curl-requests-on-osx
```


## Events:
```
kubectl get ev -w --watch-only=true | grep Pod
```
