Em GO

:= is a short variable declaration, basicamente faz a operação de declaração e atribuição

Para criar váriaveis com package scope, não podemos usar o short operator, sobrando assim o var ou o const.

```go
package main

import "fmt"

var teste = "uma string qualquer"

func main() {
   teste = "outra string"
   fmt.Println(teste)
}

```

Ela pode ser mudada pelas funções, qualquer uma no mesmo package (não sei se consigo mudar em packages diferentes)

O short variable declaration pode ser usado se ao menos uma variavel na mesma linha seja declarada, apesar de ser muito feio.
ex:

```go

func main() {
   x := 10
   fmt.Println(x)
   x, z := 11, "teste"

}

```

No Go, boolean == bool

statement -> uma ou mais expressões
expressão -> qualquer coisa que produz um resultado

variaveis dentro de um code block sao restritas ao codeblock ORIGINAAAAL
variaveis podem ser declaradas no escopo de package (não me parece boa ideia a não ser que sejam constantes)

o var funciona em qualquer lugar.

Boas praticas do GO, sempre utilizar :=, a menos que não de, ai usa outras paradas

TIPOS

Tipo em go são estaticos (tipado)
se eu declarar a seguinte variavel em um code block

`x := 10`

Ela é do tipo int, não posso colocar uma string nela, pq da erro
Go é uma linguagem de tipagem estatica

Para fazer uma declaração com o tipo, é da seguinte forma

`var x int;`

palavra reservada var, nome da variavel, tipo
Ao ser declarada da forma do exemplo acima, ela é iniciada com 0.

O tipo geralmente é deduzido pelo compilador, mas pode ser declarado pelo dev, como nos exemplos

```go

var x string = "teste1";
y := 10.2

```

Ao rodar o seguinte comando

```go

fmt.Printf("valor: %v tipo: %T", x, x)
fmt.Printf("valor: %v tipo: %T", y, y)

```

o resultado será

valor: teste1 tipo: string
valor: teste2 tipo: float64

Se a variavel for declarada em package scope, só podemos inserir um valor nela em codeblock

exemplo

```go
 package main

import "fmt"

var test int

func main() {
	test = 10
	fmt.Println(test)
}

```

se eu tentar atribuir valor a variavel a nivel de package, da o seguinte erro

```

./prog.go:8:1: syntax error: non-declaration statement outside function body

```

Tipos de dados primitivo:

- int
- string
- bool

Tipos de dados compostos: (tipos feitos pelo dev, compostos de tipos primitivos)

- slice
- array
- struct
- map

O ato de definir, criar, estruturar tipo compostos chama-se composição.

VALOR ZERO

declaração -> em analogia, é como comprar uma caixa postal
inicialização -> receber a primeira carta na caixa postal
atribuição -> na analogia, é como receber outro valor na caixa postal

O short variable declaration faz os 3 juntos

valor zero é o valor que se encontra presente em uma váriavel antes de ela ser inicializada pelo dev

valores zeros no go

int: 0
floats: 0.0
bool: false
strings: ""
pointer, funcions, interfaces, slices, channels, maps : nil

user o := sempre que possivel
use var para package-level scope

Com tipos primitivos, segue exemplo:

```go

package main

import "fmt"

var a int
var b float64
var c string
var d bool

func main() {
	fmt.Printf("%v, %T\n", a, a)
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", c, c)
    fmt.Printf("%v, %T\n", d, d)
}

```

resultado

```

0, int
0, float64
, string
false, bool

```

Pacote fmt

Print
Printf
Println

mas antes, sobre string, existem dois tipos de string que podem ser declaradas, interpreted string literals e raw string literals

em ciencia de computação, um literal é uma notação para representar um valor fixo no código fonte

Ou seja, um int, um bool, um string são literals

interpreted string literal

```go
x := "oi, bom dia\nComo vai? \t espero que tudo bem"
```

Essa é uma string interpretada, por exemplo, onde o go acha o \n, ele não mostra um \n, e sim interpretar como uma quebra de linha

raw string literals

```go

x := `Oi, bom dia \n`

```

Nesse caso, como é um raw string literal, o go vai printar o \n

Nada vai ser interpretado, podendo incluir new lines com enter, e assim vai, parece o do javascript

Cada caractere da string é chamado de rune literal.

Agra, voltando para o pacote fmt, são usados principalmente 3 categorias de print

Print -> printa na tela, retorna o numero de bytes trabalho e se ocorreu algum erro
Sprint -> retorna string e não coloca nada na tela, ele somente retorna o que coloquei dentro dele. Só coloca espaço quando nenhum dos valores é um string.
Esse S na frente de print vem de stream.
ex:

