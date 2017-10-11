# Docker info:

## docker documentation:
[docker doc](https://docs.docker.com/reference/api/docker_remote_api_v1.15/#inspect-a-container)

## keep container running
```
docker run -d centos tail -f /dev/null
```

## run command `date` inside the container and exit:
```
docker run ubuntu:14.04 date --help
docker run ubuntu:latest date
docker run ubuntu:14.04 ls /
```

## create docker and get inside:
```
docker run -it ubuntu:14.04 /bin/bash
```


## running docker APIs:
## inspect container
```
echo -e "GET /containers/5e7ed147/json HTTP/1.1\r\n" | sudo nc -U /var/run/docker.sock
```

## list containers:
```
echo -e "GET /containers/json HTTP/1.1\r\n" | sudo nc -U /var/run/docker.sock
HTTP/1.1 200 OK
Content-Type: application/json
Date: Mon, 03 Nov 2014 16:32:24 GMT
Content-Length: 406

[{"Command":"/bin/bash","Created":1414681947,"Id":"5e7ed147c897347a871941297aa1953d1ac30fc73274c5e8bcae259774fdbd45","Image":"ubuntu:latest","Names":["/trusting_fermi"],"Ports":[],"Status":"Up 4 days"}
,{"Command":"/bin/bash","Created":1414681902,"Id":"7fecef45bd78c59e2351c9669c8d0525cbc08711419377a93b869a65a0c42cb9","Image":"ubuntu:latest","Names":["/tender_stallman"],"Ports":[],"Status":"Up 4 days"}
```

## create a container
```
docker run -p HOST:CONT -d --name test -e "KEY=value" image
```

## running docker CLI
```
docker ps
docker inspect <id_prefix>
```

## tail the console
```
docker logs -f <id_prefix>
```

## get inside the docker
```
docker exec -it <id> /bin/sh
```

## docker logs:
```
docker ps
```
## get id
```
ls -larth /var/lib/docker/containers/*/*.log | grep $ID
```

## create and add hosts to /etc/hosts
```
docker run --add-host host1:1.2.3.4 --add-host host2:4.3.2.1 -it image  /bin/bash
docker run -e "CLOUDANT_ENV_PREFIX=dev-" --restart="always" -t -p 8081:8081 --name="ccsapinum" -d image
```

## run docker client without TLS:
```
DOCKER_TLS_VERIFY="" docker -H 10.140.181.167:2379 ps
# or:
unset DOCKER_TLS_VERIFY
# or:
export DOCKER_TLS_VERIFY=""
```

## remove all containers
```
docker rm $(docker ps -qa)
```
## remove all recent containers
```
docker rm $(docker ps -n=10 -q)
```

## push image to docker hub
```
docker login
docker tag web-ms docker.io/mrsabath/web-ms
docker push docker.io/mrsabath/web-ms
```

##Get to docker files from the host:
```
#### version 1.10.x:
container_long_id=$1
cat `find /var/lib/docker -name  "${container_long_id}*" | grep mounts`/init-id | cut -d'-' -f1

### version 1.9.1:
container_long_id=$1
cat `find /var/lib/docker -name  "${container_long_id}*" | grep mounts`/init-id | cut -d'-' -f1​
```
## for aufs  take the output of the above command and join with the prefix
```
/var/lib/docker/aufs/mnt/
```
## to the the root file system
24414060544
24414650368

## Image Cleanup (remove images)
```
# delete all the temp images:
docker images --no-trunc| grep none | awk '{print $3}' | xargs docker rmi
# delete all unused images:
docker rmi $(docker images --filter dangling=true -q 2>/dev/null) 2>/dev/null
# delete exited containers
docker rm -v $(docker ps --filter status=exited -q 2>/dev/null) 2>/dev/null
# delete all images for specific name
docker rmi $(docker images | grep api-proxy | awk '{ print $3}' )
```


## docker exec:
```
docker exec [options] CONTAINER COMMAND
  -d, --detach        # run in background
  -i, --interactive   # stdin
  -t, --tty           # interactive
```
## get inside the container:
```
docker exec -it <container_id> /bin/bash
```

```
$ docker exec app_web_1 tail logs/development.log
$ docker exec -t -i app_web_1 rails c
docker exec remoteabac /opt/kubernetes/ruser --authorization-policy-file=etcd@http://10.140.171.205:4001/abac-policy --type=add \
--user=<apikey> --namespace=<spaceid>
docker exec etcd etcdctl get /abac-policy`
```

## setup custom registry:
```
vi /etc/default/docker

DOCKER_OPTS='--insecure-registry 9.0.0.0/8 --insecure-registry 10.0.0.0/8'
service docker restart
```
## move image between hosts:
```
# You will need to save the docker image as a tar file:
docker save -o <save image to path> <image name>

# Then copy your image to a new system with regular file transfer tools such as cp or scp. After that you will have to load the image into docker:
docker load -i <path to image tar file>
```

## debug
```
sudo su
vi /etc/default/docker
# add -D
DOCKER_OPTS="$DOCKER_OPTS -D -H tcp://0.0.0.0:2375 -H unix:///var/run/docker.sock"
service docker restart
# get docker version
cat /etc/lsb-release
# get logs:
```
