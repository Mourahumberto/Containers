## Queries simples

- promql tem wildcards porém com a anoptação deles. as queries são assim.

```
time_epoch
it responds with this result and labels:
time_epoch{instance="myhost245",job="dev"}
time_epoch{instance="myhost119",job="qa_aws"}
time_epoch{instance="myhost119",job="production_gc"}
time_epoch{instance="myhost119",job="production_aws"}
time_epoch{instance="myhost119",job="production_azure"}
```
- se você quer quer procurar só pelos jobs igual aws não tem a opção '*' então você usa o REGEX “.+”.
```
time_epoch{job=~".+aws"}
```

- se tiver no meio você pode usar
```
time_epoch{job=~".+aws.+"}
```