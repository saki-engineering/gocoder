package gocoder

import (
	"fmt"
	"math"
)

func main() {
	MOD := int(math.Pow10(9)) + 7

	var N, M int
	fmt.Scan(&N, &M)
	if M > N {
		N, M = M, N
	}

	ans := 1
	if N-M > 1 {
		fmt.Println(0)
		return
	} else if N-M == 0 {
		ans = 2
	}

	for i := 1; i <= N; i++ {
		ans = (ans * i) % MOD
	}
	for i := 1; i <= M; i++ {
		ans = (ans * i) % MOD
	}

	fmt.Println(ans)
}
