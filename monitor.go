package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

const cicloMonitoramentos = 5
const delaySegundos = 5
const sucessoMarcador = "[SUCCESS] - "
const erroMarcador = "[ERRO] - "

func main() {

	exibeIntroducao()

	for {
		exibeMenu()
		comando := leComando()

		switch comando {
		case 1:
			iniciarMonitoramento()
		case 2:
			imprimeLogs()
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
	sites := leSitesDoArquivo()

	for i := 0; i < cicloMonitoramentos; i++ {
		for _, site := range sites {
			testaSite(site)
		}
		fmt.Println("----------------------------------------")
		time.Sleep(delaySegundos * time.Second)
	}

}

func testaSite(site string) {
	resp, err := http.Get(site) //-- underscore ignora variavel de retorno
	trataErro(err)

	if resp.StatusCode == 200 {
		fmt.Println("Site", site, "acessado com sucesso!!!")
		registraLog(site, true, resp.StatusCode)
	} else {
		fmt.Println("Site", site, "apresentou algo estranho, Status Code:", resp.StatusCode)
		registraLog(site, false, resp.StatusCode)
	}
}

func leSitesDoArquivo() []string {
	var sites []string

	arquivo, errOpen := os.Open("sites.txt")
	trataErro(errOpen)

	reader := bufio.NewReader(arquivo)

	for {
		linha, err := reader.ReadString('\n')
		linha = strings.TrimSpace(linha)
		sites = append(sites, linha)
		if err == io.EOF {
			break
		}
	}

	arquivo.Close()

	return sites
}

func registraLog(site string, flagStatus bool, statusCode int) {
	arquivo, err := os.OpenFile("log.txt", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	trataErro(err)

	var status string
	var linha string

	if flagStatus {
		status = sucessoMarcador
		linha = status + time.Now().Format("02/01/06 15:04:05") + " - " + site + "\n"
	} else {
		status = erroMarcador
		linha = status + time.Now().Format("02/01/06 15:04:05") + " - " + site + " Status Code:" + strconv.Itoa(statusCode) + "\n"
	}

	arquivo.WriteString(linha)

	arquivo.Close()
}

func imprimeLogs() {
	fmt.Println("Logando...")

	arquivo, err := ioutil.ReadFile("log.txt") //-- ioutil.ReadFile ja fecha o arquivo sozinho
	trataErro(err)

	fmt.Println(string(arquivo))
	fmt.Println("")

}

func trataErro(err error) {
	if err != nil {
		fmt.Println("[ERRO]:", err)
		os.Exit(-1)
	}
}
