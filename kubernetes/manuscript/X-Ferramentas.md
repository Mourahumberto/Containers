## Sumário

<!-- TOC -->

- [Popeye](#popeye)
  - [Install](#instalação)
- [Velero](#velero)


<!-- TOC -->

# Popeye
- https://github.com/derailed/popeye
- Popeye é uma aplicação readonly que não altera seu cluster kubernetes. Ela tem a principal finalidade de escanear seu cluster e alertar de possíveis erros futuros. Ele tem a função de ajudar o administrador a sanitizar seu cluster. Para que você siga as melhores práticas com seu cluster kubernetes.

## Instalação
-  Baixe o pacote

```
wget https://github.com/derailed/popeye/releases/download/v0.9.8/popeye_Linux_x86_64.tar.gz
```

- Extract
```
tar -xzf popeye_Linux_x86_64.tar.gz
```

- Movendo pacote
```
sudo mv popeye /usr/local/bin/
```
- Dando execução 
```
chmod +x /usr/local/bin/popeye
```

## Execução
- fazendo scan e salvando em html a saída
```
POPEYE_REPORT_DIR=$(pwd) popeye --save --out html --output-file report.html
```

# Velero