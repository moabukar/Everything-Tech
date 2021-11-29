# Required tools for running Local K8's:
* Docker
* Minikube
* Kubernetes client (kubectl)


# To start kubernetes cluster using minikube
## Configure worker node on your Laptop:
```
minikube start
```
## Check the status worker node:
```
minikube status
```
## Check k8's master cluster information:
### Normal Info:
```
kubectl cluster-info
```
### Detailed Info
```
kubectl cluster-info dump
```

## List down any resource of k8's cluster:
```
kubectl get <resource_type>,<another_resource_type>
E.g:
kubectl get nodes || kubectl get no
```
## Check the details of any resource:
```
kubectl describe <resource_type> <resource_name>
E.g:
kubectl describe node <node_name>
```
## Make longer k8's commands in shorter commands using Linux utility "alias":
```
alias kgn = "kubectl get no"
```
## Create any resource from YAML file:
```
kubectl create -f <file_name.yaml>
E.g:
kubectl create -f myfirstpod.yaml
```
## List any resource with monitoring mode:
```
kubectl get <resource_type> -w
                            ||
                             ===> watch
E.g:
kubectl get pods -w
```

## Delete any resource:
```
kubectl delete <resource_type> <resource_name>
E.g:
kubectl delete pod <pod_name>
```

## Get the details of existing resources:
```
kubectl get <resource_type> <resource_name> -o yaml
OR
kubectl get <resource_type> <resource_name> -o json
```


# Port Forwarding:
## Forward internal port to external world:
```
kubectl port-forward <pod_name> 6500:80
                                 ||  ||
       Host port (External) <<====    ====>> Internal port that you have given in YAML file
```

## Get into Pod's container:
```
kubectl exec <pod_name> -c <container_name> -it -- <shells or commands>
```
## Creating from CLI with YAML:
```
kubectl run <pod_name> --image=<image_name> --port=80 --restart=Never
							      ||
							       ====>> must use this flag for pod creation
```
## List resources with Labels:
```
kubectl get <resource_type> --show-labels
```
## List resources with Labels columns:
```
kubectl get <resource_type> -L <label_name>,<other_label_name>,...
```
## Label a resource at runtime:
```
kubectl label <resource_type> <resource_name> <label_key>=<label_value> <other_label_key>=<other_label_value>...
```
## Change the Label's value:
```
kubectl label <resource_type> <resource_name> <label_key>=<label_value> --overwrite
```
## Remove the Labels from resource:
```
kubectl label <resource_type> <resource_name> <label_key>-
```


# Label Selector:
## Filtering Resources:
```
kubectl get <resource_type> -l <key>=<value> --show-labels || -L <key>
E.g:
kubectl get pod -l type=frontend, --show-labels
kubectl get pod -l type!=frontend -L type
kubectl get pod -l type!=frontend,env=production --show-labels
``` 
## List resource that contain <Key>:
```
kubectl get <resource_type> -l env --show-labels
kubectl get <resource_type> -l '!env' --show-labels
```
## List resource who matched with any of value:
```
kubectl get <resource_type> -l '<key> in (<value>,<value>)' --show-labels
E.g:
kubectl get pod -l 'type in (frontend,backend)' --show-labels
kubectl get pod -l 'type notin (frontend,backend)' --show-labels
```
## Pod Scheduling with Node selector:
### First get the nodes
```
kubectl get nodes
```
### Then assign label to node name
```
kubectl label node <node_name> <label_key>=<label_value>
```


# Annotations:
## Assign annotation using CLI:
```
kubectl annotate <resource_type> <resource_name> appdescription="This is example of annotation in k8's"
```


# Namespaces (NS): 
```
--namespace == -n
```
## Creating Namespaces:
```
kubectl create namespace <ns_name>
```
## List Namespaces
```
kubectl get ns
```
## List the resource of specific NS:
```
kubectl get <resource_type> -n=<ns_name>
``` 
## Put resource in NS using CLI:
```
kubectl run <resource_name> --image=<image_name> --port=80 --restart=Never --namespace=development
```
## List the resource from all NS:
```
kubectl get <resource_type> --all-namespaces
```
## Delete the resource from NS:
```
kubectl delete <resource_type> <resource_name> -n=<ns_name>
```
## Delete resource using Labels:
```
kubectl delete <resource_type> -l <label_key>=<label_value>
```
## Delete resource from current NS:
```
kubectl delete <resource_type> --all
```


