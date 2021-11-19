#### Falco Documentation 

- [Falco docs](https://falco.org/docs/)
- [Falco supported field](https://falco.org/docs/rules/supported-fields/)

#### Steps for installing Falco

```sh

curl -s https://falco.org/repo/falcosecurity-3672BA8F.asc | apt-key add -
echo "deb https://dl.bintray.com/falcosecurity/deb stable main" | tee -a /etc/apt/sources.list.d/falcosecurity.list
apt-get -y install linux-headers-$(uname -r)
apt-get update && apt-get install -y falco

```

#### Start falco

```sh

falco

```

#### Test simple falco rule

```sh

kubectl run nginx-pod --image=nginx
kubectl exec -it nginx-pod -- bash

```

```sh

mkdir /bin/tmp-1
cat /etc/shadow

```
