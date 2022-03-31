# GO

:= is a short variable declaration, basicamente faz a operação de declaração e atribuição

Para criar variáveis com package scope, não podemos usar o short operator, sobrando assim o var ou o const.

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

O short variable declaration pode ser usado se ao menos uma variável na mesma linha seja declarada, apesar de ser muito feio.
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

variáveis dentro de um code block sao restritas ao codeblock ORIGINAAAAL
variáveis podem ser declaradas no escopo de package (não me parece boa ideia a não ser que sejam constantes)

o var funciona em qualquer lugar.

Boas praticas do GO, sempre utilizar :=, a menos que não de, ai usa outras paradas

TIPOS

Tipo em go são estáticos (tipado)
se eu declarar a seguinte variável em um code block

`x := 10`

Ela é do tipo int, não posso colocar uma string nela, pq da erro
Go é uma linguagem de tipagem estática

Para fazer uma declaração com o tipo, é da seguinte forma

`var x int;`

palavra reservada var, nome da variável, tipo
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

Se a variável for declarada em package scope, só podemos inserir um valor nela em codeblock

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

se eu tentar atribuir valor a variável a nível de package, da o seguinte erro

```log
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

valor zero é o valor que se encontra presente em uma variável antes de ela ser inicializada pelo dev

valores zeros no go

int: 0
floats: 0.0
bool: false
strings: ""
pointer, functions, interfaces, slices, channels, maps : nil

user o := sempre que possível
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

```log

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

em ciência de computação, um literal é uma notação para representar um valor fixo no código fonte

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

Fprint -> É um file print, não necessariamente um arquivo, pq em go não tem diferença entre colocar bytes em um arquivo ou uma conexão ao servidor, é tudo um writer interface, qualquer coisa que tenha como entrada um writer, pode-se usar o Fprint

E também, print, somente printa na tela
println, coloca um new line depois do que foi printado
printf, permite formar o print com algumas coringas

Como criar meu próprio tipo

tipos são imutáveis

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

Quando criamos nossos próprios tipos, podemos trabalhar com eles de forma que com os tipos primitivos nós não conseguimos.

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

Repare que ocorreu uma conversão. Conseguimos transformar uma variável do tipo hotdog em int.
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
    x = 10 < 100 // bool como resultado de operadores relacionais
    fmt.Println(x)
}
```

- Cap 4 aula 2 -> Como os computadores funcionam

- Computadores digitais funcionam com eletricidade e podem estar ligados ou desligados
  - estremas de codificação:
    - Com 1 lampada na varanda, que pode estar ligada ou desligada, podemos ter 2 mensagens.
    - Com 2 lampadas, 4 mensagens
    - Com 3 lampadas, 8 mensagens

A formula, é sempre 2 elevado a n

Sabendo disso, podemos ter um esquema de codificação, atribuindo significado a cada uma
dessas mensagens
2 lampadas
on on -> festa
off off -> a mimir

esses on off, podem ser substituídos por uns e zeros, e serem chamados de BInary digiTS
bits =)

Quantificando bits
1 bit
8 bits -> 1 byte
1024 bytes -> 1 kb
1024 kb -> 1 mb
1024 mb -> 1 gb
1024 gb -> tb

Cap 4 aula 3 -> Tipos numéricos

existem dois tipos numéricos principais
int inteiros
float ponto flutuante

- Os tipos numerais no Go são

- uint -> unsigned, não tem sinal de menos
- int
- float
- complex
- byte -> alias para uint8, no caso, int sem sinal negativo com 8 bites (0 a 255)
- rune -> alias para int32, no caso, um int com 32 bits

Não escolhemos com quantos bits vamos trabalhar, depende do computador que está rodando o programa.

Todos os tipos numéricos são distintos, menos byte e rune.

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

```log
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

Cap 4 aula 4 -> Overflow

Se eu estourar a capacidade do type, o que acontece?
Erro, como em quase todas as linguagens, só fica ligado com isso e tá de boa.

Porém, se eu usar a atribuição com ++, ele volta a ser 0.

Cap 4 aula 5 -> Strings

Strings são sequencias de bytes imutáveis.
Slice of byte em go.
Em go, pode ser feita entre "" ou ``

Pra converter em slice de bytes

```go
s := "Teste"
sb := []byte(s);
fmt.Println(sb)

for _, v := range sb {
    fmt.Printf("%v - %T - %#U - %#x \n", v, v, v, v);
}


```

Resultado

```log

[84 101 115 116 101]
84 - uint8 - U+0054 'T' - 0x54
101 - uint8 - U+0065 'e' - 0x65
115 - uint8 - U+0073 's' - 0x73
116 - uint8 - U+0074 't' - 0x74
101 - uint8 - U+0065 'e' - 0x65

```

Também pode ser feito dessa maneira esse loop

```go
s := "teste é assim mesmo e é 香 😂"


for _, v := range s {
    fmt.Printf("%v - %T - %#U - %#x \n", v, v, v, v);
}

fmt.Println();

for i := 0; i < len(s); i++ {
    fmt.Printf("%v - %T - %#U - %#x \n", s[i], s[i], s[i], s[i]);
}

```

Resultado

```log
116 - int32 - U+0074 't' - 0x74
101 - int32 - U+0065 'e' - 0x65
115 - int32 - U+0073 's' - 0x73
116 - int32 - U+0074 't' - 0x74
101 - int32 - U+0065 'e' - 0x65
32 - int32 - U+0020 ' ' - 0x20
233 - int32 - U+00E9 'é' - 0xe9
32 - int32 - U+0020 ' ' - 0x20
97 - int32 - U+0061 'a' - 0x61
115 - int32 - U+0073 's' - 0x73
115 - int32 - U+0073 's' - 0x73
105 - int32 - U+0069 'i' - 0x69
109 - int32 - U+006D 'm' - 0x6d
32 - int32 - U+0020 ' ' - 0x20
109 - int32 - U+006D 'm' - 0x6d
111 - int32 - U+006F 'o' - 0x6f
115 - int32 - U+0073 's' - 0x73
109 - int32 - U+006D 'm' - 0x6d
111 - int32 - U+006F 'o' - 0x6f
32 - int32 - U+0020 ' ' - 0x20
101 - int32 - U+0065 'e' - 0x65
32 - int32 - U+0020 ' ' - 0x20
233 - int32 - U+00E9 'é' - 0xe9
32 - int32 - U+0020 ' ' - 0x20
39321 - int32 - U+9999 '香' - 0x9999
32 - int32 - U+0020 ' ' - 0x20
128514 - int32 - U+1F602 '😂' - 0x1f602

116 - uint8 - U+0074 't' - 0x74
101 - uint8 - U+0065 'e' - 0x65
115 - uint8 - U+0073 's' - 0x73
116 - uint8 - U+0074 't' - 0x74
101 - uint8 - U+0065 'e' - 0x65
32 - uint8 - U+0020 ' ' - 0x20
195 - uint8 - U+00C3 'Ã' - 0xc3
169 - uint8 - U+00A9 '©' - 0xa9
32 - uint8 - U+0020 ' ' - 0x20
97 - uint8 - U+0061 'a' - 0x61
115 - uint8 - U+0073 's' - 0x73
115 - uint8 - U+0073 's' - 0x73
105 - uint8 - U+0069 'i' - 0x69
109 - uint8 - U+006D 'm' - 0x6d
32 - uint8 - U+0020 ' ' - 0x20
109 - uint8 - U+006D 'm' - 0x6d
111 - uint8 - U+006F 'o' - 0x6f
115 - uint8 - U+0073 's' - 0x73
109 - uint8 - U+006D 'm' - 0x6d
111 - uint8 - U+006F 'o' - 0x6f
32 - uint8 - U+0020 ' ' - 0x20
101 - uint8 - U+0065 'e' - 0x65
32 - uint8 - U+0020 ' ' - 0x20
195 - uint8 - U+00C3 'Ã' - 0xc3
169 - uint8 - U+00A9 '©' - 0xa9
32 - uint8 - U+0020 ' ' - 0x20
233 - uint8 - U+00E9 'é' - 0xe9
166 - uint8 - U+00A6 '¦' - 0xa6
153 - uint8 - U+0099 - 0x99
32 - uint8 - U+0020 ' ' - 0x20
240 - uint8 - U+00F0 'ð' - 0xf0
159 - uint8 - U+009F - 0x9f
152 - uint8 - U+0098 - 0x98
130 - uint8 - U+0082 - 0x82

```

range em um stream não dá bite por bite, e sim caractere por caractere

no segundo for,ele converte para bite

Cap 4 Aula 6 - Sistemas numéricos

Decimal, numéricos, Hexadecimal

decimal - base 10 - 0-9
binário - base 2 - 0-1
hexadecimal - base 16 - 0-f

Vi isso na faculdade, só verificando

Para ver esse valores em GO

```go
package main

import "fmt"

func main() {

    valor := 100;

    fmt.Printf("%d, %b, %#x\n", valor, valor, valor)

}


```

resultado

```log

100, 1100100, 0x64

```

Cap 4 Aula 7 - Constantes (aeeeeee)

Constantes são variáveis imutáveis, que não pode mudar a instancia nunca

Podem ou não ser tipada, por exemplo

```go
const oi := "bom dia"
const oiTipado string := "BOOOM DIA"
```

As constantes não tipadas, em go, só ganham o tipo quando forem usadas,
Isso é feito para ganhar flexibilidade, pois o go atribui tipo dependendo do host que está rodando o programa

```go
package main

import "fmt"

const b = 10

var a float64

func main() {
    a = b
    fmt.Println(a)
}
```

No momento que o programa roda, e a interação do a recebendo b acontece, a contante b vira um float, antes disso, ela não tem tipo atribuído
O tipo de uma variável é definido em tempo de atribuição.

Pode-se declarar constantes assim também

```go
package main

import "fmt"

const (
    x = 10
    y = 10
    z = 10
)


func main() {
    fmt.Println(a, y, z)
}
```

