# Notes on Dockerfiles

Build an image from a Dockerfile (default is 'PATH/Dockerile')
```
docker build -t $(IMAGE_NAME) .
docker build -t $(IMAGE_NAME) -f Dockerfile1.server
```


To keep the container alive forever:

```
CMD ["sh", "-c", "tail -f /dev/null"]
# or
CMD ["/bin/bash", "-c", "while true; do sleep 10; done;"]
```

To keep the layer thin, but it takes long to build:

```
RUN apt-get update \
 && apt-get -y install --no-install-recommends ca-certificates curl git \
 && curl -sSLO https://storage.googleapis.com/golang/go1.7.3.linux-amd64.tar.gz \
 && tar -C /usr/local -xzf go1.7.3.linux-amd64.tar.gz \
 && rm -f go1.7.3.linux-amd64.tar.gz \
 && export PATH=$PATH:/usr/local/go/bin \
 && export GOPATH=/api-proxy \
 && go get "github.com/golang/glog" \
 && go get "github.com/spf13/viper" \
 && go get "github.com/coreos/etcd/client" \
 && go get "github.com/coreos/pkg/capnslog" \
 && go get "github.com/coreos/etcd/clientv3" \
 && go get "github.ibm.com/fr8r-secure/etcd/encconfig" \
 && go get "github.ibm.com/fr8r-secure/etcd/skeysapiv3" \
 && go get "github.ibm.com/fr8r-secure/etcd/encwrap" \
 && go build -o bin/api-proxy src/api-proxy/api-proxy.go \
 && go build -o bin/etcd_client src/misc/etcdmain/etcd_client.go \
 && rm -rf /usr/local/go \
 && apt-get -y remove git \
 && apt-get clean autoclean \
 && apt-get -y autoremove \
 && rm -rf /var/lib/apt/lists/* \
 && rm -rf ./src/
 ```

Install and cleanup:
```
 RUN wget https://releases.hashicorp.com/vault/1.0.3/vault_1.0.3_linux_amd64.zip && \
    unzip vault_1.0.3_linux_amd64.zip && \
    mv vault /usr/local/bin/ && \
    rm -f vault_1.0.3_linux_amd64.zip
```

Pass default values for env. variables:
```
 # Default values for vault client setup
ARG DEFAULT_VAULT_ADDR="http://vault:8200"
ARG DEFAULT_VAULT_ROLE="demo"
ENV VAULT_ADDR=${DEFAULT_VAULT_ADDR}
ENV VAULT_ROLE=${DEFAULT_VAULT_ROLE}
```
Then when building, you can pass the ARG
```
export VAULT_ADDR="http://myvault:8211"
docker build --build-arg DEFAULT_VAULT_ADDR=${VAULT_ADDR}
docker build --build-arg var_name=${VARIABLE_NAME} (...)
```

Include files in the image:
```
COPY script1.sh script2.py gen-token.py run-server.sh startup_server.sh validate-token.py /usr/local/bin/
```

Run a sanity check on your Dockerfiles using (hadolint)[https://github.com/hadolint/hadolint]

```
brew install hadolint
hadolint <Dockerfile>
hadolint --ignore DL3003 --ignore DL3006 <Dockerfile> # exclude specific rules
hadolint --trusted-registry my-company.com:500 <Dockerfile> # Warn when using untrusted FROM images
```

Docker comes to the rescue to provide an easy way how to run hadolint on most platforms. Just pipe your Dockerfile to docker run:

```
docker run --rm -i hadolint/hadolint < Dockerfile
```
