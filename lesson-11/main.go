package main

import (
	"fmt"
	"math/big"
)

func main() {
	bigInt()
	primeCheck()
	bigFloat()
}

func bigInt() {
	n := new(big.Int)
	n.SetString("34500000000000000000", 10)
	fmt.Printf("n= %v -%T\n", n, n)

	m := big.NewInt(128)
	fmt.Printf("m = %v - %T\n", m, m)

	n.Add(n, m)
	fmt.Printf("o = %v -%T\n", n, n)

	o := new(big.Int).Mul(n, m)
	fmt.Printf("o = %v - %T\n", o, o)

	fmt.Printf("n.Cmp(o): %d\n", n.Cmp(o))
	fmt.Printf("n.Cmp(m): %d\n", n.Cmp(m))
	fmt.Printf("n.Cmp(n): %d\n", n.Cmp(n))
}

func primeCheck() {
	primes := []*big.Int{
		big.NewInt(329),
		big.NewInt(337),
		big.NewInt(347),
		big.NewInt(349),
		big.NewInt(350),
		big.NewInt(358),
	}
	for _, p := range primes {
		fmt.Printf("%v a prime? %t\n", p, p.ProbablyPrime(1))
	}
}

func bigFloat() {
	var pi big.Float
	pi.Parse("3.14332648732473243247328473287432847328473289473284327983274", 10)
	fmt.Printf("pi = %.10g\n", &pi)
}
