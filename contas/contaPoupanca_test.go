package contas

import (
	"testing"

	"github.com/GuilhermeGGM/projetosGo.git/clientes"
)

func TestContaPoupanca_Sacar(t *testing.T) {
	type fields struct {
		Titular       clientes.Titular
		NumeroAgencia int
		NumeroConta   int
		Operação      int
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
			name: "Teste_saque_sucesso",
			fields: fields{
				Titular: clientes.Titular{
					Nome:      "Albertina",
					CPF:       "03973861099",
					Profissao: "Empresária",
				},
				NumeroAgencia: 1234567,
				NumeroConta:   12345,
				Saldo:         100000,
			},
			args: args{
				valorDoSaque: 1000,
			},
			want:    "Saque reaizado com sucesso",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ContaPoupanca{
				Titular:       tt.fields.Titular,
				NumeroAgencia: tt.fields.NumeroAgencia,
				NumeroConta:   tt.fields.NumeroConta,
				Saldo:         tt.fields.Saldo,
			}
			got, err := c.Sacar(tt.args.valorDoSaque)
			if (err != nil) != tt.wantErr {
				t.Errorf("ContaPoupanca.Sacar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ContaPoupanca.Sacar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContaPoupanca_Depositar(t *testing.T) {
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
			name: "Teste_depositar_sucesso",
			fields: fields{
				Titular: clientes.Titular{
					Nome:      "Neide",
					CPF:       "91109266073",
					Profissao: "Policial",
				},
				NumeroAgencia: 1234567,
				NumeroConta:   12345,
				Saldo:         10000,
			},
			args: args{
				valorDoDeposito: 1000,
			},
			want:    "Depósito feito com sucesso",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ContaPoupanca{
				Titular:       tt.fields.Titular,
				NumeroAgencia: tt.fields.NumeroAgencia,
				NumeroConta:   tt.fields.NumeroConta,
				Saldo:         tt.fields.Saldo,
			}
			got, err := c.Depositar(tt.args.valorDoDeposito)
			if (err != nil) != tt.wantErr {
				t.Errorf("ContaPoupanca.Depositar() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ContaPoupanca.Depositar() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestContaPoupanca_ObterSaldo(t *testing.T) {
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
			name: "Teste_obterSaldo_sucesso",
			fields: fields{
				Titular: clientes.Titular{
					Nome:      "Elisangela",
					CPF:       "02010311086",
					Profissao: "Estagiária",
				},
				NumeroAgencia: 1234567,
				NumeroConta:   12345,
				Saldo:         100,
			},
			want: 100,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &ContaPoupanca{
				Titular:       tt.fields.Titular,
				NumeroAgencia: tt.fields.NumeroAgencia,
				NumeroConta:   tt.fields.NumeroConta,
				Saldo:         tt.fields.Saldo,
			}
			if got := c.ObterSaldo(); got != tt.want {
				t.Errorf("ContaPoupanca.ObterSaldo() = %v, want %v", got, tt.want)
			}
		})
	}
}
