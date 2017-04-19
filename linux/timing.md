```
while [ 1 ]; do date; time curl https://containers-api.ng.bluemix.net/v3/containers/version; sleep 1; done

command="curl https://containers-api.ng.bluemix.net/v3/containers/version"
command="time curl localhost:8081/v3/containers/version"


NOW=$(date +%s); while true; do $command; NOW2=$(date +%s);echo $(($NOW2-$NOW)); sleep 2; done
```
