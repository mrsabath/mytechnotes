# radiant24:
VIP is 10.140.146.7

`openssl x509 -in $cert -text | grep CN`

## On-boarding Processes

```
# login to dev-mon01-ricerca-radiant03-pd-01
ssh radiant@10.140.171.98

radiant@dev-mon01-ricerca-radiant03-pd-01:~$ env | grep KUB
KUBECONFIG=/etc/kubernetes/admin-kubeconfig
cat /etc/kubernetes/admin-kubeconfig
# TLC certs:
/var/run/kubernetes/

# put the cert and key to
~/workspace/kubectl/radiant03/
# then reference them in
~/workspace/kubectl/radiant03/admin-kubeconfig
# then
export KUBECONFIG=~/workspace/kubectl/radiant03/admin-kubeconfig


# setup and run:
. ./setup_radiant03.sh
./kubectl --v=10 get pods
```

create namespace:

`kubectl create namespace <name>`

```
export space=
kubectl get namespace
kubectl create namespace s${space}-default
kubectl --namespace=s${space}-default create -f quota.yaml
kubectl --namespace=s${space}-default create -f limits.yaml

# to delete:
kubectl delete namespace s${space}-default

```
create user:

```
curl -XPUT -k https://localhost:8888/user/$user/s${space}-default

curl -XPUT -k https://localhost:8888/user/<user name>/<namespace>
curl -v -XPUT -k https://localhost:8888/user/6f709ffde60778edbcaa0d7afdedfdd806d0a6a733d8bc5a/s4975c8b9-a876-4e74-bebe-2f02f7e2b314-default
curl -XGET -k


# delete
curl -XDELETE -k https://localhost:8888/user/$user
# or if several namespaces:
curl -XDELETE -k https://localhost:8888/user/$user/s${space}-default

```


Raw calls:
```
kubectl --v=10 create namespace test2

I0513 20:24:30.750416   20872 request.go:555] Request Body: {"kind":"Namespace","apiVersion":"v1","metadata":{"name":"test2","creationTimestamp":null},"spec":{},"status":{}}
I0513 20:24:30.750546   20872 round_trippers.go:267] curl -k -v -XPOST  -H "Content-Type: application/json" -H "User-Agent: kubectl/v1.2.0 (linux/amd64) kubernetes/968370b" -H "Accept: application/json, */*" https://10.140.171.205:443/api/v1/namespaces
I0513 20:24:30.754018   20872 round_trippers.go:286] POST https://10.140.171.205:443/api/v1/namespaces 201 Created in 3 milliseconds


kubectl --v=10 --namespace=test2 create -f quota.yaml

I0513 20:28:06.712015   21141 request.go:555] Request Body: {"kind":"ResourceQuota","apiVersion":"v1","metadata":{"name":"quota","namespace":"test2","creationTimestamp":null},"spec":{"hard":{"resourcequotas":"1","secrets":"10","services":"10"}},"status":{}}
I0513 20:28:06.712114   21141 round_trippers.go:267] curl -k -v -XPOST  -H "Accept: application/json, */*" -H "Content-Type: application/json" -H "User-Agent: kubectl/v1.2.0 (linux/amd64) kubernetes/968370b" https://10.140.171.205:443/api/v1/namespaces/test2/resourcequotas
I0513 20:28:06.718436   21141 round_trippers.go:286] POST https://10.140.171.205:443/api/v1/namespaces/test2/resourcequotas 201 Created in 6 milliseconds

new Body
{"kind":"ResourceQuota","apiVersion":"v1","metadata":{"name":"quota","namespace":"test2","creationTimestamp":null},"spec":{"hard":{"resourcequotas":"1","secrets":"10","services":"10"}},"status":{}}
{"kind":"ResourceQuota","apiVersion":"v1","metadata":{"namespace":"s6bb3","name":"quota","creationTimestamp":null},"spec":{"hard":{"resourcequotas":"1","secrets":"10","services":"10"},"status":{}}


kubectl --v=10 --namespace=test2 create -f limits.yaml

I0513 20:29:22.739107   21212 request.go:555] Request Body: {"kind":"LimitRange","apiVersion":"v1","metadata":{"name":"limits","namespace":"test2","creationTimestamp":null},"spec":{"limits":[{"type":"Pod","max":{"cpu":"8","memory":"112Gi"},"min":{"cpu":"200m","memory":"64Mi"}},{"type":"Container","max":{"cpu":"8","memory":"112Gi"},"min":{"cpu":"200m","memory":"64Mi"},"default":{"cpu":"400m","memory":"128Mi"},"defaultRequest":{"cpu":"200m","memory":"64Mi"},"maxLimitRequestRatio":{"cpu":"2","memory":"2"}}]}}
I0513 20:29:22.739194   21212 round_trippers.go:267] curl -k -v -XPOST  -H "Accept: application/json, */*" -H "Content-Type: application/json" -H "User-Agent: kubectl/v1.2.0 (linux/amd64) kubernetes/968370b" https://10.140.171.205:443/api/v1/namespaces/test2/limitranges
I0513 20:29:22.743928   21212 round_trippers.go:286] POST https://10.140.171.205:443/api/v1/namespaces/test2/limitranges 201 Created in 4 milliseconds
```

