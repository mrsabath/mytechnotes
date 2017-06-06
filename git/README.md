## Collection of git commands

## List (see) global setup:
```
git config --global -l
```
## List config setup:

```
git config -l
```

## Git global setup:
```
git config --global user.name "Mariusz Sabath"
git config --global user.email "sabath@us.ibm.com"
```

## Create Repository:
```
mkdir ccsapi
cd ccsapi
git init
touch README
git add README
git commit -m 'first commit'
git remote add origin git@github.rtp.raleigh.ibm.com:sabath-us/ccsapi.git
git push -u origin master
git status
```

## reset git to match master
```
git fetch origin
git reset --hard origin/master
```

## list all the Users
```
git log --format='%aN' | sort -u
```

## list all the users and their emails:
```
git log --all --format='%aN <%cE>' | sort -u
```

## list all the commits and users:
```
git shortlog -s -n
   472  Mariusz Sabath
   214  Mike Spreitzer
```
## count all the files
```
git ls-files | wc -l
```

## count all the commits in one branch
```
git rev-list --count master
```
## count all the commits accross all branches
```
git rev-list --all --count
```

# Contribution Process:
## Fork the repo to your own account
```
git clone ... && cd <dir>
git remote add upstream <url to real/upstream repo>

# work on master
git checkout master
git fetch upstream master --tags
git merge --ff-only upstream/master
git push origin master

# work on branch
git checkout issue-
git fetch upstream
git rebase upstream/master
git push -f origin issue-
```

## Existing Git Repo?
```
cd existing_git_repo
git remote add origin git@github.rtp.raleigh.ibm.com:sabath-us/ccsapi.git
git push -u origin master

git init
# all the stuff is in .git

git status

# add a file to tracking
git add <file>

# commit tracked files
git commit -m "message"
# add and commit tracked files
git commit -am "message"
git commit -a -m "message1"

#after modifying a file, you need to add again
git add <file>
# then commit

# gui:
gitk --all

git branch
git checkout -b experiment
git remote -v

# combines fetch and merge
git pull

git fetch
git merge

# push to the server:
git push

# switch to branch1
git checkout branch1

# to merge with another branch
git merge master

# delete branch
git branch -d experiment

github.rtp.raleigh.ibm.com

# Ignore local changes:
git stash save
git stash drop # this will remove the stashed files

# Change the Git Repo?
git remote set-url origin <new URL>
```

## view history of a single file:
```
git log <filename>
```

## see who update the lines:
```
git blame <filename>
```

## setup the env.
```
. ~/os_config/setup.gitssh.sh

 cd /Users/sabath/Documents/workspace/workspace_git2
 git clone git@github.rtp.raleigh.ibm.com:sabath-us/cluster-manager.git
Cloning into 'cluster-manager'...
remote: Counting objects: 1698, done.
remote: Compressing objects: 100% (616/616), done.
remote: Total 1698 (delta 1080), reused 1663 (delta 1060)
Receiving objects: 100% (1698/1698), 411.18 KiB | 0 bytes/s, done.
Resolving deltas: 100% (1080/1080), done.
Checking connectivity... done
 cd cluster-manager/
 git checkout cluster
Branch cluster set up to track remote branch cluster from origin.
Switched to a new branch 'cluster'
```

