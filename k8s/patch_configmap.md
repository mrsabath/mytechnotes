## create a configmap, then patch it with a new value
```shell
kubectl create configmap test8 --from-literal=special.how=very --from-literal=special.type=charm
kubectl get configmap test8 -ojson
{
    "apiVersion": "v1",
    "data": {
        "special.how": "very",
        "special.type": "charm"
    },
    "kind": "ConfigMap",
    "metadata": {
        "creationTimestamp": "2017-11-06T15:22:39Z",
        "name": "test8",
        "namespace": "default",
        "resourceVersion": "10411004",
        "selfLink": "/api/v1/namespaces/default/configmaps/test8",
        "uid": "534e3574-c306-11e7-9190-fe68e41328f1"
    }
}

kubectl patch configmap test8 -p '"data": { "special.how" : "very", "special.type": "charmxx"}'
kubectl get configmaps test8 -ojson
{
    "apiVersion": "v1",
    "data": {
        "special.how": "very",
        "special.type": "charmxx"
    },
    "kind": "ConfigMap",
    "metadata": {
        "creationTimestamp": "2017-11-06T15:22:39Z",
        "name": "test8",
        "namespace": "default",
        "resourceVersion": "10411249",
        "selfLink": "/api/v1/namespaces/default/configmaps/test8",
        "uid": "534e3574-c306-11e7-9190-fe68e41328f1"
    }
}
#update only one field:
Silesia:update-service-deployment sabath$ kubectl patch configmap test8 -p '"data": { "special.type": "charmxyy"}'
configmap "test8" patched
Silesia:update-service-deployment sabath$ kubectl get configmaps test8 -ojson
{
    "apiVersion": "v1",
    "data": {
        "special.how": "very",
        "special.type": "charmxyy"
    },
    "kind": "ConfigMap",
    "metadata": {
        "creationTimestamp": "2017-11-06T15:22:39Z",
        "name": "test8",
        "namespace": "default",
        "resourceVersion": "10411785",
        "selfLink": "/api/v1/namespaces/default/configmaps/test8",
        "uid": "534e3574-c306-11e7-9190-fe68e41328f1"
    }
}
# or add a new field:
kubectl patch configmap test8 -p '"data": { "special.typeNEW": "charmyNEW"}'
configmap "test8" patched
Silesia:update-service-deployment sabath$ kubectl get configmaps test8 -ojson
{
    "apiVersion": "v1",
    "data": {
        "special.how": "very",
        "special.type": "charmxyy",
        "special.typeNEW": "charmyNEW"
    },
    "kind": "ConfigMap",
    "metadata": {
        "creationTimestamp": "2017-11-06T15:22:39Z",
        "name": "test8",
        "namespace": "default",
        "resourceVersion": "10411874",
        "selfLink": "/api/v1/namespaces/default/configmaps/test8",
        "uid": "534e3574-c306-11e7-9190-fe68e41328f1"
    }
}
```

## Another trick to update the configmap:

```shell
kubectl create configmap test4 --from-file=config.json -ojson
# update config.json file
kubectl create configmap test4 --from-file=config.json -ojson --dry-run | kubectl replace -f -
configmap "test4" replaced
```

## Create configmap from directory of files
```
kubectl create configmap test9 --from-file=configdir/
configmap "test9" created
Silesia:update-service-deployment sabath$ kubectl get configmap test9
NAME      DATA      AGE
test9     3         16s
Silesia:update-service-deployment sabath$ kubectl get configmap test9 -ojson
{
    "apiVersion": "v1",
    "data": {
        "test1": "T1=V1\nT2=V2\n",
        "test2": "XXX=YYY\nZZZ=PPP\n",
        "test3": "A:1\nB:2\n"
    },
    "kind": "ConfigMap",
    "metadata": {
        "creationTimestamp": "2017-11-06T15:35:16Z",
        "name": "test9",
        "namespace": "default",
        "resourceVersion": "10412218",
        "selfLink": "/api/v1/namespaces/default/configmaps/test9",
        "uid": "1697865f-c308-11e7-9190-fe68e41328f1"
    }
}

ls configdir/
test1	test2	test3

```
