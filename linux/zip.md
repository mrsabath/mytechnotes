## Create Zip Archive

```
tar -czvf archive.tgz dir_name
tar -czvf archive.tgz dir_name1 dir_name2
tar -czvf archive.tgz dir_name --exclude=dir_name/tmp
tar -czvf archive.tgz dir_name --exclude=*.mp4
```

## Unzip it
```
tar -xzvf archive.tgz
```

## Encrypt, password protected
```
# encrypt:
tar cz folder_to_encrypt | openssl enc -aes-256-cbc -e > out.tar.gz.enc
# decrypt:
openssl enc -aes-256-cbc -d -in out.tar.gz.enc | tar xz
# or
zip -e <file_name>.zip <list_of_files>
# or 
gpg --encrypt out.tar.gz
```
