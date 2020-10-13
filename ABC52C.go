package gocoder

import (
	"fmt"
	"math"
)

func main() {
	MOD := int(math.Pow10(9)) + 7

	var N int
	fmt.Scan(&N)

	pList := make([]int, N+1)
	for i := 1; i <= N; i++ {
		n, q := i, 2
		for n > 1 {
			if n%q == 0 {
				pList[q]++
				n /= q
			} else {
				q++
			}
		}
	}

	ans := 1
	for _, p := range pList {
		ans = (ans * (p + 1)) % MOD
	}
	fmt.Println(ans)
}
