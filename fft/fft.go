// translated & modifyed from https://www.geeksforgeeks.org/fast-fourier-transformation-poynomial-multiplication/

package main

import (
	"math"
)

func fft(a []complex) []complex {
	n := float64(len(a))
	if len(a) == 1 {
		return a
	}
	y := []complex{}
	//nth roots of unity
	w := []complex{}
	for i := range a {
		//prefill slice with null values
		y = append(y, complex{})
		//nth roots of unity calculations
		alpha := 2 * math.Pi * float64(i) / n
		w = append(w, complex{math.Cos(alpha), math.Sin(alpha)})
	}
	A0, A1 := []complex{}, []complex{}
	for i := 0; i < len(a)/2; i++ {
		A0 = append(A0, a[i*2])
		A1 = append(A1, a[i*2+1])
	}
	y0 := fft(A0)
	y1 := fft(A1)
	for i := 0; i < len(a)/2; i++ {
		y[i] = w[i].Mulc(y1[i]).Addc(y0[i])
		y[i+len(a)/2] = y0[i].Subc(w[i].Mulc(y1[i]))
	}
	return y
}
