# Doc Helm

Explicação completa sobre helm e exemplos para instalação
- Doc Oficial: https://helm.sh/docs/

> Author: **[Humberto Moura](https://github.com/Mourahumberto)**

## Summary

1. [install](#01)
1. [Adicione um repositório](#02)
1. [instale um projeto](#03)
1. [Customizando um Chart](#04)
1. [Customizando com linha de comando](#05)
1. [Veridicando Values](#06)
1. [Helm upgrade e help rollback](#07)
1. [Helm uninstall](#08)
1. [Criando o próprio chart](#09)
1. [Baixando um chart do hub](#10)
1. [Criando uma estrutura chart](#11)



## Instalação do Helm <div id='01'/>

- From Apt (Debian/Ubuntu)
```
curl https://baltocdn.com/helm/signing.asc | sudo apt-key add -
sudo apt-get install apt-transport-https --yes
echo "deb https://baltocdn.com/helm/stable/debian/ all main" | sudo tee /etc/apt/sources.list.d/helm-stable-debian.list
sudo apt-get update
sudo apt-get install helm
```

## Adicione um repositório. <div id='02'/>
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

## instale um projeto <div id='03'/>

```
$ helm install <xpto> bitnami/wordpress
```

## Customizando um chart antes de instalar. <div id='04'/>

```
$ helm show values bitnami/wordpress > values.yaml
```
- redireciona o valor padrão do repo, faz as alterações desejadas e instala.
- é importante instalar com a mesma versão.

```
helm install -f values.yaml  <stak-xpto> bitnami/wordpress --version 1.0.1
```

## Customizando com linha de comando <div id='05'/>


```
helm install --set mariadb.auth.username: user1  <stak-xpto> bitnami/mariadb --version 1.0.1
```
- Esse set no yaml équivale a:

```yaml
mariadb:
  auth:
    username: user1
```

## Verificando os valores depois de instalado <div id='06'/>
- neste comando já mostra todos os valores que estão inclusive os que foram passados com --set
```
helm get values <chartxpto>
```

## Helm upgrade e help rollback <div id='07'/>
- você pode passar vários values e o arquivo não precisa ter o nome de values.yaml, pode ser um nome qualquer.yaml

```
upgrade
$ helm upgrade -f values.yaml -f xpto.yaml happy-panda bitnami/wordpress

rollback
$ helm rollback [RELEASE] [REVISION]
$ helm rollback xpto 2
```

## Helm uninstall <div id='08'/>

```
helm uninstall xpto
```

- verificar quais helms estão instalados ainda
```
$ helm list
```

## Criando o próprio chart <div id='09'/>

```
$ helm create deis-workflow
```

## Baixando um chart do hub <div id='10'/>

```
helm pull chartrepo/chartname
```

# Criando uma estrutura chart <div id='11'/>

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
#values.yaml
imageRegistry: "quay.io/deis"
dockerTag: "latest"
pullPolicy: "Always"
storage: "s3"
```
- ou

```
helm install --set imageRegistry="quay.io/deis" 
```
