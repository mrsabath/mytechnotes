## Remote ABAC (Attribute-Based Access Control)
[ABAC](https://kubernetes.io/docs/admin/authorization/)

```
root@api-proxy:~/.fr8r/envs/dev-vbox/admin-certs# pwd
/root/.fr8r/envs/dev-vbox/admin-certs
root@api-proxy:~/.fr8r/envs/dev-vbox/admin-certs# output=$(curl -s --cert /root/.fr8r/envs/dev-vbox/admin-certs/admin.pem --key /root/.fr8r/envs/dev-vbox/admin-certs/admin-key.pem  --cacert /root/.fr8r/envs/dev-vbox/admin-certs/ca.pem https://192.168.10.2:8888/user ); echo $output | wc -c

echo $output | wc -c

# on master:
# access to etcd:
export HOST=192.168.10.2

# bootstrap etcd with remoteabac:
curl -L -XPUT --data-urlencode value@/etc/kubernetes/bootstrap-abac-policy --cacert /etc/kubernetes/cert/ca.pem --key /etc/kubernetes/cert/remoteabac-key.pem --cert /etc/kubernetes/cert/remoteabac.pem https://$HOST:4001/v2/keys/abac-policy

# via curl: (shows nothing)
curl --cacert /etc/kubernetes/cert/ca.pem --key /etc/kubernetes/cert/remoteabac-key.pem --cert /etc/kubernetes/cert/remoteabac.pem https://$HOST:4001/abac-policy

curl --cacert /etc/kubernetes/cert/ca.pem --key /etc/kubernetes/cert/remoteabac-key.pem --cert /etc/kubernetes/cert/ca.pem https://$HOST:4001/v2/keys/abac-policy

curl -s --cert /etc/kubernetes/cert/admin.pem --key /etc/kubernetes/cert/admin-key.pem  --cacert /etc/kubernetes/cert/ca.pem https://192.168.10.2:8888/user

# via etcdctl: (shows nothing)
export ETCDCTL_API=3
/opt/bin/etcdctl --cert=/etc/kubernetes/cert/remoteabac.pem --key=/etc/kubernetes/cert/remoteabac-key.pem --cacert=/etc/kubernetes/cert/ca.pem --endpoints=$HOST:4001 --keys-only=true --from-key=true get /


root@api-proxy:~# curl -s --cert /root/.fr8r/envs/dev-vbox/admin-certs/admin.pem --key /root/.fr8r/envs/dev-vbox/admin-certs/admin-key.pem  --cacert /root/.fr8r/envs/dev-vbox/admin-certs/ca.pem https://192.168.10.2:8888/user
{"status":"OK"}
```

## test script:
```
root@api-proxy:~# cat prasad.sh
#!/bin/bash
exec > /tmp/output.log 2>&1
#set -x
NUMBER_OFFERING=1

for i in `seq 1 $NUMBER_OFFERING`;
do
tenantid="preddy$i"
var=`./create_user.sh  dev-vbox  $tenantid shard1 192.168.10.2 192.168.10.4:6969 CEcKKSLLIXpupUz2ALPgMwd6ZnymfkSRHddrYKbIthUQ7eMd`
if echo $var | grep -q 'ERROR'; then
 echo "Error reported for $tenantid"
else
 echo "Success for $tenantid"
fi
echo $var

output=$(curl -s --cert /root/.fr8r/envs/dev-vbox/admin-certs/admin.pem --key /root/.fr8r/envs/dev-vbox/admin-certs/admin-key.pem  --cacert /root/.fr8r/envs/dev-vbox/admin-certs/ca.pem https://192.168.10.2:8888/user )
echo $output | wc -c
done
root@api-proxy:~#
```


on master:
```
service etcd stop
# rm /etc/init/etcd-proxy.conf
rm -R /var/etcd-proxy/
#rm /var/log/upstart/etcd-proxy.log
service etcd start
```
