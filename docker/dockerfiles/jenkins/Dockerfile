FROM jenkins/jenkins:latest
USER root
RUN apt-get update && apt-get install -y make git openjdk-8-jdk
USER jenkins
#copia o arquivo para a pasta em questão esse arquivo serve para fazer download dos plugins
COPY plugins.txt /usr/share/jenkins/ref/plugins.txt
#salva o que ta escrito no no txt para esse arquivo.sh que é um script do jenkins
RUN /usr/local/bin/install-plugins.sh < /usr/share/jenkins/ref/plugins.txt
#seta a variável de ambiente
ENV JAVA_HOME /usr/lib/jvm/java-8-openjdk-amd64
