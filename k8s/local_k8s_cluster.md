# Install K8s cluster locally, using Vagrant. 1 master + 3 workers
Follow the `system-admin` steps https://github.ibm.com/alchemy-containers/fr8r-ansible/tree/master/examples

create the VMs
```
git clone git@github.ibm.com:alchemy-containers/fr8r-ansible.git
cd fr8r-ansible/examples/vagrant
vi ../envs/dev-vbox/vagrant_config.rb #change to 3 workers
vi ../examples/envs/dev-vbox/shard1.hosts # add 2 more worker IPs
./quickstart.sh system-admin
```

install Fr8r (k8s cluster):
```
vagrant ssh installer
cd fr8r/ansible/
ansible-playbook -v -i ../examples/envs/dev-vbox/shard1.hosts env-basics.yml -e "envs=../examples/envs env_name=dev-vbox"
ansible-playbook -v -i ../examples/envs/dev-vbox/shard1.hosts shard.yml -e "envs=../examples/envs cluster_name=dev-vbox-shard1 network_kind=calico"
```

test the k8s cluster:
```
vagrant ssh shard1-master
kubectl get nodes
```

move the certs to your local system from `master`:
```
/etc/kubernetes/creds to
/Users/sabath/.fr8r/envs/dev-vbox/admin-certs
export KUBECONFIG=/Users/sabath/.fr8r/envs/dev-vbox/admin-certs/kube-config
```

Cleanup the Vagrant VMs:

```
vagrant destroy -f api-proxy client installer
vagrant status

# create a vagrant snapshot:
vagrant snapshot save shard1-master1 k8s-master
vagrant snapshot save shard1-worker1 k8s-worker1
vagrant snapshot list
vagrant snapshot restore k8s-master
vagrant snapshot delete k8s-master
```

## Other docs that don't work well:

git@github.com:jeremievallee/kubernetes-vagrant-ansible.git
cd kubernetes-vagrant-ansible

# install ansible
virtualenv venv
source venv/bin/activate
pip install ansible

vagrant up
ansible-playbook playbook.yml -i inventory -e @vars.yml


1. create Vagrant file: http://jeremievallee.com/2017/01/31/kubernetes-with-vagrant-ansible-kubeadm/
```
nodes = [
  { :hostname => 'kubernetes-master',  :ip => '10.0.2.10', :ram => 4096 },
  { :hostname => 'kubernetes-node1',  :ip => '10.0.2.11', :ram => 2048 },
  { :hostname => 'kubernetes-node2',  :ip => '10.0.2.12', :ram => 2048 },
  { :hostname => 'kubernetes-node3',  :ip => '10.0.2.13', :ram => 2048 }
]

Vagrant.configure("2") do |config|
  nodes.each do |node|
    config.vm.define node[:hostname] do |nodeconfig|
      nodeconfig.vm.box = "bento/ubuntu-16.04";
      nodeconfig.vm.hostname = node[:hostname] + ".box"
      nodeconfig.vm.network :private_network, ip: node[:ip]
      memory = node[:ram] ? node[:ram] : 256;
      nodeconfig.vm.provider :virtualbox do |vb|
        vb.customize [
          "modifyvm", :id,
          "--memory", memory.to_s,
          "--cpus", "4"
        ]
      end
    end
  end
end
```

```console
vagrant up
vagrant status
vagrant ssh kubernetes-master
```

1. Instal kubeadm: https://kubernetes.io/docs/setup/independent/install-kubeadm/

```console
sudo su -
apt-get update
apt-get install -y \
    apt-transport-https \
    ca-certificates \
    curl \
    software-properties-common
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
apt-key fingerprint 0EBFCD88
add-apt-repository \
   "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
   $(lsb_release -cs) \
   stable"
apt-get install -y docker-ce

# install kubectl
curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl
chmod +x ./kubectl
mv ./kubectl /usr/local/bin/kubectl


# install kubelet and kubeadm:
apt-get update && apt-get install -y apt-transport-https
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add -
cat <<EOF >/etc/apt/sources.list.d/kubernetes.list
deb http://apt.kubernetes.io/ kubernetes-xenial main
EOF
apt-get update
# Install docker if you don't have it already.
#apt-get install -y docker-engine
apt-get install -y kubelet kubeadm kubernetes-cni

```

# on Master node:
```
kubeadm init
kubectl taint nodes --all node-role.kubernetes.io/master-
kubectl apply -f http://docs.projectcalico.org/v2.3/getting-started/kubernetes/installation/hosted/kubeadm/1.6/calico.yaml
```
And capture the output from the above command. e.g.
```
You can now join any number of machines by running the following on each node
as root:

  kubeadm join --token 6695a0.97dc4f6169c5567b 10.0.2.15:6443
```
Then run it on the Worker nodes.
