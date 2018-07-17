# Installing Istio in IBM Kubernetes Service

Following https://console.bluemix.net/docs/containers/cs_tutorials_istio.html#istio_tutorial

Installing in IKS paid_1.9.3

```
cd ibmcloud/
curl -L https://git.io/getLatestIstio | sh -
Add /Users/sabath/ibmcloud/istio-0.8.0/bin to your path; e.g copy paste in your shell and/or ~/.profile:
export PATH="$PATH:/Users/sabath/ibmcloud/istio-0.8.0/bin"
```

Deploy Istio and validate the deployment:
```
kubectl apply -f install/kubernetes/istio-demo.yaml
kubectl get pod -n istio-system
```

Deploy BookInfo app
```
kubectl apply -f samples/bookinfo/kube/bookinfo.yaml -n istio-system
# validate: 
kubectl get svc -n istio-system
kubectl get pods -n istio-system

```
