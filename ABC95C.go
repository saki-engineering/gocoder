package gocoder

import (
	"fmt"
	"math"
)

func intMin(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

func main() {
	var A, B, C, X, Y int
	fmt.Scan(&A, &B, &C, &X, &Y)

	ans := intMin(A+B, 2*C) * intMin(X, Y)
	if X > Y {
		ans += intMin(2*C, A) * (X - Y)
	} else {
		ans += intMin(2*C, B) * (Y - X)
	}
	fmt.Println(ans)
}
