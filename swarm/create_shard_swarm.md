## Add new Shard

Follow the steps that are using the old CCSAPIs:

```bash
pyenv activate venv277
source ~/workspace/alchemy-1337/environment-dev-mon01/cloudantrc
export PYTHONPATH=~/workspace/alchemy-OLD/ccsapi/:~/workspace/alchemy-OLD/containers-credentialstore-python/:~/workspace/alchemy-OLD/containers-service-utils/
cd ~/workspace/alchemy-OLD/ccsapi/lib/ts
python cloudant_manage.py
```

if you get error:
```
File "/Users/sabath/workspace/alchemy_new/api/lib/auth/tokens.py", line 7, in <module>
  import cloud.CloudBackend as CloudBackend
AttributeError: 'module' object has no attribute 'CloudBackend'
```
disable in lib/auth/tokens.py:
`#import cloud.CloudBackend as CloudBackend`

Execute steps below:
```
python cloudant_manage.py
Using cloudant credentials/setup from the environment variables for user alchemycontainers(dev-mon01-tenants,dev-mon01-cloud_instances)
{'https://10.140.174.27:5000/v2.0': u'dev-mon01-kraken2'}
Enter a number:
1) Show all state
2) Manage tenants
3) Manage shards (cloud instance)
4) Get shard (cloud instance) for tenant
5) Create collections from ENV vars
6) Inside-Outside address mapping
7) Manage address mapping
8) Manage routers
9) Manage subnets
10) Manage ports
0) Exit
>
3
Pick an operation ['create', 'state','auth']
create
Enter new cloud id (e.g. cloud1):
swarm1
Enter the backend type (e.g. openstack, swarm)
swarm
Enter Auth URL for Cloud (sample: http://108.168.168.10:5000/v2.0):
#### USE JUST IP ######
10.140.28.132
Enter ADMIN Tenant Name:
admin
Enter ADMIN User Name:
admin
Enter ADMIN Password:
admin
Enter initial state [active, oos]:
oos
Enter capacity total [integer value]:
1000
Enter capacity used [integer value]:
0
Enter a number:
1) Show all state
2) Manage tenants
3) Manage shards (cloud instance)
4) Get shard (cloud instance) for tenant
5) Create collections from ENV vars
6) Inside-Outside address mapping
7) Manage address mapping
8) Manage routers
9) Manage subnets
10) Manage ports
0) Exit
>
0
Exit selected. Goodbye
```

The data is stored at `dev-mon01-cloud_instances` database
NOTE: need to add:
`"cloud_backend": "swarm",`
to the DB record.

## Assign space to the new shard
```bash
export API_HOST=http://10.140.155.229:8081
source  ~/workspace/alchemy-1337/environment-dev-mon01/deploy/envrc.dev-mon01
# list all the shards:
curl -u $API_USER:$API_PASSWORD -X GET -H "Content-Type: application/json" $API_HOST/v3/admin/tenant_sharder/clouds
```

Ask user to run:
```bash
cf ic login --host containers-api-dev.stage1.ng.bluemix.net
cf ic info
```
then continue running as admin:
```bash
space=
curl -v -u $API_USER:$API_PASSWORD -X GET $API_HOST/v3/admin/space/$space
curl -v -u $API_USER:$API_PASSWORD -X GET $API_HOST/v3/admin/tenant_sharder/override/$space
curl -v -u $API_USER:$API_PASSWORD -X PUT  -H "Content-Type: application/json" -d '{"space_uuid": "xxx", "cloud_id": "dev-mon01-radiant1"}' $API_HOST/v3/admin/tenant_sharder/override
curl -v -u $API_USER:$API_PASSWORD -X PUT  -H "Content-Type: application/json" -d '{"space_uuid": "xxx", "cloud_id": "dev-mon01-swarm1"}' $API_HOST/v3/admin/tenant_sharder/override
curl -v -u $API_USER:$API_PASSWORD -X GET $API_HOST/v3/admin/tenant_sharder/override/$space
curl -v -u $API_USER:$API_PASSWORD -X DELETE $API_HOST/v3/admin/space/$space/$space-instance
# as a user, execute:
cf ic login --host containers-api-dev.stage1.ng.bluemix.net
# this is when the doc is recreated.
# run `cf ic info` to validate then:
curl -v -u $API_USER:$API_PASSWORD -X GET $API_HOST/v3/admin/space/$space
```

Alternative solution:
check current association:
```
space=xxx
curl -u $API_USER:$API_PASSWORD -X GET -H "Content-Type: application/json" $API_HOST/v3/admin/space/$space
curl -v -u $API_USER:$API_PASSWORD -X GET $API_HOST/v3/admin/tenant_sharder/override/$space

# use Jenkins to unprovision space:
http://alchemy.hursley.ibm.com:8080/job/Containers-CCSAPI/view/All/job/space-sharder-unprovisioner/
# check again the space via curl

# execute 'cf ic info' this will prime the account,
# check again the space via curl
# go directl to Cloudant and check the DB:
### in cloudant
<env>-credential-service
design/users index, by Field
search:   space_uuid:<space_id>

## get the API key from curl and run:
curl -v  -H "X-Tls-Client-Dn: /CN=client/emailAddress=dd231978436708483f29de364d17d928b656ce3ffa4ec8d0" -X GET $API_HOST/v3/admin/getHost/NoneContainer
```

CCSAPI logs:
```bash
ls -larth /var/lib/docker/containers/*/*.log | grep $docker_ID


# login to the API_HOST
ice login -u sabath@us.ibm.com -o sabath@us.ibm.com -a api.stage1.ng.bluemix.net -H $API_HOST -R $REG
ice info # to get spaceid
```

space=xxx
curl -u $API_USER:$API_PASSWORD -X GET -H "Content-Type: application/json" $API_HOST/v3/admin/space/$space
space_uuid:"5fa0581e-8a39-42d9-bacc-8cb5784eddc3"
