Em GO

:= is a short variable declaration, basicamente faz a operação de declaração e atribuição


Para criar váriaveis com package scope, não podemos usar o short operator, sobrando assim o var ou o const.

```package main

import "fmt"

var teste = "uma string qualquer"

func main() {
    teste = "outra string"
    fmt.Println(teste)    
} ```


Ela pode ser mudada pelas funções, qualquer uma no mesmo package (não sei se consigo mudar em packages diferentes)


O short variable declaration pode ser usado se ao menos uma variavel na mesma linha seja declarada, apesar de ser muito feio.
ex:

```func main() {
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

```x := 10```
Ela é do tipo int, não posso colocar uma string nela, pq da erro
Go é uma linguagem de tipagem estatica

Para fazer uma declaração com o tipo, é da seguinte forma

```var x int;```

palavra reservada var, nome da variavel, tipo
Ao ser declarada da forma do exemplo acima, ela é iniciada com 0.

O tipo geralmente é deduzido pelo compilador, mas pode ser declarado pelo dev, como nos exemplos

```var x string = "teste1";
y := 10.2```

Ao rodar o seguinte comando

```fmt.Printf("valor: %v tipo: %T", x, x)
fmt.Printf("valor: %v tipo: %T", y, y)```

o resultado será

valor: teste1 tipo: string
valor: teste2 tipo: float64


Se a variavel for declarada em package scope, só podemos inserir um valor nela em codeblock

exemplo

```package main

import "fmt"

var test int

func main() {
	test = 10
	fmt.Println(test)
}```

se eu tentar atribuir valor a variavel a nivel de package, da o seguinte erro 

```./prog.go:8:1: syntax error: non-declaration statement outside function body
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
