#### Service account list

```sh

kubectl get sa

```

#### When a namespace is created, a service account is created automatically

```sh

kubectl create namespace jupiter
kubectl get sa -n jupiter

```

#### Verify SA Token Mount in Pods

```sh

kubectl run nginx-pod --image=nginx
kubectl exec -it nginx-pod -- bash
cd /run/secrets/kubernetes.io/serviceaccount
cat token

```

#### Create a pod with a custom service account

```sh

kubectl create sa test-sa
kubectl run custom-sa-pod --image=nginx --serviceaccount="test-sa"

```
