# enable SUDO on ubuntu:
```
export USER_NAME=radiant
sudo mkdir -p /etc/sudoers.d/
sudo -E bash -c 'echo "$USER_NAME ALL=(ALL) NOPASSWD:ALL" > /etc/sudoers.d/$USER_NAME'
sudo chmod 0440 /etc/sudoers.d/$USER_NAME
sudo usermod -a -G sudo $USER_NAME
```
