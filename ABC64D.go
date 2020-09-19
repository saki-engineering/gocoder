package gocoder

import (
	"fmt"
	"math"
	"strings"
)

func intMin(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

func main() {
	var (
		N int
		S string
	)
	fmt.Scan(&N, &S)

	x, d := 0, 0
	for i := 0; i < N; i++ {
		if S[i] == '(' {
			d++
		} else {
			d--
		}
		x = intMin(x, d)
	}

	ans := strings.Repeat("(", -x) + S + strings.Repeat(")", d-x)
	fmt.Println(ans)
}
