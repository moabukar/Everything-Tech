#### 1 - Installing kube-bench (NEEDS TO BE UPDATED)

```sh

curl -L https://github.com/aquasecurity/kube-bench/releases/download/v0.6.5/kube-bench_0.6.5_linux_amd64.tar.gz -o kube-bench_0.6.5_linux_amd64.tar.gz

tar -xvf kube-bench_0.6.5_linux_amd64.tar.gz

```

The kube-bench version used here is v0.6.5
Make sure to change the version depending on your needs.

#### 2 - Running a kube-bench test

```sh

./kube-bench --config-dir `pwd`/cfg --config `pwd`/cfg/config.yaml

```

#### 3 - check results and fix any failiures

kube-bench will analyse 5 main configurations and policies as listed below:

```sh

[INFO] 1 Master Node Security Configuration
[INFO] 2 Etcd Node Configuration
[INFO] 3 Control Plane Configuration
[INFO] 4 Worker Node Security Configuration
[INFO] 5 MKubernetes Policies

```

For each section it will give you a summary, breakdown and remedies similar to this:

```sh

[INFO] 3 Control Plane Configuration
[INFO] 3.1 Authentication and Authorization
[WARN] 3.1.1 Client certificate authentication should not be used for users (Manual)
[INFO] 3.2 Logging
[WARN] 3.2.1 Ensure that a minimal audit policy is created (Manual)
[WARN] 3.2.2 Ensure that the audit policy covers key security concerns (Manual)

== Remediations ==
3.1.1 Alternative mechanisms provided by Kubernetes such as the use of OIDC should be
implemented in place of client certificates.

3.2.1 Create an audit policy file for your cluster.

3.2.2 Consider modification of the audit policy in use on the cluster to include these items, at a
minimum.

== Summary ==
0 checks PASS
0 checks FAIL
3 checks WARN
0 checks INFO

```
