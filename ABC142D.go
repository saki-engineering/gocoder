package gocoder

import (
	"fmt"
)

func gcd(x, y int) int {
	if x < y {
		x, y = y, x
	}

	if y == 0 {
		return x
	}
	return gcd(y, x%y)
}

func abc142d() {
	var A, B int
	fmt.Scan(&A, &B)
	G := gcd(A, B)

	primeList := map[int]int{}
	for p := 2; p*p <= G; p++ {
		for G%p == 0 {
			G /= p
			primeList[p]++
		}
	}
	if G > 1 {
		primeList[G]++
	}

	fmt.Println(len(primeList) + 1)
}
