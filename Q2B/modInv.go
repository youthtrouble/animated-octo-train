package main

import (
	"flag"
	"fmt"
	"math/big"
	"os"
	"strconv"
)


func GCD(a, b *big.Int) (q, s, t *big.Int) {

	r := big.NewInt(1) // since r cannot be  equal to zero
	q = new(big.Int).Set(a)
	prevQ := new(big.Int).Set(b)
	
	S0 := new(big.Int)
	s = big.NewInt(1)
	t = new(big.Int)
	T1 := big.NewInt(1)

	Instance := new(big.Int)

	for r.Sign() != 0 {

		//Using the Euclidean Algorithm,
		q, r := q.QuoRem(q, prevQ, r)

		Instance.Set(S0)
		S0.Mul(q, S0)
		S0.Sub(s, S0)
		s.Set(Instance)

		Instance.Set(T1)
		T1.Mul(q, T1)
		T1.Sub(t, T1)
		t.Set(Instance)

		q.Set(prevQ)
		prevQ.Set(r)

	}

	return
}

func ModularMultiplicativeInverse(a, b *big.Int) *big.Int {

	//The function GCD returns q, s, and t where q is the greatest common divisor,
	//GCD while s and t are the Bezout coefficients. it is well known that if the gcd(a, b) = r
	//then there exist integers p and s so that:  q = a*s + b*t

	_, s, _ := GCD(a, b)
	if s.Sign() == -1 {
		s.Add(s, b)
	}
	return s
}


func main() {

	flag.Parse()
	args := flag.Args()

	// requesting for a number if none is specified as a flag
	if len(args) < 1 {
		fmt.Println("Please specify a number")
		os.Exit(1)
	}

	aint64, err := strconv.ParseInt(args[0], 10, 64)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}

	bint64, err := strconv.ParseInt(args[1], 10, 64)
	if err != nil {
		// handle error
		fmt.Println(err)
		os.Exit(2)
	}

	a := big.NewInt(aint64)
	b := big.NewInt(bint64)

	k := ModularMultiplicativeInverse(a, b)
	fmt.Println(a, "and", b, "have a Modular Multiplicative Inverse of", k)
}