package gocoder

import (
	"fmt"
	"math"
)

func main() {
	MOD := int(math.Pow10(9)) + 7

	var N int
	fmt.Scan(&N)

	S := [2]string{}
	fmt.Scan(&S[0], &S[1])

	var ans int
	if S[0][0] == S[1][0] {
		ans = 3
	} else {
		ans = 6
	}
	for i := 1; i < N; i++ {
		if S[0][i] == S[1][i] {
			if S[0][i-1] == S[1][i-1] {
				ans = (ans * 2) % MOD
			}
		} else {
			if S[0][i] == S[0][i-1] {
				continue
			}
			if S[0][i-1] == S[1][i-1] {
				ans = (ans * 2) % MOD
			} else {
				ans = (ans * 3) % MOD
			}
		}
	}
	fmt.Println(ans)
}
