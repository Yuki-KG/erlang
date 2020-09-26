package main

import "fmt"

func main() {
	var s int
	var a, es float64

	fmt.Println("Erlang B Formula (positive integer server)")
	fmt.Printf("s=")
	fmt.Scanf("%d", &s)
	fmt.Printf("a=")
	fmt.Scanf("%f", &a)

	es = Es(s, a)

	fmt.Printf("s=%d\ta=%f\tEs(a)=%f\n", s, a, es)
}

func Es(s int, a float64) float64 {
	es := 1.0
	for i := 1; i <= s; i++ {
		es = a * es / (float64(i) + a*es)
	}
	return es
}
