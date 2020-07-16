# Advanced commands used for practicing the CKAD Exam (July 2020)
https://medium.com/bb-tutorials-and-thoughts/practice-enough-with-these-questions-for-the-ckad-exam-2f42d1228552

## Core Concepts
## Events:
```
kubectl get ev -w --watch-only=true | grep Pod
```

```
# delete without delay
k delete po nginx --grace-period=0 --force

# get yaml for the Pod
kubectl run nginx --image=nginx --restart=Never --dry-run -o yaml > nginx-pod.yaml

# List all the pods showing name and namespace with a json path expression
kubectl get pods -o=jsonpath="{.items[*]['metadata.name', 'metadata.namespace']}"

# Create the nginx pod with version 1.17.4 and expose it on port 80
kubectl run nginx --image=nginx:1.17.4 --restart=Never --port=80

# change the image of the existing pod
kubectl set image pod/nginx nginx=nginx:1.15-alpine

# watch the pod
kubectl get po nginx -w # watch it

# get pod image from jsonpath
k get po nginx -o jsonpath='{.spec.containers[].image}{"\n"}'

# get IP of the pod
k get po -owide

# If pod crashed check the previous logs of the pod
kubectl logs busybox -p

# Create a busybox pod with command sleep 3600
kubectl run busybox --image=busybox --restart=Never -- /bin/sh -c "sleep 30"

# Check the connection of the nginx pod from the busybox pod
kubectl get po nginx -o wide// check the connection
kubectl exec -it busybox -- wget -o- <IP Address>

# create a busybox pod and echo message ‘How are you’ and delete it manually
kubectl run busybox --image=nginx --restart=Never -it -- echo "How are you"
kubectl delete po busybox

# create a pod and list different level of verbosity:
// create a pod
kubectl run nginx --image=nginx --restart=Never --port=80
// List the pod with different verbosity
kubectl get po nginx --v=7
kubectl get po nginx --v=8
kubectl get po nginx --v=9

# list pods using custom columns:
kubectl get po -o=custom-columns="POD_NAME:.metadata.name, POD_STATUS:.status.containerStatuses[].state"

# sort pods by name, then create time
kubectl get pods --sort-by=.metadata.name
kubectl get pods --sort-by=.metadata.creationTimestamp

```

## Change the context

```
# list contexts
k config get-contexts

# switch to another one:
k config use-context <context-name>

# list clusters:
k config get-clusters

```

## Multi-containers

```
# Create a Pod with three busy box containers with commands “ls; sleep 3600;”, “echo Hello World; sleep  3600;” and “echo this is the third container; sleep 3600” respectively and check the status

// first create single container pod with dry run flag
kubectl run busybox --image=busybox --restart=Never --dry-run -o yaml -- bin/sh -c "sleep 3600; ls" > multi-container.yaml
// edit the pod by adding 2 additonal containers  busybox1,2,3 then create:
kubectl create -f multi-container.yaml

// Check the previous logs of the second container busybox2 if any
kubectl logs busybox -c busybox2 --previous

// Run command ls in the third container busybox3 of the above pod
kubectl exec busybox -c busybox3 -- ls

// Show metrics of the above pod containers and put them into the file.log and verify

kubectl top pod busybox --containers
// putting them into file
kubectl top pod busybox --containers > file.log
cat file.log

// list metrics for nodes and pods
kubectl top nodes
kubectl top pods

```

Create a Pod with main container busybox and which executes this “while true; do echo ‘Hi I am from Main container’ >> /var/log/index.html; sleep 5; done” and with sidecar container with nginx image which exposes on port 80. Use emptyDir Volume and mount this volume on path /var/log for busybox and on path /usr/share/nginx/html for nginx container. Verify both containers are running.

```console
// create an initial yaml file with this
kubectl run multi-cont-pod --image=busbox --restart=Never --dry-run -o yaml > multi-container.yaml
// edit the yml as below and create it
kubectl create -f multi-container.yaml

// or simply:
kubectl create -f- <<EOF
apiVersion: v1
kind: Pod
metadata:
  creationTimestamp: null
  labels:
    run: multi-cont-pod
  name: multi-cont-pod
spec:
  volumes:
  - name: var-logs
    emptyDir: {}
  containers:
  - image: busybox
    command: ["/bin/sh"]
    args: ["-c", "while true; do echo 'Hi I am from Main container' >> /var/log/index.html; sleep 5;done"]
    name: main-container
    resources: {}
    volumeMounts:
    - name: var-logs
      mountPath: /var/log
  - image: nginx
    name: sidecar-container
    resources: {}
    ports:
      - containerPort: 80
    volumeMounts:
    - name: var-logs
      mountPath: /usr/share/nginx/html
  dnsPolicy: ClusterFirst
  restartPolicy: Never
status: {}
EOF

kubectl get po multi-cont-pod
```
Exec into both containers and verify that main.txt exist and query the main.txt from sidecar container with curl localhost

```console
// exec into main container
kubectl exec -it  multi-cont-pod -c main-container -- cat /var/log/index.html

// exec into sidecar container
kubectl exec -it  multi-cont-pod -c sidecar-container -- sh

cat /usr/share/nginx/html/index.html

// install curl and get default page
apt-get update && apt-get install -y curl
curl localhost
```
