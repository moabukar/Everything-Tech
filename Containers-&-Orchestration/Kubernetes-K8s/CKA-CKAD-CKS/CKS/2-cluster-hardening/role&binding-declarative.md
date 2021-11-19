### Question - Role & binding

Create a service account name seminar-sa on the namespace seminar

Create a new role named k8s-seminar in the namespace seminar which only allows create and update operations only on resources of type pods and deployments

Create a new rolebinding name k8s-seminar-bind binding to the newly created role to the service account created previously named seminar-sa.

### Option 2 - Declarative

#### 1 - Create a namespace

```sh

kubectl create ns seminar

```

#### 2 - Create service account

```sh

kubectl -n seminar create sa seminar-sa

```

#### 3 - Create role

```sh

apiVersion: rbac.authorization.k8s.io/v1
kind: Role
metadata:
  namespace: seminar
  name: k8s-seminar
rules:
- apiGroups: [""]
  resources: ["pods","deployments"]
  verbs: ["create","update"]

```

#### 4 - Create rolebinding

```sh

apiVersion: rbac.authorization.k8s.io/v1
kind: RoleBinding
metadata:
  name: k8s-seminar-pod
  namespace: seminar
subjects:
roleRef:
  kind: Role
  name: k8s-seminar
  apiGroup: rbac.authorization.k8s.io

```
