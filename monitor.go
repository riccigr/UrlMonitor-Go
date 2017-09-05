package main

import "fmt"

func main() {

	exibeIntroducao()
	exibeMenu()

	comando := leComando()

	switch comando {
	case 1:
		fmt.Println("Monitorando...")
	case 2:
		fmt.Println("Logando...")
	case 0:
		fmt.Println("Saindo...")
	default:
		fmt.Println("Não conheço esse comando.")
	}
}

func exibeIntroducao() {
	nome := "Frodo"
	versao := 1.1
	fmt.Println("Olá, mr.", nome)
	fmt.Println("Essa é a versao:", versao)
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair da Aplicação")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido)

	return comandoLido
}
