### Question - PSP

Create a Pod Security Policy with the following conditions:

- PSP name : pod-psp
- Not allow privleged pods
- Allow volumes to mount on pod: secret, configMap
- seLinux, runAsUser, fsGroup: RunAsAny

### Solution

- [PSP K8s docs](https://kubernetes.io/docs/concepts/policy/pod-security-policy/)

#### 1 - Enable PSP in kube-api server

Edit the /etc/kubernetes/manifests/kube-api-server.yaml file and add the below config.

```sh

  --enable-admission-plugins=NodeRestriction,PodSecurityPolicy ## add in "PodSecurityPolicy"

```

#### 2 - Create the PSP

```sh

vi ~/psp.yaml

apiVersion: policy/v1beta1
kind: PodSecurityPolicy
metadata:
  name: pod-psp
spec:
  privileged: false ## ensure this is set to false
  seLinux:
    rule: RunAsAny
  runAsUser:
    rule: RunAsAny
  supplementalGroups:
    rule: RunAsAny
  fsGroup:
    rule: RunAsAny
  volumes:  ## apply the correct volumes as specified in the question
  - configMap
  - secret

kubectl apply -f ~/psp.yaml

```

#### 3 - Apply PSP on pod

```sh

vi ~/pod.yaml

apiVersion: v1
kind: Pod
metadata:
    name: psp-pod
spec:
    containers:
    - name: app
      image: ubuntu
      command: ["sleep" , "3600"]
    securityContext:
      privileged: False ## non-privileged
      runAsUser: 0
    volumes:
    - name: data-volume
      hostPath:
        path: '/data'
        type: Directory

kubectl apply -f ~/pod.yaml

```
