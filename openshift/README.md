# Openshift
Orquestrador de containers baseado no Kubernetes, o openshift é um orquestrador de containers da Redhat


### Instalação do Minishift
- Minishift é uma plataforma local do openshift, serve para desenvolver de forma local. um concorrente do minikube.
https://computingforgeeks.com/how-to-run-local-openshift-cluster-with-minishift/

Username: system
Password: admin

#### Comandos Minishift
$ minishift status|stop|start

$ minishift console --url

### Instalação o CodeReady Containers
- versão mais próxima do openshift 4.x
https://console.redhat.com/openshift/create/local
1) faz o download
1.1) copia o pull secret do download, vai precisar no crc setup
2) extrai tudo da pasta tar -xvJf crc-linux-amd64.tar.xz
3) verifica alguma pasta que esteja no $PATH e mv o binário para essa pasta
4) instala dependências sudo apt install qemu-kvm libvirt-daemon libvirt-daemon-system network-manager
4.1) crc setup
4.2) crc start

vendo se está ok
crc --help

crc console --url
crc console --credentials