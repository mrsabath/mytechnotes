# Automated push to git
These steps are for Mac system.

Create a new key dedicated for GitHub:
```console
ssh-keygen -b 4096 -t rsa -f ~/.ssh/id_rsa_github -q -C "your@email" -N ""
```

Create/update `~/.ssh/config`. If you are using separate GitHub hosts, follow the example:

```shell
# this is for my public GitHub
Host github.com
     HostName github.com
     User git
     UserKnownHostsFile /dev/null
     StrictHostKeyChecking no
     IdentityFile /Users/user/.ssh/id_rsa_github

# this is for my IBM GitHub
Host github.ibm.com
     HostName github.ibm.com
     User git
     UserKnownHostsFile /dev/null
     StrictHostKeyChecking no
     IdentityFile /Users/user/.ssh/id_rsa_githubibm
```
Add the newly created key(s) to your system, then list

```
ssh-add ~/.ssh/id_rsa_github
ssh-add -l
```

Publish your **public** key to github: [https://github.com/settings/keys](https://github.com/settings/keys)

Test the connection (must use user 'git'):
```
ssh -T git@github.com
# or more verbose to see the errors:
ssh -vT git@github.com
```

Change the directory to your local clone of the repo and run:
```
git remote set-url origin git@github.com:username/your-repository.git
```

If you want fully automated `git push` daily, you need following:

create `push_git.sh`
```shell
cd "/Users/username/workspace/mytechnotes"
git add *
git status
MSG=$(date +'%Y-%m-%d %H:%M:%S')
git commit  -m "Auto-committed at $MSG"
HOME=/Users/username git push origin master
```
create `crontab.txt`
```
MAILTO=username@email
# run during lunch
15 12 * * * "/Users/username/Box Sync/projects/mytechnotes/push_git.sh" >> /tmp/cron.log 2>&1
```

create or update crontab:
```
crontab -l
crontab crontab.txt
# or add to existing one
crontab -e
```
