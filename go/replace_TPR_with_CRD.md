## Good examples from Paolo:
https://github.ibm.com/seed/seedcr
https://github.ibm.com/seed/seedcrs

Make sure the GOPATH has only one entry, otherwise this error:
```
silesia:crsample sabath$ hack/update-codegen.sh
hack/..
Generating deepcopy funcs
vendor/k8s.io/code-generator/generate-groups.sh: line 65: /Users/sabath/workspace/go-work:/Users/sabath/workspace/go-tools/bin/deepcopy-gen: No such file or directory
silesia:crsample sabath$ pwd
/Users/sabath/workspace/go-work/src/github.ibm.com/seed/crsample
silesia:crsample sabath$ echo $GOPATH
/Users/sabath/workspace/go-work:/Users/sabath/workspace/go-tools
```

## Load the dependencies
```
silesia:crsample sabath$ GOPATH=/Users/sabath/workspace/go-work

silesia:crsample sabath$ rm -rf vendor/*
silesia:crsample sabath$ glide cache-clear
[INFO]	Glide cache has been cleared.

silesia:crsample sabath$ glide up -v
```
## Regenerating deep copy

Deep copy helpers are required for the data schema of your custom resource. A data copy helper (of the form zz_generated.deepcopy.go) already exists for this sample app under apis/cr/v1, however if change the schema you will need to regenerate the helper.

To regenerate the helper, simply run the following script from the root of the project:

```
hack/update-codegen.sh
```
If all goes well, you should see a message like the one below:

```
I1112 22:05:57.707809   16615 execute.go:67] Assembling file "$GOPATH/src/github.ibm.com/seed/crsample/a
pis/cr/v1/zz_generated.deepcopy.go"
```
where GOPATH depends on your environment.

## Rebuild the project

```
cd cmd
#
cd executor
go build -o executor
```

## Running
```
./executor --kubeconfig=$KUBECONFIG

# from another console:
cd ../crd
kubectl create -f inventory.yaml
kubectl replace -f inventory-44L.yaml
```


## Fixing the compilation problems
```
create a new project e.g. kompass/update-service-test and move all the
files that compile there.
pkg, glide.yml etc
remove glide.lock
run glide up -v
move the 'vendors' back to original project
run the hack/update-codegen.sh script
