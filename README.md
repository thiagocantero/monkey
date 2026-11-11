# Monkey Language

Este repositório é um projeto de estudo da linguagem Monkey, implementado em Go. O código atual cobre a definição dos tokens, a análise léxica de entrada e um REPL simples que exibe os tokens reconhecidos.

> **Escopo atual:** o projeto ainda não implementa parser, AST, avaliação de expressões ou execução de programas. Por isso, comandos digitados no REPL são tokenizados e listados; eles não são avaliados.

## Estrutura do projeto

```text
.
├── main.go
├── go.mod
├── token/
│   └── token.go
├── lexer/
│   ├── lexer.go
│   └── lexer_test.go
└── repl/
    └── repl.go
```

### Diretórios e arquivos

- **`token/`** — define os tipos de token, os literais e as palavras reservadas da linguagem. `LookupIdent` diferencia palavras-chave, como `let` e `fn`, de identificadores comuns.
- **`lexer/`** — transforma o texto de entrada em uma sequência de tokens. `Lexer` mantém a posição de leitura, reconhece operadores, delimitadores, identificadores, números e espaços em branco.
- **`lexer/lexer_test.go`** — verifica os tokens produzidos para um trecho representativo da linguagem, incluindo declarações, funções, condicionais e operadores.
- **`repl/`** — implementa o loop interativo: lê uma linha, cria um lexer e imprime cada token até encontrar o fim da entrada.
- **`main.go`** — ponto de entrada do programa. Exibe uma saudação e inicia o REPL.
- **`go.mod`** — identifica o módulo Go (`monkeylang`) e a versão da linguagem configurada.

## Fluxo de leitura do código

Para acompanhar como uma linha chega à saída, siga este percurso:

1. `main.go` chama `repl.Start`.
2. `repl.Start` lê uma linha e passa o texto para `lexer.New`.
3. O lexer avança pela entrada e `NextToken` retorna um token por chamada.
4. O REPL imprime os tokens até receber `EOF`.
5. Os tipos e as palavras reservadas usados pelo lexer são definidos em `token/token.go`.

Uma boa ordem de estudo é começar por `token/token.go`, seguir para `lexer/lexer.go`, conferir os casos esperados em `lexer/lexer_test.go` e então ler `repl/repl.go` e `main.go`.

## Linguagem reconhecida pelo lexer

O lexer reconhece, entre outros elementos:

- Declarações: `let`, `return`
- Funções e condicionais: `fn`, `if`, `else`
- Booleanos: `true`, `false`
- Identificadores e números inteiros
- Operadores: `=`, `+`, `-`, `!`, `*`, `/`, `<`, `>`, `==`, `!=`
- Delimitadores: `,`, `;`, `(`, `)`, `{`, `}`

Exemplo de entrada:

```monkey
let soma = fn(x, y) {
  x + y;
};
```

O REPL mostra os tokens correspondentes a essa entrada, em vez de calcular o resultado de `x + y`.

## Executar

Na raiz do projeto, use:

```sh
go run .
```

Digite expressões ou declarações para visualizar os tokens. Para encerrar, use `Ctrl+C` ou envie o fim da entrada do terminal.

## Testes

Execute os testes do módulo com:

```sh
go test ./...
```

Atualmente, os testes cobrem o lexer. Ao adicionar novos componentes, crie testes junto ao pacote correspondente para documentar e validar seu comportamento.
