# check etcd on master:
```
curl -k --cacert /etc/kubernetes/cert/ca.pem --cert /etc/kubernetes/cert/admin.pem --key  /etc/kubernetes/cert/admin-key.pem  https://$MASTERIP:4001/v2/stats/self
curl -k --cacert /etc/kubernetes/cert/ca.pem --cert /etc/kubernetes/cert/admin.pem --key  /etc/kubernetes/cert/admin-key.pem  https://$MASTERIP:4001/v2/keys
service etcd restart
```

# make sure calico config points at master etcd:
```
cat /etc/cni/net.d/10-calico.conf
service kubelet restart

vi /usr/lib/systemd/system/calico-node.service:
-e ETCD_ENDPOINTS="https://9.47.168.128:4001"
vi /etc/calico/calicoctl.cfg:
etcdEndpoints: https://9.47.168.128:4001
systemctl daemon-reload
service calico-node restart
```
