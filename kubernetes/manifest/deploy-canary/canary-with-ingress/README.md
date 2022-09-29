# Deploy Canario
- Desta forma você consegue fazer várias estratégias de canário, existe três estratégias nesse manifesto
- No caso você cria um deployment stable e um service apontando para o deploy stable com o arquivo deploymentstable.yaml, depois cria um deployment e um service apontando
para esse deployment com o deploymentcanary.yaml. e no arquivo ingress.yaml você tem o ingress normal e as estratégias de canary, por porcentagem, header e cokies
caso queira usar uma delas comente uma e descomente a outra. 

## formas de fazer a chamada.

```
#usando apenas o redirect
curl -H "Host: meu.dns.com.br" http://<ingress>:<port>/
#usando a opção com header
curl -H "Host: meu.dns.com.br" -H "x-region: us-east" http://<ingress>:<port>/
```