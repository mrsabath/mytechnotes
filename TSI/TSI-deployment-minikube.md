# Trusted Service Identity
# Deployment of TSI in minikube

I will show you today how to deploy TSI in Minikube. If you like to follow the
documentation for this demo, please visit our open-sourced page in github.

This demo assumes you already have a minkube running, with kubectl and helm
installed as instructed in our
Readme page on GitHub.
To make typing more efficient let's setup an alias KK to reference the Namespace
trusted-identity that we are using in this deployment.
OK, step 1 create a namespace
This script removes existing namespace, if there is any, and creates a new namespace
- Trusted Identity.

2. TSI requires Vault to store the secrets. We can use an existing Vault service,
if one is available, or we can create one, running as a container. For the purpose
of this demo, let's create one quickly. This configurations starts a container and
a vault service. Let's get the minikube access to this service and test it using
CURL.
This is an expected result at this point. Let's assign this endpoint to the environment
variable.

Step 3, Setup all cluster nodes. In our minikube there is only one node. This helm
deployment would create a daemonset, with one container per each node. This container
creates a private key associated with this node instance.

Step 4, install TSI environment, Using the most recent Helm chart. Here we need
to define the cluster region, let's make it EU-Germany and name of the cluster,
for example MINIKUBE. We need also the Vault endpoint that we set earlier.

Step 5, setup Vault. For this step we need the Vault Admin token, that can be obtained
from the vault container log. Using this Root admin vault credential, we will deploy
the Vault authentication plugin, and create new Root CA, that would be securely
stored in the Vault

Step 6, register all the nodes. This script creates intermediate CAs, for all
the nodes, that are signed by the main ROOT CA stored in the Vault. Once the node
is registered, the public interface for this node is disabled, so no more changes can
be done without resetting the node.

Step 7, remove the Setup containers. Since the environment is considered to be securely
bootstrapped, and the worker nodes are registered with Vault, we don't need the
daemonset on the workers, so let's remove them now by removing the helm deployment.

Step 8, We can now load few sample policies to the Vault, so we can test the end
to end interaction. Let me show you an example of the policy that contains a
cluster region, cluster name, a namespace and the name of the image. This policy
grants the read access.
To load the policies, we execute 'demo.load-sample-policies' script.

Step 9. Let's load few sample keys to the Vault, so various test containers can
retrieve them. To run this script, we need to provide the sample region name, Germany, and the
cluster name, in our case minikube.


Step 10, let's deploy a sample container. We need to annotate what secrets we want
 and where they should be mounted on this container. Here is the
sample deployment configuration with several secrets that we would attempt to access.

Let me dump the deployment details, to show the annotated secrets. The
TSI sidecar attempts to retrieve the secrets and if they exist and the policies
on vault allow this container to access them, the secrets will be mounted at
specified locations.

We can exec into this container and see if we can read the secret.
OH, I need a secret name to follow the secret path.

Here we are, here is the secret.

Just for fun, I can show you the script that deploys the sample keys. And here
is the entry that creates this sample secret for this specific environment, Germany
minkube cluster and specific image that we used for this demo.

This concludes this demonstration.