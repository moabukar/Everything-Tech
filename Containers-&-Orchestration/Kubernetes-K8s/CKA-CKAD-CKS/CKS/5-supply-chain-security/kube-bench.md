### Documentation

[Kube-bench binary](https://github.com/aquasecurity/kube-bench#download-and-install-binaries)

#### Installation steps

```sh

curl -L https://github.com/aquasecurity/kube-bench/releases/download/v0.6.5/kube-bench_0.6.5_linux_amd64.tar.gz -o kube-bench_0.6.5_linux_amd64.tar.gz

tar -xvf kube-bench_0.6.5_linux_amd64.tar.gz

```

```sh

sudo apt install ./kube-bench_0.6.5_linux_amd64.deb -f

```

#### Running kube-bench on master node

```sh

run below command to view security benchmakrs (failed & passed configs)

kube-bench master

```
