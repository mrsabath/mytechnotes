# Git Stash
Stashing takes the dirty state of your working directory — that is, your modified tracked files and staged changes — and saves it on a stack of unfinished changes that you can reapply at any time.

```
git stash
Saved working directory and index state \
  "WIP on master: 049d078 added the index file"
HEAD is now at 049d078 added the index file
(To restore them type "git stash apply")
```

To get the stashed files, first list the stash instances:

```
git stash list
stash@{0}: GitHub: stashing before switching to jwt_expire
stash@{1}: WIP on master: e9ac9e3 Merge pull request #29 from kompass/version_cleanup
stash@{2}: GitHub: stashing before switching to vtpm-sidecar
stash@{3}: WIP on sidecar: a148517 Parametrize the address of the Vault server
stash@{4}: WIP on master: 0512a67 Merge pull request #8 from kompass/move-repo
```

View the details of specific stash:
```
git stash show stash@{1}
 .gitignore                   |  2 ++
 jwt-sidecar/gen-jwt.py       | 99 ---------------------------------------------------------------------------------------------------
 jwt-sidecar/requirements.txt |  1 -
 3 files changed, 2 insertions(+), 100 deletions(-)
 ```

Select the stash and recreate the stashed files into a branch:
```
git stash branch testchanges stash@{1}
Switched to a new branch "testchanges"
# On branch testchanges
# Changes to be committed:
#   (use "git reset HEAD <file>..." to unstage)
#
#      modified:   index.html
#
# Changes not staged for commit:
#   (use "git add <file>..." to update what will be committed)
#
#      modified:   lib/simplegit.rb
#
Dropped refs/stash@{0} (f0dfc4d5dc332d1cee34a634182e168c4efc3359)
```