Cap 4 Aula 8 - IOTAs

```go
package main

import "fmt"

const (
    x = iota
    y = iota
    z = iota
)

var a float64

func main() {
    a = x

    fmt.Println(a, y, z)
}
```

Resultado:

```log
0 1 2
```

Números inteiros sucessivos na declaração de constantes, que podem ser usados como float também, pois são constantes

Serve para quando não ligamos para o valor da constante, só queremos que o valor seja diferente das outras (ENUM LIKE?)

É possível descartar valores usando underline(\_)

Também é possível fazer isso:

```go
import "fmt"

const (
    a = iota * 2
    b
    _
    d
)

func main() {
    fmt.Println(a,b,d)
}
```

Resultado

```log

0 1 3

```

Podemos colocar uma formula depois do iota da primeira linha, e o go vai repetir a formula para os outros, para usar só com + 1, pode deixar só o primeiro iota que ele repete nos outros

Cap 4 Aula 9 - Deslocamento de bits

Deslocamento de bites é literalmente deslocar dígitos binários da direita pra esquerda ou da esquerda para a direita

Como fazer isso?

```go

package main

import "fmt"

func main() {
    x := 1
    y := x << 1

    a := 10
    b := a >> 1

    fmt.Printf("%v - %b\n", x, x)
    fmt.Printf("%v - %b\n", y, y)
    fmt.Printf("%v - %b\n", a, a)
    fmt.Printf("%v - %b\n", b, b)

}

```

Não precisa ser só 1, pode ser um deslocamento de mais casas

- CAP 6 - Entendendo fluxo de controle

Os tipos de fluxo de controle:

- Sequencial -> uma linha depois da outra, da esquerda para a direita
- Condicional -> if, switch
- Repetição -> while, do while, for

- Cap. 6 – Fluxo de Controle – 2. Loops: inicialização, condição, pós

Exemplo pratico

```go
        for var i int = 0; i < 10; i++ {
        fmt.Println(i)
    }

```

Pra mim, que estou vindo do java, a diferença é, depois da palavra reservada for, não tenho parenteses, apenas chaves para o corpo do for

De resto, é bem igual, tirando a forma que o short variable declaration funciona.

O seguinte exemplo não funciona

```go
    // ISSO DÁ ERRO
    for var i int = 0; i < 10; i++ {
        fmt.Println(i)
    }
```

Dá o seguinte erro

```log
./prog.go:10:10: syntax error: var declaration not allowed in for initializer
```

Aparentemente, pra variável contadora usada em um for, não é aceito a palavra reservada var para inicialização

O seguinte exemplo funciona:

```go

func main() {
    var i int = 0
    for ; i < 10; i++ {
        fmt.Println(i)
    }
}

```

- IMPORTANTE:

  - não existe while :)
  - No final de todo statement (relembrando, uma ou mais expressões), o compilador do go adiciona o ponto e virgula pra gente, porém, quando temos mais de um statement na mesma linha, temos que adicionar manualmente

- Cap. 6 – Fluxo de Controle – 4. Loops: a declaração for

O for funciona do Go pode funcionar como um while

```go

func main() {
    i := 0
    for i < 10 {
        fmt.Println(i)
        i++
    }
}

```

Explicando o exemplo, para funcionar como while, deve ter só um statement, diferente do for normal que tem 3 statements separados por ponto e virgula (;)

Podemos criar uma condição eterna com isso, não precisamos declarar nenhum statement para o for

```go

func main() {
    for {
        fmt.Println("Loop forever")
    }
}

```

O for também funciona com um range, que será visto mais a frente no curso.

- Cap. 6 – Fluxo de Controle – 5. Loops: break & continue

```go

func main() {

    for i := 0; i < 100; i++ {
        if i%2 != 0 {
            fmt.Println("Irei pular para o próximo, esse numero não é legal")
            continue
        }
        if i == 58 {
            fmt.Println("Cansei desse loop, vou para casa")
            break
        }
        fmt.Println(i)
    }

    fmt.Println("Indo para casa")

}

```

Funciona igual ao java
Ao bater no break vai para o próximo statement após o for
Ao bater em um continue dentro do for, imediatamente acrescenta um e volta para o começo do code block do for

O modulo (%), também não tinha sido explicado, mas funciona igual as outras linguagens, modulo de uma operação, resto da divisão

- Cap. 6 – Fluxo de Controle – 7. Condicionais: a declaração if

Funciona basicamente como em outras linguagens, porém, a expressão booleana não fica entre parenteses
Porém, dentro do if, ele aceita short declaration variable junto com o if ou seja

```go
package main

import "fmt"

func main() {
    if x := 10; x < 100 {
        fmt.Println("Hello World")
    }
}

```

Basicamente, isso seria um operador de inicialização, executado antes do if rodar.

- Cap. 6 – Fluxo de Controle – 8. Condicionais: if, else if, else

Do mesmo jeito de outras linguagens, porém temos isso:

```go
    if x := 12; x < 12 {
        fmt.Println("chis é menor que doze meoooo")
    } else if x > 12 {
        fmt.Println("chis é maior que doze meoooo")
    } else {
        fmt.Println("chis é 12 meoooo")
    }
```

- Cap. 6 – Fluxo de Controle – 9. Condicionais: a declaração switch

Switch no Go, tem sua peculiaridades, por exemplo, não é necessário falar para o go qual variável eu estou fazendo o switch

Ex:

```go
package main

import "fmt"

func main() {

    x := 10;

    switch {
        case x < 5: fmt.Println("chis é  < 5 ")
        case x == 5: fmt.Println("chis é  == 5 ")
        case x > 5: fmt.Println("chis é > 5 ")
    }

}

```

O switch em go, aceita uma variável na frente da palavra reservada switch, que vai comparar o resultado do case com aquele valor.

```go
package main

import "fmt"

func main() {

    x := 10;

    switch false{
        case x < 5: fmt.Println("chis é  < 5 ")
        case x == 5: fmt.Println("chis é  == 5 ")
        case x > 5: fmt.Println("chis é > 5 ")
    }

}

```

Resultado

```log
chis é  < 5 
```

Mas como assim ? o x é igual a 10!! Sim meu jovem, porém, o go vai pegar o resultado da expressão do case, e comparar com o valor depois da palavra chave switch, se essa comparação for verdadeira, então ele será o caso correcto.

Basicamente, é o equivalente a `if (x < 5) == false`.

Na verdade, quando não colocamos nada, aquilo é simplesmente true. Podemos colocar até string se a gente quiser.

Não existe fall through por padrão, ou seja, não é que nem java, não precisamos de vários breaks no switch

Para fall through, devemos falar pro go que queremos esse comportamento

```go
package main

import "fmt"

func main() {

    x := 10;

    switch {
        case x == 10: 
            fmt.Println("fazendo algo, pois 10 é igual 10")
            fallthrough
        case x < 10: 
            fmt.Println("fazendo algo, pois 10 é menor 10")
        case x >10: 
            fmt.Println("fazendo algo, pois 10 é maior 10")
    }

}

```

Como eu ativei o fallthrough, veja o resultado:

```log
fazendo algo, pois 10 é igual 10
fazendo algo, pois 10 é menor 10
```

Esse comportamento só acontece quando coloco a palavra chave fallthrough no bloco de um case

valor default

```go
package main

import "fmt"

func main() {

    x := "rodolfo"
    y := "maria"
    z := "alfredo"

    switch "zé"{
        case x:
            fmt.Println("Rodolfinhooooooo")
        case y:
            fmt.Println("Mariaaaaaaaaaaa")
        case z:
            fmt.Println("Alfred")
        default:
            fmt.Println("Mas que coisa, parece que nada")
    }

}

```

Switch composto

```go
package main

import "fmt"

func main() {

    x := "rodolfo"
    y := "maria"
    z := "alfredo"

    switch "zé"{
        case x:
            fmt.Println("Rodolfinhooooooo")
        case y:
            fmt.Println("Mariaaaaaaaaaaa")
        case z:
            fmt.Println("Alfred")
        default:
            fmt.Println("Mas que coisa, parece que nada")
    }

}

```

```go
package main

import "fmt"

func main() {

    x := "rodolfo"
    y := "maria"
    z := "alfredo"
    w := "zé"

    switch "zé"{
        case x, y:
            fmt.Println("Rodolfinhooooooo ou Mariaaaaaa")
        case z, w:
            fmt.Println("Alfred e zé")
        default:
            fmt.Println("Mas que coisa, parece que nada")
    }

}

```

Output

```log
Alfred e zé
```

Não irei explicar, tu já desenvolve a alguns anos, foca ai que tu consegue Geraldo

Mas lembrando que, esse switch também aceita essa sintaxe

```go
package main

import "fmt"

func main() {

    x := 1

    switch {
        case (x == 4), (x == 3):
            fmt.Println("4 ou 3")
        case (x == 2), (x == 1):
            fmt.Println("2 ou 1")
    }

}

```

Agora dá pra ir longe, pois podemos colocar várias expressão, ou statements

Lembrando que, a virgula ali é meio que um ou.

- Cap. 6 – Fluxo de Controle – 10. Condicionais: a declaração switch pt. 2 & documentação

Basicamente, acho o switch também identifica types

```go
package main

import "fmt"

var x interface{}

func main() {

    x = 10

    switch x.(type) {
        case int:
            fmt.Println(x, "é um int")
        case bool:
            fmt.Println(x, "é um bool")
        case string:
            fmt.Println(x, "é uma string")
        case float64:
            fmt.Println(x, "é um float64")
    }
}
```

Uma coisa maneirinha, é isso:

```go
package main

import "fmt"

func main() {
    switch x := 1; {
        case x == 1: fmt.Printf("O valor é %d", 1)
        case x == 2: fmt.Printf("O valor é %d", 2)
        case x == 3: fmt.Printf("O valor é %d", 3)
        case x == 4: fmt.Printf("O valor é %d", 4)
    } 
}
```

