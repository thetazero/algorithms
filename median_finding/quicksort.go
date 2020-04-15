package main

func quicksortMedian(a []int) float64 {
	qSort(a, 0, len(a)-1)
	if len(a)%2 == 1 {
		return float64(a[len(a)/2])
	}
	return 0.5 * float64(a[len(a)/2-1]+a[len(a)/2])
}

func qSort(arr []int, start, end int) {
	if (end - start) < 1 {
		return
	}

	pivot := arr[end]
	splitIndex := start

	for i := start; i < end; i++ {
		if arr[i] < pivot {
			temp := arr[splitIndex]

			arr[splitIndex] = arr[i]
			arr[i] = temp

			splitIndex++
		}
	}

	arr[end] = arr[splitIndex]
	arr[splitIndex] = pivot

	qSort(arr, start, splitIndex-1)
	qSort(arr, splitIndex+1, end)
}
