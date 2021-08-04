# Criando uma imagem NGNIX

<!-- Altere a Flag abaixo com sua URL do Travis -->


## criando uma imagem com pertiência de dados.


Para maiores informações acesse o [Site do docker hub](https://hub.docker.com/_/nginx) 

### executando o container
criando a imagem
docker build . --tag nomeimagem

criando o container

docker run --name nomecontainer -p 8081:80 nomeimagem

### explicações sobre ngnix

pasta com o site
/var/www/html
