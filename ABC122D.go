package gocoder

import (
	"fmt"
	"math"
)

func main() {
	MOD := int(math.Pow10(9) + 7)

	var N int
	fmt.Scan(&N)

	DP := make([][]int, N+1)
	for i := 0; i <= N; i++ {
		DP[i] = make([]int, 64)
	}

	for i := 0; i < 64; i++ {
		DP[3][i] = 1
	}
	DP[3][6], DP[3][9], DP[3][18] = 0, 0, 0

	for i := 4; i <= N; i++ {
		for j := 0; j < 64; j++ {
			if j == 6 || j == 9 || j == 18 {
				continue
			}
			for k := 0; k < 64; k++ {
				if j%4 == 2 && k%16 == 1 {
					continue
				} else if j%4 == 1 && k%16 == 2 {
					continue
				} else if j%4 == 2 && k%16 == 4 {
					continue
				} else if j%4 == 2 && k>>2 == 1 {
					continue
				} else if j%4 == 2 && k>>4 == 0 && k%4 == 1 {
					continue
				} else if (j >> 2) == (k % 16) {
					DP[i][j] = (DP[i][j] + DP[i-1][k]) % MOD
				}
			}
		}
	}

	ans := 0
	for i := 0; i < 64; i++ {
		ans = (ans + DP[N][i]) % MOD
	}
	fmt.Println(ans)
}
