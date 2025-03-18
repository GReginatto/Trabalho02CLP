## Trabalho 02 - Comparando C, Rust e Golang

# Integrantes do Grupo

- [Gustavo Reginatto](https://github.com/GReginatto)
- [Lucas Ludwig](https://github.com/lucas26042002)
- [Sandro Júnior](https://github.com/SandJunior)

# Objetivo

Este projeto tem como objetivo comparar as linguagens C, Rust e Golang, implementando o algoritmo de Eliminação de Gauss em cada uma delas e analisando diferenças em:
1. Tipos de dados
2. Acesso às variáveis
3. Organização de memória
4. Chamadas de função
5. Comandos de controle de fluxo
6. Métricas de código (número de linhas, número de comandos, etc.)
7. Desempenho das linguagens (tempo de execução para diferentes tamanhos de matriz)

# Metodologia

1 - [Recuperamos uma implementação do algoritmo de Eliminação de Gauss em C.](https://github.com/gmendonca/gaussian-elimination-pthreads-openmp/blob/master/gauss.c)

2 - Reimplementamos o mesmo algoritmo em Rust e Golang.

3 - Comparamos os três códigos em estrutura, organização de memória e métricas de código.

4 - Executamos testes de desempenho para avaliar tempo de execução em diferentes linguagens.

5 - Documentamos as comparações em um relatório contendo gráficos e tabelas de desempenho.

# Como Rodar os Programas

## Instalação das Linguagens

### **Rust**

#### **Windows**
1. Baixe e instale o Rust pelo terminal com:
   ```powershell
   curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
   ```
2. Reinicie o terminal e verifique a instalação:
   ```sh
   rustc --version
   ```

#### **Linux**
1. Instale pelo terminal:
   ```sh
   curl --proto '=https' --tlsv1.2 -sSf https://sh.rustup.rs | sh
   ```
2. Reinicie o terminal e verifique a instalação:
   ```sh
   rustc --version
   ```

### **Golang**

#### **Windows**
1. Baixe o instalador oficial em [https://go.dev/dl/](https://go.dev/dl/)
2. Siga as instruções do instalador.
3. Verifique a instalação:
   ```sh
   go version
   ```

#### **Linux**
1. Instale o Go executando os comandos:
   ```sh
   wget https://go.dev/dl/go1.20.6.linux-amd64.tar.gz
   sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.20.6.linux-amd64.tar.gz
   echo 'export PATH=$PATH:/usr/local/go/bin' >> ~/.bashrc
   source ~/.bashrc
   ```
2. Verifique a instalação:
   ```sh
   go version
   ```

## Como Compilar e Executar os Códigos

### **Rust**

#### **Compilar**:
```sh
rustc main.rs -o gauss_rust
```

#### **Executar**:
```sh
cargo run -- 10
```

### **Golang**

#### **Executar diretamente**:
```sh
go run gauss.go
```

#### **Compilar**:
```sh
go build gauss.go
```

#### **Executar o binário compilado**:
```sh


./gauss
```

# No caso de estar utilizando o VSCode.
1 - Utilize a extensão chamada "CodeRunner", que compila e executa automaticamente para o usuário.


