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
kubectl get pods -o jsonpath="{.items[*]['metadata.name', 'metadata.namespace']}"

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


```console
# create few sample pods to test lables
kubectl run nginx-dev1 --image=nginx --restart=Never --labels=env=dev
kubectl run nginx-prod1 --image=nginx --restart=Never --labels=env=prod

kubectl get pods -l env=prod --show-labels
kubectl get pods -L env
kubectl get pods -l 'env in (dev,prod)' --show-labels

# change the label
kubectl label pod/nginx-dev3 env=uat --overwrite
kubectl get pods --show-labels

# Remove the labels for the pods that we created now and verify all the labels are removed
kubectl label pod nginx-dev{1..3} env-
kubectl label pod nginx-prod{1..2} env-kubectl get po --show-labels

# Get pods with name range
kubectl get pod nginx-dev{1..3}

# Remove the labels for the pods that we created now and verify all the labels are removed
kubectl label pod nginx-dev{1..3} env-
kubectl label pod nginx-prod{1..2} env-
kubectl get po --show-labels

# get nodes and show their labels
kubectl get nodes --show-labels

# Label one of the nodes nodeName=nginxnode and
# create a Pod that will be deployed on this node with the label nodeName=nginxnode
kubectl label node minikube nodeName=nginxnode
kubectl run nginx --image=nginx --restart=Never --dry-run -o yaml > pod.yaml
# modify the pod.yaml with nodeSelector:
kubectl create -f- <<EOF
apiVersion: v1
kind: Pod
metadata:
  creationTimestamp: null
  labels:
    run: nginx
  name: nginx
spec:
  nodeSelector:
    nodeName: nginxnode
  containers:
  - image: nginx
    name: nginx
    resources: {}
  dnsPolicy: ClusterFirst
  restartPolicy: Never
status: {}
EOF

# get the node that was labeled
kubectl get nodes -l nodeName=nginxnode

# describe the pod and grep Node-Selector
kubectl describe po nginx | grep Node-Selectors
kubectl describe po nginx | grep Labels

# Annotate the pods with name=webapp
kubectl annotate pod nginx-dev{1..3} name=webapp
kubectl annotate pod nginx-prod{1..2} name=webapp

# Verify the pods that have been annotated correctly
kubectl describe po nginx-dev{1..3} | grep -i annotations
kubectl describe po nginx-prod{1..2} | grep -i annotations

# show all the pods that have name=webapp annotation:
???


# Remove the annotations on the pods and verify
kubectl annotate pod nginx-dev{1..3} name-
kubectl annotate pod nginx-prod{1..2} name-

# delete all the pods in this namespace
k delete po --all

# create 5 replicas of the pod
kubectl create deploy webapp --image=nginx
k scale deploy webapp --replicas=5

# Get the deployment you just created with labels
kubectl get deploy webapp --show-labels

# Output the yaml file of the deployment you just created
kubectl get deploy webapp -o yaml

# Get the pods of this deployment
// get the label of the deployment
kubectl get deploy --show-labels
// get the pods with that label
kubectl get pods -l app=webapp

# Get the deployment rollout status
kubectl rollout status deploy webapp

# Get the replicaset that created with this deployment
kubectl get rs -l app=webapp

# Delete the deployment you just created and watch all the pods are also being deleted
kubectl delete deploy webapp
kubectl get po -l app=webapp -w

# Create a deployment of webapp with image nginx:1.17.1 with container port 80 and verify the image version

kubectl create deploy webapp --image=nginx:1.17.1 --dry-run -o yaml > webapp.yaml
# add the port section and create the deployment
#         name: nginx
#         ports:
#         - containerPort: 80
kubectl create -f webapp.yaml
// verify
kubectl describe deploy webapp | grep Image

# Update the deployment with the image version 1.17.4 and verify
kubectl set image deploy/webapp nginx=nginx:1.17.4

# Check the rollout history and make sure everything is ok after the update
kubectl rollout history deploy webapp
kubectl get deploy webapp --show-labels
kubectl get rs -l app=webapp
kubectl get po -l app=webapp

# Undo the deployment to the previous version 1.17.1 and verify Image has the previous version
kubectl rollout undo deploy webapp
kubectl describe deploy webapp | grep Image
kubectl rollout history deploy webapp

# Update the deployment to the Image 1.17.1 and verify everything is ok
kubectl rollout undo deploy webapp --to-revision=3
kubectl describe deploy webapp | grep Image
kubectl rollout status deploy webapp
kubectl rollout history deploy webapp

# Update the deployment with the wrong image version 1.100 and verify something is wrong with the deployment
kubectl set image deploy/webapp nginx=nginx:1.100
kubectl rollout status deploy webapp (still pending state)
kubectl get pods (ImagePullErr)

# Undo the deployment with the previous version and verify everything is Ok
kubectl rollout undo deploy webapp
kubectl rollout status deploy webapp
kubectl get pods

# check the details of specifc revision
kubectl rollout history deploy webapp
kubectl rollout history deploy webapp --revision=7

# Pause the rollout of the deployment, then attempt to change something
kubectl rollout pause deploy webapp
kubectl set image deploy/webapp nginx=nginx:latest
kubectl describe deploy webapp | grep Image // image updated but not deployed
kubectl rollout history deploy webapp
(No new revision)

# Resume the rollout of the deployment
kubectl rollout resume deploy webapp

# Check the rollout history and verify it has the new version
kubectl rollout history deploy webapp
kubectl rollout history deploy webapp --revision=8

# Apply the autoscaling to this deployment with minimum 10 and maximum 20 replicas and target CPU of 85% and verify hpa (Horizontal Pod Autoscaler) is created and replicas are increased to 10 from 1

kubectl autoscale deploy webapp --min=10 --max=20 --cpu-percent=85
kubectl get hpa
kubectl get pod -l app=webapp

# Clean the cluster by deleting deployment and hpa you just created
kubectl delete deploy webapp
kubectl delete hpa webapp

```

# Jobs
```console
# Create a Job with an image node which prints node version and also verifies
# there is a pod created for this job
kubectl create job nodeversion --image=node -- node -v
kubectl get job -w
kubectl get pod

# Get the logs of the job just created
kubectl logs <pod name> // created from the job

# Output the yaml file for the Job with the image busybox which echos “Hello I am from job”
kubectl create job hello-job --image=busybox --dry-run -o yaml -- echo "Hello I am from job" > hello-job.yaml
kubectl create -f hello-job.yaml

# Verify the job and the associated pod is created and check the logs as well
kubectl get job
kubectl get pokubectl logs hello-job-*

# Delete the job we just created
kubectl delete job hello-job

# Create the same job and make it run 10 times one after one
kubectl create job hello-job --image=busybox --dry-run -o yaml -- echo "Hello I am from job" > hello-job.yaml
// edit the yaml file to add completions: 10
//  spec:
//     completions: 10
kubectl create -f hello-job.yaml
kubectl get job -w
kubectl get po
// get log from the first container of the job:
k logs job/hello-job
