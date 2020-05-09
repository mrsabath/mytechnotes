# Create Intermediate CAs that include the identity of the Node

## Create RootCA in Vault for signing CSRs and generating Intermediate CAs
```bash
# setup vault creds:
export VAULT_ADDR=
export ROOT_TOKEN=
COMMON_NAME="trusted-identity.ibm.com"
CONFIG="plugin-config.json"

vault status
vault secrets enable pki

# Increase the TTL by tuning the secrets engine. The default value of 30 days may
# be too short, so increase it to 1 year:
vault secrets tune -max-lease-ttl=8760h pki
vault delete pki/root

# create internal root CA
# expire in 100 years
export OUT
OUT=$(vault write pki/root/generate/internal common_name=${COMMON_NAME} \
    ttl=876000h -format=json)
echo "$OUT"

# capture the public key as plugin-config.json
CERT=$(echo "$OUT" | jq -r '.["data"].issuing_ca'| awk '{printf "%s\\n", $0}')
echo "{ \"jwt_validation_pubkeys\": \"${CERT}\" }" > ${CONFIG}
```


## On each node, create CSR
```console
# generate node private key:
openssl genrsa -out node.key 2048

# get the node information: cluster-name, regionName, datacenter
cat > "openssl.cnf" << EOF
[req]
req_extensions = v3_req
distinguished_name	= req_distinguished_name

[ req_distinguished_name ]
countryName      = Country Name (2 letter code)
countryName_min  = 2
countryName_max  = 2
stateOrProvinceName = State or Province Name (full name)
localityName        = Locality Name (eg, city)
0.organizationName  = Organization Name (eg, company)
organizationalUnitName = Organizational Unit Name (eg, section)
commonName       = Common Name (eg, fully qualified host name)
commonName_max   = 64
emailAddress     = Email Address
emailAddress_max = 64

[v3_req]
subjectAltName= @alt_names

[alt_names]
URI.1 = TSI:cluster-name:$CLUSTER_NAME
URI.2 = TSI:cluster-region:$CLUSTER_REGION
# To assert additional claims about this intermediate CA
# add new lines in the following format:
# URI.x = TSI:<claim>
# where x is a next sequencial number and claim is
# a `key:value` pair. For example:
# URI.3 = TSI:datacenter:fra02
EOF

# create CSR:
openssl req -new -sha256 -key node.key -out node.csr -subj  "/CN=jss-jwt-server" -reqexts v3_req -config <(cat openssl.cnf)

# alternatively: -config <(cat /etc/ssl/openssl.cnf openssl.cnf)
# review the CSR:
openssl req -in node.csr -noout -text
```
Sample output of CSR format:
```
Certificate Request:
    Data:
        Version: 0 (0x0)
        Subject: CN=jss-jwt-server
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                Public-Key: (2048 bit)
                Modulus:
                    00:ec:21:be:9a:9d:a8:7d:2b:69:69:57:de:eb:7d:
                    87:27:4f:89:62:1a:c9:8b:e1:a0:5d:e0:2a:92:28:
                    fb:6f
                Exponent: 65537 (0x10001)
        Attributes:
        Requested Extensions:
            X509v3 Basic Constraints:
                CA:FALSE
            X509v3 Key Usage:
                Digital Signature, Non Repudiation, Key Encipherment, Key Agreement
            X509v3 Extended Key Usage:
                TLS Web Server Authentication, TLS Web Client Authentication
            X509v3 Subject Alternative Name:
                URI:TSI:cluster-name:mycluster, URI:TSI:regionName:eu-de, URI:TSI:datacenter:fra02
    Signature Algorithm: sha256WithRSAEncryption
         09:1d:ba:99:cb:db:b2:b4:81:49:9e:6b:5b:d3:54:f9:db:27:
         84:ca:7b:7e:9d:7d:ac:4c:63:5e:3c:a0:af:d8:b3:48:21:49:
```

## On Vault - sign the CSR with Vault rootCA to create intermediate CA:

First, extract the node identity:
```bash
openssl req -in node.csr -noout -text | grep "URI.TSI"

  URI:TSI:cluster-name:mycluster, URI:TSI:regionName:eu-de, URI:TSI:datacenter:fra02
```

