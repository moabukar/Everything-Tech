### Question - Network Policy

Create a network policy named all-deny and it should deny all ingress and egress traffic.

### Solution

- [Network Policy K8s docs](https://kubernetes.io/docs/concepts/services-networking/network-policies/)

#### 1 - Enable PSP in kube-api server

```sh

vi ~/netpol.yaml

apiVersion: networking.k8s.io/v1
kind: NetworkPolicy
metadata:
  name: all-deny
spec:
  podSelector: {} ## selects all pods
  policyTypes: ## all pods are selected - both ingress and egress traffic not allowed to all pods
  - Ingress
  - Egress

kubectl apply -f ~/netpol.yaml

```
