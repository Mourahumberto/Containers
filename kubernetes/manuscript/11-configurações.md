# Configurações

## ConfigMaps
- Um ConfigMap é um objeto da API usado para armazenar dados não-confidenciais em pares chave-valor. Pods podem consumir ConfigMaps como variáveis de ambiente, argumentos de linha de comando ou como arquivos de configuração em um volume.
- Um ConfigMap ajuda a desacoplar configurações vinculadas ao ambiente das imagens de contêiner, de modo a tornar aplicações mais facilmente portáveis.
- O ConfigMap não oferece confidencialidade ou encriptação.

Existem algumas formas de usar os configmaps.
- Como env
```yaml
      env:
        # Define as variáveis de ambiente
        - name: PLAYER_INITIAL_LIVES # Note que aqui a variável está definida em caixa alta,
                                     # diferente da chave no ConfigMap.
          valueFrom:
            configMapKeyRef:
              name: game-demo           # O ConfigMap de onde esse valor vem.
              key: player_initial_lives # A chave que deve ser buscada.
```
- Como arquivo
```yaml
      volumeMounts:
      - name: config
        mountPath: "/config"
        readOnly: true
  volumes:
    # Volumes são definidos no escopo do Pod, e os pontos de montagem são definidos
    # nos contêineres dentro dos pods.
    - name: config
      configMap:
        # Informe o nome do ConfigMap que deseja montar.
        name: game-demo
        # Uma lista de chaves do ConfigMap para serem criadas como arquivos.
        items:
        - key: "game.properties"
          path: "game.properties"
        - key: "user-interface.properties"
          path: "user-interface.properties"
```
[Exemplo de Manifest](../manifest/configmap/pod-configmap.yaml)

## Secrets
- Um Secret é um objeto que contém uma pequena quantidade de informação sensível, como senhas, tokens ou chaves. Este tipo de informação poderia, em outras circunstâncias, ser colocada diretamente em uma configuração de Pod ou em uma imagem de contêiner. O uso de Secrets evita que você tenha de incluir dados confidenciais no seu código.
- Secrets são semelhantes a ConfigMaps, mas foram especificamente projetados para conter dados confidenciais.
- Secrets podem ser montados como volumes de dados ou expostos como variáveis de ambiente para serem utilizados num container de um Pod. Secrets também podem ser utilizados por outras partes do sistema, sem serem diretamente expostos ao Pod.
- O secret pod ser lido caso consiga entrar no pod
- Caso você Atualize uma secret ela pode demorar um pouco pra atualizar, quem determina esse TTL é o kubelet ele tem um tempo de cache configurado.
    
### Para utilizar Secrets de forma segura, siga pelo menos as instruções abaixo:
  - Habilite encriptação em disco para Secrets.
  - Habilite ou configure regras de RBAC que restrinjam o acesso de leitura a Secrets (incluindo acesso indireto).
  - Quando apropriado, utilize mecanismos como RBAC para limitar quais perfis e usuários possuem permissão para criar novos Secrets ou substituir Secrets existentes.

### Alternativas a Secrets
Ao invés de utilizar um Secret para proteger dados confidenciais, você pode escolher uma maneira alternativa. Algumas das opções são:
  - Se for recurso dentro da sua própria nuvem você pode usar serviceaccount dentro do próprio pod. ex: https://www.youtube.com/watch?v=bu0M2y2g1m8
  - se o seu componente cloud native precisa autenticar-se a outra aplicação que está rodando no mesmo cluster Kubernetes, você pode utilizar uma ServiceAccount e seus tokens para identificar seu cliente.
  - existem ferramentas fornecidas por terceiros que você pode rodar, no seu cluster ou externamente, que providenciam gerenciamento de Secrets. Por exemplo, um serviço que Pods accessam via HTTPS, que revelam um Secret se o cliente autenticar-se corretamente (por exemplo, utilizando um token de ServiceAccount).
  - para autenticação, você pode implementar um serviço de assinatura de certificados X.509 personalizado, e utilizar CertificateSigningRequests para permitir ao serviço personalizado emitir certificados a pods que os necessitam.
  - você pode utilizar um plugin de dispositivo para expor a um Pod específico um hardware de encriptação conectado a um nó. Por exemplo, você pode agendar Pods confiáveis em nós que oferecem um Trusted Platform Module, configurado em um fluxo de dados independente.

### Criando Por linha de comando

