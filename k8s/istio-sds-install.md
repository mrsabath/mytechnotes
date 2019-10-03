# Installing Istio with SDS on Minikube

Istio Minikube setup: https://istio.io/docs/setup/platform-setup/minikube/

```console
minikube start --vm-driver virtualbox --memory=8192 --cpus=4 --kubernetes-version v1.14.0
```

Download Istio release: https://istio.io/docs/setup/#downloading-the-release

Use Helm install: https://istio.io/docs/setup/install/helm/

```
kubectl create namespace istio-system
helm template install/kubernetes/helm/istio-init --name istio-init --namespace istio-system | kubectl apply -f -
# wait few seconds to allow the registry, then verify:
kubectl get crds | grep 'istio.io' | wc -l
```

Follow the SDS install:
```
helm template install/kubernetes/helm/istio --name istio --namespace istio-system \
    --values install/kubernetes/helm/istio/values-istio-sds-auth.yaml | kubectl apply -f -
```

Verfify:

```
kubectl get svc -n istio-system
kubectl get pods -n istio-system
```
