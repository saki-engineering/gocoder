package gocoder

import (
	"fmt"
	"strconv"
)

func digit(N int) int {
	return len(strconv.Itoa(N))
}

func main() {
	var N int
	fmt.Scan(&N)

	ans := digit(N)
	for i := 1; i*i <= N; i++ {
		if N%i != 0 {
			continue
		}
		ans = digit(N / i)
	}
	fmt.Println(ans)
}
