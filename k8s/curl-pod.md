# Create a Deployment with container to execute curl commands

Create a deployment first: e.g. `deploy-curl.yaml`
```
apiVersion: extensions/v1beta1
kind: Deployment
metadata:
   name: sleep-us
spec:
  replicas: 1
  template:
    metadata:
      labels:
        app: sleep-us
    spec:
      containers:
      - name: sleep
        image: tutum/curl
        command: ["/bin/sleep","infinity"]
        imagePullPolicy: IfNotPresent
```


```
kubectl create -f deploy-curl.yaml
kubectl exec -it $(kubectl get pod -l app=sleep-us -o jsonpath='{.items[0].metadata.name}') -c sleep -- bash -c 'curl http://ratings:9080/ratings/0'
```
