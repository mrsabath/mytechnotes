# Update Service Helm Charts
Update Service is a tool that automates and orchestrates the process of updating
all the worker nodes in your cluster.
Update Service optimally selects the first set of nodes, drains and reloads them
 on behalf of a user. Draining makes workloads migrate to other nodes.
After a reload, the node is automatically tested and brought back to the cluster.
As nodes complete the update process, the Update Service selects new nodes for
update and automatically initiates the update process.

An optimized update process selects nodes to minimize the update time and minimize
the number of disruptions while making sure that workload constraints are always
met, and there is enough available capacity in the cluster to host all workloads.

This chart deploys Update Service microservices in the following order:
* update-planner - Planner microservice that decides what is done and when
* update-executor - Executor microserivce that executes requested tasks
* update-req - update request, initializes update procedure. Contains the update
procedures. Type of the request is selected by user and currently supports
IBM Container Cloud worker reboot or reload.

Additionally, there is a `synthetic-load` chart that provides the workloads for
testing this service.

Update Planner is responsible for discovering the nodes, their current workloads
and making decisions about the sequence of updates.
Planner execution is triggered by Update CR. Planner creates, watches and updates
Inventory CRs. When Planner is deployed, it will register Custom Resource Definitions for the following:
* updates.updatesrv.ibm.com - update request details
* inventories.updatesrv.ibm.com - individual inventory details (e.g. for each node instance)

Update Executor is watching Inventory CRs and executes required transactions.
During the execution it updates the Inventories with the execution result.

### Update Service Steps

