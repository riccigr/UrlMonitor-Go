package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

const MONITORAMENTOS = 5
const DELAY = 10

func main() {

	exibeIntroducao()

	for {
		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			fmt.Println("Logando...")
		case 0:
			fmt.Println("Saindo...")
			os.Exit(0)
		default:
			fmt.Println("Não conheço esse comando.")
		}
	}

}

func exibeIntroducao() {
	nome := "Frodo"
	versao := 1.1
	fmt.Println("Olá, mr.", nome)
	fmt.Println("Essa é a versao:", versao)
	fmt.Println("")
}

func exibeMenu() {
	fmt.Println("1 - Iniciar Monitoramento")
	fmt.Println("2 - Exibir Logs")
	fmt.Println("0 - Sair da Aplicação")
	fmt.Println("")
}

func leComando() int {
	var comandoLido int
	fmt.Scan(&comandoLido) //-- & remete ao endereco da variavel na memoria

	return comandoLido
}

func iniciarMonitoramento() {
	fmt.Println("Monitorando...")
	sites := []string{"https://random-status-code.herokuapp.com/", "http://www.uol.com.br", "http://www.facebook.com"}
	sites = append(sites, "http://www.instagram.com") //-- usando append para ver como funciona

	for i := 0; i < MONITORAMENTOS; i++ {
		for _, site := range sites {
			testaSite(site)
		}
		fmt.Println("----------------------------------------")
		time.Sleep(DELAY * time.Second)
	}

}

func testaSite(site string) {
	resp, _ := http.Get(site) //-- underscore ignora variavel de retorno

	if resp.StatusCode == 200 {
		fmt.Println("Site", site, "acessado com sucesso!!!")
	} else {
		fmt.Println("Site", site, "apresentou algo estranho, Status Code:", resp.StatusCode)
	}
}
