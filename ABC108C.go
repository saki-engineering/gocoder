package gocoder

import (
	"fmt"
	"math"
)

func main() {
	var N, K int
	fmt.Scan(&N, &K)

	ans := int(math.Pow(float64(N/K), 3))
	if K%2 == 0 {
		c := N / K
		if N%K >= (K / 2) {
			c++
		}
		ans += int(math.Pow(float64(c), 3))
	}
	fmt.Println(ans)
}
