# Start daemonset that monitors container labels.
If container label changes to paused=true, container will be paused.
using [kube-pause-daemonset.yaml](kube-pause-daemonset.yaml)
```
kubect create -f kube-pause-daemonset.yaml
kubectl get po
kubectl run nginx --image-nginx --restart=Never
kubectl get po -owide
kubectl exec -it nginx /bin/bash
# works
kubectl label po nginx paused=true
kubectl exec -it nginx /bin/bash
# does not works
kubectl label po nginx paused=false --overwrite
kubectl exec -it nginx /bin/bash
# works again
``
