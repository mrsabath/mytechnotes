# How to hack containers
When running on Mac.

The location of files is typically here:

`/Users/$USER/Library/Containers/com.docker.docker/Data/`

List all the containers and their PIDs:

```
docker ps
...
docker ps -q | xargs docker inspect --format '{{.State.Pid}}, {{.ID}}'
55419, 12a755f37c0939ce544a13c2043cfc58c8c47dd7abdec11129803e2e26d59910
5117, 0ff1b2e613d584363608054dc2bf4aadc86acb30d5a90179f0c3c56c85dc1b05
4945, 9aaaa5bb411b00dc8500b395217760ea83de7d56e9af3d22852952fb55ca30dd
```


When running Docker Desktop for Mac, get inside the Docker VM
https://www.bretfisher.com/docker-for-mac-commands-for-getting-into-local-docker-vm/

```
docker run -it --rm --privileged --pid=host justincormack/nsenter1
```

Now, using process id from above, go to `/proc/$PID/` directory that has all the
details of the container.
`environ` - env. variables of the container

Get location of the log files
```
ls -larth /var/lib/docker/containers/*/*.log | grep cont
```
