package main

import (
	"math/big"
	"reflect"
	"testing"
)

func TestGCD(t *testing.T) {
	type args struct {
		a *big.Int
		b *big.Int
	}
	tests := []struct {
		name  string
		args  args
		wantQ *big.Int
		wantS *big.Int
		wantT *big.Int
	}{

		{
			"Greatest Common Divisor(GCD) Test",
			args{
				big.NewInt(81),
				big.NewInt(57),
			},
			big.NewInt(3),
			big.NewInt(-7),
			big.NewInt(10),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotQ, gotS, gotT := GCD(tt.args.a, tt.args.b)
			if !reflect.DeepEqual(gotQ, tt.wantQ) {
				t.Errorf("GCD() gotQ = %v, want %v", gotQ, tt.wantQ)
			}
			if !reflect.DeepEqual(gotS, tt.wantS) {
				t.Errorf("GCD() gotS = %v, want %v", gotS, tt.wantS)
			}
			if !reflect.DeepEqual(gotT, tt.wantT) {
				t.Errorf("GCD() gotT = %v, want %v", gotT, tt.wantT)
			}
		})
	}
}

func TestModularMultiplicativeInverse(t *testing.T) {
	type args struct {
		a *big.Int
		b *big.Int
	}
	tests := []struct {
		name string
		args args
		want *big.Int
	}{

		{
			"Modular Multiplicative Inverse Test",
			args{
				big.NewInt(42),
				big.NewInt(2017),
			},
			big.NewInt(1969),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ModularMultiplicativeInverse(tt.args.a, tt.args.b); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ModularMultiplicativeInverse() = %v, want %v", got, tt.want)
			}
		})
	}
}
