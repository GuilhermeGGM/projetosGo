package contas

import (
	"errors"

	"github.com/GuilhermeGGM/projetosGo.git/clientes"
)

type ContaCorrente struct {
	Titular                    clientes.Titular
	NumeroAgencia, NumeroConta int
	Saldo                      float64
}

func (c *ContaCorrente) Sacar(valorDoSaque float64) (string, error) {

	podeSacar := valorDoSaque <= c.Saldo && valorDoSaque > 0

	if !podeSacar {
		return "", errors.New("Erro, valor de saque inválido")
	} else {
		c.Saldo -= valorDoSaque
		return "Sucesso", nil
	}
}

func (c *ContaCorrente) Depositar(valorDoDeposito float64) (string, error) {

	if valorDoDeposito > 0 {
		c.Saldo += valorDoDeposito
		return "Depósito feito com sucesso", nil
	} else {
		return "", errors.New("Erro! Valor de depósito inválido")
	}

}

func (c *ContaCorrente) Transferir(valorTransferencia float64, contaDestino *ContaCorrente) (string, error) {

	if valorTransferencia <= c.Saldo && valorTransferencia > 0 {
		contaDestino.Depositar(valorTransferencia)
		c.Saldo -= valorTransferencia
		return "Transferência feita com sucesso", nil
	} else {
		return "", errors.New("Erro")
	}
}

func (c *ContaCorrente) ObterSaldo() float64 {

	saldo := c.Saldo
	return saldo
}
