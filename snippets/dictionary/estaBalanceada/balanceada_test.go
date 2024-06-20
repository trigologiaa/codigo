package main

import "testing"

func TestBalanceada(t *testing.T) {
	type args struct {
		cadena string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{name: "caso 1: enunciado balanceada", args: args{"[()]{}{[()()]()}"}, want: true},
		{name: "caso 2: enunciado desbalanceada", args: args{"[(])"}, want: false},
		{
			name: "caso 3: desbalanceada no deja la pila vacia",
			args: args{"[()]{}{[()()]()}{[]"},
			want: false,
		},
		{name: "caso 4: cadena vacia", args: args{""}, want: true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Balanceada(tt.args.cadena); got != tt.want {
				t.Errorf("Balanceada() = %v, want %v", got, tt.want)
			}
		})
	}
}