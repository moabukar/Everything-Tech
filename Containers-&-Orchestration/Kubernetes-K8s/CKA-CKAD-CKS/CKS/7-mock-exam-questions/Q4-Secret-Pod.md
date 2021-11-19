### Question - Secret mounted on pod (volume)

Create a secret named "secret1" in the "seminar" namespace using details:

user: admin
pass : P455W0RD

Create a pod named “secretpod” in the namespace "seminar" using the image nginx and mount the secret as a volume with "readOnly" option.

### Solution

- [Secret K8s docs](https://kubernetes.io/docs/concepts/configuration/secret/)

#### 1 - Create namespace

```sh

kubectl create ns seminar

```

#### 1 - Create a generic secret

```sh

kubectl -n seminar create secret generic secret1 --from-literal=user=admin --from-literal=pass=P455W0RD

```

#### 2 - Create a pod and mount secret as volume

```sh

vi ~/secret-pod.yaml

apiVersion: v1
kind: Pod
metadata:
  name: secretpod ## pod name
  namespace: seminar ## correct namespace
spec:
  containers:
  - name: secretpod
    image: nginx
    volumeMounts:
    - name: vol1 
      mountPath: /etc/foo
      readOnly: true ## select this option as true
  volumes:
  - name: vol1
    secret:
      secretName: secret1 ## use created secret here

kubectl apply -f ~/secret-pod.yaml

```
