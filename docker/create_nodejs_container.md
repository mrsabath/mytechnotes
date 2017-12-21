# Create a sample Node.js container.

## install npm
```
brew install npm
```
Follow steps [here](https://nodejs.org/en/docs/guides/nodejs-docker-webapp/)
Mike's github https://github.ibm.com/msava/k8-test/tree/master
My playground: /Users/sabath/projects/node_js/k8-test/

Mike's note:
```
Here is a repo for Node.js:  https://github.ibm.com/msava/k8-test/tree/master
you need to cd to each `/app` and `/service` and run `npm install` first to load up necessary modules (after cloning of course) (edited)
you need to put a hidden .env file into the tree locally after cloning. (edited)
for /app, create a file `.env` and paste:
PORT=3001
NODE_ENV=development
SERVICE_URL=http://localhost:3002

for /service same file but paste:
PORT=3002
NODE_ENV=development

`npm start` should work for each.  http://localhost:3002 should bring up a service with json returning.  http://localhost:3001/service is the app calling that service url..
```

516  vi .dockerignore
 517  docker build -t mrsabath/node-web-app .
 518  docker images
 519  docker run -p 49160:8080 -d mrsabath/node-web-app
 520  docker ps
 521  docker logs f59bcfb5c3fd
 522  curl -i localhost:49160
 523  ls
 524  pwd
 525  git clone git@github.ibm.com:msava/k8-test.git
 526  cd k8-test/
 527  ls
 528  cd app
 529  mkdir .env
 530  iPORT=3001
 531  NODE_ENV=development
 532  rm -rf .env/
 533  vi .env
 534  npm install
 535  cd ..
 536  vi service/
 537  ls
 538  cd service/
 539  ls
 540  vi .env
 541  npm install
 542  ls
 543  cat package.json
 544  cat service.js
 545  vi Dockerfile
 546  vi .dockerignore
 547  docker build -t mrsabath/node-test-service
 548  docker build -t mrsabath/node-test-service .
 549  docker run -p 49170:3002
 550  docker run -p 49170:3002 -d mrsabath/node-test-service
 551  docker ps
 552  docker log ab
 553  docker logs ab
 554  curl -i localhost:49170
 555  docker ps
 556  docker exec -it ab /bin/bash
 557  cd ..
 558  ls
 559  cd app
 560  cp ../service/Dockerfile .
 561  cp ../service/.dockerignore
 562  cp ../service/.dockerignore .
 563  vi Dockerfile
 564  docker build -t mrsabath/node-test-app .
 565  docker ps
 566  docker run -p 49171:3001 -d mrsabath/node-test-app
 567  docker ps
 568  docker logs -f b75

# build image
docker build -t mrsabath/node-test-service .
# test locally
docker run -p 49171:3007 -e PORT=3007 -e NODE_ENV=development --name service -d mrsabath/node-test-service
curl -i localhost:49171
{"title":"Express"}
# push the image to the registry:
docker push mrsabath/node-test-service


Silesia:service sabath$ kubectl create -f webService-depl.yaml
poddisruptionbudget "web-ms-budget" created
deployment "web-ms-deployment" created
service "web-ms-service" created
Silesia:service sabath$ kubectl get po
NAME                                 READY     STATUS        RESTARTS   AGE
web-ms-deployment-1001510701-5682p   0/1       Running       0          5s
web-ms-deployment-1001510701-7t9b7   0/1       Running       0          5s
web-ms-deployment-1001510701-dzj2p   0/1       Running       0          5s
web-ms-deployment-1001510701-nh3b9   0/1       Terminating   0          2m
Silesia:service sabath$ kubectl get po
NAME                                 READY     STATUS    RESTARTS   AGE
web-ms-deployment-1001510701-5682p   0/1       Running   0          10s
web-ms-deployment-1001510701-7t9b7   0/1       Running   0          10s
web-ms-deployment-1001510701-dzj2p   0/1       Running   0          10s
Silesia:service sabath$ kubectl get service
NAME             CLUSTER-IP       EXTERNAL-IP   PORT(S)          AGE
kubernetes       172.21.0.1       <none>        443/TCP          39d
web-ms-service   172.21.184.160   <nodes>       3005:31005/TCP   20s
Silesia:service sabath$ curl -l 172.21.184.160:31005
^C
Silesia:service sabath$ bx cs cluster-get kompass-dev
Retrieving cluster kompass-dev...
OK

Name:			kompass-dev
ID:			0db57ad77d95406e9ef2ccbfada9503e
State:			normal
Created:		2017-11-10T15:39:43+0000
Datacenter:		dal12
Master URL:		https://169.47.70.10:30907
Ingress subdomain:	kompass-dev.us-south.containers.mybluemix.net
Ingress secret:		kompass-dev
Workers:		4
Version:		1.7.4_1504
Silesia:service sabath$ curl -l kompass-dev.us-south.containers.mybluemix.net:31005
{"title":"Express"}

# cd to app
docker build -t mrsabath/node-test-app .
# test locally
Silesia:app sabath$ docker run -p 49172:3001 -e PORT=3001 -e NODE_ENV=development -e SERVICE_URL=http://kompass-dev.us-south.containers.mybluemix.net:31005 --name app -d mrsabath/node-test-app
5a805d2f226c1539da8ed9a269fccdd3a42d7bc2887f5fd52e0b1d7356e86363
Silesia:app sabath$ docker ps
CONTAINER ID        IMAGE                        COMMAND             CREATED             STATUS              PORTS                               NAMES
5a805d2f226c        mrsabath/node-test-app       "npm start"         4 seconds ago       Up 3 seconds        0.0.0.0:49172->3001/tcp             app
Silesia:app sabath$ curl -i localhost:49172
HTTP/1.1 200 OK
X-Powered-By: Express
Content-Type: text/html; charset=utf-8
Content-Length: 170
ETag: W/"aa-SNfgj6aecdqLGkiTQbf9lQ"
Date: Wed, 20 Dec 2017 15:05:48 GMT
Connection: keep-alive

<!DOCTYPE html><html><head><title>Express</title><link rel="stylesheet" href="/stylesheets/style.css"></head><body><h1>Express</h1><p>Welcome to Express</p></body></html>Silesia:app sabath$

## Final testing:
```
silesia:crd sabath$ curl -i http://kompass-dev.us-south.containers.mybluemix.net:31001
HTTP/1.1 200 OK
X-Powered-By: Express
Content-Type: text/html; charset=utf-8
Content-Length: 170
ETag: W/"aa-SNfgj6aecdqLGkiTQbf9lQ"
Date: Thu, 21 Dec 2017 13:46:12 GMT
Connection: keep-alive

<!DOCTYPE html><html><head><title>Express</title><link rel="stylesheet" href="/stylesheets/style.css"></head><body><h1>Express</h1><p>Welcome to Express</p></body></html>silesia:crd sabath$
silesia:crd sabath$
silesia:crd sabath$ curl -i http://kompass-dev.us-south.containers.mybluemix.net:31001/service
HTTP/1.1 200 OK
X-Powered-By: Express
Content-Type: application/json; charset=utf-8
Content-Length: 19
ETag: W/"13-T/cRC9ueNuYVFeL2rMyC1w"
Date: Thu, 21 Dec 2017 13:46:17 GMT
Connection: keep-alive

{"title":"Express"}
```
