# Kind hosts a local K8s cluster
https://kind.sigs.k8s.io/

## install:
```
GO111MODULE="on" go get sigs.k8s.io/kind@v0.9.0
# or
brew install kind
```

## create cluster:
```
kind create cluster
king get clusters
kind delete cluster

# create cluster-name, verbosity 5:
kind create cluster --name cluster-name -v 5

# using config file:
cat <<EOF | kind create cluster --name spire-example -v 5 --config=-
kind: Cluster
apiVersion: kind.x-k8s.io/v1alpha4
containerdConfigPatches:
- |-
  [plugins."io.containerd.grpc.v1.cri".registry.mirrors."localhost:${reg_port}"]
    endpoint = ["http://${reg_name}:${reg_port}"]
nodes:
- role: control-plane
  kubeadmConfigPatches:
  - |
EOF

#
docker build . -t $image-name
# load the image to kind registry
kind load docker-image $image-name

```

## create a cluster with multiple nodes
```
cat <<EOF | kind create cluster --name multi-node-cluster -v 5 --config=-
kind: Cluster
apiVersion: kind.sigs.k8s.io/v1alpha3
nodes:
- role: control-plane
- role: worker
- role: worker
EOF
```
