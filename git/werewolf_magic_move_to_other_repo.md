## move a whole repo to another location

```console
 505  git clone git@github.ibm.com:alchemy-containers/proxy.git
 507  git branch -a
 508  cd proxy/
 509  git branch -a
 510  git checkout master
 523  git checkout -b proxy
 513  mkdir proxy
# move it all to proxy dir
 516  open .
 517  git status
 520  git add .
 522  git status
 git commit -m "moved to proxy dir"
 524  git branch -a
 527  git remote add upstream git@github.ibm.com:alchemy-containers/openradiant.git
 528  git remote -v
 529  git status
 530  git push upstream proxy
# move the .gitigorne
   582  mv .gitignore proxy/
   584  git status
   585  git add .
   586  git status
   587  git commit -m "moved gitignore"
   588  git push upstream proxy
   589  git fetch upstream master
# it has to be merge first with the existing
   590  git merge upstream/master
   591  git status
   592  git logs
   594  git push upstream proxy
```


## split of containercafe to fr8r-proxy
```container
# clone containercafe   
520  mv containercafe/ fr8r-proxy/
525  cd fr8r-proxy/
527  git status
528  git checkout -b fr8r-proxy
531  git status
# remove all that is not needed:
533  git rm -r remoteabac/
534  git rm -r crawler/
535  git rm -r ansible/
537  git rm README.md
# move things around the way you like it
538  git status
539  git add .
540  git status
541  git commit -m "reorganized for proxy repo"
542  git remote add proxyupstream git@github.com:fr8r/fr8r-proxy.git
543  git remote -v
544  git branch -a
545  git push proxyupstream fr8r-proxy
```
## migrate repo to another repo, including all the branches:
```console
for brn in $(git branch --list -r "origin/*" | grep -v HEAD); do lbr=${brn#origin/}; echo lbr=$lbr; if git branch | grep -w $lbr; then echo $lbr already here; continue; fi; echo making $lbr; git checkout -b $lbr $brn; done
```
