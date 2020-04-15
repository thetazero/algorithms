package main

import (
	"math/rand"
	"sort"
)

func detQuickSelectMedian(a []int) float64 {
	if len(a)%2 == 1 {
		return float64(detQuickSelect(a, len(a)/2))
	}
	return 0.5 * float64(detQuickSelect(a, len(a)/2-1)+quickSelect(a, len(a)/2))
}

func detQuickSelect(a []int, k int) int {
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

func pick(a []int) int {
	if len(a) < 5 {
		sort.Ints(a)
		if len(a)%2 == 1 {
			return a[len(a)/2]
		}
		return (a[len(a)/2-1] + a[len(a)/2]) / 2
	}
	medians := []int{}
	for i := 0; i < len(a); i++ {
		for j := i; j < i+5 && j < len(a); j++ {
			b := i
			for k := i; k < i+5 && k < len(a); k++ {
				if a[k] > a[b] {
					b = k
				}
			}
			a[j], a[b] = a[b], a[j]
		}
	}
	for i := 0; i < len(a)-5; i += 5 {
		medians = append(medians, a[i+2])
	}
	mofm := int(detQuickSelectMedian(a))
	return mofm
}
