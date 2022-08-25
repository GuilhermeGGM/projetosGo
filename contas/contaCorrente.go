package contas

import (
	"guilherme.moreira/clientes"
)

type ContaCorrente struct {
	Titular       clientes.Titular
	NumeroAgencia int
	NumeroConta   int
	saldo         float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) string {

	podeSacar := valorDoSaque <= c.saldo && valorDoSaque > 0

	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque reaizado com sucesso"
	} else {
		return "Erro! saldo insuficiente"
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) string {

	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Depósito feito com sucesso"
	} else {
		return "Erro! Valor de depósito inválido"
	}

}

func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) bool {

	if valorTransferencia < c.saldo && valorTransferencia > 0 {
		contaDestino.Depositar(valorTransferencia)
		c.saldo -= valorTransferencia
		return true
	} else {
		return false
	}
}

func (c *ContaCorrente) ObterSaldo() float64 {

	saldo := c.saldo
	return saldo
}
