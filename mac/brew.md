# Brew - install mangager for Mac

Homebrew installs packages to their own directory and then symlinks their files into /usr/local.


Install (https://brew.sh/)
```
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/HEAD/install.sh)"
```

Most common commands:
```
brew install <app>
# list all applications installed by brew
brew list

# update the formulas
brew update

# list outdated apps:
brew outdated

# upgrade all the outdated apps:
brew upgrade
brew upgrade <app>

```
