#### 1 - Create simple pod with nginx image

```sh

kubectl run pod1 --image=nginx -l=app=nginx

```

#### 2 - Create service

```sh

kubectl expose pod pod1 --name=svc --port=80

```

#### 3 - Install nginx controller

```sh

kubectl apply -f https://raw.githubusercontent.com/kubernetes/ingress-nginx/controller-v0.41.2/deploy/static/provider/baremetal/deploy.yaml

```

#### 4 - Creating ingress

```sh

vi ingress1.yaml

```

```sh

apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: ingress-example
  annotations:
     nginx.ingress.kubernetes.io/rewrite-target: /
spec:
  rules:
  - host: ingress.external
    http:
      paths:
      - pathType: Prefix
        path: "/"
        backend:
          service:
            name: svc
            port:
              number: 80

```

```sh

kubectl apply -f ingress1.yaml

```

#### 5 - Verify ingress (DOUBLE CHECK THIS PART)

```sh

kubectl get svc -n ingress-nginx

ifconfig eth0

```

Add hosts

```sh

vi /etc/hosts

```

```sh

curl http://ingeress.external:30433

```