x := "oi"
y := "bom dia"

z := fmt.Sprint(x, y)
fmt.Println(z);

resultado:

oibom dia

Fprint -> É um file print, não necessáriamente um arquivo, pq em go não tem diferença entre colocar bytes em um arquivo ou uma conexão ao servidor, é tudo um writer interface, qualquer coisa que tenha como entrada um writer, pode-se usar o Fprint

E também, print, somente printa na tela
println, coloca um new line depois do que foi printado
printf, permite formar o print com algumas coringas

Como criar meu proprio tipo

tipos são imutaveis

Seguindo o exemplo

```go

package main

import "fmt"

type hotdog int

var b hotdog

func main() {

	fmt.Printf("%v, %T\n", b, b)
}


```

o resultado será 0, main.hotdog

Esse hotdog é meu tipo, que tem um tipo subjacente, que vem por trás como base

Quando criamos nossos proprios tipos, podemos trabalhar com eles de forma que com os tipos primitivos nós não conseguimos.

Apesar de hotdog nesse exemplo tenha o tipo subjacente como int, nós não podemos atribuir um int a ele

Conversão de tipos

No exemplo anterior, foi criado um tipo chamado hotdog, e foi visto que ele não aceita um int, mesmo sendo este seu tipo subjacente

Para converter, para por exemplo um tipo hot dog ser aceito em um int

```go
package main

import "fmt"

type hotdog int

var b hotdog = 10

func main() {
	x := 10
	fmt.Printf("%v, %T\n", b, b)
	fmt.Printf("%v, %T\n", x, x)

	x = int(b)
	fmt.Printf("%v, %T\n", x, x)

}

```

Repare que ocorreu uma conversão. Conseguimos transformar uma variavel do tipo hotdog em int.
(também chamado de casting, ou coercion)

Para converter um valor para o tipo que eu quero, eu insiro i tipo e o valor que eu quero entre parentese, como no exemplo

`int(12)`

Porém, tem que ter cuidado, dá exceção em tempo de execução

Cap 4 aula 1 -> aula de boolean

```go
package main

import "fmt"

var x bool

func main() {
	fmt.Println(x) // zero value == false
	x = true
	fmt.Println(x) // atribuindo valor
	x = 10 < 100 // bool como resultado de operadores ralacionais
	fmt.Println(x)
}
```

Cap 4 aula 2 -> Como os computadores funcionam

Computadores digitais funcionam com eletrecidade e podem estar ligados ou desligados
estremas de codificação:
Com 1 lampada na varanda, que pode estar ligada ou desligada, podemos ter 2
mensagens.
Com 2 lampadas, 4 mensagens
Com 3 lampadas, 8 mensagens

    A formula, é sempre 2 elevado a n

Sabendo disso, podemos ter um esquema de codificação, atribuindo significado a cada uma
dessas mensagens
2 lampadas
on on -> festa
off off -> a mimir

esses on off, podem ser substituidos por uns e zeros, e serem chamados de BInary digiTS
bits =)

Quantificando bits
1 bit
8 bits -> 1 byte
1024 bytes -> 1 kb
1024 kb -> 1 mb
1024 mb -> 1 gb
1024 gb -> tb

Cap 4 aula 3 -> Tipos numericos

existem dois tipos numericos principais
int inteiros
float ponto flutuante

Os tipos numerais no Go são

uint -> unsigned, não tem sinal de menos
int
float
complex
byte -> alias para uint8, no caso, int sem sinal negativo com 8 bites (0 a 255)
rune -> alias para int32, no caso, um int com 32 bits

Não escolhemos com quantos bits vamos trabalhar, depende do computador que está rodando o programa.

Todos os tipos numericos são distintos, menos byte e rune.

Para verificar os bits usados por um caractere

```go
package main

import "fmt"

func main() {

	a := "e"
	b := "é"
	c := "香"
	g := "😂"

	fmt.Printf("%v %v %v %v\n", a, b, c, g)

	d := []byte(a)
	e := []byte(b)
	f := []byte(c)
	h := []byte(g)
	fmt.Printf("%v %v %v %v", d, e, f, h)
}

```

resultado 

```
e é 香 😂
[101] [195 169] [233 166 153] [240 159 152 130]
```

Da pra ver que cada byte é um item do "array", podendo ter 4, cada um indo até 255

int e int32 não são a mesma coisa, para misturar, é necessário conversão

Float point: racionais ou reais

Não da pra colocar um float em um int 

Pra verificar qual o OS e a arquitetura do pc onde está rodando o go, usar seguinte script

```go
package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)
}
```