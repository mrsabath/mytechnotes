# ERROR:
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
