### Question - RuntimeClass

Create a RuntimeClass named "gVisor" using the handler "runsc"

Create a pod named "gVisor-pod" that uses the nginx image in the namespace "test" which utilises the runtime class "gVisor"

### Solution

- [RuntimeClass K8s docs](https://kubernetes.io/docs/concepts/containers/runtime-class/)

#### 1 - Install RuntimeClass for gVisor

```sh

apiVersion: node.k8s.io/v1  # RuntimeClass is defined in the node.k8s.io API group
kind: RuntimeClass
metadata:
  name: gVisor  # The name the RuntimeClass will be referenced by
  # RuntimeClass is a non-namespaced resource
handler: runsc  # The name of the corresponding CRI configuration

```

#### 2 - Create pod using the gVisor runtimeClass

```sh

vi ~/gVisor-pod.yaml

apiVersion: v1
kind: Pod
metadata:
  labels:
    run: gVisor-pod
  name: gvisor-pod
  namespace: test
spec:
  runtimeClassName: gvisor ## use created runtimeclass here
  containers:
  - image: nginx
    name: gVisor-pod


kubectl apply -f ~/gVisor-pod.yaml

```

#### Create another pod to compare differences in dmesg ouput

```sh

kubectl run nginx-test --image=nginx

```

#### Verifying dmesg output

##### RuntimeClass pod

```sh

kubectl exec -it gVisor-pod -- bash dmesg

```

##### Test pod

```sh

kubectl exec -it nginx-test --bash dmesg

```
