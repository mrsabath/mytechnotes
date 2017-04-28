## to restart the service, login to both master and worker
```
# master:
service kubelet status
service kubelet restart

# then worker:
service kubelet status
service kubelet restart
service kube-proxy status
service kube-proxy restart`
```
