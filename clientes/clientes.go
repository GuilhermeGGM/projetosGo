package clientes

type Titular struct {
	Nome      string
	CPF       string
	Profissao string
}

func (c *Titular) ValidarCPF(cpf string) bool {

	if len(cpf) == 11 {
		return true
	} else {
		return false
	}
}
