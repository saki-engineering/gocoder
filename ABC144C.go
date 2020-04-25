package gocoder

import (
	"fmt"
	"math"
)

func abc144c() {
	var N int
	fmt.Scanf("%d", &N)

	ans := N - 1
	for i := 1; i <= int(math.Sqrt(float64(N))); i++ {
		if N%i == 0 {
			ans = int(math.Min(float64(ans), float64(i+N/i-2)))
		}
	}

	fmt.Println(ans)
}
