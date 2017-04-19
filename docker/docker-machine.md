## install [docker toolbox](https://www.docker.com/docker-toolbox)


## list virtual
```
docker-machine ls
NAME      ACTIVE   DRIVER       STATE     URL                         SWARM
default   *        virtualbox   Running   tcp://192.168.99.100:2376
```

## list active
```
docker-machine active
```

## start docker-machine
```
docker-machine start default
```

## show IP
```
docker-machine ip default
192.168.99.100
```

## setup
```
docker-machine env default
export DOCKER_TLS_VERIFY="1"
export DOCKER_HOST="tcp://192.168.99.100:2376"
export DOCKER_CERT_PATH="/Users/sabath/.docker/machine/machines/default"
export DOCKER_MACHINE_NAME="default"
# Run this command to configure your shell:
# eval "$(docker-machine env default)"
```

## update the docker level:
```
docker-machine upgrade default
```

## re-create docker-machine:
```
docker-machine kill default
docker-machine rm default
docker-machine create default --driver virtualbox
```
## Modify your local docker configuration to run in insecure registry mode adjust the IP as needed:
```
docker-machine ssh default "echo $'EXTRA_ARGS=\"--insecure-registry 10.140.132.215:5001\"' | sudo tee -a /var/lib/boot2docker/profile && sudo /etc/init.d/docker restart"
```

## pull images
```
docker pull mrsabath/web-ms
```
