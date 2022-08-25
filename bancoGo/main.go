package main

import (
	"fmt"

	"github.com/GuilhermeGGM/projetosGo.git/contas"
)

func PagarBoleto(conta verificarConta, valorBoleto float64) {
	conta.Sacar(valorBoleto)
}

type verificarConta interface {
	Sacar(valor float64) string
}

func main() {

	contaTesteP := contas.ContaPoupanca{}
	contaTesteP.Depositar(100)
	PagarBoleto(&contaTesteP, 60)

	fmt.Println(contaTesteP.ObterSaldo())
}
