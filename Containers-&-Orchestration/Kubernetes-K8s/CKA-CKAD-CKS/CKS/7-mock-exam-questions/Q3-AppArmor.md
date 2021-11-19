### Question - AppArmor 

A pod has been created in the "spectacle" namespace. However, there are a couple of issues with it:

- The pods has been created with privileged permissions
- it allows read access

- Use the AppArmor profile created at /etc/apparmor.d/spectacleapp
- create a service account in the spectacle namespace called "test-sa" and use this service account on the pod

### Solution

- [AppArmor K8s docs](https://kubernetes.io/docs/tutorials/clusters/apparmor/)

#### 1 - Load the AppArmor profile

```sh

apparmor_parser -q /etc/apparmor.d/spectacleapp

```

#### 2 - Create servive account

```sh

kubectl create ns spectacle ## create namespace if not already created

kubectl -n spectacle create sa test-sa

```

#### 3 - Create a pod using the AppArmor profile backend

```sh

vi ~/app-armor-pod.yaml

apiVersion: v1
kind: Pod
metadata:
  annotations:
    container.apparmor.security.beta.kubernetes.io/nginx: localhost/spectacleapp #Apply profile 'spectacleapp' on 'nginx' container
  labels:
    run: nginx
  name: apparmor-pod
  namespace: spectacle
spec:
  serviceAccount: test-sa ## use the created service account
  containers:
  - image: nginx:alpine
    name: nginx ## this container name needs to match the annotation "container.apparmor.security.beta.kubernetes.io/nginx"
    volumeMounts:
    - mountPath: /usr/share/nginx/html
      name: test-volume
  volumes:
  - name: test-volume
    hostPath:
      path: /data/pages
      type: Directory

kubectl apply -f ~/app-armor-pod.yaml

```
