package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	exibeIntroducao()

	for {
		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciaMonitoramento()
		case 2:
			fmt.Println("Exibindo Logs...")
		case 0:
			fmt.Println("Saindo do programa...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço este comando")
			os.Exit(-1)
		}
	}
}

func iniciaMonitoramento() {
	fmt.Println("Monitorando...")
	sites := []string{"https://random-status-code.herokuapp.com/", "https://www.alura.com.br"}

	for i := 0; i < len(sites); i++ {

		resp, _ := http.Get(sites[i])

		if resp.StatusCode == 200 {
			fmt.Println("Site: ", sites[i], " carregado com sucesso!")
		} else {
			fmt.Println("Site: ", sites[i], " esta com problemas. Status code ", resp.StatusCode)
		}

	}
}

func exibeIntroducao() {
	nome := "Douglas"
	versao := 1.1
	fmt.Println("Olá, sr(a).", nome)
	fmt.Println("Este programa está na versão", versao)
}

func exibeMenu() {
	fmt.Println("1- Iniciar Monitoramento")
	fmt.Println("2- Exibir Logs")
	fmt.Println("0- Sair do Programa")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)
	fmt.Println("O valor da variável comando é:", comandoLido)
	return comandoLido
}
