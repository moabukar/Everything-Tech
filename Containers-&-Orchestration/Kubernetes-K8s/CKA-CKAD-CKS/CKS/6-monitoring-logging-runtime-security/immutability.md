#### Create a pod with ReadOnlyRootFileSystem 

```sh

vi immutable-pod.yaml

```

```sh

apiVersion: v1
kind: Pod
metadata:
  name: immutable-pod
spec:
  containers:
  - name: ubuntu
    image: ubuntu
    command: ["sleep","3600"]
    securityContext:
      readOnlyRootFilesystem: true

```

#### Verify

```sh

kubectl exec -it immutable-pod -- bash
touch file.txt

```