# testing the kubeinit:
curl -XPOST -H "X-Tls-Client-Dn: /CN=5790bcf3279d852cc233e9ec00064c7819bb13f548cbdfab" -H "Content-Type: application/json"  https://localhost:8087/kubeinit

/usr/local/Cellar/curl/7.46.0/bin/curl -k -v -XPOST --cert cert.pem --key key.pem --cacert ca.pem -H "Content-Type: application/json"  https://localhost:8087/kubeinit

# old:
 curl -XPOST -H "X-Tls-Client-Dn: /CN=5790bcf3279d852cc233e9ec00064c7819bb13f548cbdfab" -H "Content-Type: application/json"  localhost:8087/kubeinit

export user=
curl -XPOST -H "X-Tls-Client-Dn: /CN=$user" -H "Content-Type: application/json"  localhost:8087/kubeinit

kubectl get namespaces
kubectl delete namespace test2



REST calls for 10.140.171.205:8888
```
// Show user authorization information
 GET /user                             // Show all user authz info
 GET /user/<user>                      // Show authz info for one user
 GET /user/<user>/<namespace>          // Show authz info for one user in one namespace
// Add user
PUT /user/<user>[?privileged=true]     // Create an admin user
PUT /user/<user>/<namespace>           // Create a normal user in a particular namespace
// Delete user
DELETE /user/<user>                    // Revoke authz of a user completely
DELETE /user/<user>/<namespace>        // Revoke authz of a user in a particular namespace
```

## new user onboarding:
```
ssh radiant@10.140.171.205
sudo su -
# make sure the user policy manager is running:
docker ps | grep remoteabac
docker exec remoteabac /opt/kubernetes/ruser --help

# show the policies:
docker exec remoteabac /opt/kubernetes/ruser --authorization-policy-file=etcd@http://10.140.171.205:4001/abac-policy --type=show

# setup new policy:
docker exec remoteabac /opt/kubernetes/ruser --authorization-policy-file=etcd@http://10.140.171.205:4001/abac-policy --type=add \
--user=<apikey> --namespace=<spaceid>

# view directly in etcd:
docker exec etcd etcdctl get /abac-policy`

```

## testing KUBE locally. (inject the labels)
```
* radiant01(f7f413cb-a678-412d-b024-8e17e28bcb88)
# with label:
curl -k -v -XPOST  -H "X-Tls-Client-Dn: /CN=d7eae25d39f061dd40937d3839b96fc34d4401391823160f" -H "Content-Type: application/json" -H "User-Agent: kubectl/v1.3.0 (darwin/amd64) kubernetes/df66b90" -H "Accept: application/json, */*" -d '{"kind":"Pod","apiVersion":"v1","metadata":{"name":"kube-web-ms","namespace":"sf7f413cb-a678-412d-b024-8e17e28bcb88-default","creationTimestamp":null,"labels":{"app":"web-ms-demo"},"annotations":{"aaa":"bbbb", "containers-label.alpha.kubernetes.io/com.swarm.tenant.0":"sf7f413cb-a678-412d-b024-8e17e28bcb88-default"}},"spec":{"containers":[{"name":"kube-web-server","image":"mrsabath/web-ms:v3","ports":[{"containerPort":80,"protocol":"TCP"}],"env":[{"name":"TEST","value":"web-server"}],"resources":{},"terminationMessagePath":"/dev/termination-log","imagePullPolicy":"IfNotPresent"}],"restartPolicy":"Always","terminationGracePeriodSeconds":30,"dnsPolicy":"ClusterFirst","securityContext":{}},"status":{}}' localhost:8087/api/v1/namespaces/sf7f413cb-a678-412d-b024-8e17e28bcb88-default/pods


