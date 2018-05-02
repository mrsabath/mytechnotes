
To add permission for the vagrant user
```console
ssh -o UserKnownHostsFile=/dev/null \
    -o StrictHostKeyChecking=no \
    -i ~/.vagrant.d/insecure_private_key \
    vagrant@192.168.10.4 "sudo chmod -R o+rx /root"
```

To copy the files locally:
```console
scp -r \
    -o UserKnownHostsFile=/dev/null \
    -o StrictHostKeyChecking=no \
    -i ~/.vagrant.d/insecure_private_key \
    vagrant@192.168.10.4:/root/.fr8r ~/
```

if you want to copy the files to the client VM you need to copy `~/.vagrant.d/insecure_private_key` to the client
