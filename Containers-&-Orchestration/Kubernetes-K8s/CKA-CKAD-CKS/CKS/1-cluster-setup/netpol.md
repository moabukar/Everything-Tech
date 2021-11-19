#### 1 - Creating Network Policy based on given requirements

A number of pods have been deployed in the "static" namespace. There is a pod called "nginx-pod" and its the backend for all three of our applications "application 1, application 2 and application 3". Create a network policy named "allow-net" that only allows incoming trafffic from all three applications to the "nginx-pod"

```sh

kubectl -n static get pods --show-labels ## this will display the labels allocated to application 1,2 and 3 and we will use these labels on our Network Policy.

kind: NetworkPolicy
apiVersion: networking.k8s.io/v1
metadata:
  name: allow-net
  namespace: static
spec:
  podSelector:
    matchLabels:
      tier: backend
      role: app-logic
  ingress:
  - from:
    - podSelector:
        matchLabels:
          name: application1
          tier: frontend
    - podSelector:
        matchLabels:
          name: application2
          tier: frontend
    - podSelector:
        matchLabels:
          name: application3
          tier: frontend

  ```
