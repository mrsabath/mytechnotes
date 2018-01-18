# Simple Python Version Mangagment

https://github.com/pyenv/pyenv
```
pyenv versions
pyenv version

pyenv activate venv351
```

## make version default
```
pyenv versions
pyenv global 3.6.4
pyenv versions
  system
  2.7.7
  2.7.9
  3.5.1
* 3.6.4 (set by /Users/sabath/.pyenv/version)

python --version
Python 3.6.4
```

## install a new version:
```
pyenv install --list

#If the version you need is missing, try upgrading pyenv:
brew update && brew upgrade pyenv

pyenv install 3.6.4
```

## pyenv-virtualenv
https://github.com/pyenv/pyenv-virtualenv

### list virtualenvs:
pyenv virtualenvs

### create a new virtualenv
 pyenv virtualenv 3.6.4 venv364


### activate pyenv-virtualenv

pyenv versions
pyenv activate venv364
