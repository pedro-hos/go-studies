package main

import (
	"fmt"
)

func main() {

	contaPedro := ContaCorrente{titular: "Pedro", numeroAgencia: 1234, numeroConta: 1234567, saldo: 100.3}
	fmt.Println(contaPedro.Sacar(100.))

}
