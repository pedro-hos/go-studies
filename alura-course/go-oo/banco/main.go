package main

import (
	"fmt"

	"github.com/pedro-hos/go-studies/tree/main/alura-course/go_oo/banco/contas"
)

func main() {

	contaPedro := contas.ContaCorrente{
		Titular:       "Pedro",
		NumeroAgencia: 1234,
		NumeroConta:   1234567,
		Saldo:         10000}

	fmt.Println(contaPedro.Saldo)
	fmt.Println(contaPedro.Sacar(100.))
	fmt.Println(contaPedro.Saldo)

}
