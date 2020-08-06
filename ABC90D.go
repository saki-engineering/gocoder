package gocoder

import (
	"fmt"
)

func main() {
	var N, K int
	fmt.Scan(&N, &K)

	ans := 0
	for b := K + 1; b <= N; b++ {
		q := N / b
		r := N % b

		ans += q * (b - K)
		if r-K+1 >= 0 {
			ans += r - K + 1
		}
	}
	if K == 0 {
		ans -= N - K
	}
	fmt.Println(ans)
}
