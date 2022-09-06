package utils

func ValidarCPF(cpf string) bool {

	if len(cpf) == 11 {
		return true
	} else {
		return false
	}
}
