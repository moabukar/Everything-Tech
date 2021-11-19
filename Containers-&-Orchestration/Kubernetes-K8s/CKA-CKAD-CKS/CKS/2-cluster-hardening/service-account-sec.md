### K8s Documentation 

[Configuring Service Accounts](https://kubernetes.io/docs/tasks/configure-pod-container/configure-service-account/)

#### Method 1 - Opting out of automounting at Service Account level

```sh

apiVersion: v1
kind: ServiceAccount
metadata:
  name: example-sa
automountServiceAccountToken: false

```

#### Method 2 - Opting out of automounting at Pod level

```sh

apiVersion: v1
kind: ServiceAccount
metadata:
  name: example-sa
automountServiceAccountToken: false

```

```sh

vi pod-sa.yaml

```

```sh

apiVersion: v1
kind: Pod
metadata:
  name: pod-sa
spec:
  automountServiceAccountToken: false
  containers:
  - image: nginx
    name: pod-sa

```

```sh

kubectl apply -f pod-sa.yaml

```