* [Important](#important)
* [Pre-reqs and setup](#pre-reqs-and-setup)
* [Deploying from staging charts](#deploying-from-staging-charts)
* [Deploying from production charts](#deploying-from-production-charts)
* [Deploying Synthetic Load](#deploying-synthetic-load)
* [Deploying the Update Service](#deploying-the-update-service)
* [Track or debug](#track-or-debug)
* [Cleanup](#cleanup)

---

### Important
As of this moment, all these charts are available only from Staging Helm repository.
We are working on publishing them into production. This means you can still use
them in your IBM Cloud Kubernetes cluster (either free of paid one, however using
  free cluster, with only one worker node, is not very useful here...), but you will
  not be able to see them when logged to [IBM Cloud UI](https://console.bluemix.net/containers-kubernetes/solutions/helm-charts)

Here are the typical steps for setting up
the Update Service deployment in IBM Container Cloud:

### Pre-reqs and Setup
In order the deploy the Update Service charts, you must use [IBM Cloud CLI](https://console.bluemix.net/docs/containers/cs_cli_install.html#cs_cli_install)
and have [a Kubernetes cluster](https://console.bluemix.net/docs/containers/cs_tutorials.html#cs_cluster_tutorial)

1.  Login to IBM Cloud and target your cluster with [Kubernetes CLI](https://console.bluemix.net/docs/containers/cs_cli_install.html#cs_cli_configure)

  *  Log in to the IBM Cloud CLI client. If you have a federated account, use `--sso`.
      ```console
      $ bx login [--sso]
      ```

  *  Get your IBM Cloud account ORG name that you want to use.
      ```console
      $ bx account orgs
      ```

  * Target the specific ORG
      ```console
      $ bx target -o <ORG_NAME>
      ```

1.  Get the name of the cluster where you want to install Update Service. This values is needed to deploy the service.
    ```console
    $ bx cs clusters
    ```

1. Setup the KUBECONFIG for selected cluster
    ```console
    $ bx cs cluster-config <CLUSTER_NAME>
    ```

1. Export the KUBECONFIG as per message obtained above
     ```
     $ export KUBECONFIG=~/.bluemix/plugins/container-service/clusters/<CLUSTER_NAME>/<CLUSTER_CONFIG_FILE>
     ```

1. Test the kubectl access to your cluster
    ```console
    $ kubectl get nodes
    ```

1. Create API key as secret. Simply save the text below as `init.sh` script

  ```bash
    #!/bin/bash

    # check if API key was created and assigned as K8s secret
    SECRET="update-srv-secret"
    KEYNAME="update-key"

    # check if secret with this name exists already
    kubectl get secret $SECRET -ojson
    if [ "$?" -eq "0" ]; then
     echo "Secret already exists!"
    else
      echo "Secret needs to be created"
      # create the new API key
      KEY=$(bx iam api-key-create $KEYNAME -d "Key for Update Service" | grep "API Key" | awk '{ print $3 }')
      KEY64=$(echo -n "$KEY" | base64)
      # create k8s secret using base64 encoding of the API key
      cat >updateSrv-secret.yaml <<EOF
    apiVersion: v1
    kind: Secret
    metadata:
      name: $SECRET
    data:
      api-key: $KEY64
    EOF
      kubectl create -f updateSrv-secret.yaml
      if [ "$?" -eq "0" ]; then
       echo "Everything OK!"
      else
        echo "Error. Exiting"
        exit 1
      fi
    fi
    kubectl get secret $SECRET
  ```

  And then execute it
  ```console
    $ chmod +x ./init
    $ ./init
  ```

1.  Install the [Helm CLI](https://docs.helm.sh/using_helm/#installing-helm) and configure Helm in your cluster using the `helm init` command. IMPORTANT: be sure to follow the best practices on [securing the helm deployment](https://docs.helm.sh/using_helm/#securing-your-helm-installation) ([and here](https://github.ibm.com/alchemy-containers/documentation/issues/1375))

### Deploying from Staging Charts
Right now we are only supporting charts in Staging Incubator.

1.  If you haven’t already, add the repo for this chart as given in the instructions [IBM Console Helm Charts Page](https://console.bluemix.net/containers-kubernetes/solutions/helm-repos)
  * List your Helm repos and add this one, if missing:
     * https://registry.stage1.ng.bluemix.net/helm/ibm (Staging stable)
     * https://registry.stage1.ng.bluemix.net/helm/ibm-incubator/ (Staging Incubator)
     * https://registry.bluemix.net/helm/ibm (IBM Production stable)
     * https://registry.bluemix.net/helm/ibm-incubator/ (IBM Production Incubator)

    ```console
    $ helm repo list

    # assign env variable for this repo name
    $ export HELM_REPO_NAME=ibm-staging-incubator
    $ helm repo add $HELM_REPO_NAME https://registry.stage1.ng.bluemix.net/helm/ibm-incubator/
    ```

### Deploying from Production Charts
As of this moment, all these charts are available only from Staging Helm repository.
We are working on publishing them into production.

### Deploying Synthetic Load
To test the Update Service, you can deploy some test workloads first. This step is not
needed if you already have some workloads that you like to use to evacuate during
the update.

* Decide on parameters for the deployment:
  * namespace - name of the namespace to deploy the workloads
  * webService.podCount - number of pods in the deployment that is load balancing the `web-ms-service`
  * webDeplGroup.podCount - number of pods in each group. There are always 10 groups deployed.

* Deploy the `synthetic-load` chart, where `webService.port` value must be unique:

    ```console
    $ helm install --name=synthetic-load --namespace=default --set webService.port=30001 \
    --set webService.podCount=5 --set webDeplGroup.podCount=3 $HELM_REPO_NAME/synthetic-load
    ```
The example above will create 10 groups with 3 pods each and one deployment with 5 pods and a service running on port 30001.

### Deploying the Update Service

1. Decide on parameters for the deployment:
  * namespace - name of the namespace to deploy the Update Service
  * deployer.image.name - name of the Deployer image
  * deployer.image.tag - Deployer image tag
  * planner.type - name of the planner to orchestrate the update:
    * SEQ - sequential planner - one node executed at the time, sequentially
    * SIMOPT - simple optimized - planner decides what is the most efficient sequence of executions. Nodes can be updated concurrently
  * planner.ClusterAvailPodPerc - percent value of pods in the cluster that must stay alive during the update (e.g. 0.3)
  * executor.pods - number of executor pods to run
  * repository.name - name of the helm repository; used to identify the charts (e.g. ibm-staging-incubator)
  * repository.url - url of the above repository
  * updateReq.repository - name of the update request chart. Currently supporting:
    * "ibm-staging-incubator/update-req-armada-reboot" - to reboot IBM Container Cloud worker nodes
    * "ibm-staging-incubator/update-req-armada-reload" - to reload IBM Container Cloud worker nodes
  * cluster.name - name of the IBM Container Cloud cluster (see steps above)

1. Deploy the `update-deployer` chart using reboot option (reload takes much longer,
 so reboot is better for testing...) Make sure your HELM_REPO_NAME is set in
 steps above. Assign your cluster name to CLUSTER_NAME variable:

    ```console
    $ export CLUSTER_NAME=
    
    $ helm install --name=update-deployer --namespace=default \
    --set planner.type=SIMOPT --set planner.ClusterAvailPodPerc=0.3 \
    --set updateReq.repository=$HELM_REPO_NAME/update-req-armada-reboot \
    --set cluster.name=$CLUSTER_NAME $HELM_REPO_NAME/update-deployer
    ```

1.  If you like to create a new values files instead of passing the values directly, use the example embedded in the chart.
    ```
    $ helm inspect values $HELM_REPO_NAME/update-deployer > config.yaml
    ```
    Where `$HELM_REPO_NAME` is the name of the repository containing `update-deployer` (added above or from `helm repo list`).

1.  Edit the config.yaml file with values based on your needs.

1.  Deploy the `update-deployer` into your cluster.
    ```console
    $ helm upgrade -i --values=config.yaml update-deployer $HELM_REPO_NAME/update-deployer
    ```
1. List all the Helm deployments
  ```
  helm list --all
  ```
  Following charts should be deployed and running:
  * update-deployer
  * update-executor
  * update-planner
  * update-req
  * synthetic-load (if started separately)

### Track or debug

To track the deployment progress, open separate console for each activity.
Make sure all the consoles have KUBECONFIG setup per instructions above.

* Track the status of each inventory (node):
  ```console
  watch -n 5 kubectl get upd,inv -o custom-columns=TYPE:kind,NAME:metadata.name,STATUS:status.state,REQ-STEP:status.requested-step
  ```
* Get a detailed view of status of all steps of the inventories
   ```
   kubectl get inventories -ojson | jq '.items[] | ([.metadata.name, "Status: " + .status["state"], "Requested: " +  .status["requested-step"], (.spec.steps | to_entries[] | (  "Step " + .key  + ": " + .value.status["step-state"]) )  ])'
   ```
* Get logs from various micro-services
   ```
   kubectl exec $(kubectl  get po | grep update-deployer | awk '{print $1}') /update-deployer/getLogs.sh <type> <n>

   # Where:
   # type - deployer, planner or executor
   # n    - 1 (first executor), 2 (second) etc [optional]
   ```

* Track the status of micro-services
```
watch "kubectl get po -owide | grep update"
```

* Track the status of the worker nodes via `bx cs` cli
```
watch "bx cs workers <CLUSTER_NAME>"
```

## Cleanup
Cleanup the Update Service deployment (remove it all)
```console
kubectl exec $(kubectl  get po | grep update-deployer | awk '{print $1}') /update-deployer/remove_all.sh
# then delete the update-deployer deployment:
helm delete --purge update-deployer
```
