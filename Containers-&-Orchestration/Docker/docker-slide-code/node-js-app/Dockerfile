FROM alpine
LABEL maintainer="ali@gmail.com"
RUN apk add --update nodejs npm
RUN apk add --update npm
COPY . /src
WORKDIR /src
ENV CREATEDBY="Ali"
EXPOSE 8080
ENTRYPOINT ["node", "./app.js"]
