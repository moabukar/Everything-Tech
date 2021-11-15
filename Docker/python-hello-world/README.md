# Python Hello World Docker Image Example

This demonstrates how to dockerize a simple flask app

## How to build

- Navigate to the directory containing the Dockerfile
- Run `docker build -t <image_name> .`
This builds the docker image and tags it with the image name you specified

## How to run
Run the command `docker run -p 5000:5000 <image_name>` to run the docker image. Visit the URL [http://localhost:5000](http://localhost:5000) to view the running application.