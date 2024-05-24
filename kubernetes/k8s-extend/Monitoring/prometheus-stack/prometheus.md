# PROMETHEUS

### O que é o prometheus
- É um sistema de alerta e monitoração open source. Ele salva os dados com seus timestamp em forma de ke-value chamada label.
- Ele usa a linguagem PromQL, ele suporta pushing via um gateway intermediário.
- Targets são descobertos via service discovery o configuração estática.

## Componentes Principais
- A arquitetura do Prometheus é composta por vários componentes inter-relacionados que trabalham em conjunto para coletar, armazenar, consultar e alertar com base em métricas. Os principais componentes incluem:

- **Servidor Prometheus:** Este é o núcleo do sistema, responsável por coletar e armazenar métricas de diferentes fontes. Ele oferece uma interface HTTP para consulta de métricas e uma linguagem de consulta poderosa chamada PromQL.
- **Targets:** São os sistemas, serviços ou endpoints que o Prometheus monitora. Eles expõem suas métricas em um formato específico (geralmente no formato de endpoint HTTP) para que o Prometheus possa coletá-las.
- **Exportadores: São processos auxiliares que ajudam o Prometheus a coletar métricas de sistemas que não fornecem nativamente métricas no formato adequado. Existem exportadores para diversas tecnologias, como bancos de dados, servidores web, sistemas de armazenamento, entre outros.
- **Banco de Dados de Séries Temporais (TSDB):** Todas as métricas coletadas pelo Prometheus são armazenadas em um banco de dados de séries temporais. Isso permite a consulta eficiente e rápida das métricas ao longo do tempo.
- **API HTTP:** O Prometheus disponibiliza uma API HTTP para consulta das métricas armazenadas. Essa API é usada por ferramentas de visualização, como Grafana, para criar dashboards e gráficos com as métricas coletadas.
- **Alertmanager:** É responsável por gerenciar alertas com base em regras configuradas pelo usuário. Ele recebe alertas do Prometheus e os encaminha para canais de notificação, como e-mail, Slack, PagerDuty, entre outros.

- imagem arquitetura prometheus.
- o Prometheus coleta as métricas de aplicações ou de push gateways que servem de intermediário pra jobs de vida curta.

### Anotação
- Procurando uma métrica
```
<metric name>{<label name>=<label value>, ...}
```
- vamos supor que você tem uma métrica ```api_http_requests_total``` e as seguintes labels ```method="POST"``` e ```handler="/messages"```:
```
api_http_requests_total{method="POST", handler="/messages"}
```

## Entendo a métrica
- Suponha que temos um serviço web que fornece informações sobre o estado de uma aplicação. Vamos considerar uma métrica que conta o número de solicitações HTTP recebidas por esse serviço em um determinado período de tempo.

```
# HELP http_requests_total The total number of HTTP requests received.
# TYPE http_requests_total counter
http_requests_total{method="GET", endpoint="/api/v1/status", status="200"} 100
http_requests_total{method="POST", endpoint="/api/v1/data", status="500"} 20
```

### Estrutura
- **Nome da Métrica:** http_requests_total - No caso, estamos monitorando o número total de solicitações HTTP recebidas.
- **HELP:** # HELP http_requests_total The total number of HTTP requests received. - Esta linha fornece uma breve descrição da métrica, explicando o que ela representa. Ajuda os usuários a entenderem o significado da métrica.
- **TYPE:** # TYPE http_requests_total counter - Este campo especifica o tipo da métrica. Neste caso, é um contador (counter), o que significa que é uma métrica que aumenta ao longo do tempo e nunca diminui.
- **Labels:** As labels fornecem metadados adicionais sobre a métrica. Por exemplo:
    - method="GET": Indica o método HTTP usado na solicitação.
    - endpoint="/api/v1/status": Indica o endpoint da solicitação.
    - status="200": Indica o código de status HTTP retornado.
- **Valor da Métrica:** No exemplo foram recebidas 100 solicitações HTTP GET para o endpoint /api/v1/status com sucesso (código de status 200) e 20 solicitações POST para o endpoint /api/v1/data com erro (código de status 500).

### Explicação
- A métrica http_requests_total é do tipo contador, o que significa que aumenta conforme mais solicitações são recebidas.
- Cada linha representa uma instância da métrica com diferentes valores de label. Por exemplo, podemos ter contadores separados para diferentes métodos HTTP, endpoints e códigos de status.
- As labels permitem uma segmentação eficaz dos dados, facilitando a análise e a identificação de padrões específicos.

### Funcionamento
- Inicialmente, os targets são identificados e configurados para expor suas métricas. O Prometheus, então, periodicamente consulta esses targets, coletando métricas e armazenando-as no banco de dados de séries temporais. Os dados armazenados podem ser consultados por meio da API HTTP ou utilizados para gerar alertas.
- Abaixo um exemplo de arquivo de configuração (prometheus.yml) onde é possível configurar o intervalo de scrape, os targets e demais configurações de discovery etc.

```yaml
global:
  scrape_interval: 15s

scrape_configs:
  - job_name: 'node_exporter'
    static_configs:
      - targets: ['localhost:9100']

  - job_name: 'app_service'
    static_configs:
      - targets: ['app1.example.com:8080', 'app2.example.com:8080']
```
- **global:** Esta seção define configurações globais que se aplicam a todos os jobs de coleta definidos no arquivo de configuração.
- **scrape_interval:** Define o intervalo de tempo entre cada scrape (raspagem) dos targets. No exemplo, está definido como 15s, o que significa que o Prometheus irá coletar métricas de cada target a cada 15 segundos, o que pode ser considerado muito agressivo para aplicações com carga elevada.
- **scrape_configs:** Esta seção define os diferentes jobs de coleta e os targets associados a cada job.
- **job_name:** Um nome descritivo para o job de coleta. Isso ajuda a identificar e organizar os diferentes conjuntos de targets e configurações.
- **static_configs:** Define os targets estáticos que o Prometheus irá monitorar para este job.
- **targets:** Uma lista de endereços e portas dos targets a serem monitorados. No exemplo, temos dois jobs de coleta: node_exporter, que coleta métricas do node_exporter (um exportador do Prometheus para máquinas Unix/Linux), e app_service, que monitora dois aplicativos (app1.example.com:8080 e app2.example.com:8080).

Cada job de coleta pode ter configurações adicionais, como relabel_configs para ajustar/adicionar labels de métricas, metrics_path para especificar um caminho específico para as métricas, e outras opções específicas do job.         
Este é apenas um exemplo básico de um arquivo de configuração do Prometheus. Na prática, os arquivos de configuração podem ser mais complexos, incluindo configurações avançadas, discovery de serviços, relatórios de alertas e muito mais, dependendo dos requisitos específicos de monitoramento de cada ambiente.  

- agregando targets com o mesmo job, exemplo deploy canário
```yaml
scrape_configs:
  - job_name:       'node'

    # Override the global default and scrape targets from this job every 5 seconds.
    scrape_interval: 5s

    static_configs:
      - targets: ['localhost:8080', 'localhost:8081']
        labels:
          group: 'production'

      - targets: ['localhost:8082']
        labels:
          group: 'canary'
```