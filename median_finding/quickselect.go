package main

import "math/rand"

func quickSelectMedian(a []int) float64 {
	if len(a)%2 == 1 {
		return float64(quickSelect(a, len(a)/2))
	}
	return 0.5 * float64(quickSelect(a, len(a)/2-1)+quickSelect(a, len(a)/2))
}

func quickSelect(a []int, k int) int {
	if len(a) == 1 {
		return a[0]
	}
	p := a[rand.Intn(len(a))]
	lows, highs, pivots := []int{}, []int{}, []int{}
	for i := range a {
		if a[i] < p {
			lows = append(lows, a[i])
		} else if a[i] > p {
			highs = append(highs, a[i])
		} else {
			pivots = append(pivots, a[i])
		}
	}
	if k < len(lows) {
		return quickSelect(lows, k)
	} else if k < len(lows)+len(pivots) {
		return pivots[0]
	} else {
		return quickSelect(highs, k-len(lows)-len(pivots))
	}
}
