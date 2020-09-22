package gocoder

import "fmt"

func main() {
	var N, a int
	fmt.Scan(&N)

	colors := [8]bool{}
	numOver := 0
	for i := 0; i < N; i++ {
		fmt.Scan(&a)
		if a >= 3200 {
			numOver++
		} else {
			colors[a/400] = true
		}
	}

	numUnder := 0
	for _, color := range colors {
		if color {
			numUnder++
		}
	}

	min, max := numUnder, numUnder+numOver
	if numUnder == 0 {
		min = 1
	}
	fmt.Println(min, max)
}
