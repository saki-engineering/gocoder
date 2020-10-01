package gocoder

import (
	"fmt"
	"math"
)

func intAbs(x int) int {
	return int(math.Abs(float64(x)))
}

func main() {
	var X, Y int
	fmt.Scan(&X, &Y)

	if intAbs(X-Y) <= 1 {
		fmt.Println("Brown")
		return
	}
	fmt.Println("Alice")
}
