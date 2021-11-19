#### Creating a sample apparmor profile

```sh

apparmor_parser -q <<EOF
#include <tunables/global>

profile k8s-apparmor-example-deny-write flags=(attach_disconnected) {
  #include <abstractions/base>

  file,

  # Deny all file writes.
  deny /** w,
}
EOF

```

#### Verifying the profile status

```sh

aa-status

```

#### Create a demo YAML based on Host

```sh

cd /root/apparmor1

```

```sh

vi pod-app-armor.yaml

```

```sh
apiVersion: v1
kind: Pod
metadata:
  name: pod-app-armor
  annotations:
    container.apparmor.security.beta.kubernetes.io/hello: localhost/k8s-apparmor-example-deny-write
spec:
  containers:
  - name: hello
    image: busybox
    command: [ "sh", "-c", "echo 'Hello AppArmor!' && sleep 1h" ]
```

```sh

kubectl apply -f pod-app-armor.yaml

```

#### Verifying

```sh

kubectl exec -it pod-app-armor -- sh
touch /tmp/file.yml

```
