package main

import "sort"

func defsortMedian(a []int) float64 {
	sort.Ints(a)
	if len(a)%2 == 1 {
		return float64(a[len(a)/2])
	}
	return 0.5 * float64(a[len(a)/2-1]+a[len(a)/2])
}
