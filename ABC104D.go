package gocoder

import (
	"fmt"
	"math"
)

func main() {
	var S string
	fmt.Scan(&S)
	MOD := int(math.Pow10(9) + 7)
	N := len(S)
	T := "ABC"

	DP := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		DP[i] = make([]int, 4)
	}
	DP[N][3] = 1

	for i := N - 1; i >= 0; i-- {
		for j := 0; j < 3; j++ {
			m1, m2 := 1, 0
			if S[i] == '?' {
				m1, m2 = 3, 1
			} else if S[i] == T[j] {
				m2 = 1
			}
			DP[i][j] = (m1*DP[i+1][j] + m2*DP[i+1][j+1]) % MOD
		}
		if S[i] == '?' {
			DP[i][3] = (3 * DP[i+1][3]) % MOD
		} else {
			DP[i][3] = DP[i+1][3]
		}
	}
	fmt.Println(DP[0][0])
}
