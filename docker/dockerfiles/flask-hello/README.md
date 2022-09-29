# Simple Python Flask Dockerized Application#

Build the image using the following command

```bash
$ docker build -t netomoura10/helloflask:latest .
```

Run the Docker container using the command shown below.

```bash
$ docker run -d -p 5000:5000 netomoura10/helloflask:latest
```

The application will be accessible at http:127.0.0.1:5000 or if you are using boot2docker then first find ip address using `$ boot2docker ip` and the use the ip `http://<host_ip>:5000`

## Docker Hub Image

- Uma aplicação com um sleep de 3 segundos antes de retornar o response
$ docker push netomoura10/helloflask:sleep3

- Uma aplicação com o retorno igual a V1
docker push netomoura10/helloflask:v1

- Uma aplicação com o retorno igual a V2
docker push netomoura10/helloflask:v2

## Rodando a imagem do dockerhub


```bash
$ docker run -d -p 5000:5000 netomoura10/helloflask:1.0
```