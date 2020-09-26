package main

import (
	"fmt"
	"math"
)

func main() {
	var s int
	var a float64
	var es, psi, dea, des, dsa float64

	fmt.Println("Derivations of Erlang B formula")
	fmt.Print("s=")
	fmt.Scanf("%d", &s)
	fmt.Print("a=")
	fmt.Scanf("%f", &a)

	psi, es = Psi(s, a, es)
	dea, des, dsa = Der(s, a, es, psi, dea, des, dsa)

	fmt.Printf("s=%d\ta=%f\tEs(a)=%f\n", s, a, es)
	fmt.Printf("Psi=%g\n", psi)
	fmt.Printf("dEs/da=%g\tdEs/ds=%f\n", dea, des)
	fmt.Printf("ds/da=%f\n", dsa)
}

func Psi(s int, a float64, es float64) (float64, float64) {
	var psi, ps1 float64

	if a > 1 {
		ps1 = Ps1a(a)
	} else {
		ps1 = Ps1c(a)
	}
	psi = ps1
	es = 1.0
	for i := 1; i <= s; i++ {
		es = a * es / (float64(i) + a*es)
		psi = (1 - es) / (psi + 1/float64(i))
	}
	return psi, es
}

func Ps1a(a float64) float64 {
	var a1 = 8.5733287401
	var a2 = 18.059016973
	var a3 = 8.6347608925
	var a4 = 0.2677737343
	var b1 = 9.5733223454
	var b2 = 25.6329561486
	var b3 = 21.0996530827
	var b4 = 3.9584969228

	psn := 1 + a1/a + a2/math.Pow(a, 2) + a3/math.Pow(a, 3) + a4/math.Pow(a, 4)
	psd := 1 + b1/a + b2/math.Pow(a, 2) + b3/math.Pow(a, 3) + b4/math.Pow(a, 4)
	ps1 := psn / psd / a

	return ps1
}

func Ps1c(a float64) float64 {
	var c0 = -0.57721566
	var c1 = 0.99999193
	var c2 = -0.24991055
	var c3 = 0.05519968
	var c4 = -9.76004e-3
	var c5 = 1.07857e-3

	ps1 := math.Exp(a) * (c0 + c1*a + c2*math.Pow(a, 2) + c3*math.Pow(a, 3) + c4*math.Pow(a, 4+c5*math.Pow(a, 5)-math.Log10(a)))
	return ps1
}

func Der(s int, a float64, es float64, psi float64, dea float64, des float64, dsa float64) (float64, float64, float64) {
	dea = (float64(s)/a - 1.0 + es) * es
	des = -psi * es
	dsa = (float64(s)/a - 1.0 + es) / psi
	return dea, des, dsa
}
