http://nmap.org/book/man-port-scanning-basics.html
```
-Pn is "no ping"  

# use nmap to scan port range
nmap -Pn -p 22-22 9.2.226.125  

# or nc
nc <ip> <port>  

# or sinlge port
nmap -Pn -p 22 9.2.226.125  

# full scan is
-p 1-65535


# home scan:
silesia-mac:~ sabath$ nmap -Pn -p 1-65535 32.216.129.147

Starting Nmap 6.47 ( http://nmap.org ) at 2014-12-10 17:49 EST
Nmap scan report for 32.216.129.147
Host is up (0.0045s latency).
Not shown: 65521 filtered ports
PORT      STATE  SERVICE
25/tcp    closed smtp
80/tcp    open   http
135/tcp   closed msrpc
139/tcp   closed netbios-ssn
445/tcp   closed microsoft-ds
1214/tcp  closed fasttrack
3128/tcp  open   squid-http
5554/tcp  closed sgi-esphttp
6346/tcp  closed gnutella
8080/tcp  open   http-proxy
8866/tcp  closed unknown
8998/tcp  closed unknown
9996/tcp  closed unknown
50001/tcp closed unknown

Nmap done: 1 IP address (1 host up) scanned in 163.65 seconds
```
