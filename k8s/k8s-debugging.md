# Debugging K8s

https://thenewstack.io/living-with-kubernetes-debug-clusters-in-8-commands/

A few useful commands for debugging Kubernetes:

```console
kubectl version --short
kubectl cluster-info
kubectl get componentstatus
kubectl api-resources -o wide --sort-by name
kubectl get events -A
kubectl get nodes -o wide
kubectl get pods -A -o wide
kubectl run a --image alpine --command -- /bin/sleep 1d
```

## kubectl cluster-info
To get all the information about the cluster:
```console
kubectl cluster-info dump
```

Your favorite `kubectl` debugging/troubleshooting commands:
https://twitter.com/I_saloni92/status/1513710771087155204/photo/1
Main track: https://twitter.com/saiyampathak/status/1513572111721271298

busybox for telnet, service connection, nslookup etc.
```console
kubectl run --generator=run-pod/v1 -i --tty busybox --image=radial/busyboxplus:curl --restart=Never -- sh
```
Also as an ephemeral container:
```console
kubectl debug -it demo --image=busybox --share-processes --copy-to=demo-debug
```
How many things do I have in etcd? is that number growing crazily?
```console
kubectl get --raw /metrics | grep apiserver_storage_objects
```
and
Is etcd being slow?
```
kubectl get --raw /metrics | grep etcd.*.sum
```

## ERROR:
```
{kubelet 9.47.168.134}            Warning        FailedSync    Error syncing pod, skipping: failed to "SetupNetwork" for "web-ms-deployment-1563362432-zg245_default" with SetupNetworkError: "Failed to setup network for pod \"web-ms-deployment-1563362432-zg245_default(caef25c2-94d6-11e7-90ae-52548b38cc72)\" using network plugins \"cni\": client: etcd cluster is unavailable or misconfigured; error #0: dial tcp 127.0.0.1:2379: getsockopt: connection refused\n; Skipping pod"
```
## check etcd on master:
```
curl -k --cacert /etc/kubernetes/cert/ca.pem --cert /etc/kubernetes/cert/admin.pem --key  /etc/kubernetes/cert/admin-key.pem  https://$MASTERIP:4001/v2/stats/self
curl -k --cacert /etc/kubernetes/cert/ca.pem --cert /etc/kubernetes/cert/admin.pem --key  /etc/kubernetes/cert/admin-key.pem  https://$MASTERIP:4001/v2/keys
service etcd restart
```

## make sure calico config points at master etcd:
```
vi /etc/cni/net.d/10-calico.conf
"etcd_endpoints": "https://9.47.168.128:4001",
# service kubelet restart

vi /usr/lib/systemd/system/calico-node.service
-e ETCD_ENDPOINTS=https://9.47.168.128:4001 \

vi /etc/calico/calicoctl.cfg
etcdEndpoints: https://9.47.168.128:4001

service kubelet restart
systemctl daemon-reload
service calico-node restart
```
