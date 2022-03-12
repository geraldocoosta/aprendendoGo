package main

import (
	"fmt"

	"golang.org/x/crypto/bcrypt"
)

func main() {

	senha := "20julho1980"
	// usando o bcrypt para gerar um hash de um pass
	hash, error := bcrypt.GenerateFromPassword([]byte(senha), 10)

	// comparando hash com um slice de bytes, se der erro tá errado
	err := bcrypt.CompareHashAndPassword(hash, []byte(senha))

	// printando resultados para não dar erro
	fmt.Println(error, string(hash))
	fmt.Println(err)
}