Then use the node idenitity as `uri_sans` parameter:
```console
# setup vault creds:
export VAULT_ADDR=
export ROOT_TOKEN=

vault write pki/root/sign-intermediate csr=@node.csr ttl=438000h uri_sans="TSI:cluster-name:mycluster,TSI:regionName:eu-de,TSI:datacenter:fra02"  > out


```
Sample output of Intermediate CA:
```json
{
  "request_id": "4a7e40ad-767a-608b-7b9c-9dcdacefe56c",
  "lease_id": "",
  "lease_duration": 0,
  "renewable": false,
  "data": {
    "certificate": "-----BEGIN CERTIFICATE-----\nMIIDSDCCAjCgAwIBAgIUSrHIC/0jBBVidkKkN2WttiL+fLswDQYJKoZIhvcNAQEL\nBQAwIzEhMB8GA1UEAxMYdHJ1c3RlZC1pZGVudGl0eS5pYm0uY29tMB4XDTIwMDUw\nMTE1MDEyMl
    ........
    n1WZOaY0FUyraQ==\n-----END CERTIFICATE-----",
    "expiration": 1619881312,
    "issuing_ca": "-----BEGIN CERTIFICATE-----\nMIIDXjCCAkagAwIBAgIUCBnOr1vUIQ9YwpzmVtMviFRzWNwwDQYJKoZIhvcNAQEL\nBQAwIzEhMB8GA1UEAxMYdHJ1c3RlZC1pZGVudGl0eS5pYm0uY29tMB4XDTIwMDQy\nM
    .........
    kzT\nxwM=\n-----END CERTIFICATE-----",
    "serial_number": "4a:b1:c8:0b:fd:23:04:15:62:76:42:a4:37:65:ad:b6:22:fe:7c:bb"
  },
  "warnings": [
    "The expiration time for the signed certificate is after the CA's expiration time. If the new certificate is not treated as a root, validation paths with the certificate past the issuing CA's expiration time will fail."
  ]
}
```

Sample view of the certificate:
```
openssl x509 -in cert.pem -noout -text
Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            4a:b1:c8:0b:fd:23:04:15:62:76:42:a4:37:65:ad:b6:22:fe:7c:bb
    Signature Algorithm: sha256WithRSAEncryption
        Issuer: CN=trusted-identity.ibm.com
        Validity
            Not Before: May  1 15:01:22 2020 GMT
            Not After : May  1 15:01:52 2021 GMT
        Subject: CN=jss-jwt-server
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                Public-Key: (2048 bit)
                Modulus:
                    00:ec:21:be:9a:9d:a8:7d:2b:69:69:57:de:eb:7d:
                    a5:39:16:00:40:d3:5a:0e:87:0c:5a:87:90:5c:4f:
                    fb:6f
                Exponent: 65537 (0x10001)
        X509v3 extensions:
            X509v3 Key Usage: critical
                Certificate Sign, CRL Sign
            X509v3 Basic Constraints: critical
                CA:TRUE
            X509v3 Subject Key Identifier:
                95:63:03:F2:7C:DE:AA:18:53:64:AB:41:15:9F:49:BC:AF:8A:A0:41
            X509v3 Authority Key Identifier:
                keyid:FE:65:5D:5D:41:03:22:CD:DF:48:61:06:F8:02:C1:E3:7F:ED:CF:51

            X509v3 Subject Alternative Name:
                DNS:jss-jwt-server, URI:TSI:cluster-name:mycluster, URI:tsi:regionName:eu-de, URI:tsi:datacenter:fra02
    Signature Algorithm: sha256WithRSAEncryption

         42:ba:21:9f:cc:cc:22:ba:12:6b:23:3c:3a:1c:70:6c:71:96:
```

Sample view of the issuing CA chain:
```
openssl x509 -in chain.pem -noout -text
Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number:
            08:19:ce:af:5b:d4:21:0f:58:c2:9c:e6:56:d3:2f:88:54:73:58:dc
    Signature Algorithm: sha256WithRSAEncryption
        Issuer: CN=trusted-identity.ibm.com
        Validity
            Not Before: Apr 23 16:07:43 2020 GMT
            Not After : Apr 23 16:08:13 2021 GMT
        Subject: CN=trusted-identity.ibm.com
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                Public-Key: (2048 bit)
                Modulus:
                    00:b6:bf:d9:bd:52:26:20:bc:63:3d:1e:db:f2:3f:
                    2c:7f:dd:1d:10:66:5e:6e:e9:36:b7:bd:31:06:ed:
                    aa:00:14:95:a7:26:39:7d:94:bd:ba:e1:99:fe:35:
                    81:8a:3d:b4:0f:e9:41:4d:be:04:54:f9:37:00:8b:
                    88:01
                Exponent: 65537 (0x10001)
        X509v3 extensions:
            X509v3 Key Usage: critical
                Certificate Sign, CRL Sign
            X509v3 Basic Constraints: critical
                CA:TRUE
            X509v3 Subject Key Identifier:
                FE:65:5D:5D:41:03:22:CD:DF:48:61:06:F8:02:C1:E3:7F:ED:CF:51
            X509v3 Authority Key Identifier:
                keyid:FE:65:5D:5D:41:03:22:CD:DF:48:61:06:F8:02:C1:E3:7F:ED:CF:51

            X509v3 Subject Alternative Name:
                DNS:trusted-identity.ibm.com
    Signature Algorithm: sha256WithRSAEncryption
         21:1d:33:0a:a1:f5:b4:32:f5:6d:ee:85:55:4c:7a:99:63:0b:
         73:ff:6f:77:d7:b9:f6:1d:85:03:18:3a:71:c3:a7:06:45:d8:
```

