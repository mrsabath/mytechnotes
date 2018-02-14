# Create an image and sign it.
Key is stored on my computer

Sample Docker file:
```yaml
FROM ubuntu:16.04
RUN echo "hello" > /data
ENTRYPOINT ["/bin/bash","-c", "while true; do sleep 5; cat /data; done"]
```

```
docker build . -t mrsabath/test:latest
```

## Sign the image:

```
silesia:kube sabath$ export DOCKER_CONTENT_TRUST=1
silesia:kube sabath$ docker push mrsabath/test:latest
The push refers to repository [docker.io/mrsabath/test]
c8194ce6f1c6: Pushed
6f4ce6b88849: Mounted from library/ubuntu
92914665e7f6: Mounted from library/ubuntu
c98ef191df4b: Mounted from library/ubuntu
9c7183e0ea88: Mounted from library/ubuntu
ff986b10a018: Mounted from library/ubuntu
latest: digest: sha256:4bf123ff597dc46ddeb96dc4aec8e814375b544f050725528e2a20740ebde7f1 size: 1564
Signing and pushing trust metadata
You are about to create a new root signing key passphrase. This passphrase
will be used to protect the most sensitive key in your signing system. Please
choose a long, complex passphrase and be careful to keep the password and the
key file itself secure and backed up. It is highly recommended that you use a
password manager to generate the passphrase and keep it safe. There will be no
way to recover this key. You can find the key in your config directory.
Enter passphrase for new root key with ID 113aac7:
Passphrase is too short. Please use a password manager to generate and store a good random passphrase.
Enter passphrase for new root key with ID 113aac7:
Repeat passphrase for new root key with ID 113aac7:
Enter passphrase for new repository key with ID e78b0e0:
Repeat passphrase for new repository key with ID e78b0e0:
Finished initializing "docker.io/mrsabath/test"
Successfully signed docker.io/mrsabath/test:latest

silesia:kube sabath$ ./notary -s https://notary.docker.io list docker.io/mrsabath/test
Enter username: mrsabath (docker userid)
Enter password:
NAME      DIGEST                                                              SIZE (BYTES)    ROLE
----      ------                                                              ------------    ----
latest    4bf123ff597dc46ddeb96dc4aec8e814375b544f050725528e2a20740ebde7f1    1564            targets

silesia:kube sabath$ ./notary -s https://notary.docker.io list docker.io/mrsabath/4demo
NAME      DIGEST                                                              SIZE (BYTES)    ROLE
----      ------                                                              ------------    ----
latest    097b39c70b441cd4f517131b328436f4c92e9cbb8dcdbc6d9c0a5f9e5b1cf3f7    4087            targets
v2        097b39c70b441cd4f517131b328436f4c92e9cbb8dcdbc6d9c0a5f9e5b1cf3f7    4087            targets

# list public images:
./notary -s https://notary.docker.io list docker.io/library/<image name> e.g.
./notary -s https://notary.docker.io list docker.io/library/ubuntu

```


## Delete deployed image:
```
./notary -s https://notary.docker.io delete docker.io/mrsabath/test --remote
```

## Deploy signed and unsigned images:
```
silesia:kube sabath$ kubectl -n signed run test1 --image=mrsabath/test:latest
deployment "test1" created
silesia:kube sabath$ kubectl -n signed run test1 --image=mrsabath/web-ms:latest
Error from server (Error rewriting img: Notary failed to verify
):
silesia:kube sabath$ kubectl -n signed run test2 --image=mrsabath/web-ms:latest
Error from server (Error rewriting img: Notary failed to verify
):
```

## Reset the Docker trust root key and everything.
You will not be able to update, modify any pushed images with the old keys, so delete them
first
```
mv ~/.docker/trust/ ~/.docker/trust.BKP
# push again with the DOCKER_CONTENT_TRUST=1, new key prompts will happen.
```
