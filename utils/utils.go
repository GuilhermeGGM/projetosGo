package utils

import (
	"errors"
)

func ValidarCPF(cpf string) (bool, error) {

	if len(cpf) == 11 {
		return true, nil
	} else {
		return false, errors.New("CPF inválido")
	}
}
