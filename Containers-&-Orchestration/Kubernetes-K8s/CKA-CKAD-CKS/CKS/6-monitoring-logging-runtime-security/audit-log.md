#### 1 - Create a directory for storing the audit logs and audit policy 

```sh

cd /etc/kubernetes/
mkdir audit

```

#### 2 - Mount the directory as HostPath Volumes in kube-apiserver

```sh

  - hostPath:
      path: /etc/kubernetes/audit
      type: DirectoryOrCreate
    name: auditing

    - mountPath: /etc/kubernetes/audit
      name: auditing

```

#### 3 - Add auditing related configuration in kube-apiserver

```sh

- --audit-policy-file=/etc/kubernetes/audit/audit-policy.yaml
- --audit-log-path=/etc/kubernetes/audit/audit.log

```

#### 4 - Audit policy

```sh

apiVersion: audit.k8s.io/v1
kind: Policy
omitStages:
  - "RequestReceived"
rules:
  - level: None
    resources:
    - group: ""
      resources: ["secrets"]
    namespaces: ["kube-system"]

  - level: None
    users: ["system:kube-controller-manager"]
    resources:
    - group: ""
      resources: ["secrets","configmaps"]

  - level: RequestResponse
    resources:
    - group: ""
      resources: ["secrets"]
  - level: Metadata
    omitStages:
      - "RequestReceived"

```
