#### Step 1 - Configuring audit from kube-api

```sh

Set parameters here based on requirements

--audit-policy-file=/root/certs/log-audit.yaml \
--audit-log-path=/var/log/api-audit.log \
--audit-log-maxage=20 \
--audit-log-maxbackup=10 \
--audit-log-maxsize=300 \

```

#### Step 2 - Create Audit Policy File

```sh

vi ~/certs/log.yaml

```

```sh

Simple Audit policy that logs everything at the Metadata level

apiVersion: audit.k8s.io/v1
kind: Policy
rules:
- level: Metadata

```

#### Step 3: Verifying audit policy

```sh

systemctl daemon-reload
systemctl restart kube-apiserver
less /var/log/api-audit.log
cd /var/log
grep -E test api-audit.log

```
