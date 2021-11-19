### Question - Kube-bench (1) 

You have just read the kube-bench assessment report. Fix the tests that have FAIL status for the worker node configuration.

The kube-bench report gives issues such as (or close to):

- "authorisation mode is set as Always allowed"
- "Kernel defaults are not protected"

Make changes to the /var/lib/kubelet/config.yaml

### Solution

- [Kube-bench GitHub](https://github.com/aquasecurity/kube-bench)

#### 1 - Update the kubelet config

```sh

vi /var/lib/kubelet/config.yaml

change "authorization mode"

authorization:
  mode: Webhook

```

#### 2 - Update protectKernelDefaults in kubelet config

```sh

vi /var/lib/kubelet/config.yaml

protectKernelDefaults: true

now you may exit the kubelet with :wq! and do "systemctl kubelet restart"

```
