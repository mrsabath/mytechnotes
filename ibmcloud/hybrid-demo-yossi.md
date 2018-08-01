# ISTIO for multi-cloud

Newer version (from Jossi)
https://github.com/ymesika/istio_multiclusters_demos
My repo: ~/workspace/github.com/istio_multiclusters_demos

Originally based on: https://github.com/ymesika/hybrid-demo
My repo: ~/workspace/github.com/hybrid-demo

Cluster_A:
```
export KUBECONFIG=/Users/sabath/.bluemix/plugins/container-service/clusters/paid_1.9.3/kube-config-dal13-paid_1.9.3.yml
```

Cluster_B:
```
export KUBECONFIG=/Users/sabath/.bluemix/plugins/container-service/clusters/paid-1.9.8/kube-config-fra05-paid-1.9.8.yml
```

Set both contexts, test them:
```
export KUBECONFIG=/Users/sabath/.bluemix/plugins/container-service/clusters/paid_1.9.3/kube-config-dal13-paid_1.9.3.yml:/Users/sabath/.bluemix/plugins/container-service/clusters/paid-1.9.8/kube-config-fra05-paid-1.9.8.yml

kubectl config get-contexts
kubectl get nodes -n istio-system --context=paid_1.9.3
kubectl get nodes -n istio-system --context=paid-1.9.8
```

Execute the installation script:

```
./install.sh
```

The scripts will pause after installing Istio to both cluster and wait for you to press any key. Before continuing make sure all Istio pods are running on both clusters:

```
kubectl get pods -n istio-system --context=paid_1.9.3
kubectl get pods -n istio-system --context=<paid-1.9.8
```
The script continues and the Bookinfo web page URL will be printed at the end. You can open this URL in your browser but as the deployment of the app takes few seconds you might need to refresh to see all information for that app.

```
http://169.60.191.202/productpage
```
Make sure to refresh the page several times to see the different Book Reviews (versions). Note: it might not work right away
