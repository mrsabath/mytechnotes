# Swarm stuff:

```
# test OpenVPN
ping 10.140.155.217

# private registry (Paolo)
10.143.129.245:5000

# k8s-swarm-01 (dev, Paolo)
[array, Paolo]
10.143.197.112  k8s-swarm-01-1 // Swarm manager
10.143.197.68   k8s-swarm-01-2
10.143.129.244  k8s-swarm-01-3
10.143.197.75   k8s-swarm-01-4
SWARM=10.143.197.112:2375

# Swarm dev-mon-vizio1
[OpenVPN]
Dev-mon01-swarm01-cont.alchemy.ibm.com	 	10.140.28.132
Dev-mon01-swarm01-host01.alchemy.ibm.com 	10.140.28.156
Dev-mon01-swarm01-host02.alchemy.ibm.com	10.140.28.165
Dev-mon01-swarm01-host03.alchemy.ibm.com	10.140.28.163
Dev-mon01-swarm01-host04.alchemy.ibm.com  10.140.28.171
Dev-mon01-swarm01-host05.alchemy.ibm.com	10.140.28.166
port: 2375, 2379
DOCKER_TLS_VERIFY="" docker -H 10.140.28.132:2375 ps
SWARM=10.140.28.132:2375 # docker
SWARM=10.140.28.132:2379 # swarm
SSH Users:  doron, ezras, paolo, alaa, hai, alek, vinod, dominik1

# Swarm dev-mon-vizio-3 (kuryr)
[OpenVPN]
10.140.181.167    dev-mon01-vizio3-host-01
10.140.181.152    dev-mon01-vizio3-host-02
10.140.181.209    dev-mon01-vizio3-host-03
port: 2379
DOCKER_TLS_VERIFY="" docker -H 10.140.181.167:2379 ps
SWARM=10.140.181.167:2379

# Swarm Mesos
Swarm standalone (cluster-swarm4whisk):
10.143.197.123      - Mesos master/swarm manager
10.143.197.102      - Mesos slave/ swarm node
10.143.197.109      - Mesos slave/ swarm node
10.143.197.117      - Mesos slave/ swarm node
10.143.197.118      - Mesos slave/ swarm node
10.143.197.126      - Mesos slave/ swarm node
Swarm Port: 2375

e.g. docker -H 10.143.197.123:2375 info
Example docker run:
 docker -H 10.143.197.123:2375 run -itd -p 3000:3000 -m 16m busybox httpd -f -p 3000

# Swarm on Mesos (cluster-swarm-mesos4whisk​):
10.143.197.120    - Mesos master/swarm manager
10.143.129.216    - Mesos slave/ swarm node
10.143.129.218    - Mesos slave/ swarm node
10.143.129.222    - Mesos slave/ swarm node
10.143.129.224    - Mesos slave/ swarm node
10.143.197.122    - Mesos slave/ swarm node
Swarm port: 4375

e.g. docker -H 10.143.197.120:4375 info
Example docker run:

docker -H 10.143.197.120:4375 run -itd -p 10000:3000 -m 16m busybox httpd -f -p 3000
Note that with mesos you must specify at least memory constrains as above. And if you expose ports they must be in the 10000-90000 range.

Mesos: port 5050
access mesos UI at: http://10.143.197.120:5050/#/

# Swarm w/ Messos (Ezra by Paolo)
10.143.129.XX

# Swarm DAL (OLD), arrayVPN
/etc/hosts
10.120.39.40    vizio-devswarm-cont m
10.120.39.41    vizio-devswarm-host1 s1
10.120.39.34    vizio-devswarm-host2 s2
10.120.39.61    vizio-devswarm-host3 s3
m - manager
s1, s2, s3

ssh to m
source ip.txt

docker $H ps
docker $H images


cat /tmp/my_cluster
10.120.39.41:2375
10.120.39.34:2375

#this is the command to start swarm the way we want ...
docker $H run -v /tmp:/tmp -p 2377:2377 --rm swarm manage -H tcp://0.0.0.0:2377 file:///tmp/my_cluster

# this is deamon:
docker $H run -d -v /tmp:/tmp -p 2377:2377 swarm manage -H tcp://0.0.0.0:2377 file:///tmp/my_cluster


# starting the new swarm manager
cd /root/gocode/src/github.com/docker/swarm
nohup ./swarm --debug manage -H tcp://0.0.0.0:2379 file:///tmp/my_cluster &
# nohup ./swarm --debug manage -H tcp://0.0.0.0:2379 127.0.0.1:2375 &
echo "Running Swarm Manager on port 2379"



# get swarm details
docker $S info


# now you can run swarm:

docker $S images

docker $S run --net=none -d redis
docker $S run --name=MST1 --expose=80 -d redis

docker $S ps


# running local:
boot2docker up
export DOCKER_HOST=m:2377
Silesia:~ sabath$ docker $S ps
CONTAINER ID        IMAGE               COMMAND                CREATED             STATUS              PORTS               NAMES
87bb6122d144        redis:latest        "/entrypoint.sh redi   8 minutes ago       Up 8 minutes                            vizio-devswarm-host2/high_sinoussi
83c4f0851de7        redis:latest        "/entrypoint.sh redi   9 minutes ago       Up 9 minutes        6379/tcp            vizio-devswarm-host2/hopeful_ptolemy



# restarting docker daemon
service docker stop
vi /etc/default/docker
service docker start

# delete exited containers
# test
docker $S ps -f status=exited -q | xargs -L1 echo docker rm
docker $S ps -f status=exited -q | xargs -L1 docker rm
docker -H s1:2375 ps -a -q | xargs -L1  docker rm

# swarm requests:
# list
curl -H "Content-Type: application/json" -H "Accept: application/json" -H "X-Auth-Project-Id": "999999"  -X GET http://0.0.0.0:2377/containers/json?all=0

# create:
curl -H "Content-Type: application/json" -H "Accept: application/json" -H "X-Auth-Project-Id": "999999" -X POST -d '{"Name": "T1", "Volumes": [], "Image": "redis", "Cmd": [], "WorkingDir": "", "HostConfig": {"Links": []}, "Env": [], "Memory": 256, "CpuShares": 1024, "NumberCpus": 4, "Public_ports": []}' http://0.0.0.0:2377/containers/create

###### New:
curl --data-binary @redis1.json -H "X-Auth-Token:$UserToken" -H "Content-type: application/json" http://m:2379/containers/create

curl -X POST -d '{ "Image": "redis", "Cmd": ["redis-server"],"CpuShares": 0,"ExposedPorts": { "6379/tcp": {}}}' -H "X-Auth-Token:$UserToken" -H "Content-type: application/json" http://m:2379/containers/create

curl -X POST -d '{ "Image": "redis", "Cmd": [],"CpuShares": 0,"ExposedPorts": { "6379/tcp": {}}}' -H "X-Auth-Token:$UserToken" -H "Content-type: application/json" http://m:2379/containers/create

curl -X GET -H "X-Auth-Token:$UserToken" -H "Content-type: application/json" http://m:2379/containers/json?all=1

curl -X DELETE -H "X-Auth-Token:Space1" http://m:2379/containers/7862932c

curl -X POST -d '{}' -H "X-Auth-Token:$UserToken" -H "Content-type: application/json" http://m:2379/containers/32eebec6d74a/start

# list directly on the node
docker -H s1:2375 ps -a

# Swarm commands
cd /Users/sabath/workspace/alchemy_new/swarm
. configure/setup_env.sh
DOCKER_TLS_VERIFY="" docker -H $SWARM --config configure info
DOCKER_TLS_VERIFY="" docker -H $SWARM --config configure ps
DOCKER_TLS_VERIFY="" docker -H $SWARM --config configure images
DOCKER_TLS_VERIFY="" docker -H $SWARM --config configure run hello-world

# Events
DOCKER_TLS_VERIFY="" docker -H $SWARM events --since '1h'
DOCKER_TLS_VERIFY="" docker -H $SWARM events --since '2016-01-01'
curl -X GET -H  "Content-type: application/json" $SWARM/events?since=1452537669

# delete all containers
# last 50
DOCKER_TLS_VERIFY="" docker -H $SWARM rm -f $(DOCKER_TLS_VERIFY="" docker -H $SWARM ps -n=50 -q)
DOCKER_TLS_VERIFY="" docker -H $SWARM rm -f $(DOCKER_TLS_VERIFY="" docker -H $SWARM ps -q)

# recreate swarm agents:
# run this on every docker engine
export engine_ip=
export docker_engine_port=2377
export manager_ip=
export etcd_listen_client_port1=2379
docker run -d --name swarm-node swarm join --advertise=$engine_ip:docker_engine_port etcd://$manager_ip:$etcd_listen_client_port1/swarm
# swarm_manager_port: 2375
# etcd_listen_client_port1: 2379
# etcd_listen_client_port2: 4001
# etcd_advertise_peer_port: 2380
docker run -d --name swarm-node swarm join --advertise=$engine_ip:2377 etcd://10.143.197.112:2379/swarm

DOCKER_TLS_VERIFY="" docker -H $SWARM info
Containers: 3
Images: 9
Role: primary
Strategy: spread
Filters: health, port, dependency, affinity, constraint
Nodes: 3
 k8s-swarm-01-2: 10.143.197.68:2377
  └ Status: Healthy
  └ Containers: 1
  └ Reserved CPUs: 0 / 4
  └ Reserved Memory: 0 B / 4.102 GiB
  └ Labels: executiondriver=native-0.2, kernelversion=4.2.0-25-generic, operatingsystem=Ubuntu 14.04.3 LTS, storagedriver=aufs
 k8s-swarm-01-3: 10.143.129.244:2377
  └ Status: Healthy
  └ Containers: 1
  └ Reserved CPUs: 0 / 4
  └ Reserved Memory: 0 B / 4.102 GiB
  └ Labels: executiondriver=native-0.2, kernelversion=4.2.0-25-generic, operatingsystem=Ubuntu 14.04.3 LTS, storagedriver=aufs
 k8s-swarm-01-4: 10.143.197.75:2377
  └ Status: Healthy
  └ Containers: 1
  └ Reserved CPUs: 0 / 4
  └ Reserved Memory: 0 B / 4.102 GiB
  └ Labels: executiondriver=native-0.2, kernelversion=4.2.0-25-generic, operatingsystem=Ubuntu 14.04.3 LTS, storagedriver=aufs
CPUs: 12
Total Memory: 12.31 GiB
Name: 41491a68c418

# debug swarm, after reboot check the logs at: /var/log/upstart/docker.log


# list containers that belong to groups:
DOCKER_TLS_VERIFY="" docker -H $SWARM ps --filter "label=ibm.alchemy.ng.group.name"
# list containers that belong to a specific group:
DOCKER_TLS_VERIFY="" docker -H $SWARM ps --filter "label=ibm.alchemy.ng.group.name=group4"


# restart swarm agents remotely:
ssh sabath@10.140.155.151 "sudo docker start swarm-node; sudo docker ps --all"
```
