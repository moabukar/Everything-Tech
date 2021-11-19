#### K8s Documentation

[Ingress K8s docs](https://kubernetes.io/docs/concepts/services-networking/ingress/)

#### 1 - Create self signed certificate for domain

```sh

mkdir ingress-test
cd ingress-test
openssl req -x509 -nodes -days 300 -newkey rsa:2048 -keyout ingress.key -out ingress.crt -subj "/CN=ingress.external/O=security"

```

#### 2 - Create k8s TLS based secret

```sh

kubectl create secret tls tls-cert --key ingress.key --cert ingress.crt

kubectl get secret tls-cert -o yaml

```

#### 3 - Create Kubernetes Ingress with TLS

```sh

vi ingress-tls.yaml

```

```sh

apiVersion: networking.k8s.io/v1
kind: Ingress
metadata:
  name: ingress-test
  annotations:
     nginx.ingress.kubernetes.io/rewrite-target: /
spec:
  tls:
  - hosts:
      - ingress.external
    secretName: tls-cert
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

kubectl apply -f ingress-tls.yaml

```

#### 4 - Make a request to controller

```sh

kubectl get svc -n svc

curl -kv https://ingress.external:32794

```
