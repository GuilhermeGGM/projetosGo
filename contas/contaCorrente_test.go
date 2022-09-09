package contas

import (
	"testing"

	"github.com/GuilhermeGGM/projetosGo.git/clientes"
)

func TestContaCorrente_Sacar(t *testing.T) {
	type fields struct {
		Titular       clientes.Titular
		NumeroAgencia int
		NumeroConta   int
		Saldo         float64
	}
	type args struct {
		valorDoSaque float64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Teste_saque_com_sucesso",
			fields: fields{
				Titular: clientes.Titular{
					Nome:      "Roberto",
					CPF:       "82123978078",
					Profissao: "Pedreiro",
				},
				NumeroAgencia: 12345,
				NumeroConta:   1234567,
				Saldo:         10000,
			},
			args: args{
				valorDoSaque: 1000,
			},
			want:    "Sucesso",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ContaCorrente{
				Titular:       tt.fields.Titular,
				NumeroAgencia: tt.fields.NumeroAgencia,
				NumeroConta:   tt.fields.NumeroConta,
				Saldo:         tt.fields.Saldo,
			}
			got, err := c.Sacar(tt.args.valorDoSaque)
			if (err != nil) != tt.wantErr {
				t.Errorf("ContaCorrente.Sacar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ContaCorrente.Sacar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContaCorrente_Depositar(t *testing.T) {
	type fields struct {
		Titular       clientes.Titular
		NumeroAgencia int
		NumeroConta   int
		Saldo         float64
	}
	type args struct {
		valorDoDeposito float64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Teste_depósito_sucesso",
			fields: fields{
				Titular: clientes.Titular{
					Nome:      "Cleiton",
					CPF:       "14839048061",
					Profissao: "Médico",
				},
				NumeroAgencia: 1234567,
				NumeroConta:   12345,
				Saldo:         100000,
			},
			args: args{
				valorDoDeposito: 10000,
			},
			want:    "Depósito feito com sucesso",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ContaCorrente{
				Titular:       tt.fields.Titular,
				NumeroAgencia: tt.fields.NumeroAgencia,
				NumeroConta:   tt.fields.NumeroConta,
				Saldo:         tt.fields.Saldo,
			}
			got, err := c.Depositar(tt.args.valorDoDeposito)
			if (err != nil) != tt.wantErr {
				t.Errorf("ContaCorrente.Depositar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ContaCorrente.Depositar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContaCorrente_Transferir(t *testing.T) {
	type fields struct {
		Titular       clientes.Titular
		NumeroAgencia int
		NumeroConta   int
		Saldo         float64
	}
	type args struct {
		valorTransferencia float64
		contaDestino       *ContaCorrente
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Teste_transferência_sucesso",
			fields: fields{
				Titular: clientes.Titular{
					Nome:      "cíntia",
					CPF:       "43667270054",
					Profissao: "Caminhoneira",
				},
				NumeroAgencia: 1234567,
				NumeroConta:   12345,
				Saldo:         10000,
			},
			args: args{
				valorTransferencia: 1000,
				contaDestino: &ContaCorrente{
					Titular: clientes.Titular{
						Nome:      "Maria",
						CPF:       "18815867074",
						Profissao: "Desempregrada",
					},
					NumeroAgencia: 7654321,
					NumeroConta:   54321,
					Saldo:         0,
				},
			},
			want:    "Transferência feita com sucesso",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ContaCorrente{
				Titular:       tt.fields.Titular,
				NumeroAgencia: tt.fields.NumeroAgencia,
				NumeroConta:   tt.fields.NumeroConta,
				Saldo:         tt.fields.Saldo,
			}
			got, err := c.Transferir(tt.args.valorTransferencia, tt.args.contaDestino)
			if (err != nil) != tt.wantErr {
				t.Errorf("ContaCorrente.Transferir() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ContaCorrente.Transferir() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContaCorrente_ObterSaldo(t *testing.T) {
	type fields struct {
		Titular       clientes.Titular
		NumeroAgencia int
		NumeroConta   int
		Saldo         float64
	}
	tests := []struct {
		name   string
		fields fields
		want   float64
	}{
		{
			name: "Teste_ObterSaldo_Sucesso",
			fields: fields{
				Titular: clientes.Titular{
					Nome:      "Marina",
					CPF:       "30402279042",
					Profissao: "Streamer",
				},
				NumeroAgencia: 1234567,
				NumeroConta:   12345,
				Saldo:         1000000,
			},
			want: 1000000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ContaCorrente{
				Titular:       tt.fields.Titular,
				NumeroAgencia: tt.fields.NumeroAgencia,
				NumeroConta:   tt.fields.NumeroConta,
				Saldo:         tt.fields.Saldo,
			}
			if got := c.ObterSaldo(); got != tt.want {
				t.Errorf("ContaCorrente.ObterSaldo() = %v, want %v", got, tt.want)
			}
		})
	}
}
