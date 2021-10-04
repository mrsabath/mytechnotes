# Kubernets Secrets

Create secret from a file or files:

```console
# create a file or directory e.g. /tmp/kubeconfigs
kubectl -n tornjak create secret generic kubeconfigs --from-file=/tmp/kubeconfigs
```

Read existing secrets:
```console
# get the secret data structure:
$kubectl -n tornjak get secret kubeconfigs -o jsonpath="{.data}"
kubect -n tornjak get secret kubeconfigs -oyaml

# decode the value e.g. data.space-x01
kubectl -n tornjak get secret kubeconfigs -o jsonpath="{.data.space-x01}" | base64 -D
kubectl -n tornjak get secret kubeconfigs -o jsonpath={.data.space-x01} | base64 -D

kubectl get secrets/db-user-pass --template={{.data.password}} | base64 -D
```
