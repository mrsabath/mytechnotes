# Automated push to git

## if you want automated `git push` daily, you need following:

create `push_git.sh`
```
cd "/Users/sabath/Box Sync/projects/mytechnotes"
git add *
git status
MSG=$(date +'%Y-%m-%d %H:%M:%S')
git commit  -m "Auto-committed at $MSG"
HOME=/Users/sabath git push origin master
```
create `crontab.txt`
```
MAILTO=sabath@us.ibm.com
# run during lunch
15 12 * * * "/Users/sabath/Box Sync/projects/mytechnotes/push_git.sh" >> /tmp/cron.log 2>&1
```

update ``~/.ssh/config`

```
Host github.ibm.com
     HostName github.ibm.com
     User git
     UserKnownHostsFile /dev/null
     StrictHostKeyChecking no
     IdentityFile /Users/sabath/.ssh/id_rsa
```
