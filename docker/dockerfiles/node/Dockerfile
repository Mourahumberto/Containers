#speciffy a base image
FROM node:alpine

#evitar o problema de cache do docker
COPY ./packege.json ./
#neste caso ele coloca o container pra iniciar desde diretório e trabalhar neste diretório
WORKDIR /usr/app

COPY ./ ./
#Install some depedencies
RUN npm install

#Defaault comand

CMD ["npm", "start"]
