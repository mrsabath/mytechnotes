# MAC hints:


## open a file from a terminal
```
open -a Smultron test.yaml
```

##frozen Menubar (w/ Spotlight):
```
killall SystemUIServer
```

##Lotus Notes files:
```
cd ~/Library/Application\ Support/IBM\ Notes\ Data/
 ~/Library/Application Support/Lotus Notes Data/sabath.nsf
"/Users/sabath/Library/Application Support/Lotus Notes Data/SABATH.id"
```

##List large directories:
```
du -m | sort -n
```

## List large files:
```
ls -Slhr
```
##list ports:   
```
lsof -i -P   # list open ports
lsof -i -P | grep LISTEN  
netstat -a
```

## useful key shortcuts:
```
cmd + shift + 4     # snapshot to the file
cmd + ctr + shift + 4    # snapshot to memory
ctr + shift + eject     # lock screen
cmd + tab    # scroll through apps
cmd +     # zoom in terminal
cmd opt + # zoom in screen, cmd opt 8 # toggle in/out
```

## zip:
```
zip -er Docs.zip Docs
```

##vi arrows problem   
```
:set term=cons25
```

## run forever:
```
NOW=$(date +%s); while true; do $command; NOW2=$(date +%s);echo $(($NOW2-$NOW)); sleep 2; done
```

##Estimate disk size on directories   
```
du -d 2 -m -P -x -c
```
