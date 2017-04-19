This document is based on https://help.github.com/articles/remove-sensitive-data/
Moving to a new repository: https://www.smashingmagazine.com/2014/05/moving-git-repository-new-server/


## list all the sensitive files
git log --all -- **/redis.go

## execute the purge:
git filter-branch --force --index-filter 'git rm --cached --ignore-unmatch proxy/src/conf/redis.go' --prune-empty --tag-name-filter cat -- --all
or
git filter-branch --force --index-filter 'git rm --cached --ignore-unmatch **src/conf/redis.go' --prune-empty --tag-name-filter cat -- --all

## try listing again, once happy, push to remote

git push origin --force --all
# removed from tagged releases:
git push origin --force --tags

# force all objects in your local repository to be dereferenced and garbage collected
git for-each-ref --format='delete %(refname)' refs/original | git update-ref --stdin
git reflog expire --expire=now --all
git gc --prune=now


################
############### OpenRadiant cleanup:
##############
review:
https://github.ibm.com/alchemy-containers/openradiant/search?utf8=%E2%9C%93&q=password
https://github.ibm.com/alchemy-containers/openradiant/search?utf8=%E2%9C%93&q=passwd&type=Code
https://github.ibm.com/alchemy-containers/openradiant/search?utf8=%E2%9C%93&q=DyJtZPNjkMeY7zgk
# list deleted files
git log --diff-filter=D --summary | grep delete

# search history:
git grep docker4ever $(git rev-list --all)
git grep DyJtZPNjkMeY7zgk $(git rev-list --all)

# files to fix and put back after cleanup
proxy/create_tenant.sh
proxy/mk_user_cert.sh
proxy/gen_server_certs.sh
proxy/src/conf/conf.go
proxy/proxy-test-container/lib/mk_test_user_cert.sh

# list of files that were deleted
git log --diff-filter=D --summary | grep delete

# purge the history (need to run twice each batch to take effect)
# batch 1
git filter-branch --force --index-filter 'git rm --cached --ignore-unmatch */conf/conf.go' --prune-empty --tag-name-filter cat -- --all
git filter-branch --force --index-filter 'git rm --cached --ignore-unmatch **/*.pem' --prune-empty --tag-name-filter cat -- --all
git filter-branch --force --index-filter 'git rm --cached --ignore-unmatch **/*.key' --prune-empty --tag-name-filter cat -- --all
git filter-branch --force --index-filter 'git rm --cached --ignore-unmatch **/ca.*' --prune-empty --tag-name-filter cat -- --all
git filter-branch --force --index-filter 'git rm --cached --ignore-unmatch user_certificates*' --prune-empty --tag-name-filter cat -- --all
git filter-branch --force --index-filter 'git rm --cached --ignore-unmatch *mk_user_cert*' --prune-empty --tag-name-filter cat -- --all
git filter-branch --force --index-filter 'git rm --cached --ignore-unmatch **hjproxyup.sh' --prune-empty --tag-name-filter cat -- --all
git filter-branch --force --index-filter 'git rm --cached --ignore-unmatch *make_TLS_certs.sh' --prune-empty --tag-name-filter cat -- --all

# batch 2
git filter-branch --force --index-filter 'git rm --cached --ignore-unmatch *proxy-test/*' --prune-empty --tag-name-filter cat -- --all
git filter-branch --force --index-filter 'git rm --cached --ignore-unmatch *gen_server_certs*' --prune-empty --tag-name-filter cat -- --all
git filter-branch --force --index-filter 'git rm --cached --ignore-unmatch *conf/redisclient.go' --prune-empty --tag-name-filter cat -- --all
git filter-branch --force --index-filter 'git rm --cached --ignore-unmatch *mk_test_user_cert*' --prune-empty --tag-name-filter cat -- --all
git filter-branch --force --index-filter 'git rm --cached --ignore-unmatch *create_tenant.sh' --prune-empty --tag-name-filter cat -- --all
git filter-branch --force --index-filter 'git rm --cached --ignore-unmatch *redis.go' --prune-empty --tag-name-filter cat -- --all


# post cleanup verification. Should be empty;
git log --all -- *proxy-test/*
git log --all -- **/redis.go
git log --all -- **/*.key
git log --all -- **/*.cert
git log --all -- **/ca.*
git log --all -- **/*.pem
git log --all -- */conf.go
git log --all -- */user_certificates*
git log --all -- user_certificates/*
git log --all -- **user_certificates**
git log --all -- **mk_user_cert**
git log --all -- **hjproxyup.sh**
git log --all -- *make_TLS_certs*
git log --all -- *proxy-test/*
git log --all -- **test_containers/*
git log --all -- **gen_server_certs*
