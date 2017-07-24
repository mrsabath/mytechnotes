```
# To list all the available images (in JSON format):
curl $REG:5001/v1/search
# or pretty printed:
curl $REG:5001/v1/search | python -m json.tool

# Modify your local docker configuration to run in insecure registry mode adjust the IP as needed:

docker-machine ssh default "echo $'EXTRA_ARGS=\"--insecure-registry 10.140.132.215:5001\"' | sudo tee -a /var/lib/boot2docker/profile && sudo /etc/init.d/docker restart"

# Tag the image with the $REG location and push it

docker tag mrsabath/web-ms $REG:5001/mrsabath/web-ms
docker push $REG:5001/mrsabath/web-ms

Pull image from the Infrastructure Image Registry

Connect to swarm and list your current images:

export SWARM=<your_swarm_manager>
DOCKER_TLS_VERIFY="" docker -H $SWARM images

Pull the image from Infrastructure Image Registry:

DOCKER_TLS_VERIFY="" docker -H $SWARM pull $REG:5001/mrsabath/web-ms

Once the image is pulled, tag it with a shorter name:

DOCKER_TLS_VERIFY="" docker -H $SWARM tag $REG:5001/mrsabath/web-ms mrsabath/web-ms

Now you can use the new image directly on your swarm:

DOCKER_TLS_VERIFY="" docker -H $SWARM run -d -p 8080:80 -e TEST=my_test --name WEB mrsabath/web-ms
949612b5700bd4607a517d3e6a1eb0195a5d87b14be3da753280e25bf65cce48
```
