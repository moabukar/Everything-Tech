```
Exit container's interactive mode without stopping the container.
              Ctrl+P+Q
```
## List Image:
```
docker image ls
```

## Remove Image: 
```
docker rmi <container_id>
```

## Interactive Mode:
```
docker run -it alihasanahmedkhan/helloworld:latest sh
```

## Detach Mode:
```
docker run -d alihasanahmedkhan/helloworld:latest
```

## Attach to container:
```
docker attach <container_name>
```

## Start the exisiting container:
```
docker start <container_name>
```

## List Containers:
```
docker ps -a
```

## Execute an existing running container
```
docker container exec -it <container_id or name>
```

## Stop Containers:
```
docker container stop <container_id or name>
```

## Start existing Containers:
```
docker container start <container_id or name>
```

## Publishing port (For browser)
```
docker container run -d -p 5010:80 ammirpinger/helloworld:latest

                              ||  ||
                              ||  || 
                              ||  ||===> Container's Internal Port 
                              ||  
                              \/
                         External Port
                            (HOST)
```

## Rename the Image:
```
docker image tag <old_image_name or id> <new_name_of_image>
```

## Custom Container's name:
```
docker container run -d --name docker_app -p 5020:80 alihasanahmedkhan/helloworld:latest
```

## Build Dockerfile:
```
docker build -t <custom_name> .
```

## Run that docker image:
```
docker container run --name=<container_name> -d -p 8500:80 <image_id or image_name>
```

## Pushing Image to DockerHub:
```
docker image push <image_name or image_id>

OR
```
### First change the name of the image
```
docker image tag <old_tag> <docker_hub_account_name>/<old_tag>
docker image push <username>/<image_name>
```

## Make copy of existing image with different name:
```
docker tag <image_name> <new_image_name>
```

## Docker History: ====> NORMAL INFORMATION
```
docker history <image_name>
```

## Docker Inspect: ====> DETAILED INFORMATION
```
docker inspect <image_name>
```

## BIND MOUNT:
### To map a host directory to container directory
```
docker container run -it --name=test-app -v <host_directory>:<container_directory> <image_name> sh

1. This "Bind Mount" is used to save your container files even after removing the container.
2. used to map host directories to a container directories.
```
