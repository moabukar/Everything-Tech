#### 1 - Create certificate for Will

```sh

mkdir /root/certs
cd /root/certs

```

```sh

openssl genrsa -out will.key 2048
openssl req -new -key will.key -subj "/CN=will/O=developers" -out will.csr

```

#### 2 - Create CertificateSigningRequest (CSR)

```sh

cat <<EOF | kubectl apply -f -
apiVersion: certificates.k8s.io/v1
kind: CertificateSigningRequest
metadata:
  name: will
spec:
  groups:
  - system:authenticated
  request: $(cat will.csr | base64 | tr -d '\n')
  signerName: kubernetes.io/kube-apiserver-client
  usages:
  - client auth
EOF

```

```sh

kubectl certificate approve will
kubectl get csr will -o jsonpath='{.status.certificate}' | base64 --decode > will.crt

```

#### 3 - Create user for Will and move certificates:

```sh

useradd -m will -s /bin/bash
cp will.crt will.key /home/will
cp /etc/kubernetes/pki/ca.crt /home/will
chown -R will.will /home/will

```

```sh

kubectl get pods --server=https://109.365.122.141:6443 --client-certificate /home/will/will.crt --certificate-authority /home/will/ca.crt --client-key /home/will/will.key

```
