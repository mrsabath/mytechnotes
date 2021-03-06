# Bash shell Syntax

Header:

```bash
#!/bin/bash -x
# -x debug mode
```

## create help menu:
```bash
helpme()
{
  cat <<HELPMEHELPME

Syntax: ${0} <volume id> <VRT id> <test id>
Where:
  volume id   - Volume id (OpenStack)
  VRT id  - VRT id (WCC) being tested
  test id - test id associated with this test (optional)

HELPMEHELPME
}

# validate the arguments
if [[ "$1" == "-?" || "$1" == "-h" || "$1" == "--help" || "$2" == "" ]] ; then
  helpme
  exit 1
fi
```

## input parameters

```bash
echo "Select Cloud:"
read CL
echo "Selected $CL"
```

## some other things:

```bash
#check if directory exists:
if [ -d "$DIRECTORY" ]; then
  echo "Directory $DIRECTORY exists."
fi

# validate the external parameters
if [[ "$PROBE_DIR" == "" || "$WCC_DIR" == "" ]] ; then
  echo "ERROR: environment setup required. For example, run the following: "
  echo "    . /root/wccma/setwccmaruntime.sh"
  helpme
  exit 1
fi

# while infinite
while :; do
   curl $1
   sleep 1
   echo -e "\n"
done


# test integers:
# gather the return code from previous command
status=$?
if [ "$status" -eq "0" ]; then
	echo "($status) validation status: SUCCESS!"
elif [ "$status" -eq "1" ]; then
        echo "($status) validation status: FAIL!"
else
       echo "($status) validation status: UNKNOWN"
fi

func() {
	echo $1
	echo $2
	return 0
}

rc=func arg1 arg2
echo $rc
```

## line parsing:

```
# remove end of line
 sed 's/ *$//g'
```


## Complete example with parsing
```bash
 function get_admin_TLS {
   if [ ! -f "$1" ]; then
       echo "Missing creds file: $stub_auth_file"
       exit 1
   fi
   stub_auth_file=$1
   local line
   # find the line matching api_key, make sure it's admin
   line=$(grep "$ADMIN_KEY" "$stub_auth_file" |grep "\"Role\":\"admin\"")
   if [[ "$line" == "" ]] ; then
       echo "No admin account found for provided key: $ADMIN_KEY"
       exit 1
   fi
   # get 11th field after "," get the value of it after ":", and strip out doublequotes
   USERNAME=$(echo "$line" |grep "\"Role\":\"admin\"" | cut -d "," -f11 | cut -d ":" -f2 | sed 's/\"//g')
   SHARDNAME=$(echo "$line" |grep "\"Role\":\"admin\"" | cut -d "," -f13 | cut -d ":" -f2 | sed 's/\"//g')
   if [[ "$USERNAME" == "" || "$SHARDNAME" == "" ]] ; then
       echo "Missing 'Userid' or 'Shard_name' value in admin account"
       echo "Line: $line"
       exit 1
   fi
} # end of example
```


## create random name diretory using process id
```bash
mkdir temp.$$
```

## check the current directory and cd there if needed.
```bash
if [ ! $(basename "$PWD") = "vagrant" ]; then
	cd vagrant || exit
fi
```

## check if dir exist, otherwise create one
```bash
TARGET="$GOPATH/src/github.ibm.com/alchemy-containers"
if [ ! -d "$TARGET" ]; then
  mkdir -p "$TARGET"
fi
```

## execute complex SSH call
```bash
# parse the long string, ADMIN_OUT, result of running a command, and find the
# last occurrence of ADMIN_KEY=value, using 'rev' function:
ssh -o LogLevel=quiet -o StrictHostKeyChecking=no -o UserKnownHostsFile=/dev/null -i 	~/.vagrant.d/insecure_private_key vagrant@$master_ip -t 'export ADMIN_OUT=$(sudo docker exec api-proxy /api-proxy/create_admin.sh admin1 shard1); export ADMIN_KEY=$(echo $ADMIN_OUT | grep "ADMIN_KEY" | rev | cut -d "=" -f1 | rev);sudo docker exec api-proxy /api-proxy/create_user.sh dev-vbox '$tenant' '$shard' '$master_ip' '$master_ip' $ADMIN_KEY'
```

# get home directory:
```bash
export KUBECONFIG=$(cd;pwd)/first-user/kube-config
```

# some for loop operations
```shell
for N in $(kubectl get node | grep SchedulingDisabled | awk '{ print $1 }');do kubectl uncordon $N; done
```

# get first value in second row
```shell
kubectl logs  $(kubectl get pods | grep executor | awk -F ' ' '{print $1}' | sed -n 2p) | grep "###EXEC" > $dir/executor2.log
```

## insert multi-line text to the file
```
cat >~/.config/jenkins_jobs/jenkins_jobs.ini <<EOF
[job_builder]
ignore_cache=True
keep_descriptions=False
include_path=.:scripts:~/git/
recursive=False
exclude=.*:manual:./development
allow_duplicates=False

[stash]
username=user
password=pass
EOF
```


## source profiles from inside the script:
Run the shell script invoking DB2

```sh
#!/bin/bash
# run some stuff as root
# then switch to db2inst1 instance

su - db2inst1 -c '. /home/db2inst1/sqllib/db2profile; db2 catalog db ORDERDB on /db2; db2start'
```
