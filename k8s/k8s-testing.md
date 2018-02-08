# Sample kubectl calls to test k8s:

```
kubectl run bb --image=busybox sleep 864000

kubectl run hello --image=busybox  -- /bin/sh -c "while true; do sleep 10; echo test; done;"
kubectl logs -f $(kubectl get po --selector=run=hello --output=jsonpath={.items..metadata.name})

kubectl run nginx1 --image=nginx --port=80  --replicas=2  --command --

kubectl run web --image=mrsabath/web-ms  --replicas=2 --env="TEST=test-web-ms"
kubectl get po
kubectl exec -it web-2950442834-72td9 curl localhost

kubectl run web-ms --image=mrsabath/web-ms --port=80  --replicas=2 --env="TEST=test-web-ms" --command --
kubectl exec -it web-ms-in-2950442834-72td9 sh
```

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
