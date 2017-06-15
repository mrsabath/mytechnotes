# Install Kubeadm locally, using Vagrant. 1 master + 3 workers

1. create Vagrant file: http://jeremievallee.com/2017/01/31/kubernetes-with-vagrant-ansible-kubeadm/
```
nodes = [
  { :hostname => 'kubernetes-master',  :ip => '172.16.66.2', :ram => 4096 },
  { :hostname => 'kubernetes-node1',  :ip => '172.16.66.3', :ram => 2048 },
  { :hostname => 'kubernetes-node2',  :ip => '172.16.66.4', :ram => 2048 },
  { :hostname => 'kubernetes-node3',  :ip => '172.16.66.5', :ram => 2048 }
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
sudo apt-get update
sudo apt-get install -y \
    apt-transport-https \
    ca-certificates \
    curl \
    software-properties-common
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo apt-key add -
sudo apt-key fingerprint 0EBFCD88
sudo add-apt-repository \
   "deb [arch=amd64] https://download.docker.com/linux/ubuntu \
   $(lsb_release -cs) \
   stable"
sudo apt-get install -y docker-ce

# install kubectl
curl -LO https://storage.googleapis.com/kubernetes-release/release/$(curl -s https://storage.googleapis.com/kubernetes-release/release/stable.txt)/bin/linux/amd64/kubectl
chmod +x ./kubectl
sudo mv ./kubectl /usr/local/bin/kubectl


# install kubelet and kubeadm:
sudo su -
apt-get update && apt-get install -y apt-transport-https
curl -s https://packages.cloud.google.com/apt/doc/apt-key.gpg | apt-key add -
cat <<EOF >/etc/apt/sources.list.d/kubernetes.list
deb http://apt.kubernetes.io/ kubernetes-xenial main
EOF
apt-get update
# Install docker if you don't have it already.
apt-get install -y docker-engine
apt-get install -y kubelet kubeadm kubernetes-cni

```

# on Master node:
```
kubeadm init
```
And capture the output from the above command. e.g.
```
You can now join any number of machines by running the following on each node
as root:

  kubeadm join --token 6695a0.97dc4f6169c5567b 10.0.2.15:6443
```s
Then run it on the Worker nodes.
