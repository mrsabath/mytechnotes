# create token:
https://help.github.com/articles/creating-an-access-token-for-command-line-use/



# use this script:
 https://github.com/IQAndreas/github-issues-import


 # create config.ini file:
 ```
 [source]
server = github.com
repository = fr8r/fr8r-proxy
username = tommaso@madonia.me
password =

[target]
server = github.ibm.com
repository = alchemy-containers/fr8r-proxy
username = tmadonia@us.ibm.com
password =
```

# to move the issues:
pyenv activate venv351
python3 gh-issues-import.py --issues 6
python3 gh-issues-import.py --issues 84 89 90 91 92
# or `python3 gh-issues-import.py --open` if you want to move all the open issues
