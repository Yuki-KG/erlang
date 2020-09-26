package main

import (
	"fmt"
	"math"
)

func main() {
	var k, n int
	var s, a, x, ex, es float64

	fmt.Println("Erlang B Formula (continuous server)")

	fmt.Printf("s=")
	fmt.Scanf("%f", &s)
	fmt.Printf("a=")
	fmt.Scanf("%f", &a)

	x, k, n = Xk(s, a, k, n)
	ex = Ex(s, a, k, x, n)
	es = Es(a, ex, x, n)

	fmt.Printf("s=%f\ta=%f\tEs(a)=%f\n", s, a, es)
}

func Xk(s float64, a float64, k int, n int) (float64, int, int) {
	var x float64
	n = int(s)
	x = s - float64(n)
	k = int(5/4*math.Sqrt(x+500) + 4/a)
	return x, k, n
}

func Ex(s float64, a float64, k int, x float64, n int) float64 {
	var esk float64
	var ex float64

	esk = a
	for i := k; i >= 1; i-- {
		esk = a + (-x+float64(i)-1)/(1+float64(i)/esk)
	}
	ex = esk / a
	return ex
}

func Es(a float64, ex float64, x float64, n int) float64 {
	var es float64
	es = ex
	for j := 1; j <= n; j++ {
		es = a * es / (x + float64(j) + a*es)
	}
	return es
}
