# Projeto stress

- Projeto simples que usa o comando stress, para emular cenários em seu servidor. Um exemplo seria criar um stress em suas máquinas para simular um auto-scaling.

## Como utilizar o container

buildando a imagem
```bash
$ docker build -t stress:latest .
```

agora rode o container com os seguintes argumentos

```bash
$ docker run -d stress:latest --cpu 2 --vm 1 --vm-bytes 1G
```

--cpu 2 utilizará dois cpu's
--vm 1 apenas um processo
--vm-bytes 1G 1G por processo

## Usando a imagem do dockerhub desse projeto

```
$ docker run -ti netomoura10/stress:1.0 --cpu 2 --vm 1 --vm-bytes 1G
```