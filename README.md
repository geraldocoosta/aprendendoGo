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
s := "teste é assim mosmo e é 香 😂"


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

o range significa alcance, faixa, extenção, ou seja, ela atravessa toda a extenção do slice

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

O range retorna duas variáveis, pra usar só uma, tem que jogar o valor da outra fora com _

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

Ele funciona da seguinte forma, o primeiro parametro entre colchetes, será o indice de um elemento.

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

Cap. 8 – Agrupamentos de Dados – 8. Slice: a surpresa do array subjacente

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
