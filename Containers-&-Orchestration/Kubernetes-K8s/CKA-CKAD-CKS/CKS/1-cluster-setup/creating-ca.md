#### 1 - Creating a private key for CA

```sh

mkdir /root/certs
cd /root/certs

```

```sh

openssl genrsa -out ca1.key 2048

```

#### 2 -  Creating CSR

```sh

openssl req -new -key ca1.key -subj "/CN=KUBERNETES-CA" -out ca1.csr

```

#### 3 - Self-sign CSR

```sh

openssl x509 -req -in ca1.csr -signkey ca1.key -CAcreateserial  -out ca1.crt -days 300

```

#### 4 - Optional - Remove the CSR

```sh

rm -f ca1.csr

```
