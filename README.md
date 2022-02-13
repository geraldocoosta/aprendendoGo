Em GO

:= is a short variable declaration, basicamente faz a operação de declaração e atribuição


Para criar váriaveis com package scope, não podemos usar o short operator, sobrando assim o var ou o const.

 ``` package main

import "fmt"

var teste = "uma string qualquer"

func main() {
    teste = "outra string"
    fmt.Println(teste)    
}  ``` 


Ela pode ser mudada pelas funções, qualquer uma no mesmo package (não sei se consigo mudar em packages diferentes)


O short variable declaration pode ser usado se ao menos uma variavel na mesma linha seja declarada, apesar de ser muito feio.
ex:

 ``` func main() {
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

 ``` x := 10 ``` 
Ela é do tipo int, não posso colocar uma string nela, pq da erro
Go é uma linguagem de tipagem estatica

Para fazer uma declaração com o tipo, é da seguinte forma

 ``` var x int; ``` 

palavra reservada var, nome da variavel, tipo
Ao ser declarada da forma do exemplo acima, ela é iniciada com 0.

O tipo geralmente é deduzido pelo compilador, mas pode ser declarado pelo dev, como nos exemplos

 ``` var x string = "teste1";
y := 10.2 ``` 

Ao rodar o seguinte comando

 ``` fmt.Printf("valor: %v tipo: %T", x, x)
fmt.Printf("valor: %v tipo: %T", y, y) ``` 

o resultado será

valor: teste1 tipo: string
valor: teste2 tipo: float64


Se a variavel for declarada em package scope, só podemos inserir um valor nela em codeblock

exemplo

 ``` package main

import "fmt"

var test int

func main() {
	test = 10
	fmt.Println(test)
} ``` 

se eu tentar atribuir valor a variavel a nivel de package, da o seguinte erro 

 ``` ./prog.go:8:1: syntax error: non-declaration statement outside function body
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

```0, int
0, float64
, string 
false, bool```


Pacote fmt

Print
Printf
Println

mas antes, sobre string, existem dois tipos de string que podem ser declaradas, interpreted string literals e raw string literals

em ciencia de computação, um literal é uma notação para representar um valor fixo no código fonte

Ou seja, um int, um bool, um string são literals

 interpreted string literal

 x := "oi, bom dia\nComo vai? \t espero que tudo bem"

 Essa é uma string interpretada, por exemplo, onde o go acha o \n, ele não mostra um \n, e sim interpretar como uma quebra de linha

raw string literals

x :=  `Oi, bom dia \n`

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

Fprint ->  É um file print, não necessáriamente um arquivo, pq em go não tem diferença entre colocar bytes em um arquivo ou uma conexão ao servidor, é tudo um writer interface, qualquer coisa que tenha como entrada um writer, pode-se usar o Fprint

E também, print, somente printa na tela
println, coloca um new line depois do que foi printado
printf, permite formar o print com algumas coringas

