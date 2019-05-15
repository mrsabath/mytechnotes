# Various Useful Hacks

## Run a local web server on specific directory
```
cd ~/Downloads
pyenv activate venv277
python -m SimpleHTTPServer
Serving HTTP on 0.0.0.0 port 8000 ...

# connect there with a browser :)
```

## dump the content of the file to Clipboard
```
cat myfile | pbcopy
```

## Move any binary file as text
you take any binary and you do
```
cat ${BIN_PATH} | base64 -w0
```

then you base64 decode it on the host and store it

```
echo -n "f0VMRgIBAQAAAAAA...AAAAAAAAAAAAAAAAAAA" | base64 -d > xxd.bin && chmod +x xxd.bin

```
