package contas

import "guilherme.moreira/clientes"

type ContaPoupanca struct {
	Titular                              clientes.Titular
	NumeroAgencia, NumeroConta, Operação int
	saldo                                float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) string {

	podeSacar := valorDoSaque <= c.saldo && valorDoSaque > 0

	if podeSacar {
		c.saldo -= valorDoSaque
		return "Saque reaizado com sucesso"
	} else {
		return "Erro! saldo insuficiente"
	}
}

func (c *ContaPoupanca) Depositar(valorDoDeposito float64) string {

	if valorDoDeposito > 0 {
		c.saldo += valorDoDeposito
		return "Depósito feito com sucesso"
	} else {
		return "Erro! Valor de depósito inválido"
	}

}

func (c *ContaPoupanca) ObterSaldo() float64 {

	saldo := c.saldo
	return saldo
}