## Some old stuff:
```
# choir
git clone git@github.rtp.raleigh.ibm.com:lindj-us/choir.git
git checkout -b ms_test
git push -u origin ms_test

# create a new branch:
git checkout -b non_nested_client

# review remote access
git remote -v
origin	git@github.rtp.raleigh.ibm.com:sabath-us/cluster-manager.git (fetch)
origin	git@github.rtp.raleigh.ibm.com:sabath-us/cluster-manager.git (push)

# check the submission status
git status

# add files
git add <file>

# commit the changes
git commit -m "reason for changes"

# push the data to the remote server
git push -u origin non_nested_client

# remove (delete) file
git rm myfile
git commit -m "Deleting file"
git push
# if you still see 'deleted' file in git status:
git add -u
git commit -m "Deleting file"

# GIT choir
git clone git@github.rtp.raleigh.ibm.com:lindj-us/choir.git
git clone http://github.rtp.raleigh.ibm.com/lindj-us/choir.git
git push origin master
# or if another branch
git push -u origin ms_test

sudo git push origin master
# merge branch to master
git checkout master
git pull origin master
git merge test
git push origin master

# remove temporary changes
git reset --hard
git pull origin master
# or preserve it for later:
git stash
git pull
git stash pop


# fetcher
git clone http://github.rtp.raleigh.ibm.com/lindj-us/fetcher.git
git clone http://github.rtp.raleigh.ibm.com/lindj-us/python-fetcherclient.git



# old stuff
 659  git remote add -t cluster cluster-manager git@github.rtp.raleigh.ibm.com:sabath-us/cluster-manager.git
  660  git fetch cluster-manager
  661  git status


http://git-scm.com/book/en/v2
It has a nice section on branching



---------------------------------------------------- git branching
git branch -a   #list all
git branch v3
git checkout v3
git add .
git commit -m "init v3"
git push -u origin v3

--------------------------------------------------------- git merge
git checkout master
git pull
git merge origin/v3


to merge master into v3 branch:
git checkout v3
git pull  (to get the latest v3)
git merge master  (to get master latest updates)
----> ready to work on v3
-------------------------------------------------


##### rebase branch to master
git checkout swarm
git rebase master


## capture changes to new branch
git branch    # list local branches
git branch force  # create a new branch
git checkout force
git commit -am "text" # with add


# combine several commits, squash
git commit -am "test1"
git commit -am "test2"
git log
git rebase -i HEAD~3
# edit the file to list pick, s, s
pick #
s #
s #
# save
# then modify the message
git log

# to abort any squash above:
git rebase --abort


# To delete a local branch
git branch -d the_local_branch

# check the remote config
git remote -v
origin	git@github.rtp.raleigh.ibm.com:project-alchemy/ccsapi.git (fetch)
origin	git@github.rtp.raleigh.ibm.com:project-alchemy/ccsapi.git (push)

# Legit
http://www.git-legit.org/

# manual merge:
# Step 1. Update the repo and checkout the branch we are going to merge
git fetch origin
git checkout -b vpnaas origin/vpnaas

# Step 2. Merge the branch and push the changes to GitLab
git checkout master
git pull
git merge --no-ff vpnaas
git push origin master


##### copy from one repo to another  ###
# in new repo:
git remote set-url origin git@github.ibm.com:alchemy-containers/swarm.git
git push origin master


###### other useful tricks #####

# revert temporarly to different commit:
git checkout -b old-state <SHA>

# undo things in git
https://github.com/blog/2019-how-to-undo-almost-anything-with-git

# revert commit that was pushed
git revert <SHA>

# modify the message for executed commit
git commit --amend -m "New message"

# undo local commit:
git checkout -- <bad filename>

# go back to old commit (it removes all the changes since)
git reset --hard <last good SHA>

# git reflog is an amazing resource for recovering project history.
# You can recover almost anything—anything you've committed—via the reflog.

# list all my commits:
git log --committer="Mariusz Sabath"

# get list of commits between 2 SHAs:
git --no-pager shortlog --oneline --no-merges <git-id-first-SHA>..<git-id-last-SHA>

#### git tagging
https://git-scm.com/book/en/v2/Git-Basics-Tagging

# list tags:
git tag
git tag -l "v0.*"  # list tags that start with v0.

# add a tag:
git tag -a v0.1 -m "CRUD w/ working status updates"
git push

git show v0.1

### create PR (pull request)
To do a PR, you create a branch with git checkout -b ‘issue-X’ for example, work on the issue, test, and push back,
then you will see the option to create the PR from your latest commit/branch
after you create it, you can merge from the UI

## Rename the branch
1. Rename your local branch.
# If you are on the branch you want to rename:
git branch -m new-name

# If you are on a different branch:
git branch -m old-name new-name

2. Delete the old-name remote branch and push the new-name local branch.
git push origin :old-name new-name

3. Reset the upstream branch for the new-name local branch.
# Switch to the branch and then:
git push origin -u new-name
```
