package main

import (
	"log"

	"github.com/GuilhermeGGM/projetosGo.git/contas"

	"github.com/GuilhermeGGM/projetosGo.git/clientes"

	"fmt"

	""
)

func PagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {

	testeContaC := contas.ContaCorrente{
		Titular: clientes.Titular{
			Nome:      "",
			CPF:       "02387122062",
			Profissao: "",
		},
		NumeroAgencia: 0,
		NumeroConta:   0,
		Saldo:         1000,
	}

	if testeContaC.Titular.ValidarCPF(testeContaC.Titular.CPF) {

		fmt.Println(testeContaC.Saldo)

		result, err := testeContaC.Sacar(10000)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
		fmt.Println(err)
		fmt.Println(testeContaC.Saldo)
	} else {
		fmt.Println("Erro! CPF informado inválido")
	}
}
