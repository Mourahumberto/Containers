## Usando Loki Padrão parser
- Você pode dinamicamente criar novas labels com seu log estruturado usando um pattern parser.
- você pode dividir o log do nginx em várias outras labels.
exemplo de linha
```
213.205.198.138 - - [29/Nov/2021:08:23:54 +0000] "GET /public/build/grafanaPlugin.9293a56f182a84c40c07.js HTTP/1.1" 200 11042 "https://grafana.sbcode.net/?orgId=1" "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/1.2 (KHTML, like Gecko) Chrome/1.2.3.4 Safari/537.36 Edg/1.2.3.4"
```
exemplo de query.
```
{job="nginx"} | pattern `<_> - - <_> "<method> <_> <_>" <status> <_> <_> "<_>" <_>`
```

- Doc usada: https://sbcode.net/grafana/nginx-promtail/
- video: https://www.youtube.com/watch?v=kR5ay4lX0OM