```
$ echo -n 'admin' > ./username.txt
$ echo -n '1f2d1e2e67df' > ./password.txt
$ kubectl create secret generic db-user-pass \
  --from-file=./username.txt \
  --from-file=./password.txt
```
ou
```
$ kubectl create secret generic db-user-pass \
  --from-literal=username=devuser \
  --from-literal=password='S!B\*d$zDsb='
```

- verificando as secrets
```
$ kubectl get secrets
$ kubectl describe secrets/db-user-pass
```

- verificando os valores da secret
```
$ kubectl get secret db-user-pass -o jsonpath='{.data}'
```
- Saída
```
{"password":"MWYyZDFlMmU2N2Rm","username":"YWRtaW4="}
```
- Decodificando
```
$ echo 'MWYyZDFlMmU2N2Rm' | base64 --decode
```
### Criando por arquivo de configuração
- Codifique o secret
```
echo -n 'admin' | base64
echo -n '1f2d1e2e67df' | base64
```
- Coloque os no .yaml e aplique
```yaml
apiVersion: v1
kind: Secret
metadata:
  name: mysecret
type: Opaque
data:
  username: YWRtaW4=
  password: MWYyZDFlMmU2N2Rm
```

- Ou usando um stringdata no lugar de data
```yaml
apiVersion: v1
kind: Secret
metadata:
  name: mysecret
type: Opaque
stringData:
  config.yaml: |
    apiUrl: "https://my.api.com/api/v1"
    username: teste
    password: testando
```

- Usando o kustomiza
https://kubernetes.io/pt-br/docs/tasks/configmap-secret/managing-secret-using-kustomize/

### Consumindo as secrets

[Exemplo de Manifest](../manifest/secrets/file-secrets.yaml)

- Usando secrets para baixar imagens de um repositório externo
https://kubernetes.io/pt-br/docs/concepts/configuration/secret/#secrets-de-configura%C3%A7%C3%A3o-do-docker

- Usando secrets para chaves ssh
https://kubernetes.io/pt-br/docs/concepts/configuration/secret/#caso-de-uso-pod-com-chaves-ssh

- Usando secrets atraves de token de service account
[Secrets de token de service account](https://kubernetes.io/pt-br/docs/concepts/configuration/secret/#secrets-de-token-de-service-account-conta-de-servi%C3%A7o)

- Usando TLS
https://kubernetes.io/pt-br/docs/concepts/configuration/secret/#secrets-de-configura%C3%A7%C3%A3o-do-docker

### Secrets imutáveis
- protege você de alterações acidentais (ou indesejadas) que poderiam provocar disrupções em aplicações.
- em clusters com uso extensivo de Secrets (pelo menos dezenas de milhares de montagens únicas de Secrets a Pods), utilizar Secrets imutáveis melhora o desempenho do seu cluster através da redução significativa de carga no kube-apiserver. O kubelet não precisa manter um watch em Secrets que são marcados como imutáveis
- Uma vez que um Secret ou ConfigMap seja marcado como imutável, não é mais possível reverter esta mudança, nem alterar os conteúdos do campo data. Você pode somente apagar e recriar o Secret. Pods existentes mantém um ponto de montagem referenciando o Secret removido - é recomendado recriar tais Pods.

### Pontos importantes nos secrets
- Embora ConfigMaps e Secrets funcionem de formas similares, o Kubernetes aplica proteções extras aos objetos Secret.
- Um Secret só é enviado a um nó se um Pod naquele nó precisa do Secret em questão. Para montar Secrets em Pods, o kubelet armazena uma cópia dos dados dentro de um sistema de arquivos tmpfs, de modo que os dados confidenciais não sejam escritos em armazenamento durável.
- Aplicações ainda devem proteger o valor da informação confidencial após lê-la de uma variável de ambiente ou volume. Por exemplo, sua aplicação deve evitar imprimir os dados do Secret sem encriptação ou transmitir esta informação para aplicações terceiras de confiabilidade não-estabelecida.
- Ao instalar aplicações que interagem com a API de Secrets, você deve limitar o acesso utilizando políticas de autorização, como por exemplo RBAC.
- Na API do Kubernetes, requisições watch e list em Secrets dentro de um namespace são extremamente poderosas. Evite fornecer este acesso quando possível, já que listar Secrets permite aos clientes inspecionar os valores de todos os Secrets naquele namespace.
- Um usuário que pode criar um Pod que utiliza um Secret pode também ver o valor daquele Secret. Mesmo que as permissões do cluster não permitam ao usuário ler o Secret diretamente, o mesmo usuário poderia ter acesso a criar um Pod que então expõe o Secret.
- 