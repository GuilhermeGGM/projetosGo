package contas

import (
	"errors"

	"github.com/GuilhermeGGM/projetosGo.git/clientes"
)

type ContaPoupanca struct {
	Titular                    clientes.Titular
	NumeroAgencia, NumeroConta int
	Saldo                      float64
}

func (c *ContaPoupanca) Sacar(valorDoSaque float64) (string, error) {

	podeSacar := valorDoSaque <= c.Saldo && valorDoSaque > 0

	if podeSacar {
		c.Saldo -= valorDoSaque
		return "Saque reaizado com sucesso", nil
	} else {
		return "", errors.New("Valor de saque inválido")
	}
}

func (c *ContaPoupanca) Depositar(valorDoDeposito float64) (string, error) {

	if valorDoDeposito > 0 {
		c.Saldo += valorDoDeposito
		return "Depósito feito com sucesso", nil
	} else {
		return "", errors.New("Valor de depósito inválido")
	}
}

func (c *ContaPoupanca) ObterSaldo() float64 {

	saldo := c.Saldo
	return saldo
}
