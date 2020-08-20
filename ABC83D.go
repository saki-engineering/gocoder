package gocoder

import (
	"fmt"
	"math"
)

func intMax(x, y int) int {
	return int(math.Max(float64(x), float64(y)))
}

func intMin(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

func main() {
	var S string
	fmt.Scan(&S)

	N := len(S)
	ans := N
	for i := 0; i < N-1; i++ {
		if S[i] != S[i+1] {
			ans = intMin(ans, intMax(i+1, N-i-1))
		}
	}
	fmt.Println(ans)
}
