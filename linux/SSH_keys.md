# SSH Key management

## create SSH key:
```console
# locally:
ssh-keygen -b 2048 -t rsa -f ~/.ssh/radiant -q  -N ""

# on another host:
USER_NAME=
sudo -u $USER_NAME ssh-keygen -b 2048 -t rsa -f /home/$USER_NAME/.ssh/id_rsa  -q -N ""
sudo -u $USER_NAME cp /home/$USER_NAME/.ssh/id_rsa.pub /home/$USER_NAME/.ssh/authorized_keys
```

## Add SSH key to the remote server
```
ssh-keygen -b 2048 -t rsa -f ~/.ssh/id_rsa -q  -N ""
ssh-copy-id -i ~/.ssh/id_rsa.pub root@IP
```

## Add and list the default keys:
```
ssh-add ~/.ssh/MS99
ssh-add -l
2048 SHA256:LHiW/yxxxxxxxxxxxxx/2QLZZO8hg8jKnk /Users/sabath/.ssh/id_rsa (RSA)
2048 SHA256:OLQR3xxxxxxxxxxxxxM982AF7nRFXZCGRs /Users/sabath/.ssh/MS99 (RSA)
```

## example use:
```
ssh -i ~/.ssh/key user@server.watson.ibm.com  "whoami;hostname"
scp -i <key> <source_path>/<source_file> user@server:/~/<target_file>
```

If you put the above private key in the following location ~/.ssh/id_dsa
e.g. cp -a ig_eppgr.priv  ~/.ssh/id_dsa
and ensure that the file is not world readable etc., then
ssh will use this as your "identity". This is how I set it up on
grid.watson.ibm.com - and avoids the need for the tedious -i ig_eppgr.priv flag above.

## read the certificate:
```
openssl x509  -text -in cert.pem
```

## The entire flow
```console
$ ssh-keygen -t rsa -C "sabath@us.ibm.com"
Generating public/private rsa key pair.
Enter file in which to save the key (/Users/sabath/.ssh/id_rsa):
Enter passphrase (empty for no passphrase):
Enter same passphrase again:
Your identification has been saved in /Users/sabath/.ssh/id_rsa.
Your public key has been saved in /Users/sabath/.ssh/id_rsa.pub.
The key fingerprint is:
bc:46:c31:9b:db:29:0e:67:04
```

## SSH Key injection
``
KEY="ssh-rsa AAAA....FMozuo0tr93HIsfb8gcw9R6JUZvnmQ== me@dyn9002026087.watson.ibm.com"
root@s1 "echo $KEY >> ~/.ssh/authorized_keys"
```

## SSH suppress prompts
```
ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -i ~babaYaga-conf/weaver_sample.pem ec2-user@

# SSH debug mode (level 3)
ssh -vvv

# SSH delays:
http://injustfiveminutes.com/2013/03/13/fixing-ssh-login-long-delay/
ssh -o GSSAPIAuthentication=no


# list all the default keys:
ssh-add -l

# add a new default key:
ssh-add ~/.ssh/radiant

# add a new default key permanently in file `~/.ssh/config`:
IdentityFile ~/.ssh/radiant

# execute complex SSH call
# parse the long string, ADMIN_OUT, result of running a command, and find the
# last occurrence of ADMIN_KEY=value, using 'rev' function:
ssh -o LogLevel=quiet -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null -i 	~/.vagrant.d/insecure_private_key vagrant@$master_ip -t 'export ADMIN_OUT=$(sudo docker exec api-proxy /api-proxy/create_admin.sh admin1 shard1); export ADMIN_KEY=$(echo $ADMIN_OUT | grep "ADMIN_KEY" | rev | cut -d "=" -f1 | rev);sudo docker exec api-proxy /api-proxy/create_user.sh dev-vbox '$tenant' '$shard' '$master_ip' '$master_ip' $ADMIN_KEY'


# Add SSH key remotely:
ssh root@9.12.249.151 "echo ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQDUcM7dJ2e4T4KIZNxpToVaHeeKICjHU36vVqNrtb6KrWp8tL2PQ4bHQLh8esc8HyA4RiO1vsKjZBlrMbnzSQdW0jMh/JMsgls0YHhM6E0/sRuCrtGPlZrO1216X62YtSgnov8vjgXGwogYAyXrlxgACf/uMi7vVCuOWJqDNxSKphFYn6kkGP9bzu3Hmkkr6BPV8oZ3QUjRoZ1gTlLS1cR9KlrwAbVSy5N/Y98fs9vvEpAjOKY5K+fkF5a3YfqtE0ghFicWMC9gkGl6/QIDCcFvad63VK9hT1fnqHrclJLmgvThrfWws6c0RcPFlzPhVcgoXJ4EUxI3q9zk+l0t/E3P darroyo@us.ibm.com >> /root/.ssh/authorized_keys; cat /root/.ssh/authorized_keys"


# complex SSH call:
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
