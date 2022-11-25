package main

import (
	"fmt"

	"go-studies/alura-course/go-oo/banco/contas"
)

func PagarBoleto(conta verificarConta, valorDoBoleto float64) {
	conta.Sacar(valorDoBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {
	contaDoBruno := contas.ContaPoupanca{}
	contaDoBruno.Depositar(1100)
	fmt.Println(contaDoBruno)
	PagarBoleto(&contaDoBruno, 60)
	fmt.Println(contaDoBruno.ObterSaldo())
}
