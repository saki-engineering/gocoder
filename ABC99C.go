package gocoder

import (
	"fmt"
	"math"
)

func intMin(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

func main() {
	var N int
	fmt.Scan(&N)

	DP := make([]int, N+1)
	for i := 1; i <= N; i++ {
		DP[i] = int(math.Pow10(6))
	}

	for i := 1; i <= N; i++ {
		for j := 1; j <= i; j *= 6 {
			DP[i] = intMin(DP[i], DP[i-j]+1)
		}
		for j := 1; j <= i; j *= 9 {
			DP[i] = intMin(DP[i], DP[i-j]+1)
		}
	}
	fmt.Println(DP[N])
}
