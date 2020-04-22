package main

//Complex number stored in real,imaginary form
type complex [2]float64

func (c1 complex) Addc(c2 complex) complex {
	return complex{c1[0] + c2[0], c1[1] + c2[1]}
}

func (c1 complex) Subc(c2 complex) complex {
	return complex{c1[0] - c2[0], c1[1] - c2[1]}
}

func (c1 complex) Mulc(c2 complex) complex {
	return complex{c1[0]*c2[0] - c1[1]*c2[1], c1[0]*c2[1] + c1[1]*c2[0]}
}

func (c1 complex) Mul(a float64) complex {
	return complex{c1[0] * a, c1[1] * a}
}
