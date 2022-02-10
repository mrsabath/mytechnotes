# Using Podman instead of Docker Desktop

Install
```
brew upgrade podman
podman machine init --cpus 2 --disk-size 20 --memory 4096
podman machine start
podman info
```

if errors
```
podman system connection list
Name                         Identity                                   URI
podman-machine-default*      /Users/sabath/.ssh/podman-machine-default  ssh://core@localhost:59393/run/user/1000/podman/podman.sock
podman-machine-default-root  /Users/sabath/.ssh/podman-machine-default  ssh://root@localhost:59393/run/podman/podman.sock
```

Setup Env. variables, replace "localhost" with "127.0.0.1"
```
export CONTAINER_SSHKEY=$HOME/.ssh/podman-machine-default
export CONTAINER_HOST=ssh://core@127.0.0.1:59393/run/user/1000/podman/podman.sock
```

Test it:
```
podman version
podman info
```

when all good, setup the alias,
add to `~/.bash_profile`
```
alias docker="/usr/local/bin/podman"
```
Now you can use `docker` commands.

## Image management
To login to docker.io registry,
use CLI/API Token created on https://hub.docker.com/settings/security

```
# interactive:
podman login docker.io -u mrsabath
Password: <API Token>
# or all in one:
podman login docker.io -u mrsabath -p <API Token>
```

Once logged-in successfully, search or pull image:
```
podman search docker.io/tsidentity
podman pull docker.io/tsidentity/tornjak-spire-server:new
```

## Debugging
```
podman pull docker.io/tsidentity/tornjak-spire-server:new --log-level=debug

podman machine ssh --log-level=debug

```

## Cleanup. Stop/Remove the VM
```
podman machine list
podman machine stop
podman machine rm
```
