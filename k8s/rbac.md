## RBAC - Role Based Access Control
[https://kubernetes.io/docs/admin/authorization/](https://kubernetes.io/docs/admin/authorization/)

```
kubectl create -f clusterRole-read.yaml
kubectl get clusterroles
kubectl get clusterroles pod-writer -o yaml
kubectl delete clusterroles pod-reader
kubectl delete rolebindings 54hRbGJUN7oQlmuhDYD1aQPKWiJIjm8PT5FLPlB10BAYKQTLpod-writer --namespace stest-tenant-default

# verify rbac is enabled
ps -ef | grep rbac

kubectl create -f roleBinding.yaml

kubectl get rolebindings --all-namespaces
kubectl get rolebinding NZfUUePeMFk6iNUGDTmsmmSLhxvFadfm4H7MEe67RR7UMR2Wpod-writer --namespace stest-tenant-default -o yaml

kubectl replace -f clusterRole-write.yaml

vagrant@shard1-master1:~$ cat clusterRole-write.yaml
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1alpha1
metadata:
  # "namespace" omitted since ClusterRoles are not namespaced.
  name: pod-writer
rules:
  - apiGroups: [""]
    resources: ["pods", "replicationcontrollers", "replicasets", "replicasets.extensions"]
    verbs: ["get", "watch", "list", "create", "update", "patch", "delete"]
    nonResourceURLs: []

vagrant@shard1-master1:~$ cat clusterRole-read.yaml
kind: ClusterRole
apiVersion: rbac.authorization.k8s.io/v1alpha1
metadata:
  # "namespace" omitted since ClusterRoles are not namespaced.
  name: pod-reader
rules:
  - apiGroups: [""]
    resources: ["pods"]
    verbs: ["get", "watch", "list"]
    nonResourceURLs: []


vagrant@shard1-master1:~$ cat admin.yaml
kind: ClusterRoleBinding
apiVersion: rbac.authorization.k8s.io/v1alpha1
metadata:
 name: stest5-admin
subjects:
 - kind: User # May be "User", "Group" or "ServiceAccount"
   name: test5
roleRef:
 kind: ClusterRole
 name: admin-role-nonResourceURLSs
 apiGroup: rbac.authorization.k8s.io


cat roleBinding-ms3.yaml
kind: RoleBinding
apiVersion: rbac.authorization.k8s.io/v1alpha1
metadata:
 name: sms3-writer
 namespace: sms3-default # This binding only applies in the "development" namespace
subjects:
 - kind: User # May be "User", "Group" or "ServiceAccount"
   name: UqWLRpD3AHXNEQAkEZRmG3WzK49bof688W2rBYwSAnBwSUOK
roleRef:
 kind: ClusterRole
 name: pod-writer
 apiGroup: rbac.authorization.k8s.io
```

## Samples:

```
# role binding:
 I0209 16:32:42.916455   18246 request.go:559] Request Body: "{\"apiVersion\":\"rbac.authorization.k8s.io/v1alpha1\",\"kind\":\"RoleBinding\",\"metadata\":{\"name\":\"pod-secrets\",\"namespace\":\"sms5-default\"},\"roleRef\":{\"apiGroup\":\"rbac.authorization.k8s.io\",\"kind\":\"ClusterRole\",\"name\":\"pod-reader\"},\"subjects\":[{\"kind\":\"User\",\"name\":\"5kN7Wu4JjXe4k6LepaTete0yg7oV52DDwUJJJHEfnp6ThZlo\"}]}\n"
I0209 16:32:42.916994   18246 round_trippers.go:299] curl -k -v -XPOST  -H "Accept: application/json" -H "Content-Type: application/json" -H "User-Agent: kubectl/v1.5.1 (linux/amd64) kubernetes/82450d0" http://localhost:8080/apis/rbac.authorization.k8s.io/v1alpha1/namespaces/sms5-default/rolebindings
I0209 16:32:42.919293   18246 round_trippers.go:318] POST http://localhost:8080/apis/rbac.authorization.k8s.io/v1alpha1/namespaces/sms5-default/rolebindings 409 Conflict in 2 milliseconds
I0209 16:32:42.919575   18246 round_trippers.go:324] Response Headers:
I0209 16:32:42.919929   18246 round_trippers.go:327]     Content-Type: application/json
I0209 16:32:42.920171   18246 round_trippers.go:327]     Date: Thu, 09 Feb 2017 16:32:42 GMT
I0209 16:32:42.920561   18246 round_trippers.go:327]     Content-Length: 278
I0209 16:32:42.920929   18246 request.go:904] Response Body: {"kind":"Status","apiVersion":"v1","metadata":{},"status":"Failure","message":"rolebindings.rbac.authorization.k8s.io \"pod-secrets\" already exists","reason":"AlreadyExists","details":{"name":"pod-secrets","group":"rbac.authorization.k8s.io","kind":"rolebindings"},"code":409}
```

## list roles:
```
I0209 19:15:51.773933   22990 round_trippers.go:299] curl -k -v -XGET  -H "Accept: application/json" -H "User-Agent: kubectl/v1.5.1 (linux/amd64) kubernetes/82450d0" http://localhost:8080/apis/rbac.authorization.k8s.io/v1alpha1/clusterroles
I0209 19:15:51.776840   22990 round_trippers.go:318] GET http://localhost:8080/apis/rbac.authorization.k8s.io/v1alpha1/clusterroles 200 OK in 2 milliseconds
I0209 19:15:51.776860   22990 round_trippers.go:324] Response Headers:
I0209 19:15:51.776866   22990 round_trippers.go:327]     Content-Type: application/json
I0209 19:15:51.776872   22990 round_trippers.go:327]     Date: Thu, 09 Feb 2017 19:15:51 GMT
I0209 19:15:51.777105   22990 request.go:904] Response Body: {"kind":"ClusterRoleList","apiVersion":"rbac.authorization.k8s.io/v1alpha1","metadata":{"selfLink":"/apis/rbac.authorization.k8s.io/v1alpha1/clusterroles","resourceVersion":"225767"},"items":[{"metadata":{"name":"admin","selfLink":"/apis/rbac.authorization.k8s.io/v1alpha1/clusterrolesadmin","uid":"d9c5e692-e7d3-11e6-9a47-0800276ccb39","resourceVersion":"14","creationTimestamp":"2017-01-31T16:39:35Z"},"rules":[{"verbs":["get","list","watch","create","update","patch","delete","deletecollection"],"attributeRestrictions":null,"apiGroups":[""],"resources":["pods","pods/attach","pods/proxy","pods/exec","pods/portforward"]},{"verbs":["get","list","watch","create","update","patch","delete","deletecollection"],"attributeRestrictions":null,"apiGroups":[""],"resources":["replicationcontrollers","replicationcontrollers/scale","serviceaccounts","services","services/proxy","endpoints","persistentvolumeclaims","configmaps","secrets"]},{"verbs":["get","list","watch"],"attributeRestrictions":null,"apiGroups":[""],"resources":["limitranges","resourcequotas","bindings","events","pods/status","resourcequotas/status","namespaces/status","replicationcontrollers/status","pods/log"]},{"verbs":["get","list","watch"],"attributeRestrictions":null,"apiGroups":[""],"resources":["namespaces"]},{"verbs":["impersonate"],"attributeRestrictions":null,"apiGroups":[""],"resources":["serviceaccounts"]},{"verbs":["get","list","watch","create","update","patch","delete","deletecollection"],"attributeRestrictions":null,"apiGroups":["apps"],"resources":["statefulsets"]},{"verbs":["get","list","watch","....
```

## run installs manually
```
vagrant ssh installer
cd fr8r/ansible
ansible-playbook -v -i ../examples/envs/dev-vbox/shard1.hosts env-basics.yml -e envs=../examples/envs -e env_name=dev-vbox
ansible-playbook -v -i ../examples/envs/dev-vbox/shard1.hosts api-proxy-setup.yml -e envs=../examples/envs -e env_name=dev-vbox
ansible-playbook -v -i ../examples/envs/dev-vbox/shard1.hosts shard.yml -e envs=../examples/envs -e cluster_name=dev-vbox-shard1 -e network_kind=flannel
```

```
#ssh to all the proxies:
export ETCDCTL_API=3
export HOST=192.168.10.4 # 5 VMs
export HOST=192.168.10.4 # all_in_one
/opt/bin/etcdctl --cert=/opt/admin-certs/etcd.pem --key=/opt/admin-certs/etcd-key.pem --cacert=/opt/admin-certs/ca.pem --endpoints=https://$HOST:4011 --keys-only=true --from-key=true get /
```
