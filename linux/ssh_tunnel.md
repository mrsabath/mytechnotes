```
# start ssh tunnel on local host port 8888 to pvespa:
ssh -C -L 8888:192.168.10.2:8097 ubuntu@pvespa016.pok.ibm.com

# then run command on the local host for this port:
curl -X GET localhost:8888/networks/
```