# without label:
curl -k -v -XPOST  -H "X-Tls-Client-Dn: /CN=d7eae25d39f061dd40937d3839b96fc34d4401391823160f" -H "Content-Type: application/json" -H "User-Agent: kubectl/v1.3.0 (darwin/amd64) kubernetes/df66b90" -H "Accept: application/json, */*" -d '{"kind":"Pod","apiVersion":"v1","metadata":{"name":"kube-web-ms","namespace":"sf7f413cb-a678-412d-b024-8e17e28bcb88-default","creationTimestamp":null,"labels":{"app":"web-ms-demo"},"annotations":{"aaa":"bbbb"}},"spec":{"containers":[{"name":"kube-web-server","image":"mrsabath/web-ms:v3","ports":[{"containerPort":80,"protocol":"TCP"}],"env":[{"name":"TEST","value":"web-server"}],"resources":{},"terminationMessagePath":"/dev/termination-log","imagePullPolicy":"IfNotPresent"}],"restartPolicy":"Always","terminationGracePeriodSeconds":30,"dnsPolicy":"ClusterFirst","securityContext":{}},"status":{}}' localhost:8087/api/v1/namespaces/sf7f413cb-a678-412d-b024-8e17e28bcb88-default/pods


* radiant02(4975c8b9-a876-4e74-bebe-2f02f7e2b314)
# with label:
curl -k -v -XPOST  -H "X-Tls-Client-Dn: /CN=74e5c25b02de81fcfc9262abd26f3d66c9913877dd3055b7" -H "Content-Type: application/json" -H "User-Agent: kubectl/v1.3.0 (darwin/amd64) kubernetes/df66b90" -H "Accept: application/json, */*" -d '{"kind":"Pod","apiVersion":"v1","metadata":{"name":"kube-web-ms","namespace":"s4975c8b9-a876-4e74-bebe-2f02f7e2b314-default","creationTimestamp":null,"labels":{"app":"web-ms-demo"},"annotations":{"aaa":"bbbb", "containers-label.alpha.kubernetes.io/com.swarm.tenant.0":"s4975c8b9-a876-4e74-bebe-2f02f7e2b314-default"}},"spec":{"containers":[{"name":"kube-web-server","image":"mrsabath/web-ms:v3","ports":[{"containerPort":80,"protocol":"TCP"}],"env":[{"name":"TEST","value":"web-server"}],"resources":{},"terminationMessagePath":"/dev/termination-log","imagePullPolicy":"IfNotPresent"}],"restartPolicy":"Always","terminationGracePeriodSeconds":30,"dnsPolicy":"ClusterFirst","securityContext":{}},"status":{}}' localhost:8087/api/v1/namespaces/s4975c8b9-a876-4e74-bebe-2f02f7e2b314-default/pods


# without label:
curl -k -v -XPOST  -H "X-Tls-Client-Dn: /CN=74e5c25b02de81fcfc9262abd26f3d66c9913877dd3055b7" -H "Content-Type: application/json" -H "User-Agent: kubectl/v1.3.0 (darwin/amd64) kubernetes/df66b90" -H "Accept: application/json, */*" -d '{"kind":"Pod","apiVersion":"v1","metadata":{"name":"kube-web-ms","namespace":"s4975c8b9-a876-4e74-bebe-2f02f7e2b314-default","creationTimestamp":null,"labels":{"app":"web-ms-demo"},"annotations":{"aaa":"bbbb"}},"spec":{"containers":[{"name":"kube-web-server","image":"mrsabath/web-ms:v3","ports":[{"containerPort":80,"protocol":"TCP"}],"env":[{"name":"TEST","value":"web-server"}],"resources":{},"terminationMessagePath":"/dev/termination-log","imagePullPolicy":"IfNotPresent"}],"restartPolicy":"Always","terminationGracePeriodSeconds":30,"dnsPolicy":"ClusterFirst","securityContext":{}},"status":{}}' localhost:8087/api/v1/namespaces/s4975c8b9-a876-4e74-bebe-2f02f7e2b314-default/pods
```


5  curl -k https://localhost:8888/user  
7  docker exec etcd0 etcdctl ls /
8  docker exec etcd0 etcdctl get /abac-policy
10  kubectl get sa
11  kubectl get secrets
12  kubectl get namespaces
14  kubectl create namespace mariusz
15  kubectl get namespaces
16  kubectl get sa
17  ps aux | grep apiserver


# works on ubuntu
curl -k -XPOST -s -o /dev/null --cert "cert.pem" --key "key.pem" --cacert "ca.pem" -H "Content-Type: application/json"  https://localhost:8087/kubeinit?test1=T&test2=2

# works on mac
/usr/local/Cellar/curl/7.46.0/bin/curl -k -XPOST -s -o /dev/null --cert "cert.pem" --key "key.pem" --cacert "ca.pem" -H "Content-Type: application/json"  "https://localhost:8087/kubeinit?test1=T&test2=2"
