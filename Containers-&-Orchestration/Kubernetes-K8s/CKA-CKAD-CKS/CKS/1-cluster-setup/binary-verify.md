#### Kubernetes GitHub repo

<https://github.com/kubernetes/kubernetes/releases>

#### 1 - Binary download

```sh
  
wget https://github.com/kubernetes/kubernetes/releases/download/v1.22.3/kubernetes.tar.gz

```

#### 2 - Verify the binary

```sh

  apt install hashalot
  sha256 kubernetes.tar.gz
  sha512sum kubernetes.tar.gz

```
