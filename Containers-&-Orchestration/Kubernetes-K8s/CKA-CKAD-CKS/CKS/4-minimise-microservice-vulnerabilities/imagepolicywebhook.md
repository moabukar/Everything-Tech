### K8s Documentation

[Admission Controllers K8s docs](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/)

Create directory for storing the configs:

```sh

mkdir /etc/kubernetes/confighook

```

```sh

vi config.yaml

```

```sh

apiVersion: apiserver.config.k8s.io/v1
kind: AdmissionConfiguration
plugins:
- name: ImagePolicyWebhook
  configuration:
    imagePolicy:
      kubeConfigFile: /etc/kubernetes/confighook/webhook.kubeconfig
      allowTTL: 50
      denyTTL: 50
      retryBackoff: 500
      defaultAllow: true

```

```sh

vi webhook.kubeconfig

```

```sh

apiVersion: v1
kind: Config
# clusters refers to the remote service.
clusters:
- name: imagepolicy-webhook
  cluster:
    certificate-authority: /etc/kubernetes/pki/ca.crt   # CA for verifying the remote service.
    server: https://images.demo.com/policy # URL of remote service to query. Must use 'https'.
contexts:
- context:
    cluster: imagepolicy-webhook
    user: api-server
  name: demo-context
current-context: demo-context

# users refers to the API server's webhook configuration.
users:
- name: api-server
  user:
    client-certificate: /etc/kubernetes/pki/apiserver.crt # cert for the webhook admission controller to use
    client-key: /etc/kubernetes/pki/apiserver.key        # key matching the cert

    - mountPath: /etc/kubernetes/confighook
      name: admission-controller
      readOnly: true

  - hostPath:
      path: /etc/kubernetes/confighook
      type: DirectoryOrCreate
    name: admission-controller
```

#### Update the kube-api server manifesst with these lines

Note: always ensure you make a backup copy of your kube-apiserver befor making any changes. Anything can go wrong and your api-server may be down if an incorrect config is changed.

```sh

vi /etc/kubernetes/manifests/kube-apiserver.yaml

```

```sh

- --enable-admission-plugins=NodeRestriction,ImagePolicyWebhook
- --admission-control-config-file=/etc/kubernetes/confighook/config.yaml

```
