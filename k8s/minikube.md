# Playing with a minikube

Install [VitualBox](https://www.virtualbox.org/wiki/Downloads) and
[Minikube](https://kubernetes.io/docs/tasks/tools/install-minikube/)

```
$minikube start --vm-driver virtualbox --memory=8192 --cpus=4 --kubernetes-version v1.16.2
😄  minikube v1.5.2 on Darwin 10.14.6
🔥  Creating virtualbox VM (CPUs=4, Memory=8192MB, Disk=20000MB) ...
🐳  Preparing Kubernetes v1.16.2 on Docker '18.09.9' ...
💾  Downloading kubelet v1.16.2
💾  Downloading kubeadm v1.16.2
🚜  Pulling images ...
🚀  Launching Kubernetes ...
⌛  Waiting for: apiserver
🏄  Done! kubectl is now configured to use "minikube"
~$kubectl get nodes
NAME       STATUS   ROLES    AGE     VERSION
minikube   Ready    master   4m12s   v1.16.2
~$minikube status
host: Running
kubelet: Running
apiserver: Running
kubeconfig: Configured
```

Create an alias:
```
alias kk="kubectl -n trusted-identity"
```

## Create a vault service
Content of the `vault.yaml`:
```yaml
kind: Service
apiVersion: v1
metadata:
  name: tsi-vault
spec:
  selector:
    app: tsi-vault
  ports:
  - protocol: TCP
    port: 8200
    targetPort: 8200
  type: NodePort
---

apiVersion: apps/v1
kind: Deployment
metadata:
  labels:
    app: tsi-vault
  name: tsi-vault
spec:
  replicas: 1
  selector:
    matchLabels:
      app: tsi-vault
  template:
    metadata:
      labels:
        app: tsi-vault
      name: tsi-vault
    spec:
      containers:
        - name: tsi-vault
          image: trustedseriviceidentity/ti-vault:v0.3
          imagePullPolicy: Always
```

```
kk create -f vault.yaml
service/tsi-vault created
deployment.apps/tsi-vault created
```
Get the URL of the Vault service

```console
~$minikube service tsi-vault -n trusted-identity --url
http://192.168.99.105:32065
~$curl http://192.168.99.105:32065
<a href="/ui/">Temporary Redirect</a>.
```
This is an expected result at this moment

Assign the Vault address obtained above:

```
export VAULT_ADDR=http://192.168.99.105:32065
helm install charts/ti-key-release-2-1.1.0.tgz --debug --name tsi --set ti-key-release-1.cluster.name=eu-de --set ti-key-release-1.cluster.name=minikube --set ti-key-release-1.vaultAddress=$VAULT_ADDR

```
