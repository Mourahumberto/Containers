# Doc Helm

Explicação completa sobre helm e exemplos para instalação
- Doc Oficial: https://helm.sh/docs/

> Author: **[Humberto Moura](https://github.com/Mourahumberto)**

## Summary

1. [Organization](manuscript/organization.md)
1. [Acknowledgements](manuscript/acknowledgements.md)
1. [Introduction](manuscript/introduction.md)
1. [Why using Docker?](manuscript/why.md)
1. [What is Docker](manuscript/whatis.md)
1. [Set up](manuscript/setup.md)
1. [Commands](manuscript/commands.md)
1. [Creating images](manuscript/creatingimages.md)
1. [Dockerhub](manuscript/dockerhub.md)
1. [Storage](manuscript/storage.md)
1. [Understanding the Docker Network](manuscript/network.md)
1. [Using docker on multiple environments](manuscript/machine.md)
1. [Managing multiple docker containers with Docker compose](manuscript/compose.md)
1. [Using Docker on Windows or OSX](manuscript/macos_and_windows.md)
1. [Dockerizing my Application](manuscript/dockerizing_app.md)
1. [Codebase](manuscript/codebase.md)
1. [Dependencies](manuscript/dependencies.md)
1. [Config](manuscript/config.md)
1. [Serviços de Apoio](manuscript/backingservices.md)
1. [Build, Release, Run](manuscript/build-release-run.md)
1. [Processes](manuscript/processes.md)
1. [Port Binding](manuscript/portbinding.md)
1. [Concurrency](manuscript/concurrency.md)
1. [Disposability](manuscript/disposability.md)
1. [Parity between Dev & Prod](manuscript/parity.md)
1. [Logs](manuscript/logs.md)
1. [Admin](manuscript/admin.md)
1. [Tips](manuscript/tips.md)
1. [Appendix](manuscript/appendix.md)
1. [Container or VM?](manuscript/container_vm.md)
1. [Useful Commands](manuscript/useful_commands.md)
1. [Can I run GUI on docker?](manuscript/gui_applications.md)
1. [Do you lint your Dockerfile? You should...](manuscript/linting_dockerfile.md)


## Instalação do Helm

- From Apt (Debian/Ubuntu)
```
curl https://baltocdn.com/helm/signing.asc | sudo apt-key add -
sudo apt-get install apt-transport-https --yes
echo "deb https://baltocdn.com/helm/stable/debian/ all main" | sudo tee /etc/apt/sources.list.d/helm-stable-debian.list
sudo apt-get update
sudo apt-get install helm
```

## Adicione o repositório desejado.
- veja os repositórios no [Atifact Hub](https://artifacthub.io/packages/search?kind=0&sort=relevance&page=1)

```
$ helm repo add bitnami https://charts.bitnami.com/bitnami
$ helm search repo bitnami
NAME                             	CHART VERSION	APP VERSION  	DESCRIPTION
bitnami/bitnami-common           	0.0.9        	0.0.9        	DEPRECATED Chart with custom templates used in ...
bitnami/airflow                  	8.0.2        	2.0.0        	Apache Airflow is a platform to programmaticall...
bitnami/apache                   	8.2.3        	2.4.46       	Chart for Apache HTTP Server
bitnami/aspnet-core              	1.2.3        	3.1.9        	ASP.NET Core is an open-source framework create...
# ... and many more
```

## instale o projeto desejado

```
$ helm install <xpto> bitnami/wordpress
```

## Customizando um chart antes de instalar.

```
$ helm show values bitnami/wordpress > values.yaml
```
- redireciona o valor padrão do repo, faz as alterações desejadas e instala.
- é importante instalar com a mesma versão.

```
helm install -f values.yaml  <stak-xpto> bitnami/wordpress --version 1.0.1
```

## Customizando com linha de comando

```
helm install --set mariadb.auth.username: user1  <stak-xpto> bitnami/mariadb --version 1.0.1
```
- Esse set no yaml équivale a:

```yaml
mariadb:
  auth:
    username: user1
```

## Verificando os valores depois de instalado
- neste comando já mostra todos os valores que estão inclusive os que foram passados com --set
```
helm get values <chartxpto>
```

## helm upgrade e help rollback
- você pode passar vários values e o arquivo não precisa ter o nome de values.yaml, pode ser um nome qualquer.yaml

```
upgrade
$ helm upgrade -f values.yaml -f xpto.yaml happy-panda bitnami/wordpress

rollback
$ helm rollback [RELEASE] [REVISION]
$ helm rollback xpto 2
```

## helm uninstall

```
helm uninstall xpto
```

- verificar quais helms estão instalados ainda
```
$ helm list
```

## Criando nosso próprio chart

```
$ helm create deis-workflow
```

## Baixando um chart do hub

```
helm pull chartrepo/chartname
```

# Criando uma estrutura chart

- chart é organizado com uma coleção de arquivos dentro de diretórios, o nome do diretório é o nome do chart

```
wordpress/
  Chart.yaml          # A YAML file containing information about the chart
  LICENSE             # OPTIONAL: A plain text file containing the license for the chart
  README.md           # OPTIONAL: A human-readable README file
  values.yaml         # The default configuration values for this chart
  charts/             # A directory containing any charts upon which this chart depends.
  crds/               # Custom Resource Definitions
  templates/          # A directory of templates that, when combined with values,
                      # will generate valid Kubernetes manifest files.
```

## 1. chart.yaml
- arquivo requerido

```yaml
apiVersion: v2
apiVersion: 0.1.0 # do chart
name: The name of the chart (required)
version: 2
kubeVersion: A SemVer range of compatible Kubernetes versions (optional)
description: A single-sentence description of this project (optional)
type: The type of the chart (optional)
keywords:
  - A list of keywords about this project (optional)
home: The URL of this projects home page (optional)
sources:
  - A list of URLs to source code for this project (optional)
dependencies: # A list of the chart requirements (optional)
  - name: The name of the chart (nginx)
    version: The version of the chart ("1.2.3")
    repository: (optional) The repository URL ("https://example.com/charts") or alias ("@repo-name")
    condition: (optional) A yaml path that resolves to a boolean, used for enabling/disabling charts (e.g. subchart1.enabled )
    tags: # (optional)
      - Tags can be used to group charts for enabling/disabling together
    import-values: # (optional)
      - ImportValues holds the mapping of source values to parent key to be imported. Each item can be a string or pair of child/parent sublist items.
    alias: (optional) Alias to be used for the chart. Useful when you have to add the same chart multiple times
    - name: The name of the chart (nginx) # se tiver mais de uma dependência
      version: The version of the chart ("1.2.3")
      repository: (optional) The repository URL ("https://example.com/charts") or alias ("@repo-name")
      condition: (optional) A yaml path that resolves to a boolean, used for enabling/disabling charts (e.g. subchart1.enabled )
      tags: # (optional)
        - Tags can be used to group charts for enabling/disabling together
      import-values: # (optional)
        - ImportValues holds the mapping of source values to parent key to be imported. Each item can be a string or pair of child/parent sublist items.
      alias: (optional) Alias to be used for the chart. Useful when you have to add the same chart multiple times   
maintainers: # (optional)
  - name: The maintainers name (required for each maintainer)
    email: The maintainers email (optional for each maintainer)
    url: A URL for the maintainer (optional for each maintainer)
icon: A URL to an SVG or PNG image to be used as an icon (optional).
appVersion: The version of the app that this contains (optional). Needn't be SemVer. Quotes recommended.
deprecated: Whether this chart is deprecated (optional, boolean)
annotations:
  example: A list of annotations keyed by name (optional).
```

## 2. Chart LICENSE, README and NOTES

o chart pode conter um pequeno texto explicativo dentro de templates/NOTES.txt. e ao final da instalção ele irá mostrar esse texto.
ou se o usuário der um helm status [releae]

## 3. Dependêncas (caso tenha dependências)

```yaml
dependencies:
- condition: kubeStateMetrics.enabled
  name: kube-state-metrics
  repository: https://prometheus-community.github.io/helm-charts
  version: 3.4.*
- condition: nodeExporter.enabled
  name: prometheus-node-exporter
  repository: https://prometheus-community.github.io/helm-charts
  version: 2.0.*
- condition: grafana.enabled
  name: grafana
  repository: https://grafana.github.io/helm-charts
  version: 6.15.*
```
- o repository é o repositório de chart da dependência
- uma vez definida as dependências você baixa elas, onde fica o arquivo chart.yaml você usa o seguinte comando:
```
helm dependency update
```

- pode ser criada tags, para que só suba as dependências caso as tags sejam verdadeiras no values.

```yaml
# parentchart/Chart.yaml

dependencies:
  - name: subchart1
    repository: http://localhost:10191
    version: 0.1.0
    condition: subchart1.enabled, global.subchart1.enabled
    tags:
      - front-end
      - subchart1
  - name: subchart2
    repository: http://localhost:10191
    version: 0.1.0
    condition: subchart2.enabled,global.subchart2.enabled
    tags:
      - back-end
      - subchart2
```
```yaml
# parentchart/values.yaml

subchart1:
  enabled: true
tags:
  front-end: false
  back-end: true
```

ou passando por linha de comando
```
helm install --set tags.front-end=true --set subchart2.enabled=false
```

## 4. Templates and Values

- Todos os recursos que serão criados para essa feature está em templates, hpa, ingress, deployments e etc.
- os valores podem ser passados de dois jeitos: através de values.yal ou com --set

```yaml
# Template file
apiVersion: v1
kind: ReplicationController
metadata:
  name: deis-database
  namespace: deis
  labels:
    app.kubernetes.io/managed-by: deis
spec:
  replicas: 1
  selector:
    app.kubernetes.io/name: deis-database
  template:
    metadata:
      labels:
        app.kubernetes.io/name: deis-database
    spec:
      serviceAccount: deis-database
      containers:
        - name: deis-database
          image: {{ .Values.imageRegistry }}/postgres:{{ .Values.dockerTag }}
          imagePullPolicy: {{ .Values.pullPolicy }}
          ports:
            - containerPort: 5432
          env:
            - name: DATABASE_STORAGE
              value: {{ default "minio" .Values.storage }}
```

```yaml
imageRegistry: "quay.io/deis"
dockerTag: "latest"
pullPolicy: "Always"
storage: "s3"
```
- ou

```
helm install --set imageRegistry="quay.io/deis" 
```