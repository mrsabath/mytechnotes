# Few tricks using SSH 
```
# start ssh tunnel on local host port 8888 to pvespa:
ssh -C -L 8888:192.168.10.2:8097 ubuntu@pvespa016.pok.com

# then run command on the local host for this port:
curl -X GET localhost:8888/networks/
```
## execute complex SSH call
```bash
# parse the long string, ADMIN_OUT, result of running a command, and find the
# last occurrence of ADMIN_KEY=value, using 'rev' function:
ssh -o LogLevel=quiet -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null -i 	~/.vagrant.d/insecure_private_key vagrant@$master_ip -t 'export ADMIN_OUT=$(sudo docker exec api-proxy /api-proxy/create_admin.sh admin1 shard1); export ADMIN_KEY=$(echo $ADMIN_OUT | grep "ADMIN_KEY" | rev | cut -d "=" -f1 | rev);sudo docker exec api-proxy /api-proxy/create_user.sh dev-vbox '$tenant' '$shard' '$master_ip' '$master_ip' $ADMIN_KEY'
```
