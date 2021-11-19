#### Creating generic secret - using literal method 

```sh

kubectl create secret generic literalsecret --from-literal=pass=P455W0RD

```

#### Creating generic secret - using file method

```sh

kubectl create secret generic filesecret --from-file=./credentials.yml

```

#### View secret using CLI

```sh

kubectl get secret literalsecret -o yaml

```

##### vi secret-encoded.yaml

```sh

apiVersion: v1
data:
  pass: UDQ1NVcwUkQ=
  user: YWRtaW4=
kind: Secret
metadata:
  creationTimestamp: "2021-11-12T17:12:55Z"
  name: literalsecret
  namespace: default
type: Opaque

```

##### vi secret-string.yaml

```sh

apiVersion: v1
kind: Secret
metadata:
  name: literalsecret
type: Opaque
stringData:
   username: admin
   password: P455W0RD

```
