```
[thor@vizio-dev-cont ~]$ NOW=$(date +%s); nova stop 5710c82d-d7ff-49c1-881d-fd56c749eb14
[thor@vizio-dev-cont ~]$ nova list | grep 5710c82;NOW2=$(date +%s);echo $(($NOW2-$NOW))
| 5710c82d-d7ff-49c1-881d-fd56c749eb14 | MS-busybox                                            | ACTIVE | powering-off | Running     | BM_ORG-dev_org10-5a7c2604b943d8f2-net=172.16.61.148, 10.120.104.137 |
7

export command=""
now=$(date +"%Y-%m-%d.%H:%M:%S")
# time since beginning of the loop
NOW=$(date +%s); while true; do $command; NOW2=$(date +%s);echo $(($NOW2-$NOW)); sleep 2; done
# time for every command execution:
while true; do NOW=$(date +%s); $command; NOW2=$(date +%s);echo $(($NOW2-$NOW)); sleep 2; done

# Linux command to get DB2 format timestamp:
currenttime=$(date +'%Y-%m-%d %H:%M:%S')

date +'%Y-%m-%d %H:%M:%S'
> 2013-12-10 14:55:43

# in DB2 table:
2013-12-10-14.55.43.000000
```
