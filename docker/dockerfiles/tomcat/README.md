# Criando uma imagem NGNIX

<!-- Altere a Flag abaixo com sua URL do Travis -->


## criando uma imagem com pertiência de dados.


Para maiores informações acesse o [Site do docker hub](https://hub.docker.com/_/tomcat) 

### executando o container
criando a imagem
docker build . --tag nomeimagem

criando o container

docker run --name nomecontainer -p 8082:8080 nomeimagem

### explicações sobre tomcat

onde fica o site no tomcat
/usr/local/tomcat/webapps/
