package main

import (
	"log"

	"github.com/GuilhermeGGM/projetosGo.git/contas"

	"github.com/GuilhermeGGM/projetosGo.git/clientes"

	"github.com/GuilhermeGGM/projetosGo.git/utils"

	"fmt"
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
			CPF:       "0238712206",
			Profissao: "",
		},
		NumeroAgencia: 0,
		NumeroConta:   0,
		Saldo:         1000,
	}

	valid, err := utils.ValidarCPF(testeContaC.Titular.CPF)

	if err != nil {
		log.Fatal(err)
	}

	if valid {

		fmt.Println(testeContaC.Saldo)

		result, err := testeContaC.Sacar(10000)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(result)
		fmt.Println(err)
		fmt.Println(testeContaC.Saldo)
	}
}
