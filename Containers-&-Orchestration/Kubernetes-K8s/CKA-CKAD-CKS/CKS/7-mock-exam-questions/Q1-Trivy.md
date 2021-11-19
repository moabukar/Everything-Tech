### Question - Trivy

There are a number of pods running in the "spectacle" namespace. Identify and delete the pods which have CRITICAL vulnerabilities.

### Solution

- [Trivy docs](https://github.com/aquasecurity/trivy)

#### 1 - get all images of pods running in the 'spectacle' namepsace

```sh

kubectl -n spectacle get pods -o yaml | grep -E "image"

```

#### 2 - scan each image using "trivy image <NAME>"

```sh

trivy image --severity CRITICAL nginx:1.16

```

Run the above command for all the images found in step 1. If the images have CRITICAL vulnerabilities, delete the pod associated with that image.
  
#### 3 - Delete the pods

```sh

kubectl -n spectacle delete pod <PODNAME>

```

Do the above procedure as many times as required to delete all pods with CRITICAL vulnerabilities.  