Se o switch não der match com nada, aparentemente não dá erro.

A única parte que não vi, é se podemos usar switch para atribuição de variável

- Cap. 6 – Fluxo de Controle – 11. Operadores lógicos condicionais

São o && e o ||, segue o mesmo padrão de outras linguagens.

- Cap. 8 – Agrupamentos de Dados – 1. Array

Existem na GoLang

Array, Slice, Maps, Structs

Array: Tipo não utilizado tanto no dia a dia, é o bloco fundamental aonde o slice é construído

Exemplo de Array:

```go
package main

import (
    "fmt"
)

var x [5]int

func main() {
    x[0] = 1
    x[1] = 10

    fmt.Println(x[0])
    fmt.Println(x[1])
    fmt.Println(x)
}
```

Resultado

```log
[1 10 0 0 0]
```

Array: Uma estrutura de dados, que pode carregar mais de um dado de um tipo, o tamanho de um array deve ser definido antes de utiliza-lo
Integralmente, o tipo de array contém o tamanho dele, logo na declaração da variável, coloca-se o tamanho, e se diz assim: "É um array de 5 ints"

Logo, ao se fazer isso:

```go
package main

import (
    "fmt"
)

var x [5]int

func main() {

    fmt.Printf("%T", x) //Lembrando que o %T mostra o tipo
}

```

A saída será

```go
[5]int

```

O array tem tamanho fixo, já cansei de falar, e se você tentar colocar um valor nele além do seu limite, dá um erro, exemplo:

```go
package main

import (
    "fmt"
)

var x [5]int

func main() {
    x[6] = 8
    fmt.Printf("%T", x)
}
```

Resultado

```log
./prog.go:10:3: invalid array index 6 (out of bounds for 5-element array)
```

Um array de 6 ints e de 5 ints não são compatíveis entre si.

Funções que podemos utilizar em arrays:

```go
package main

import (
    "fmt"
)

var x [5]int

func main() {
    fmt.Println(len(x))
}
```

```log
5
```

O bom de saber de array, é saber que não vamos utilizar array, então basicamente é isso, funciona basicamente como no Java.

Arrays são uteis quando precisamos planejar o layout da memoria em detalhes (muito baixo nível)

Use SLICE.

- Cap. 8 – Agrupamentos de Dados – 2. Slice: literal composta

Em arrays, podemos inicializar valores da seguinte forma

```go
package main

import "fmt"

func main() {
    array := [5]int{1, 2, 3, 4, 5}
    fmt.Println(array)
    slice := []int{1, 2, 3, 4, 5}
    fmt.Println(slice)
}

```

Resultado

```log
[1 2 3 4 5]
[1 2 3 4 5]
```

Qual a diferença entre o array e o slice então?

Simples, o Slice não declara o tanto de valores que pode receber.

Para mudar o tamanho do slice, adicionando um no final

```go
package main

import "fmt"

func main() {
    array := [5]int{1, 2, 3, 4, 5}
    fmt.Println(array)
    slice := []int{1, 2, 3, 4, 5}
    fmt.Println(slice)

    slice2 := append(slice, 6)
    fmt.Println(slice2)
}
```

```log
[1 2 3 4 5]
[1 2 3 4 5]
[1 2 3 4 5 6]
```

Para pegar itens do slice, é da mesma forma que no array, `slice[3]`

Pra incluir outro elemento, também é a mesma forma `slice[3] = 123123`

Porém, não conseguimos fazer isso: `slice[20] = 123123`

Dá esse erro:

```log
panic: runtime error: index out of range [20] with length 5

goroutine 1 [running]:
main.main()
    /tmp/sandbox341066696/prog.go:16 +0x1d7

Program exited.
```

O slice é feito de arrays, então temos que usar métodos para utilizar o array

- Cap. 8 – Agrupamentos de Dados – 3. Slice: for range

o range significa alcance, faixa, extensão, ou seja, ela atravessa toda a extensão do slice

```go
package main

import "fmt"

func main() {
    slice := []string{"banana", "avocado", "orange"}
    for indice, valor := range slice {
        fmt.Printf("No indice %d temos o valor %v\n", indice, valor)
    }
}
```

```log
No indice 0 temos o valor banana
No indice 1 temos o valor avocado
No indice 2 temos o valor orange

Program exited.
```

Com a função range a gente consegue passar por todos os valores, usando o for no formato acima

Se lembre, não adicione por indice no slice, é meio ruim, pode dar erro, faz isso não

Use a função append

O range retorna duas variáveis, pra usa>

- Cap. 8 – Agrupamentos de Dados – 3. Slice: fatiando ou deletando de uma fatia

A sintaxe do slice é essa:

```go
package main

import "fmt"

func main() {
    sabores := []string{"margarita", "calabresa", "peperoni", "abacaxi", "quatro queijos"}

    fatia := sabores[2:3]

    fmt.Println(fatia)
}

```

Ele funciona da seguinte forma, o primeiro parâmetro entre colchetes, será o indice de um elemento.

O segundo, será o indice do elemento que eu quero que var aquela "fatia", porém esse indice é mais um

No exemplo, vai trazer somente peperoni, pois o indice do peperoni é 2, e o indice de peperoni é 2, +1 fica 3.

Se eu quero pegar a de abacaxi até quatro queijos, seria:

```go
fatia := sabores[3:5]
```

lembrando que, o segundo argumento não pode ser menor que o primeiro, pois dá erro

```log
./prog.go:10:18: invalid slice index: 3 > 2
```

Outro bom exemplo, que usa o método len

```go
package main

import "fmt"

func main() {
    sabores := []string{"margarita", "calabresa", "peperoni", "abacaxi", "quatro queijos"}

    fatia := sabores[3:len(sabores)]

    fmt.Println(fatia)
}

```

Ou também assim:

```go
package main

import "fmt"

func main() {
    sabores := []string{"margarita", "calabresa", "peperoni", "abacaxi", "quatro queijos"}

    fatia := sabores[3:]

    fmt.Println(fatia)
}

```

Como não temos segundo argumento, automaticamente vai ser até o final

Se não tiver o primeiro argumento, automaticamente vai ser 0

```go
package main

import "fmt"

func main() {
    sabores := []string{"margarita", "calabresa", "peperoni", "abacaxi", "quatro queijos"}

    fatia := sabores[:4]

    fmt.Println(fatia)
}

```

É fatiando que se deleta um item de um slice, usando o exemplo anterior, iria ficar assim

```go
    sabores = append(sabores[:3], sabores[4:]...)
```

- Cap. 8 – Agrupamentos de Dados – 5. Slice: anexando a uma slice

Como funciona a função append?

Sabemos que ela anexa itens a um slice, e faz parte do package builtin

Essa função recebe um slice de um tipo, e vários elementos de outro tipo como segundo argumento em diante

No exemplo passado, da aula anterior, utilizou os 3 pontos para o go entender que são vários ints.

Os 3 pontos pode ser chamado como unfurl ou enumerations. Basicamente, desenrola os itens do slice e coloca um de cada vez.

- Cap. 8 – Agrupamentos de Dados – 6. Slice: make

Slice são feitos de array

Slices são dinâmicas, ou seja, podem mudar de tamanho
Sempre que um slice muda de tamanho, um novo array é criado, e os dados são movidos do antigo para o atual

É conveniente mas tem um custo computacional

Para otimizar as coisas, podemos usar o make

`make([]T, len, cap)`

O make nos permite construir um slice que usa um array de 10 elementos, porém com a capacidade de 50 elementos

O slice tem um LENgth

O array por trás tem um CAPacidade

Para verificarmos isso, temos os métodos len(x) e cap (x)

Se eu tiver capacidade no Array que é base para meu slice, e inserir um item novo, o valor será inserido no meu array, e nele será feito um slice que resulta no meu Slice. Se eu não tiver capacidade, um novo array com capacidade maior será criado. Ou seja, aparentemente o Slice é apenas um syntax sugar para arrays.

Para criar um slice com o método make, passando seu tamanho e capacidade

```go
package main

import "fmt"

func main() {
    slice := make([]int, 5, 10)
}
```

O make vai criar o slice com valores defaults, ou seja, coloquei que tenho length de 5, ele vai preencher de 0 os que eu não preencher

```go
package main

import "fmt"

func main() {
    slice := make([]int, 5, 10)
    fmt.Println(slice)
}
```

```log
[0 0 0 0 0]
```

Se eu tenho um slice de cap 5, eu posso colocar valores assim até seu indice 4

```go
slice[0] = 123
slice[1] = 123
slice[2] = 123
slice[3] = 123
slice[4] = 123
```

O sexto, ira dar erro, out of range.

Para adicionar depois do quinto elemento, nesse caso, iremos utilizar o método append

Sobre o Capability, se eu passar do valor de capability que passei no começo, aparentemente esse valor é dobrado, e gerado um novo array de base para o slice.

Se você já sabe com quantos elementos vai trabalhar antes de criar o slice, você já podia criar o slice com um certo tamanho, para que o array de base do slice não seja recriado a todo momento.

- Cap. 8 – Agrupamentos de Dados – 7. Slice: slice multi-dimensional

Slice dentro de um slice.

Conceito parece ser o mesmo de outras linguagens, porém a sintaxe é a seguinte:

```go
package main

import "fmt"

func main() {
    ss := [][]int{
        []int{1, 2, 3},
        []int{4, 5, 6},
        []int{7, 8, 9},
    }
    fmt.Println(ss)
    fmt.Println(ss[1])
    fmt.Println(ss[1][1])

}
```

```log
[[1 2 3] [4 5 6] [7 8 9]]
[4 5 6]
5
```

Reparar como é a declaração `[][]type`.

- Cap. 8 – Agrupamentos de Dados – 8. Slice: a surpresa do array subjacente

Cuidado ao usar APPEND, pq ele faz um reslice do array subjacente.