*NOTES*

These don't work:
* alt_names - Subject Alternative Name is overwritten with `DNS:trusted-identity.ibm.com`
* other_sans - problems with setting up the role `allowed_other_sans`

```
vault write pki/roles/myrole allowed_other_sans="*;UTF8:*"
Success! Data written to: pki/roles/myrole

vault write pki/root/sign-intermediate csr=@node.csr ttl=438000h other_sans="1.3.6.1.4.1.52683;UTF8:Test"
Error writing data to pki/root/sign-intermediate: Error making API request.
URL: PUT http://ti-fra02.eu-de.containers.appdomain.cloud/v1/pki/root/sign-intermediate
Code: 400. Errors:
* other SAN OID 1.3.6.1.4.1.52683 not allowed by this role
```



Extract x5c the from intermediate CA
```bash

CERT=$(cat out1 | jq -r '.["data"].certificate' | grep -v '\-\-\-')
CHAIN=$(cat out1 | jq -r '.["data"].issuing_ca' | grep -v '\-\-\-')
echo "[\"${CERT}\",\"${CHAIN}\"]" > "$X5C"
```

-----------
## Alternative Way

Signing the CSR manually on the server with Root CA

```
# create Root CA on the server
openssl req -new -newkey rsa:2048 -nodes -x509 -subj "/CN=vault-server" -keyout ca.key -out ca.crt
# view the new CA
openssl x509 -in ca.crt -noout -text
```
Get the CSR for signing. Extract the Node identity and create node.cnf

```bash
openssl req -in node.csr -noout -text | grep "URI.TSI"
                URI:TSI:cluster-name:mycluster, URI:TSI:regionName:eu-de, URI:TSI:datacenter:fra02

# create node.cnf
cat > "node.cnf" << EOF
[req]
req_extensions = v3_req

[v3_req]
subjectAltName= @alt_names

[alt_names]
URI.1 = TSI:cluster-name:mycluster
URI.2 = TSI:regionName:eu-de
URI.3 = TSI:datacenter:fra02
EOF

# sign CSR using Root CA and the node.cnf file
```console
openssl x509 -req -in ../client/node.csr -CA ca.crt -CAkey ca.key -CAcreateserial -extensions v3_req -extfile node.cnf -out node.pem

# sample output:
openssl x509 -in node.pem -noout -text
Certificate:
    Data:
        Version: 3 (0x2)
        Serial Number: 11606569237204736990 (0xa112d13a1c909bde)
    Signature Algorithm: sha1WithRSAEncryption
        Issuer: CN=vault-server
        Validity
            Not Before: May  1 16:03:59 2020 GMT
            Not After : May 31 16:03:59 2020 GMT
        Subject: CN=jss-jwt-server
        Subject Public Key Info:
            Public Key Algorithm: rsaEncryption
                Public-Key: (2048 bit)
                Modulus:
                    00:ec:21:be:9a:9d:a8:7d:2b:69:69:57:de:eb:7d:
                    87:27:4f:89:62:1a:c9:8b:e1:a0:5d:e0:2a:92:28:
                    a5:39:16:00:40:d3:5a:0e:87:0c:5a:87:90:5c:4f:
                    fb:6f
                Exponent: 65537 (0x10001)
        X509v3 extensions:
            X509v3 Subject Alternative Name:
                URI:TSI:cluster-name:mycluster, URI:TSI:regionName:eu-de, URI:TSI:datacenter:fra02
    Signature Algorithm: sha1WithRSAEncryption
         01:be:f1:fd:45:31:7d:d8:cb:96:69:9a:8f:17:1c:cb:dc:7d:
         6f:2f:42:25:e9:8e:a6:7b:92:5d:41:1e:1c:07:0a:4d:fd:9f:
         5c:7d:a4:18:df:3a:4f:2f:3f:6c:28:1c:13:a7:dc:70:9d:6c:
```


## Other related materials
https://jvns.ca/blog/2017/01/31/whats-tls/
