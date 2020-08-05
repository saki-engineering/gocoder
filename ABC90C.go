package gocoder

import (
	"fmt"
	"math"
)

func intAbs(x int) int {
	return int(math.Abs(float64(x)))
}

func main() {
	var N, M int
	fmt.Scan(&N, &M)

	ans := intAbs(N-2) * intAbs(M-2)
	fmt.Println(ans)
}
