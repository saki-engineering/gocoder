package gocoder

import (
	"fmt"
	"math"
)

func main() {
	MOD := int(math.Pow10(9) + 7)
	var H, W, K int
	fmt.Scan(&H, &W, &K)

	DP := make([][]int, H+1)
	for i := 0; i <= H; i++ {
		DP[i] = make([]int, W)
	}
	DP[0][0] = 1

	for i := 1; i <= H; i++ {
		// jはW-1桁(0~W-2桁目まで)
		for j := 0; j < (1 << (W - 1)); j++ {
			flg := true
			for k := 0; k < (W - 2); k++ {
				if (j&(1<<k))*(j&(1<<(k+1))) > 0 {
					flg = false
					break
				}
			}

			if !flg {
				continue
			}

			for k := 0; k < W; k++ {
				if k == 0 {
					if j&(1<<k) > 0 {
						DP[i][k] += DP[i-1][k+1]
					} else {
						DP[i][k] += DP[i-1][k]
					}
				} else if k == W-1 {
					if j&(1<<(k-1)) > 0 {
						DP[i][k] += DP[i-1][k-1]
					} else {
						DP[i][k] += DP[i-1][k]
					}
				} else {
					if j&(1<<k) > 0 {
						DP[i][k] += DP[i-1][k+1]
					} else if j&(1<<(k-1)) > 0 {
						DP[i][k] += DP[i-1][k-1]
					} else {
						DP[i][k] += DP[i-1][k]
					}
				}
				DP[i][k] %= MOD
			}
		}
	}
	fmt.Println(DP[H][K-1])
}