# ReplicaSet:
### Replica Set will be create using YAML file
## Delete the RS without deleting the pods:
```
kubectl delete rs <rs_name> --cascade=orphan
```
## Edit the RS count or Edit anything from any resource:
```
KUBE_EDITOR="nano" kubectl edit <resource_type> <resource_name>
```
## Change the RS replicas count using CLI:
```
kubectl scale rs <rs_name> --replicas=5
```
## To check logs of pods:
```
kubectl logs <pod_name>
```


# Services:
### Services will be create using YAML file
## List the services:
```
kubectl get service || svc
```
### Note: External IP will not be provided when using "minikube".

## Check the IP of minikube:
```
minikube ip
```
#### Now run the <ip>:<port_num> which was provided by the service you have created. It will redirect you towards any of the Pod which was grouped.

## LoadBalancer example using CLI:
```
kubectl expose <resource_type> <resource_name> --name=<service_name> --selector=<key>=<value> --port:<pod_port> --target-port=<container_port> --type=<service_type>
```
### Note: There are many <serivce_type> supported by k8's such as, LoadBalancer, ClusterIP, NodePort, ExternalName

## Access a minikube terminal:
```
minikube ssh
```


# ConfigMap:
### ConfigMap will be create using YAML file
## Create a configmap using CLI with literals:
```
kubectl create configmap <cm_name> --from-literal=<key>=<value> --from-literal=<another_key>=<another_value>
```
## List the ConfigMap:
```
kubectl get configmap | cm
```
## Create ConfigMap from file using CLI:
### This command will make filename as a key.
```
kubectl create cm <cm_name> --from-file=<filename>
```
## Create ConfigMap from file with desire key:
```
kubectl create cm <cm_name> --from-file=<key>=<filename>
```
## Create ConfigMap from ENV file:
```
kubectl create cm <cm_name> --from-env-file=<file.env>
```
## Expose a Pod port using CLI:
```
kubectl expose pod <pod_name> --port=8080 --target=8080 --type=LoadBalancer
```
## Check port number from service list:
```
kubectl get svc
```
## Now Check IP and port of minikube or external IP if not using minikube:
```
minikube ip
```
## Now run on the Browser:
```
<minikube_ip>:<port_number_from_listing_service>
E.g:
192.168.49.2:31341
```


# Secret:
### Secrets will be create using YAML file
## Create a secret from literals uing CLI:
```
kubectl create secret generic <sec_name> --from-literals=<key>=<value> --from-literals=<another_key>=<another_value>
```
#### When you check insight of this secret, you will uable to see their values.
## Watch secrets encrypted values:
```
kubectl get secret <sec_name> -o yaml
```
## List the secrets:
```
kubectl get secret
```
## You can decrypt the base64 encrypted values:
```
echo <encrypted_value> | base64 -d
```
## Create a Secret from file:
```
kubectl create secret generic <sec_name> --from-file=<key>=<filename>
```
## List resource in order to Timestamp:
```
kubectl get <resource_type> --sort-by=metadata.creationTimestamp
```


# Deployment:
### Deployments will be create using YAML file
## List the deployments:
```
kubectl get deploy
```
## Expost the ports to deployment:
```
kubectl expose deploy <deploy_name> --port=8080 --target-port=80 --type-LoadBalancer --selector=<key>=<value>
```
## Update the image of deployed Replica Set:
```
kubectl set image deploy <deploy_name> <container_name_in_deploy>=<updated_image>
```
## Rollback the previous version of the image:
```
kubectl rollout undo deploy <deploy_name> --to-revision=<version_number>
```
#### You can get the <version_number> from below command.
## History of All deployments of deploy:
```
kubectl rollout history deploy <deploy_name>
```
## Pause the deployment in b/w:
```
kubectl rollout pause deploy <deploy_name>
```
## Resume the deployment:
```
kubectl rollout resume deploy <deploy_name>
```
## Check the status of deployment:
```
kubectl rollout status deploy <deploy_name>
```
## Check the history of specific revision of deployment:
```
kubectl rollout history deploy <deploy_name> --revision=1
```
