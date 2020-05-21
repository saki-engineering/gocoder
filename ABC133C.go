package gocoder

import (
	"fmt"
	"math"
)

func intMin(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

func main() {
	var L, R int
	fmt.Scan(&L, &R)

	if (R - L) >= 2019 {
		fmt.Println(0)
		return
	}

	L, R = L%2019, R%2019
	if L > R {
		fmt.Println(0)
		return
	}

	ans := 2019
	for i := L; i < R; i++ {
		for j := i + 1; j <= R; j++ {
			ans = intMin(ans, (i*j)%2019)
		}
	}
	fmt.Println(ans)
}
