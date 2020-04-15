package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func main() {
	f, err := os.OpenFile("times.csv",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	e(err)
	_, err = f.WriteString("size,default sort,quicksort,quick select,determanistic quick select\n")
	e(err)
	n := 1
	for i := 0; i < 23; i++ {
		n *= 2
		this := generate(n)
		try(this)
		fmt.Println("trial:", i+1, "( size", n, ")")
	}
}

func generate(size int) []int {
	var a []int
	for i := 0; i < size; i++ {
		a = append(a, rand.Int()/2)
	}
	return a
}

func copy(a []int) []int {
	r := []int{}
	for i := range a {
		r = append(r, a[i])
	}
	return r
}

func e(err error) {
	if err != nil {
		log.Println(err)
	}
}

func try(a []int) {

	f, err := os.OpenFile("times.csv",
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	f.WriteString(strconv.Itoa(len(a)) + ",")
	e(err)
	one := copy(a)
	two := copy(a)
	three := copy(a)
	four := copy(a)

	start := time.Now()
	r := defsortMedian(two)
	ellapsed := time.Now().Sub(start)
	printRes(r, ellapsed, "default sort", f)
	f.WriteString(",")

	start = time.Now()
	r = quicksortMedian(one)
	ellapsed = time.Now().Sub(start)
	printRes(r, ellapsed, "quicksort", f)
	f.WriteString(",")

	start = time.Now()
	r = quickSelectMedian(three)
	ellapsed = time.Now().Sub(start)
	printRes(r, ellapsed, "quick select", f)
	f.WriteString(",")

	start = time.Now()
	r = quickSelectMedian(four)
	ellapsed = time.Now().Sub(start)
	printRes(r, ellapsed, "determanistic quick select", f)
	f.WriteString("\n")
}

func printRes(r float64, t time.Duration, name string, f *os.File) {
	fmt.Print("\033[1m", r, "\033[0m in: ", t, " (", name, ")\n")
	f.WriteString(strconv.FormatInt(t.Nanoseconds(), 10))
}
