# ibmcloud plugin version 1.0.28 removes ability to specify KUBECONFIG files

To setup individual KOBECONFIG files:


```
ibmcloud ks clusters
export KUBECONFIG=$(mktemp).roks4.yaml
ibmcloud ks cluster config --admin -c <cluster>
or
ibmcloud ks cluster config -c <cluster>
echo $KUBECONFIG
```

## Obtain cluster name:
```
~$k config get-clusters
NAME
roks4-2/bpgj9hhf0r42tn8781q0
~$k config current-context
roks4-2/bpgj9hhf0r42tn8781q0/admin
```
