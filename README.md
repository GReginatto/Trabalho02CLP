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

# Comparando Rust, Golang e C - Eliminação de Gauss

## Comparação de Código entre as Linguagens

Abaixo está uma comparação detalhada das linguagens Rust, Golang e C com base em aspectos como tipos de dados, acesso às variáveis, organização de memória, chamadas de função e comandos de controle de fluxo.

### 1. Tipos de Dados

| Característica   | C                   | Rust               | Golang            |
|-----------------|---------------------|--------------------|-------------------|
| Inteiros        | `int`, `long`       | `i32`, `i64`, `u32`, `usize` | `int`, `int64`, `uint` |
| Ponto Flutuante | `float`, `double`   | `f32`, `f64`       | `float32`, `float64` |
| Booleanos       | `int` (0 ou 1)      | `bool`             | `bool`             |
| Strings        | `char*` (ponteiro)   | `String`, `&str`   | `string`          |
| Arrays         | `int arr[N]`        | `[i32; N]`         | `[N]Type` ou `slice` |
| Structs        | `struct`            | `struct`           | `struct`          |
| Ponteiros      | `int*`              | `&T`, `Box<T>`     | `unsafe.Pointer`  |

### 2. Acesso às Variáveis

- **C:** Utiliza variáveis globais e locais, sendo possível modificar diretamente valores através de ponteiros.
- **Rust:** Utiliza `let` para imutáveis e `let mut` para mutáveis. Segurança garantida pelo sistema de propriedade (ownership).
- **Golang:** Usa `var` para declaração explícita e `:=` para inferência. Suporte a concorrência via goroutines e channels.

### 3. Organização de Memória

| Aspecto         | C                     | Rust                 | Golang               |
|----------------|----------------------|----------------------|----------------------|
| Gerenciamento  | Manual (`malloc/free`) | Automático (Ownership) | Coletor de lixo (GC) |
| Stack/Heap     | Controle manual       | Determinado pelo compilador | GC decide           |
| Segurança      | Risco de vazamento    | Sem vazamento (borrow checker) | Possível vazamento |

### 4. Chamadas de Função

- **C:** Uso direto de ponteiros e passagem de parâmetros por valor/referência.
- **Rust:** Permite `fn` com referências (`&T`, `&mut T`) garantindo segurança.
- **Golang:** Possui passagem por valor por padrão, com suporte a ponteiros para modificações.

### 5. Controle de Fluxo

| Estrutura       | C               | Rust            | Golang          |
|---------------|----------------|---------------|---------------|
| Condicional  | `if`, `switch`   | `if`, `match` | `if`, `switch` |
| Laços        | `for`, `while`   | `for`, `loop`, `while` | `for` |
| Manipulação  | `goto`, `break`  | `break`, `continue` | `break`, `continue` |

Rust se destaca pelo uso de `match`, enquanto Go e C utilizam `switch`.
