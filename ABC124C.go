package gocoder

import (
	"fmt"
	"math"
)

func intMin(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

func main() {
	var S string
	fmt.Scan(&S)

	cnt := []int{0, 0}
	for i, rn := range S {
		s := string(rn)
		if s == "0" {
			cnt[-(i%2-1)]++
		} else {
			cnt[i%2]++
		}
	}

	fmt.Println(intMin(cnt[0], cnt[1]))
}
