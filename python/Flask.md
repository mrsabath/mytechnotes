# Flask web server
Flask is the common web server solution for Python.

## Environment setup
Flask 1. requires python 2.7.9 or higher

If you are using [pyenv](https://github.com/pyenv/pyenv):

```
pyenv versions
pyenv activate venv2715
```

Sample server file.

```
from flask import Flask, request
import os
from os.path import join, exists
import subprocess

app = Flask(__name__)

if __name__ == '__main__':
    app.run(debug=True,host='0.0.0.0')

@app.route('/')
def index():
    return "TPM wrapper"

@app.route('/getJWT')
def get():
    args = request.args.copy()
    claims = ""
    if args:
        claims = "--claims="
        for k in args:
            claims = claims + k + ":" + args[k] + "|"
    statedir = os.getenv('STATEDIR') or '/tmp'
    tpmkeyfile = join(statedir, "tpmkeyurl")
    with open(tpmkeyfile) as f:
        tpmkey = f.read().strip()
    out = subprocess.check_output(['/usr/local/bin/gen-jwt.py',tpmkey,'--iss','example-issuer', claims])
    return str(out)

# to be depricated
@app.route('/getJWKS')
def getJWKS():
    statedir = os.getenv('STATEDIR') or '/tmp'
    tpmkeyfile = join(statedir, "tpmkeyurl")
    with open(tpmkeyfile) as f:
        tpmkey = f.read().strip()
    out = subprocess.check_output(['/usr/local/bin/gen-jwt.py',tpmkey,'--jwks','/tmp/jwks.json'])
    jwksfile = join(statedir, "jwks.json")
    with open(jwksfile) as f:
        jwks = f.read().strip()
        return str(jwks)
    return str(out)

@app.route('/public/getCSR')
def getCSR():
    statedir = os.getenv('STATEDIR') or '/tmp'
    csrfile = join(statedir,"server.csr")
    with open(csrfile) as f:
        csr = f.read().strip()
        return str(csr)

@app.route('/public/postX5c', methods=["POST"])
def postX5c():
    try:
        statedir = os.getenv('STATEDIR') or '/tmp'
        x5cfile = join(statedir, "x5c")
        # if file already exists, don't all to override it
        if exists(x5cfile):
            # return 403 Forbidden, 406 Not Accesptable or 409 Conflict
            return "File already exists.", 403
        if request.method == 'POST':
            with open(x5cfile, "w+") as f:
                f.write(request.data)
                f.close()
                return "Download successful"
    except Exception as e:
        print (e)
        #flash(e)
        return ("Error %s" % e)

```

Start the server locally:
```
FLASK_APP=/server.py python -m flask run --host=0.0.0.0
```

Connect client:
```
# submit the file:
curl -X POST -H "Content-Type: application/json" -d @x5c localhost:5000/public/postX5c
# get file:
curl --request GET localhost:5000/public/getCSR
```
