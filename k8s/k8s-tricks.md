# Few K8s Useful Tricks:

## Connect locally to the running remote server:
```
kubectl port-forward my-server-pod-0 10000:10000 3000:3000
then http://localhost:3000
```

## Compile go-code inside the linux container, to get proper arch:
```
# cd to the directory with GO src code, then:
docker run --rm -v "$PWD":/usr/src/myapp -w /usr/src/myapp -e GOOS=linux -e GOARCH=amd64 golang:1.14 go build -v -ldflags '-s -w -linkmode external -extldflags "-static"' -o executable .
# this will create "excecutable" specific to the arch provided in docker image put place it locally 
```
