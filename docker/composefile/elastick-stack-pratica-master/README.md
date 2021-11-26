# Stack do elastic com aplicação hello world em Flask

- Rodando o projeto de forma simples

```bash
$ docker-compose up
```

- Verificando a aplicação 

```bash
$ curl localhost:5000
```

- Verificando se o Elasticsearch está ok

```bash
$ curl -X GET "localhost:9200/_cat/nodes?v&pretty"
```

- Verificando a interface do kibana
    - No seu navegador acesse por: http://localhost:5601/

- Acessando a Parte do apm
    - http://localhost:5601/app/apm/services?rangeFrom=now-15m&rangeTo=now&comparisonEnabled=true&comparisonType=day

- Apagando tudo

```bash
$ docker-compose down -v
```

- **DICA:** ao rodar com três nós do elasticsearch deu erro rodei o seguinte comando
```bash
sudo sysctl -w vm.max_map_count=262144
```

## Docs Oficiais
- docker-compose example: https://www.elastic.co/guide/en/apm/get-started/current/quick-start-overview.html
