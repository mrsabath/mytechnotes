## Some Go related notes.


# load dependencies:
# 1) first script:
```bash
function load_dep {
  if [ ! $(basename "$PWD") = "proxy" ]; then
    cd proxy || exit
  fi
  local SETCD_SRC="../../fr8r-secure"

  if [ ! -d "$SETCD_SRC" ]; then
    git clone "git@github.ibm.com:alchemy-containers/fr8r-secure.git" "$SETCD_SRC"
  fi

  # 'setcd_src' directory is expected by Dockerfile, so copy fr8r-secure src there
  local SETCD_LIB=setcd_src
  echo "PWD: $PWD"
  rm -rf $SETCD_LIB
  TARGET="$SETCD_LIB/github.ibm.com/alchemy-containers/"
  mkdir -p "$TARGET"
  cp -r "$SETCD_SRC" "$TARGET"
}
```


# 2) then Dockerfile:
```bash
COPY src/ ./src/
COPY setcd_src/ ./src/
COPY config.toml /api-proxy/config.toml

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
 && go get "github.ibm.com/alchemy-containers/fr8r-secure/etcd/encconfig" \
 && go get "github.ibm.com/alchemy-containers/fr8r-secure/etcd/skeysapiv3" \
 && go get "github.ibm.com/alchemy-containers/fr8r-secure/etcd/encwrap" \
 && go build -o bin/api-proxy src/api-proxy/api-proxy.go \
 && go build -o bin/etcd_client src/misc/etcdmain/etcd_client.go \
 && rm -rf /usr/local/go \
 && apt-get -y remove git \
 && apt-get clean autoclean \
 && apt-get -y autoremove \
 && rm -rf /var/lib/apt/lists/* \
 && rm -rf ./src/
```
