## create pod on specific node:
```
# label the node:
kubectl label nodes 9.47.168.131 mariusz=true

# create pod.yaml with nodeSelector:
apiVersion: v1
kind: Pod
metadata:
  name: nginx
  labels:
    env: test
spec:
  containers:
  - name: nginx
    image: nginx
    imagePullPolicy: IfNotPresent
  nodeSelector:
    mariusz: "true"

# create the pod    
kubectl create -f pod.yaml
```
