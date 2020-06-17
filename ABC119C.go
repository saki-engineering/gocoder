package gocoder

import (
	"fmt"
	"math"
)

func intMin(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

func intAbs(x int) int {
	return int(math.Abs(float64(x)))
}

func main() {
	var N, A, B, C int
	fmt.Scan(&N, &A, &B, &C)

	l := make([]int, N)
	for i := 0; i < N; i++ {
		fmt.Scan(&l[i])
	}

	ans := 10000000
	for i := 0; i < (1 << (2 * uint(N))); i++ {
		a, b, c, mp := 0, 0, 0, 0
		for j := 0; j < N; j++ {
			if k := (i >> (2 * uint(j))) & 3; k == 1 {
				a += l[j]
				mp += 10
			} else if k == 2 {
				b += l[j]
				mp += 10
			} else if k == 3 {
				c += l[j]
				mp += 10
			}
		}

		if a*b*c == 0 {
			continue
		}

		mp += intAbs(A-a) + intAbs(B-b) + intAbs(C-c) - 30
		ans = intMin(ans, mp)
	}
	fmt.Println(ans)
}
