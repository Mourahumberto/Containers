Para maiores informações acesse o [Site do docker hub](https://hub.docker.com/_/openjdk)


# Criando uma imagem openjdk v0.1
######## IMAGEM #########
FROM openjdk:7
COPY . /usr/src/myapp
WORKDIR /usr/src/myapp
RUN javac Main.java
CMD ["java", "Main"]

########################

1) Criando a imagem do dockerfile
docker build --tag jdk .

2) Criando o container com a imagem criada
docker run --name containerjdk jdk

2.1) Criando o container com a imagem criada e pertiência de dados.

docker run -it --name containerjdk -v testejdk:/usr/src/myapp jdk bash
-o bash nesse momento é para debug
<!-- Altere a Flag abaixo com sua URL do Travis -->

# Criando uma imagem openjdk v0.2

## criando uma imagem com pertiência de dados.


 

### executando o container
criando a imagem
docker build . --tag nomeimagem

criando o container

docker run -p 8080:8080 -p 50000:50000 nomeimagem

### explicações sobre jenkins


