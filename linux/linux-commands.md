# useful linxu commands:


## SSH
```
ssh -o UserKnownHostsFile=/dev/null -o StrictHostKeyChecking=no -i ~babaYaga-conf/weaver_sample.pem ec2-user@
```

## SSH debug mode (level 3)
```
ssh -vvv
```

## SSH delays:
[http://injustfiveminutes.com/2013/03/13/fixing-ssh-login-long-delay/](http://injustfiveminutes.com/2013/03/13/fixing-ssh-login-long-delay/)
```
ssh -o GSSAPIAuthentication=no
```

## history without the numbers
```
history | cut -c 8- |grep log
```

## CAT all the files in a directory. Generates a lot of Read IO
```
find . -name "*.log" -print | xargs cat
```

## find busy ports
```
sudo lsof -i TCP -P
```

## List large directories:
```
du -m | sort -n
```
## List large files:
```
ls -Slhr
open -a Smultron <file_name>
```

## watch for changes
```
watch ls -l
```

##list ports:
```
lsof -i -P   # list open ports
lsof -i -P | grep LISTEN
netstat -a
```
##### Disk Usage
```
DiskStation> du -hc -d 1 .
```

## Estimate disk size on directories:
```
du -d 2 -m -P -x -c
```

## Disk usage:
```
[root@itmm /]#  du -hc --max-depth=1 .
16K     ./lost+found
11M     ./boot
428K    ./dev
du: cannot change to directory `./proc/5629/fd': No such file or directory
du: `./proc/5681/fd/4': No such file or directory
923M    ./proc
12G     ./var
14M     ./tmp
18M     ./etc
17M     ./root

df -h
```

## disable service:
```
chkconfig iptables off
[root@rcc-pok-rem-02 bin]# chkconfig --list iptables
iptables        0:off   1:off   2:off   3:off   4:off   5:off   6:off
```

## Time synchronize:
```
[root@rcc-pok-rem-02 ~]# ntpdate pokgsa.ibm.com
13 Feb 10:48:40 ntpdate[15298]: adjust time server 9.56.248.25 offset 0.000660 sec
[root@rcc-pok-rem-02 ~]# vi /etc/ntp.conf

[root@rcc-pok-rem-02 ~]# service ntpd start
Starting ntpd:                                             [  OK  ]
[root@rcc-pok-rem-02 ~]# date
Fri Feb 13 10:50:05 EST 2009
Synchronize
[root@rcc-pok-rem-02 ~]# ntpd -q
[root@rcc-pok-rem-02 ~]# date
Fri Feb 13 10:51:19 EST 2009
```

## list IPv4 IP addresses and their interaces:
```
ip -4 a
```