Exemplo:

```go
import "fmt"

func main() {
    primeiroSlice := []int{1, 2, 3, 4, 5}

    fmt.Println(primeiroSlice)

    segundoSlice := append(primeiroSlice[:2], primeiroSlice[4:]...)

    fmt.Println(segundoSlice)

    fmt.Println(primeiroSlice)
}
```

```log
[1 2 3 4 5]
[1 2 5]
[1 2 5 4 5]
```

Ou seja, ele modificou o slice de origem, meio que a segundo slice ficou com um referencia dos itens de index 0 até index 3.

Ou seja, o array subjacente do segundoSlice é o mesmo do primeiroSlice, porém, com leg == 3 e o primeiro array com length == 5

Não é útil, mas é perigoso.

Quando for gerar um novo slice a partir de outro, faz um loop for ou modifica a referencia mesmo, mas não cria outro slice a partindo do append.

- Cap. 8 – Agrupamentos de Dados – 9. Maps: introdução

Estrutura de dados key value, com ordem aleatória, porém muito performático

```go
package main

import "fmt"

func main() {
    amigos := map[string]int{
        "alfredo": 12312312,
        "joana":   123123123,
    }
    fmt.Println(amigos["joana"])

    amigos["gopher"] = 444444
    fmt.Println(amigos["joana"])

    gopher := amigos["gopher"]
    fmt.Println(gopher)

    // se o ok for false, quem eu pesquisei no map não existe
    fantasma, ok := amigos["fantasma"]
    fmt.Println(fantasma, ok)


    if será, ok :=  amigos["fantasma"]; ok {
        fmt.Println(será, "existe?")
    } else {
        fmt.Println("Não existe", ok)
    }

}
```

A sintaxe da declaração é meio estranha, mas tudo bem.

O segundo argumento que a busca de uma chave no map é um ok value, se for true, existe, se for false, não existe.

- Cap. 8 – Agrupamentos de Dados – 10. Maps: range & deletando

```go
package main

import "fmt"

func main() {
    qualquercoisa := map[int]string{
        123: "muito legal",
        456: "not cool",
        789: "this sucks",
    }

    for key, value := range qualquercoisa {
        fmt.Println(key, value)
    }

}
```

Range passa por todos os valores do map, padrão

Para deletar, a sintaxe é essa

```go
delete(variaveldoMap, chaveParaExcluir)
```

- Cap. 10 – Structs – 1. Struct

Struct é um tipo de dado cujo nome vem da palavra inglesa structure, nos permite armazenar dados com nomes diferentes

Exemplo

```go
package main

import "fmt"

type cliente struct {
    nome      string
    sobrenome string
    fumante   bool
}

func main() {
    cliente1 := cliente{
        nome:      "Geraldo",
        sobrenome: "Costa",
        fumante:   true,
    }
    cliente2 := cliente{"José", "Oliveira", false}
    fmt.Println(cliente1)
    fmt.Println(cliente2)
}

```

Aparentemente, esses {} depois do nome da estrutura é tipo um construtor

- Cap. 10 – Structs – 2. Structs embutidos

Structs embutidos é um struct dentro de outro struct.

Eu posso ter quantos structs dentro de structs eu quiser.

Exemplo

```go
package main

import "fmt"

type pessoa struct {
    nome  string
    idade int
}

type profissional struct {
    pessoa
    titulo  string
    salario int
}

func main() {
    pessoa1 := pessoa{
        nome:  "Alfredo",
        idade: 30,
    }
    pessoa2 := profissional{
        pessoa: pessoa{
            nome:  "Mario",
            idade: 31,
        },
        titulo:  "Pizzaiolo",
        salario: 2500,
    }
    fmt.Println(pessoa1, pessoa2)
}
```

- Cap. 10 – Structs – 3. Lendo a documentação

Como acessar campos da struct?

```go

func main() {
    pessoa1 := pessoa{
        nome:  "Alfredo",
        idade: 30,
    }
    pessoa2 := profissional{
        pessoa: pessoa{
            nome:  "Mario",
            idade: 31,
        },
        titulo:  "Pizzaiolo",
        salario: 2500,
    }
    fmt.Println(pessoa1, pessoa2)
    fmt.Println(pessoa1.idade)
    fmt.Println(pessoa2.pessoa.nome)
    fmt.Println(pessoa2.idade)
}
```

repare na linha pessoa2.idade. Isso funciona, se não tiver conflito entre a struct de fora e a struct de dentro, o go verifica se a propriedade ta dentro da struct filha. Em go, é dito que o campo foi "promovido" de um subcampo para um campo externo

Como fazer a forma concisa do struct composto?

```go
pessoa4 := profissional{pessoa{"Vanderlei", 70}, "Politico", 10000}
```

Essa é a forma concisa do struct composto

