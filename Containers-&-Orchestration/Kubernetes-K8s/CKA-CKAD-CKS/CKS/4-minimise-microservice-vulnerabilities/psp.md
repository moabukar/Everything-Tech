#### 1 - Create permissive pod security policy

```sh

mkdir /root/psp
cd /root/psp

```

```sh

vi permissive-psp.yaml

```

```sh

apiVersion: policy/v1beta1
kind: PodSecurityPolicy
metadata:
  name: permissive-psp
  annotations:
    seccomp.security.alpha.kubernetes.io/allowedProfileNames: '*'
spec:
  privileged: true
  allowPrivilegeEscalation: true
  allowedCapabilities:
  - '*'
  volumes:
  - '*'
  hostNetwork: true
  hostPorts:
  - min: 0
    max: 65535
  hostIPC: true
  hostPID: true
  runAsUser:
    rule: 'RunAsAny'
  seLinux:
    rule: 'RunAsAny'
  supplementalGroups:
    rule: 'RunAsAny'
  fsGroup:
    rule: 'RunAsAny'

```

```sh

kubectl apply -f permissive-psp.yaml

```

#### 2 - Create cluster role for psp

```sh

vi permissive-psp-cluster-role.yaml

```

```sh

apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRole
metadata:
  name: permissive-cluster-role
rules:
- apiGroups: ['policy']
  resources: ['podsecuritypolicies']
  verbs:     ['use']
  resourceNames:
  - permissive-psp

```

```sh

kubectl apply -f permissive-psp-cluster-role.yaml

```

#### 3 - Create cluster role binding and bind to PSP cluster role

```sh

vi permissive-psp-cluster-role-binding.yaml

```

```sh

apiVersion: rbac.authorization.k8s.io/v1
kind: ClusterRoleBinding
metadata:
    name: permissive-psp-cluster-bind
roleRef:
  apiGroup: rbac.authorization.k8s.io
  kind: ClusterRole
  name: permissive-cluster-role
subjects:
  - kind: Group
    apiGroup: rbac.authorization.k8s.io
    name: system:authenticated

```

```sh

kubectl apply -f permissive-psp-cluster-role-binding.yaml

```

#### 4: Enable PSP admission controller on kube-api server

```sh

vi /etc/kubernetes/manifests/kube-apiserver.yaml

```

```sh

add below line on the kube-api server

--enable-admission-plugins=PodSecurityPolicy

```

```sh

watch pods

kubectl get pods -n kube-system -w

```
