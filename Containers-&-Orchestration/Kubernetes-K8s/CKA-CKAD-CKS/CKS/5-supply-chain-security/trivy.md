### Trivy official documentation and repo

[Trivy GitHub repo](https://github.com/aquasecurity/trivy)

#### Trivy installation

```sh

curl -sfL https://raw.githubusercontent.com/aquasecurity/trivy/master/contrib/install.sh | sh -s -- -b /usr/local/bin

```

#### Running trivy to scan an image

```sh

trivy image nginx:1.16

```