Essa é a [documentação](https://go.dev/ref/spec#Struct_types)

- Cap. 10 – Structs – 4. Structs anônimos

É um struct feito na hora, que não é reutilizável, mas pode ser útil

```go
package main

import "fmt"

func main() {
    x := struct {
        nome  string
        idade int
    }{
        nome:  "Geraldo",
        idade: 54,
    }
    fmt.Println(x)
}
```

- Cap. 12 – Funções – 1. Sintaxe

As funções servem para abstrair funcionalidades.

Exemplo, o fmt.Println, é uma função do pacote fmt, ele abstrai a funcionalidade de imprimir coisas no output padrão (terminal no caso). Não sei como funciona dentro da função, mas sei como chamar a função, e ela faz o que eu quero por mim.

É muito utilizada pra reutilizar código.

Toda função no go segue a seguinte estrutura

```go
func (receiver) identifier (parameters) (returns) {code}
```

func é a palavra chave para declarar uma função, o receiver, identifier, literalmente o nome da função, parameters, o retorno e o code block

Diferença entre parâmetros e argumentos, a declaração da função tem parâmetros, a chamada a uma função tem argumentos.

Se eu tiver mais de um retorno em uma função em go, eu preciso usar os parenteses, se não, pode não utilizar.

Receiver tem a ver com métodos, ou seja, funções que estão linkadas a um objeto.

Quase tudo em Go é *pass by value*, (tem 1 exceção) ou seja, não é que nem java que tem o pass by reference, que você passa a referencia de um objeto para o método (dependendo do parâmetro), o go passa o valor, provavelmente, se quiser passar a referencia, a gente passa o endereço de memoria, like C.

Exemplo de função básica, que aceita argumento e retorna

```go
package main

import "fmt"

func main() {
    fmt.Println(soma(1, 2))
}

        // os parâmetros podem ser escritos das seguintes formas
        //( x int, y int) ou ( x, y int)
        // a segunda só funciona para o mesmo tipo
func soma(x, y int) int {
    return x + y
}
```

Exemplo de função com dois retornos, e recebendo um argumento variado, que tem os tres pontos antes do tipo nos parâmetros da função

```go
package main

import "fmt"

func main() {

    total, quantos := soma(10, 10, 1, 2, 3, 5)

    fmt.Println(total, quantos)
}

func soma(x ...int) (int, int) {
    soma := 0
    for _, v := range x {
        soma += v
    }
    return soma, len(x)
}

```

O parâmetro variado de um função tem que vir no ultimo parâmetro, se não tivesse essa regra, iria virar uma zona e o compilador talvez ficasse perdido.

- Cap. 12 – Funções – 2. Desenrolando (enumerando) uma slice

Para receber um slice em um função qu recebe argumento variado, é só desenrolar com o operador (...)

Em um função variada, posso passar zero valores (ter cuidado pra não dar erro)

- Cap. 12 – Funções – 3. Defer

Esse statement pode ser explicado como dar um instrução que vai ser executada na ultima hora

Exemplo

```go
package main

import "fmt"

func main() {
    defer fmt.Println("com defer, vem primeiro")
    fmt.Println("sem defer, veio depois ")
}
```

Resultado

```log
sem defer, veio depois
com defer, vem primeiro

Program exited.
```

Basicamente, é deixar pra última hora. Tendo mais de um defer, acontece o seguinte comportamento, LIFO. O primeiro, vai ser executado por ultimo.

É bom para organizar código. Exemplo, podemos usar para fechar arquivos ou conexões ao utiliza-los.

Ele é executado no fechamento do code block ou antes do return

- Cap. 12 – Funções – 4. Métodos

Método em go, é um função anexada a um tipo.

```go
package main

import "fmt"

type pessoa struct {
    nome  string
    idade int
}

func (p pessoa) oibomdia() {
    fmt.Println(p.nome, "diz bom dia!")
}

func main() {
    mauricio := pessoa{"Mauricio", 30}
    mauricio.oibomdia()
}
```

Basicamente a gente liga uma struct a uma função, com o receiver, que recebe uma struct.

O método é uma função especifica para um valor de um tipo, no caso aqui, pra cada valor de pessoa, cada pessoa terá seu método bom dia.

- Cap. 12 – Funções – 5. Interfaces & polimorfismo

Em Go, valores podem ter mais de um tipo

Uma interface permite que um valor tenha mais que um tipo

Uma interface é um conjunto de métodos, e todo typo que eu criar e ter esses métodos, vai implementar e ganhar o tipo dessa interface

Polimorfismo (várias formas), tenho a mesma ação pra coisas de tipos distintas. Meio que, tenho uma interface fazerBarulho, e dois tipos que implementam isso, cachorro e gato. Se eu chamar fazerBarulho pra um, ele faz uma coisa, pra outro, faz outra coisa, mas eu chamei fazerBarulho. Então tá ai as várias formas do polimorfismo.

Declaração é ```type nome interface{}```, e dentro disso deve ter métodos.

Exemplos

```go
package main

import "fmt"

type pessoa struct {
    nome      string
    sobrenome string
    idade     int
}

type gente interface {
    oibomdia()
}

type dentista struct {
    pessoa
    dentesarrancados int
    salário          float64
}

type arquiteto struct {
    pessoa
    tipodeconstrução string
    tamanhodaloucura string
}

func (x dentista) oibomdia() {
    fmt.Println("Meu nome é", x.nome, "e ouve só: Bom dia!")
}

func (x arquiteto) oibomdia() {
    fmt.Println("Meu nome é", x.nome, "e ouve só: Bom dia!")
}

func serhumano(g gente) {
    g.oibomdia()
    switch g.(type) {
    case dentista:
        fmt.Println("Eu ganho:", g.(dentista).salário)

    case arquiteto:
        fmt.Println("Eu construo:", g.(arquiteto).tipodeconstrução)
    }
}

func main() {
    dentista1 := dentista{pessoa{"José", "Oliveira", 24}, 24, 5000.0}

    arquiteto1 := arquiteto{pessoa{"Geraldo", "Costa", 24}, "urbana", "bem grande"}

    serhumano(dentista1)
    serhumano(arquiteto1)
}
```

Então saca só, os dois typos, dentista e arquiteto implementam a interface gente, logo, um método que receber gente, também aceita o dentista ou o arquiteto. As vezes então, sem saber você pode estar implementando uma interface.

Reparar que posso ter até mesmo um switch dentro do método que recebe gente, pra verificar que tipo de gente é que eu to passando no parâmetro da chamada da função.

Onde se utiliza?

Área de formas geométricas (gobyexample.com)

Sort

DB

Writer interface: arquivos locais, http request/response

- Cap. 12 – Funções – 6. Funções anônimas

Funções que são declaradas dentro de função que rodam imediatamente após a sua definição.

Funcionam que nem no JavaScript

```go
package main

import "fmt"

func main() {
    x := 387
    func(x int) {
        fmt.Println(x * 12333)
    }(x)
}
```

Geralmente não tem nome, pois não precisarão ser reutilizadas. Também serve para go rotine, aparentemente em concorrência.

- Cap. 12 – Funções – 7. Func como expressão

basicamente, é uma função dentro de uma função, que pode ser associada a uma variável

```go
package main

import "fmt"

func main() {
    x := 387
    y := func(x int) {
        fmt.Println(x * 12333)
    }

    y(x)
}
```

- Cap. 12 – Funções – 8. Retornando uma função

```go
package main

import "fmt"

func main() {
    x := 387
    y := retornaFuncao()

    fmt.Println(y(x))
}

func retornaFuncao() func(int) int {
    return func(x int) int {
        return x * 12333
    }
}

```

- Cap. 12 – Funções – 9. Callback

Callback é passar uma função como argumento, essa função ẽ usada dentro de outra função

```go
package main

import "fmt"

func main() {
    t := somentePares(soma, []int{50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60}...)
    fmt.Println(t)
}

func soma(x ...int) int {
    n := 0
    for _, v := range x {
        n += v
    }
    return n
}

func somentePares(f func(x ...int) int, y ...int) int {
    var slice []int
    for _, v := range y {
        if v%2 == 0 {
            slice = append(slice, v)
        }
    }
    total := f(slice...)
    return total
}

```

No exemplo anterior, eu poderia passar para a função somente pares, uma  função que verifica a média, ou subtrai um do outro.

- Cap. 12 – Funções – 10. Closure

Closure é quando capturamos o escopo é quando capturamos um escopo(contexto) para utiliza-lo quando quisermos.

```go
package main

import "fmt"

func main() {
    a := i()
    b := i()
    fmt.Printf("a is %d\n", a())
    fmt.Printf("b is %d\n", b())
    fmt.Printf("a is %d\n", a())
}

func i() func() int {
    x := 0
    return func() int {
        x++
        return x
    }
}
```

Resultado

```log
a is 1
b is 1
a is 2

Program exited.

```

Nesse exemplo, podemos verificar que o resultado da chamada da função que a função i retorna, é 1 para a e para b. Nesse caso, o contexto foi mantido para a função quando foi chamada.

Veja que o valor da função exterior fica armazenado, e toda vez que eu chamar a função retornada por i, a função interna tem acesso a aquela variável externa.

Basicamente, quando uma função interna consegue acessar o valor de uma variável de uma função externa e ler ou alterar o seu dado, é uma closure.

```text
Uma closure ocorre normalmente quando uma função é declarada dentro do corpo de outra, e a função interior referencia variáveis locais da função exterior. Em tempo de execução, quando a função exterior é executada, então uma closure é formada, que consiste do código da função interior e referências para quaisquer variáveis no escopo da função exterior que a closure necessita.
```

- Cap. 12 – Funções – 11. Recursividade

Recursividade é quando algo chama a si mesmo. Muito usado na matemática e na ciência da computação. No escopo do Go, recursividade vai ser uma função chamando ela mesma. Deve ter um statement que com uma regra para sair dessa função, se não rola recursão infinita, dando um stack overflow. Ou também podemos ter recursão de typo, onde um typo tem dentro dele mesmo o próprio tipo.

Uma coisa que tem por definição a própria coisa.

Exemplo no código, o mais clássico

```go

package main

import "fmt"

func main() {
    fmt.Println(fibonacci(5))
}

func fibonacci(a int) int {
    if a == 0 || a == 1 {
        return 1
    }
    return fibonacci(a-1) + fibonacci(a-2)
}
```

outro exemplo

```go
package main

import "fmt"

func main() {
    fmt.Println(fatorial(5))
}

func fatorial(x int) int {
    if x == 1 || x == 0 {
        return 1
    }

    return x * fatorial(x-1)
}
```

- Cap. 14 – Ponteiros – 1. O que são ponteiros?

Todos os valores que armazenamos no computador, vão para a memoria ram do computador, e cada byte da sua memoria ram tem um endereço, esse endereço pode ser acessado, para colocar valor lá dentro, ou até mesmo pegar esse endereço.

Um ponteiro, em C pelo menos, é literalmente uma variável que contém o endereço de memoria de um int, ou um char, ou o char do começo de uma string.

Para mostrar o endereço de memoria de uma variável:

```go
package main

import "fmt"

func main() {
    x := 10
    fmt.Println(&x)
}
```

É usado a notação & (E comercial) antes da variável.

Eu posso colocar o endereço de uma variável em um outra variável, dessa forma:

```go
package main

import "fmt"

func main() {
    x := 10

    y := &x
    fmt.Println(&x)
    fmt.Println(y)
}
```

Resultando nisso (O endereço varia conforme rodamos o programa)

```log
0xc000018030
0xc000018030

Program exited.
```

Outra coisa que podemos fazer é o dereference de um endereço, ou seja, ir até o endereço e buscar o valor contido lá.

```go
package main

import "fmt"

func main() {
    x := 10

    y := &x
    fmt.Println(*y)
}
```

Resultado

```log
10

Program exited.
```

O tipo de uma variável que é contém um endereço é *Type, sou seja, um ponteiro do tipo. Exemplo:

```go
package main

import "fmt"

func main() {
    x := 10

    y := &x
    fmt.Printf("%T", y)
}
```

Resultado

```log
*int
Program exited.
```

Nesse caso, é um *int, ou seja, ponteiro de int.

- Cap. 14 – Ponteiros – 2. Quando usar ponteiros

Ponteiros servem para duas coisas (em go, aparentemente)

A primeira é pra lidar com quantidades grandes de dados e eu não quero ficar copiando pra lá e pra cá, eu deixo o item em um lugar da memoria, e quem tiver que mexer com esse valor vai lá no endereço e mexe com ele lá.

Em go, tudo é pass by value, então, se eu quiser um função que mude o dado original, vou ter que utilizar ponteiro.

Exemplo

```go

package main

import "fmt"

func main() {
    x := 0
    estarecebeovalor(x)
    fmt.Println("valor de x é: ", x)

    estarecebeumponteiro(&x)
    fmt.Println("valor de x é: ", x)

}

func estarecebeovalor(x int) {
    x++
}

func estarecebeumponteiro(x *int) {
    *x++
}

```

Veja que a primeira função recebe um valor x, e incrementa, porém, ele só incrementa no scopo da função, o x que ele recebe é uma cópia do x do método main (pass by value)

Na outra função, eu recebo um endereço de memória de um int, ou seja, a referencia a memoria. Quando uso o operador de dereference e incremento, eu estou incrementando com 1 a mesma varável da função main, pois, o endereço de memoria daquela variável foi passado.

Passeio rápido

Para receber um endereço de memoria na função, é usado: ```func algo (x *type)```.

Para pegar um endereço da memoria de uma variável: ```y := &x```

Tendo um endereço de memoria e quero ir até ele: ```*y```

- Cap. 16 – Aplicações – 1. Documentação JSON

```go
package main

import (
    "encoding/json"
    "fmt"
    "os"
)

func main() {
    type ColorGroup struct {
        ID     int
        Name   string
        Colors []string
    }
    group := ColorGroup{
        ID:     1,
        Name:   "Reds",
        Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
    }
    b, err := json.Marshal(group)
    if err != nil {
        fmt.Println("error:", err)
    }
    os.Stdout.Write(b)
}
```

Pra função Marshal do pacote enconding/json funcionar, as iniciais dos campos DEVEM ser maiúsculas

Outra coisa interessante, é que tem sites que traduzem json para uma struct em go.

[exemplo](https://mholt.github.io/json-to-go)

- Cap. 16 – Aplicações – 3. JSON unmarshal (desordenação)

Para utilizar o unmarshal, temos que usar a função unmarshal com uma cadeia de bytes que tenha o json e um ponteiro pra uma struct que contenha o formato desejado.

Exemplo

```go
package main

import (
    "encoding/json"
    "fmt"
)

func main() {
    type ColorGroup struct {
        ID     int
        Name   string
        Colors []string
    }
    group := ColorGroup{
        ID:     1,
        Name:   "Reds",
        Colors: []string{"Crimson", "Red", "Ruby", "Maroon"},
    }
    b, err := json.Marshal(group)
    if err != nil {
        fmt.Println("error:", err)
    }
    fmt.Println(string(b))
    fmt.Printf("%T\n", b)
    var group2 ColorGroup
    err = json.Unmarshal(b, &group2)
    if err != nil {
        fmt.Println(err)
    }

    fmt.Println(group2)
}
```

Resultado

```log
{"ID":1,"Name":"Reds","Colors":["Crimson","Red","Ruby","Maroon"]}
{1 Reds [Crimson Red Ruby Maroon]}
Program exited.
```

Para realizar o unmarshal, precisamos dos campos do json iguais a na struct, logo, pode-se usar o json to go.

Tem diferença do encoder/decoder para marshal/unmarshal.

Para usar o Encoder, você cria um objeto encoder com o método NewEncoder passando um Writer. A partir dai, será retornado um encoder*, que tem métodos para receber uma interface e escrever um json no writer utilizado na criação do encoder.

A diferença de acordo com a teacher é, o marshal joga o resultado em uma variável, o encoder joga direto em uma interface de escrita de algo.

Cap. 16 - Aplicações - 4. A interface Writer

Pra implementar a interface writer é uma interface que tem que implementar o seguinte método.

```go
func Write(p []byte) (n int, err error);
```

Tanto um encode quanto um Fprinln aceitam um writer.

O importante aqui é entender que eu posso passar um writer, ou seja, uma struct que implementa a interface write, para várias funções dentro da go lang, inclusive usando rede, como no package http, package dial net.

Cap. 16 – Aplicações – 5. O pacote sort

Ordenando elementos de slices.

Aqui iremos criar nosso tipo, implementar a interface sort.interface e iremos ordenar nosso slice do nosso elemento.

A interface Sort.Interface determina que temos que implementar os métodos Less, Swap e Len, e ordenara com um O(n log n)

Exemplos com funções pré prontas

```go
package main

import (
    "fmt"
    "sort"
)

func main() {
    ss := []string{"atirei", "o", "pau", "no", "gato"}
    sort.Strings(ss)
    fmt.Println(ss)
}
```

Resultado

```log
[atirei gato no o pau]

Program exited.
```

Também existe a função sort.Ints, sort.Float64s, faz a mesma coisa

- Cap. 16 – Aplicações – 6. Customizando o sort

Para ordenar um slice de forma customizável, pode-se usar duas formas a seguir:

```go
package main

import (
    "fmt"
    "sort"
)

type Pessoa struct {
    Nome  string
    Idade int
}

type PorIdade []Pessoa

func (p PorIdade) Len() int {
    return len(p)
}

func (p PorIdade) Swap(a, b int) {
    p[a], p[b] = p[b], p[a]
}

func (p PorIdade) Less(a, b int) bool {
    return p[a].Idade < p[b].Idade
}

func main() {
    sp := []Pessoa{
        Pessoa{"Geraldo", 24},
        Pessoa{"Stephane", 21},
        Pessoa{"Ale", 26},
    }
    fmt.Println(sp)
    sort.Sort(PorIdade(sp))

    fmt.Println(sp)

    sort.Slice(sp, func(i, j int) bool {
        return sp[i].Idade > sp[j].Idade
    })
    fmt.Println(sp)
}

```

Resultado

```log
[{Geraldo 24} {Stephane 21} {Ale 26}]
[{Stephane 21} {Geraldo 24} {Ale 26}]
[{Ale 26} {Geraldo 24} {Stephane 21}]

Program exited.
```

Em uma forma, a gente só uso um método do package sort chamado slice, passando o array e uma Less function de forma anonima.

A outra forma é implementando a interface sort.Interface no nosso tipo de dados que é um subtipo de um slice do que a gente quer.

Cap. 16 – Aplicações – 7. bcrypt

É uma função de hashing de senha.

Forma de usar no Go

```go
package main

import (
    "fmt"
    "golang.org/x/crypto/bcrypt"
)

func main() {

    senha := "20julho1980"
    // usando o bcrypt para gerar um hash de um pass
    hash, error := bcrypt.GenerateFromPassword([]byte(senha), 14)

    // comparando hash com um slice de bytes, se der erro tá errado
    err := bcrypt.CompareHashAndPassword(hash, []byte(senha))

    // printando resultados para não dar erro
    fmt.Println(error, string(hash))
    fmt.Println(err)
}

```

- Cap. 18 – Concorrência – 2. Goroutines & WaitGroups

Para fazer duas funções sejam executadas de maneira concorrente com o go, é usado as go routines, que são "threads"

Tem o mesmo conceito de threads, mas não são threads

Threads são linhas de execução, é quando um processo pode se dividir em dois e ser executado  concorrencialmente

O suporte a thread é fornecido pelo OS.

Na pratica, iremos utilizar um ```go func```

Quando colocamos um go na frente de uma função, essa função será usada independentemente

Depois que eu dei um go nessa rotina, não tenho como mexer nela.

```go
package main

import "fmt"

func main() {
    go func1()
    func2()
}

func func1() {

    for i := 0; i < 10; i++ {
        fmt.Println("func1", i)
    }
}

func func2() {

    for i := 0; i < 10; i++ {
        fmt.Println("func2", i)
    }
}

```

Nesse exemplo, a func2 printa, e a outra não, pq o código termina antes da func1 executar

Ou seja, nós precisamos esperar ela executar pra depois encerrar o programa

Forma errada de fazer isso: ```sync.WaitGroup```

Esse WaitGroup usa as funções Add pra adicionar go routines, done é utilizada dentro da go routine falando que terminou de executar.

E também usa a função wait, pra falar pro programa esperar pq ta realizando as go rotines.

Vai ficar bem assim o código

```go
package main

import (
    "fmt"
    "runtime"
    "sync"
    "time"
)

var wg sync.WaitGroup

func main() {

    fmt.Println(runtime.NumCPU())
    fmt.Println(runtime.NumGoroutine())

    wg.Add(2)

    go func1()
    go func2()

    fmt.Println(runtime.NumGoroutine())

    wg.Wait()

}

func func1() {
    for i := 0; i < 1000; i++ {
        fmt.Println("func1:", i)
    }
    wg.Done()
}

func func2() {
    for i := 0; i < 1000; i++ {
        fmt.Println("func2:", i)
    }
    wg.Done()
}

```

Aqui dá pra ver que estão rodando de forma concorrente.

- Cap. 18 – Concorrência – 3. Discussão: Condição de corrida

O que acontece quando fazemos as coisas de uma maneira ruim?

Programar de maneira concorrente é difícil por causa do acesso a variáveis compartilhadas

Em Go a gente compartilha os valores passados através de canais, e esses valores nunca são compartilhados.

Apenas uma goroutine tem acesso ao valor em um dado momento.

Utilizando canais, é impossível ter condição de corrido (documentação que disse)

O que é yield? Yield é a troca de execução de programa que o processador faz internamente. `runtime.Gosched()`

Race condition:
        *Função 1       var     Função 2*
         Lendo: 0   →   0
         Yield          0   →   Lendo: 0
         var++: 1               Yield
         Grava: 1   →   1       var++: 1
                        1   ←   Grava: 1
         Lendo: 1   ←   1
         Yield          1   →   Lendo: 1
         var++: 2               Yield
         Grava: 2   →   2       var++: 2
                        2   ←   Grava: 2

E é por isso que vamos ver mutex, atomic e, por fim, channels.

Mutex é uma trava de exclusão, igual o synchronized do java.

Atomic é bem baixo nível, não será utilizado muito.

- Cap. 18 – Concorrência – 4. Na prática: Condição de corrida

```go
package main

import (
    "fmt"
    "runtime"
    "sync"
)

func main() {
    fmt.Println("CPUs:", runtime.NumCPU())
    fmt.Println("Goroutines:", runtime.NumGoroutine())

    var wg sync.WaitGroup
    contador := 0
    totalGoRoutine := 1000

    wg.Add(totalGoRoutine)

    for i := 0; i < totalGoRoutine; i++ {

        go func() {
            v := contador
            runtime.Gosched()
            v++
            contador = v
            wg.Done()
        }()
    }

    wg.Wait()
    fmt.Println("Goroutines:", runtime.NumGoroutine())
    fmt.Println(contador)
}

```

Ao rodar isso, como a variável está sendo compartilhada, varias routines leem e incrementam a partir do valor base. Não vai ser nada confiável utilizar isso.

O go tem uma ferramenta para analisar se o código tem condição de corrida.

`go run -race main.go`

Cap. 18 – Concorrência – 5. Mutex

O Mutex é um lock de código, que só deixa uma thread acessar um valor compartilhado por vez. (o problema com isso não é o race condition, e sim deadlock se não for bem desenvolvido )

[Exemplo e explicação do Mutex](https://gobyexample.com/mutexes)

Como isso funciona na pratica? Segue o exemplo

```go
package main

import (
    "fmt"
    "runtime"
    "sync"
)

func main() {
    fmt.Println("CPUs:", runtime.NumCPU())
    fmt.Println("Goroutines:", runtime.NumGoroutine())

    var wg sync.WaitGroup
    var mu sync.Mutex

    contador := 0
    totalGoRoutine := 1000

    wg.Add(totalGoRoutine)

    for i := 0; i < totalGoRoutine; i++ {

        go func() {
            mu.Lock()
            v := contador
            runtime.Gosched()
            v++
            contador = v
            wg.Done()
            mu.Unlock()
        }()
    }

    wg.Wait()
    fmt.Println("Goroutines:", runtime.NumGoroutine())
    fmt.Println(contador)
}

```

Em go, um trecho de código é lockado usando os métodos Lock e Unlock de um objeto Mutex.

- Cap. 18 – Concorrência – 6. Atomic

É como se fosse o AtomicInteger e suas variáveis no Java, uma variável preparada para ser thread safe, que já realiza o lock e unlock por baixo dos panos.

Exemplo

```go
package main

import (
    "fmt"
    "runtime"
    "sync"
    "sync/atomic"
)

func main() {
    fmt.Println("CPUs:", runtime.NumCPU())
    fmt.Println("Goroutines:", runtime.NumGoroutine())

    var wg sync.WaitGroup

    var contador int64 = 0
    totalGoRoutine := 1000

    wg.Add(totalGoRoutine)

    for i := 0; i < totalGoRoutine; i++ {

        go func() {
            atomic.AddInt64(&contador, 1)
            fmt.Println("Contador:\t", atomic.LoadInt64(&contador))
            wg.Done()
        }()
    }

    wg.Wait()
    fmt.Println("Goroutines:", runtime.NumGoroutine())
    fmt.Println(contador)
}

```

Diferente do Java, aqui temos ponteiros, logo, quando usamo o atomic com algum tipo primitivo, declaramos realmente uma variável e usamos o ponteiro para ela, e dentro lá do pacote ele vai cuidar pra mim dos detalhes de mutex.

- Cap. 21 – Canais – 1. Entendendo canais

Conceito único da linguagem go.

Canais é como é chamada a maneira de transmitir dados entre go routines.

Pegamos várias go funcs e transmitimos informação entre elas.

Servem para coordenar, sincronizar, orquestrar e buffering.

Na prática:

```go
package main

import "fmt"

func main() {
    canal := make(chan int)
    canal <- 42
    fmt.Println(<-canal)
}

```

Isso dá erro, mas tem um exemplo de como fazemos um canal, que é com a expressão `make(chan <type>)`, onde o type é o tipo que desejamos, exemplo, int, Pessoa, bool etc.

**Canais bloqueiam**, canais são como corredores em uma corrida de revesamento, eles devem passar o bastão de maneira sincronizada.

Dá problema se o corredor tentar passar o bastão para o próximo, e esse próximo não estar lá para receber o bastão.

Ou se um corredor ficar esperando o bastão e não tiver ninguém para entregar, também dá problema.

O problema do código acima é que, a função main (que também é uma go routine) bloqueio a execução esperando alguma outra go routine buscar o valor do canal. Como ninguém buscou, ocorre o seguinte erro

```log
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
    /tmp/sandbox309324162/prog.go:9 +0x37

Program exited.
```

**ATENÇÃO**, colocar informação em um canal e retirar informação de um canal deve ser feito de maneira concorrente. Preciso que uma go routine entregue informação pra outra go routine através de um canal, nunca através de uma mesma.

```go
package main

import "fmt"

func main() {
    canal := make(chan int)

    go func() {
        canal <- 42
    }()

    fmt.Println(<-canal)

}
```

Nesse exemplo, a go routine (func anonima), coloca uma informação no canal, e o go routine main pega a informação do canal. Ou seja, CORRECTO!!!!

Agora, com buffer. Com ele, eu posso fazer que um canal não bloqueia imediatamente, eu posso colocar um valor em um canal e outra go routine pode retirar depois.

```go
package main

import "fmt"

func main() {
    canal := make(chan int, 1)
    canal <- 42
    fmt.Println(<-canal)

}

```

Na func make, onde está o 1, é onde declaro o tamanho do buffer. esse exemplo funciona.

Mas **ATENÇÃO** buffers não bloqueiam imediatamente, mas se colocarmos mais valores no canal do que o que eles aguentam, eles bloqueiam sim.

```go
package main

import "fmt"

func main() {
    canal := make(chan int, 1)
    canal <- 42
    canal <- 43
    fmt.Println(<-canal)

}

```

```log
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan send]:
main.main()
    /tmp/sandbox388642326/prog.go:8 +0x4b

Program exited.
```

Via de regra, nós não utilizamos buffer. **Não recomendado pra uso**.

- Cap. 21 – Canais – 2. Canais direcionais & Utilizando canais

Canais podem ser direcionais

Podemos ter canais que só recebem informação e outros canais que só enviam informação.

Isso permite que os type-checking mechanisms do compilador façam com que não seja possível, por exemplo, escrever em um canal de leitura.

Canais normalmente são bidirecionais.

Para declarar um send channel é assim: `chan<-` a seta tá direcionando pro canal, ou seja, enviando para o canal

Para declarar um receive channel, é assim: `<-chan`, a seta está partindo do canal para fora, ou seja, canal envia.

Sim, só muda o lugar da seta.

Podemos declarar assim um canal send ou receive `canal := make(<-chan int)` ou `canal := make(chan<- int)`

Exemplo de uso dos canais direcionais

```go
package main

import (
    "fmt"
)


func main() {
    canal := make(chan int)

    go send(canal)
    receive(canal)
}

func send(s chan<- int) {
    s <- 42
}

func receive(r <-chan int) {
    fmt.Println("O valor recebido do canal foi:", <-r)
}

```

Basicamente, nos argumentos das funções, eu fiz com que um canal bi-direcionais obtesse o comportamento de um direcional

Exemplos:

[geral pra específico:](https://play.golang.org/p/H1uk4YGMBB)
[específico pra específico:](https://play.golang.org/p/8JkOnEi7-a)
[específico pra geral:](https://play.golang.org/p/4sOKuQRHq7)
[atribuição tipos !=:](https://play.golang.org/p/bG7H6l03VQ)

Basicamente, nos exemplo, podemos ver que podemos fazer a conversão de canal geral para especifico.

Não podemos converter um especifico para outro especifico (send para receiver ou vice versa)

Especifico não pode ser convertido para geral.

O geral não consegue receber um especifico (exemplo `channelBD = channelReceive`). Também não é possivel converter um especifico em um geral (exemplo `channel = (chan int)(channelSender)`)

Outro exemplo bom para sempre lembrar:

```go
// Only to receive data
c1:= make(<- chan bool)

// Only to send data
c2:= make(chan<-bool
```

- Cap. 21 – Canais – 3. Range e close

```go
package main

import "fmt"

func main() {
    c := make(chan int)
    go meuloop(100, c)
    for v := range c {
        fmt.Println("Recebido do canal", v)
    }

}

func meuloop(total int, s chan<- int) {
    for i := 0; i < total; i++ {
        s <- i
    }
    close(s)
}
```

Explicando primeiro o range em cima de um canal, basicamente, o range vai ficar eternamente buscando valores colocados no canal e executando algo em cima disso.

Então se liga no passo a passo

1. Função meu loop manda coisas pro canal
2. Range em cima do canal vai ficar buscando tudo que for enviado para lá

Agora vamos para o close. Pq diamos demos o close na func meuloop?

Então meu chapa raul, temos que falar pro canal que ele não vai receber mais nada, se não o range lá na outra função vai ficar eternamente esperando por isso, ou seja, ERRO!

Então aqui o negocio é, mandei todos os valores que eu tenho que enviar, fecho o canal, e o meu range, lá na outra função consegue ficar feliz e contente sem dar errinhos.

Lembrando que, enquanto for mandado instruções no canal, o que for colocado depois do for range não será executado, só quando o canal for fechado, então cuidado ao colocar mais instruções depois daquele for, talves o que tu colocou nunca seja executado.

[Exemplo disso](https://go.dev/play/p/aze7-yDa3BE)

Exemplo de um programa que não fique bloqueado, produzido por mim mesmo antes de terminar o video

```go
package main

import (
    "fmt"
    "sync"
)

var wg sync.WaitGroup

func main() {
    c := make(chan int)
    wg.Add(2)
    go meuloop(100, c)
    go consomeDoCanal(c)
    fmt.Println("Quando será executado isso?")
    wg.Wait()
}

func meuloop(total int, s chan<- int) {
    for i := 0; i < total; i++ {
        s <- i
    }
    close(s)
    wg.Done()
}

func consomeDoCanal(r <-chan int) {
    for v := range r {
        fmt.Println("Recebido do canal", v)
    }
    wg.Done()
}

```

Aqui vemos que como as duas funções `meuloop` e `consomeDoCanal` são concorrentes, o meu fmt é executado tranquilo, mas aqui tenho que usar um waitGroup.

Porém, ainda é possivel fazer de um outro jeito que vai dar uma bloqueada, mas não vai precisar do WaitGroup

```go
package main

import (
    "fmt"
)

func main() {
    c := make(chan int)

    go meuloop(100, c)
    consomeDoCanal(c)
    fmt.Println("Quando será executado isso?")

}

func meuloop(total int, s chan<- int) {
    for i := 0; i < total; i++ {
        s <- i
    }
    close(s)
}

func consomeDoCanal(r <-chan int) {
    for v := range r {
        fmt.Println("Recebido do canal", v)
    }
}

```

Desse jeito funciona sem WaitGroup, mas eu tenho que colocar a minha send channel como go routine, se eu colocar a outra, dá um errinho camarada. xD

Acredito que seja pq, se eu deixar o chan receiver de forma concorrente, ou meuloop entrega mais mensagens do que deveria ou o consome canal tenta buscar mensagens quando ainda não existe nenhuma.

- Cap. 21 – Canais – 4. Select

Select é como um switch para canais, porém não sequencial

Ele bloqueia até receber um valor que bate com algum caso especificado nele, se eu tiver mais de um caso valido, vai ser escolhido um aleatorio.

```go
package main

import "fmt"

func main() {
    numeroDeExecucoes := 500

    canalA := make(chan int)
    canalB := make(chan int)

    go func(x int) {
        for i := 0; i < x; i++ {
            canalA <- i
        }
    }(numeroDeExecucoes / 2)

    go func(x int) {
        for i := 0; i < x; i++ {
            canalB <- i
        }
    }(numeroDeExecucoes / 2)

    for i := 0; i < numeroDeExecucoes; i++ {
        select {
        case v := <-canalA:
            fmt.Println("Canal A", v)
        case v := <-canalB:
            fmt.Println("Canal B", v)
        }
    }
}
```

Como podemos ver nesse trexo de código, o case vai ter um comportamento pra cada canal, e dá pra receber de vários canais em um mesmo local.

Outro exemplo:

```go
package main

import "fmt"

func main() {
    canal := make(chan int)
    quit := make(chan int)
    go enviaParaCanais(canal, quit)
    recebeValoresNosCanais(canal, quit)
}

func recebeValoresNosCanais(canal chan int, quit chan int) {
    for i := 0; i < 50; i++ {
        fmt.Println("Recebido:", <-canal)
    }
    quit <- 0
}

func enviaParaCanais(canal chan int, quit chan int) {
    qualquercoisa := 1
    for {
        select {
        case canal <- qualquercoisa:
            qualquercoisa++
        case <-quit:
            return
        }
    }
}

```

Nesse select do envia para canal, e um case ele envia valores para o canal e no segundo case ele recebe de outro canal.

Caso ele receba um valor em quit, ele encerra a função, concluindo o código.

Segue outro exemplo:

```go
package main

import "fmt"

func main() {
    par := make(chan int)
    impar := make(chan int)
    quit := make(chan bool)
    go send(par, impar, quit)
    receive(par, impar, quit)
}

func send(canalPar chan<- int, canalImpar chan<- int, canalQuit chan<- bool) {
    for i := 1; i < 100; i++ {
        if i%2 == 0 {
            canalPar <- i
            continue
        }
        canalImpar <- i
    }
    close(canalPar)
    close(canalImpar)
    canalQuit <- true

}

func receive(canalPar <-chan int, canalImpar <-chan int, canalQuit <-chan bool) {
    for {
        select {
        case v := <-canalPar:
            fmt.Println("Par:", v)
        case v := <-canalImpar:
            fmt.Println("Impar:", v)
        case <-canalQuit:
            return
        }
    }

}


```

Nesse outro exemplo, é igual uma junção dos dois anteriores. A função send envia valores para os 3 canais, numeros pares pro canalPar, impares para o canalImpar e envia valores para o canalQuit para encerrar a função.

O receive, recebe dos 3 canais e mostra os valores, e quando recebe um valor do canalQuit, ele encerra a função.

Por algum motivo, na última execução, rola uns Par: 0 ou Impar: 0, muito estranho, mas será resolvido na próxima aula.

- Cap. 21 – Canais – 5. A expressão comma ok

Ao receber um valor de um canal, podemos fazer o comma ok, como no seguine exemplo

```go
package main

import "fmt"

func main() {
    canal := make(chan int)
    go func() {
        canal <- 42
        close(canal)
    }()

    v, ok := <- canal

    fmt.Println(v,ok)
}
```

Se eu não tiver nada no canal, eu recebo false e um valor zero

```go
package main

import "fmt"

func main() {
    canal := make(chan int)
    go func() {
        canal <- 42
        close(canal)
    }()

    v, ok := <- canal
    fmt.Println(v,ok)

    v, ok = <- canal
    fmt.Println(v,ok)
}
```

```log
42 true
0 false
```

Logo, tenho que verificar se é true o comma ok.

Com o comma ok, podemos corrigir o exercicio da aula passada

```go
package main

import "fmt"

func main() {
    par := make(chan int)
    impar := make(chan int)
    quit := make(chan bool)
    go send(par, impar, quit)
    receive(par, impar, quit)
}

func send(canalPar chan<- int, canalImpar chan<- int, canalQuit chan<- bool) {
    for i := 1; i < 100; i++ {
        if i%2 == 0 {
            canalPar <- i
            continue
        }
        canalImpar <- i
    }
    close(canalPar)
    close(canalImpar)
    canalQuit <- true

}

func receive(canalPar <-chan int, canalImpar <-chan int, canalQuit <-chan bool) {
    for {
        select {
        case v, ok := <-canalPar:
            if ok {
                fmt.Println("Par:", v)
            }
        case v, ok := <-canalImpar:
            if ok {
                fmt.Println("Impar:", v)
            }
        case <-canalQuit:
            return
        }
    }

}

```

- Cap. 21 – Canais – 6. Convergência

A convergencia é um padrão de programção que envolve pegar dados de vários pontos diferentes e convergir eles em um ponto só.

```go
package main

import (
    "fmt"
    "sync"
)

func main() {
    par := make(chan int)
    impar := make(chan int)
    converge := make(chan int)
    go send(par, impar)
    go receive(par, impar, converge)

    for v := range converge {
        fmt.Println(v)
    }
}

func send(par, impar chan int) {
    x := 100
    for i := 0; i < x; i++ {
        if i%2 == 0 {
            par <- i
            continue
        }
        impar <- i
    }
    close(par)
    close(impar)
}

func receive(par, impar, converge chan int) {

    var wg sync.WaitGroup
    wg.Add(2)

    go func() {
        for v := range par {
            converge <- v
        }
        wg.Done()
    }()

    go func() {
        for v := range impar {
            converge <- v
        }
        wg.Done()
    }()
    wg.Wait()
    close(converge)
}
```

Exemplo de convergencia, junto o resultado de dois canais em um só e utilizo esse canal que junta informações

```go
package main

import (
    "fmt"
    "math/rand"
    "time"
)

func main() {

    canalA := trabalho("A")
    canalB := trabalho("B")

    canalC := converge(canalA, canalB)

    for {
        fmt.Println(<-canalC)
    }
}

func trabalho(s string) chan string {
    canal := make(chan string)
    go func(s string, c chan string) {
        for i := 0; ; i++ {
            c <- fmt.Sprintf("Função %v diz: %v", s, i*rand.Intn(1e3))
            time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
        }
    }(s, canal)
    return canal
}

func converge(x, y chan string) chan string {
    novo := make(chan string)
    go func() {
        for {
            novo <- <-x
        }
    }()
    go func() {
        for {
            novo <- <-y
        }
    }()
    return novo
}
```

- Cap. 21 – Canais – 7. Divergência

Divergência é padrão de projeto concorrente que pega dados de um canal e enviar para mais de um canal(Ou go funcs pelo que é explicado na aula).

```go
package main

import (
    "fmt"
    "math/rand"
    "sync"
    "time"
)

const numeroCanais = 200

func main() {

    canal1 := make(chan int)
    canal2 := make(chan int)

    go manda(numeroCanais, canal1)
    go outra(canal1, canal2)

    for v := range canal2 {
        fmt.Println(v)
    }
}

func manda(n int, canal chan int) {
    for i := 0; i < n; i++ {
        canal <- i
    }
    close(canal)
}

func outra(canal1 chan int, canal2 chan int) {
    var wg sync.WaitGroup
    wg.Add(numeroCanais)
    for v := range canal1 {
        go func(x int) {
            canal2 <- trabalho(x)
            wg.Done()
        }(v)
    }
    wg.Wait()
    close(canal2)
}

func trabalho(x int) int {
    time.Sleep(time.Millisecond * 1000)
    return x
}

```

Exemplo acima une convergencia e divergencia

Um stream(channel) vira centenas de go funcs que depois convergem.

```go
package main

import (
    "fmt"
    "sync"
    "time"
)

const numeroCanais = 200

func main() {

    canal1 := make(chan int)
    canal2 := make(chan int)
    funcoes := 5

    go manda(numeroCanais, canal1)
    go outra(funcoes, canal1, canal2)

    for v := range canal2 {
        fmt.Println(v)
    }
}

func manda(n int, canal chan int) {
    for i := 0; i < n; i++ {
        canal <- i
    }
    close(canal)
}

func outra(funcoes int, canal1 chan int, canal2 chan int) {
    var wg sync.WaitGroup
    wg.Add(funcoes)
    for i := 0; i < funcoes; i++ {
        go func() {
            for v := range canal1 {
                canal2 <- trabalho(v)
            }
            wg.Done()
        }()
    }
    wg.Wait()
    close(canal2)

}

func trabalho(x int) int {
    time.Sleep(time.Millisecond * 1000)
    return x
}

```

Nesse exemplo, eu não quero uma go func pra cada item que sai do meu canal, e sim 10 funções que trabalham de maneira concorrente.

Comparando o primeiro exemplo com o segundo, enquanto o segundo abre várias go funcs quase que ao mesmo tempo, no segundo limitamos quantas go funcs iremos abrir, tendo blocos de trabalho sendo feitos, no caso, vai demorar mais tempo, mas vai usar menos memoria.

Cap. 21 – Canais – 8. Context

Context é uma feature nova da linguagem.

Pra que serve? Imagina que temos um site, e meu consumidor digitou uma procura. O meu banco de dados é fragmentado. Quando o usuário deu search, eu disparo várias go routines, uma pra cada tabela e a medida que as go routines completam a execução, eu obtenho os resultados.

Mas imagina que o usuário clicou no primeiro resultado, e não precisamos mais das outras go routines. Seria legal eu poder matar as go routines. Eu posso matar elas e economizar memoria e processamento.

Esse é um dos motivos da existencia do package context, que permite a troca de mensagens entre go routines.

[Link de um post do blog do go explicando mais sobre context](https://go.dev/blog/context)

[Documentação](https://pkg.go.dev/context)

Exemplos (Todd):

- Analisando:
  - Background: <https://play.golang.org/p/cByXyrxXUf>
    - WithCancel: <https://play.golang.org/p/XOknf0aSpx>
    - Função Cancel: <https://play.golang.org/p/UzQxxhn_fm>
  - Exemplos práticos:
    - func WithCancel: <https://play.golang.org/p/Lmbyn7bO7e>
    - func WithCancel: <https://play.golang.org/p/wvGmvMzIMW>
    - func WithDeadline: <https://play.golang.org/p/Q6mVdQqYTt>
    - func WithTimeout: <https://play.golang.org/p/OuES9sP_yX>
    - func WithValue: <https://play.golang.org/p/8JDCGk1K4P>

- Destaques:
  - ctx := context.Background
  - ctx, cancel = context.WithCancel(context.Background)
  - goroutine: select case ←ctx.Done(): return; default: continua trabalhando.
  - check ctx.Err();
  - Tambem tem WithDeadline/Timeout