## Vagrant
Make sure you have [`virtualbox`](https://www.virtualbox.org/wiki/Downloads)
installed on your machine (tested on versions >= 5.1.26r117224).

Make sure you have [`vagrant`](https://www.vagrantup.com/downloads.html)
installed (tested on 1.8 and 2.0).

Note: check if the `hostmanager` plug in is installed using:

`vagrant plugin install vagrant-hostmanager`

Even if it is not installed, the `vagrant up` command below will install it for you.



## basic VirtualBox commands
```
vboxmanage list vms
vboxmanage showvminfo $UUI
```
## basic Vagrant commands

### start VMs
```
vagrant up
vagrant up master-node
```

### get status of the VMs
```
vagrant status
```

### get inside the VM
```
vagrant ssh master-node
```

### shutdown all the VMs
```
vagrant halt
vagrant halt master-node
```


### list all the IPs for the VMs
```
vagrant ssh radiant3 -c 'ip -4 addr' | grep inet
```

### get IPs of the vms
```
vagrant status
(cd examples/vagrant; vagrant ssh $vmname -c 'ip -4 addr' | grep inet)
```

### destroy VMs
```
vagrant destroy -f api-proxy client installer
vagrant status
```

### create a vagrant snapshot
```
vagrant snapshot save shard1-master1 k8s-master
vagrant snapshot save shard1-worker1 k8s-worker1
vagrant snapshot list
vagrant snapshot restore k8s-master
vagrant snapshot delete k8s-master
```

### copy file from host to vagrant VM
```
 scp -o StrictHostKeyChecking=no -i ~/.vagrant.d/insecure_private_key hyperkube-amd64 vagrant@192.168.10.3:~
```

### sample call to vagrant VM
```
 ssh -o LogLevel=quiet \
      -o StrictHostKeyChecking=no \
      -o UserKnownHostsFile=/dev/null \
      -i ~/.vagrant.d/insecure_private_key vagrant@$client_ip \
      /bin/bash <<EOF
output=\$(curl -w "%{http_code}" -XPOST -s -k --data "node=$shard_ip" --header "X-Tls-Client-Dn: $ADMIN_KEY" "https://$api_proxy_ip:6969/admin/api/v1/shards/$shard/users/$tenant" 2>&1)
httpCode=\$(tail -n 1 <<< "\$output")
outputSize=\$(wc -l <<< "\$output")
output=\$(head -n \$((\$outputSize - 1)) <<< "\$output")
if [[ "\$httpCode" != "200" ]]; then
  echo "ERROR \$httpCode: Could not initialize kubernetes tenant"
  echo "\$output"
  exit 1
fi

mkdir -p "$tenant"
echo "\$output" > "$tenant/kube-config"
EOF
```
