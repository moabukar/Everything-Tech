### Question - Dockerfile best practice

Given a Dockerfile, analyse it and update it based on security best practices.

### Solution

#### Docker and container security docs (can't be used in exam)

- [Docker best practices by Google](https://cloud.google.com/blog/products/containers-kubernetes/7-best-practices-for-building-containers)
- [Docker tips & tricks for smaller image size by learnk8s.io](https://learnk8s.io/blog/smaller-docker-images)

#### 1 - Open Dockerfile and fix security issues

```sh

vi ~/Dockerfile

FROM ubuntu:latest

ENV CI=true

RUN apt get update
RUN apt get install -y wget
RUN apt get install -y curl

USER root

WORKDIR /code
COPY package.json package-lock.json /code/
RUN npm ci
COPY src /code/src

CMD [ "npm", "start" ]

```

#### 2 - Update Dockerfile with best practices

```sh

vi ~/Dockerfile

FROM ubuntu:20:04 ## updated image to a specific version

ENV CI=true

RUN apt get update
RUN apt get install -y wget curl ## lighter image due to docker caching

USER user ## no privileged user being used

WORKDIR /code
COPY package.json package-lock.json /code/
RUN npm ci
COPY src /code/src

CMD [ "npm", "start" ]

```
