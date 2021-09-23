FROM debian

RUN apt update && apt-get install -y stress

ENTRYPOINT [ "stress" ] 

CMD [ "--vm", "1"]
