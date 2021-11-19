### Question - ImagePolicyWebhook 

We are required to deploy an ImagePolicyWebhook admission controller to secure the deployments that are made on our cluster by ensuring the below:

- Update and fix the issue on the /etc/kubernetes/pki/admission_configuration.yaml which will be used by ImagePolicyWebhook
- Ensure that the policy is set to implicit deny.
- Ensure the plugin is on the kube-api server

### Solution

- [Kube-Admission Controllers](https://kubernetes.io/docs/reference/access-authn-authz/admission-controllers/)

#### 1 - Update and fix the issue on the admission config

```sh

apiVersion: apiserver.config.k8s.io/v1
kind: AdmissionConfiguration
plugins:
- name: ImagePolicyWebhook
  configuration:
    imagePolicy:
      kubeConfigFile: /etc/kubernetes/pki/admission_kube_config.yaml ## ensure you apply the correct path
      allowTTL: 50
      denyTTL: 50
      retryBackoff: 500
      defaultAllow: false ## ensure this is set to false

```

#### 2 - Update the kube-api server

```sh

- --enable-admission-plugins=NodeRestriction,ImagePolicyWebhook ## add ImagePolicyWebhook here
- --admission-control-config-file=/etc/kubernetes/pki/admission_configuration.yaml ## ensure correct config file path is used

Once you leave the vim editor and save, the API server will automatically restart and pick up the configuration.

```
