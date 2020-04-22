package main

import (
	"math"
)

func ifft(y []complex) []complex {
	// fmt.Println(y)
	n := float64(len(y))
	if len(y) == 1 {
		return y
	}
	a := []complex{}
	w := []complex{}
	for i := range y {
		//prefill slice with null values
		a = append(a, complex{})
		//nth roots of unity calculations
		alpha := -2 * math.Pi * float64(i) / n
		w = append(w, complex{math.Cos(alpha), math.Sin(alpha)}) //inverse
	}
	A0, A1 := []complex{}, []complex{}
	for i := 0; i < len(y)/2; i++ {
		A0 = append(A0, y[i*2])
		A1 = append(A1, y[i*2+1])
	}
	a0 := ifft(A0)
	a1 := ifft(A1)
	for i := 0; i < len(y)/2; i++ {
		a[i] = w[i].Mulc(a1[i]).Addc(a0[i])
		a[i+len(a)/2] = a0[i].Subc(w[i].Mulc(a1[i]))
	}
	return a
}
