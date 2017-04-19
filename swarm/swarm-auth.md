## Testing Swarm-auth calls:

Setup config file:

cat ~/projects/next-gen/2a39c5af-d503-4d7b-884f-eb8fcad1777f/config.json
```json
{
    "HttpHeaders": {
          "X-Auth-TenantId": "2a39c5af-d503-4d7b-884f-eb8fcad1777f"
    }
}
```

Run docker commands directly:

```bash
export SWARM=10.140.171.229:4375
DOCKER_TLS_VERIFY="" docker -H $SWARM --config ~/projects/next-gen/2a39c5af-d503-4d7b-884f-eb8fcad1777f ps

DOCKER_TLS_VERIFY="" docker -H $SWARM --config ~/projects/next-gen/2a39c5af-d503-4d7b-884f-eb8fcad1777f run -d --name test13 -m 128m --env test=1 10.140.132.215:5001/mrsabath/web-ms
```

## Check Mesos cluster:
http://10.140.171.229:5050/


Running swarm against NGINX:

Use:
X-Tls-Client-Dn header


```
docker --tlsverify
--tlscacert=/Users/sabath/.ice/certs/containers-api-dev.stage1.ng.bluemix.net/2a39c5af-d503-4d7b-884f-eb8fcad1777f/ca.pem \
--tlscert=/Users/sabath/.ice/certs/containers-api-dev.stage1.ng.bluemix.net/2a39c5af-d503-4d7b-884f-eb8fcad1777f/cert.pem \
--tlskey=/Users/sabath/.ice/certs/containers-api-dev.stage1.ng.bluemix.net/2a39c5af-d503-4d7b-884f-eb8fcad1777f/key.pem \
-H containers-api-dev.stage1.ng.bluemix.net:9443 \
--config ~/projects/next-gen/2a39c5af-d503-4d7b-884f-eb8fcad1777f/ ps

run -d --name test13 -m 128m --env test=1 mrsabath/web-ms
```

Running swarm-auth directly via curl:
```
curl -H "X-Auth-TenantId: 2a39c5af-d503-4d7b-884f-eb8fcad1777f" -X GET $SWARM/v1.21/containers/8b1ffda8db6c/json
```


DOCKER_TLS_VERIFY="" docker run -d --name nonet --net none -m 128m mrsabath/web-ms
DOCKER_TLS_VERIFY="" docker run -d --name net1 --net default -m 128m mrsabath/web-ms
