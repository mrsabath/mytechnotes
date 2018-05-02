# Various aspects of using SSH keys

## Add SSH key to the remote server

```
ssh-copy-id -i ~/.ssh/id_rsa.pub root@IP
ssh 'root@IP' "cat ~/.ssh/authorized_keys"
```

## Various SSH commands
```
# example use:
ssh -i ~/.ssh/key user@server.watson.ibm.com  "whoami;hostname"
scp -i <key> <source_path>/<source_file> user@server:/~/<target_file>

# If you put the above private key in the following location ~/.ssh/id_dsa
# e.g. cp -a ig_eppgr.priv  ~/.ssh/id_dsa
# and ensure that the file is not world readable etc., then
# ssh will use this as your "identity". This is how I set it up on
# grid.watson.ibm.com - and avoids the need for the tedious -i ig_eppgr.priv flag above.

# read the certificate:
openssl x509  -text -in cert.pem

# create SSH key:
# locally:
ssh-keygen -b 2048 -t rsa -f ~/.ssh/radiant -q  -N ""

# on another host:
USER_NAME=
sudo -u $USER_NAME ssh-keygen -b 2048 -t rsa -f /home/$USER_NAME/.ssh/id_rsa  -q -N ""
sudo -u $USER_NAME cp /home/$USER_NAME/.ssh/id_rsa.pub /home/$USER_NAME/.ssh/authorized_keys

$ ssh-keygen -t rsa -C "sabath@us.ibm.com"
Generating public/private rsa key pair.
Enter file in which to save the key (/Users/sabath/.ssh/id_rsa):
Enter passphrase (empty for no passphrase):
Enter same passphrase again:
Your identification has been saved in /Users/sabath/.ssh/id_rsa.
Your public key has been saved in /Users/sabath/.ssh/id_rsa.pub.
The key fingerprint is:
bc:46:xxxxxx
```

## SSH Key injection

```bash
KEY="ssh-rsa AAAABxxxx== me@dyn9002026087.watson.ibm.com"
root@IP "echo $KEY >> ~/.ssh/authorized_keys"
```

## Connect to SSH server using private key

```bash
ssh -i /Users/sabath/.ssh/MSKey2.key -Y root@IP
```

## SSH  suppress prompts

```console
ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -i ~babaYaga-conf/weaver_sample.pem ec2-user@
```

## SSH debug mode (level 3)

```console
ssh -vvv
```

## SSH delays
http://injustfiveminutes.com/2013/03/13/fixing-ssh-login-long-delay/
```
ssh -o GSSAPIAuthentication=no
```

## Dealing with default keys

```bash
# list all the default keys:
ssh-add -l

# add a new default key:
ssh-add ~/.ssh/radiant

# add a new default key permanently in file `~/.ssh/config`:
IdentityFile ~/.ssh/radiant
```

## Execute complex SSH calls
parse the long string, ADMIN_OUT, result of running a command, and find the
last occurrence of ADMIN_KEY=value, using 'rev' function:

```console
ssh -o LogLevel=quiet -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null -i 	~/.vagrant.d/insecure_private_key vagrant@$master_ip -t 'export ADMIN_OUT=$(sudo docker exec api-proxy /api-proxy/create_admin.sh admin1 shard1); export ADMIN_KEY=$(echo $ADMIN_OUT | grep "ADMIN_KEY" | rev | cut -d "=" -f1 | rev);sudo docker exec api-proxy /api-proxy/create_user.sh dev-vbox '$tenant' '$shard' '$master_ip' '$master_ip' $ADMIN_KEY'
```

## Another complex SSH call

```console
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
