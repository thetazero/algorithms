package main

import "fmt"

func main() {
	a := []complex{
		complex{1, 0},
		complex{2, 0},
		complex{3, 0},
		complex{4, 0},
		// complex{5, 0},
	}
	b := fft(a)
	for i, v := range b {
		fmt.Println(i, v)
	}
	c := ifft(b)
	for i, v := range c {
		fmt.Println(i, v.Mul(0.25))
	}
	fmt.Println("hallo")
}
