FROM debian

RUN apt-get update && apt-get install -y apache2 && apt-get install -y curl 
ENV APACHE_LOCK_DIR="/var/lock"
ENV APACHE_PID_FILE="/var/run/apache2.pid"
ENV APACHE_RUN_USER="www-data"
ENV APACHE_RUN_GROUP="www-data"
ENV APACHE_LOG_DIR="/var/log/apache2"

ADD index.html /var/www/html/
HEALTHCHECK --interval=30s --timeout=30s --start-period=5s --retries=3 CMD curl -f http://localhost || exit 1
LABEL description="Webserver"
LABEL version="1.0.0"

VOLUME /var/www/html/
EXPOSE 80

# comando principal
ENTRYPOINT [ "/usr/sbin/apachectl" ] 
# parametros do entrypoint
CMD [ "-D", "FOREGROUND"]