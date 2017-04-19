## list routing table
```
netstat -nr
```

## list open ports:
```
lsof -i -P | grep LISTEN
```

## list all the connections:
```
netstat -a
```

## when the iptables are blocking the traffic:
```
# list the tables:
iptables -S
# set the tables:
iptables -P INPUT ACCEPT

# Iptable in ccsapi-mon-01 is always set to be
-P INPUT DROP
-P FORWARD ACCEPT
-P OUTPUT ACCEPT

# but it should be
-P INPUT ACCEPT
-P FORWARD ACCEPT
-P OUTPUT ACCEPT
```

## list interfaces and their IPv4
```
ip -4 a
```
