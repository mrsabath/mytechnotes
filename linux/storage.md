# summary of all storage devices
```
fdisk -l
```


# list block storage devices:
```
lsblk -l
```

# list file systems:
```
df -h
```

# disk Usage
```
# root directory, one level, usage summary
du -hc -d 1 /
# usage sorted
du -hc -d 1 | sort -n
```
