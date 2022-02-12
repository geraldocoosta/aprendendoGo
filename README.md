Em GO

:= is a short variable declaration, basicamente faz a operação de declaração e atribuição


Para criar váriaveis com package scope, não podemos usar o short operator, sobrando assim o var ou o const.

package main

import "fmt"

var teste = "uma string qualquer"

func main() {
    teste = "outra string"
    fmt.Println(teste)    
} 


Ela pode ser mudada pelas funções, qualquer uma no mesmo package (não sei se consigo mudar em packages diferentes)


O short variable declaration pode ser usado se ao menos uma variavel na mesma linha seja declarada, apesar de ser muito feio.
ex:

func main() {
    x := 10
    fmt.Println(x)
    x, z := 11, "teste"

}


No Go, boolean == bool

statement -> uma ou mais expressões
expressão -> qualquer coisa que produz um resultado


variaveis dentro de um code block sao restritas ao codeblock ORIGINAAAAL
variaveis podem ser declaradas no escopo de package (não me parece boa ideia a não ser que sejam constantes)

o var funciona em qualquer lugar.


Boas praticas do GO, sempre utilizar :=, a menos que não de, ai usa outras paradas