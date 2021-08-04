#imagem usada do docker hub
FROM debian
#informações de quem criou
MAINTAINER Humberto Moura <humbertomoura.n@gmail.com>
#ao subir o container rodar os seguintes comandos, -y é para dar yes
RUN apt-get update && apt-get install -y nginx
#adicionar no container o seguinte o seguinte arquivo no seguinte diretório
#ADD index.html /usr/share/nginx/html
ADD index.html /var/www/html
#usar a porta 80
EXPOSE 80
#quando tu estiver pronto rodar os seguintes comandos
CMD ["nginx", "-g", "daemon off;"]
