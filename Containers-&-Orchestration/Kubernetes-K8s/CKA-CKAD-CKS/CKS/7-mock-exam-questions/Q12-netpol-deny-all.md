## Question - Network Policy

Create a network policy called "np-restriction" to the pod "nginx-pod" in the namespace "moon"

Only allow pods to connect to the pod "nginx-pod":

- Pods in the namespace "hello" (kubectl get ns --show-labels)
- Pods with label "app:backend" in any namespace

### Solution

- [Network Policy K8s docs](https://kubernetes.io/docs/concepts/services-networking/network-policies/)

#### 1 - Apply network policy based on conditions

```sh

kubectl get ns hello --show-labels ## Namespace hello has label "ns: test"

vi ~/netpol.yaml

apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: np-restriction
  namespace: moon
spec:
  podSelector:
    matchLabels:
      run: nginx-pod
  policyTypes:
  - Ingress
  ingress:
  - from:
    - namespaceSelector:
        matchLabels:
          ns: test ## use the label that the namespace uses
    - podSelector:
        matchLabels:
          app: backend ## use the same label as specified on the 2nd spec
    ports:
    - protocol: TCP
      port: 6379

kubectl apply -f ~/netpol.yaml

```
