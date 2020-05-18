package gocoder

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var S string
	fmt.Scan(&S)

	N := len(S)
	MOD := int(math.Pow10(9)) + 7

	ans := make([][]int, 2)
	for i := 0; i < 2; i++ {
		ans[i] = make([]int, 13)
	}
	ans[0][0] = 1

	pow10Mod := [6]int{}
	for i := 0; i < 6; i++ {
		pow10Mod[i] = int(math.Pow10(i)) % 13
	}

	for i, s := range S {
		s := string(s)
		if s == "?" {
			for j := 0; j < 10; j++ {
				plus := (pow10Mod[(N-1-i)%6] * j) % 13
				for k := 0; k < 13; k++ {
					ans[1][(k+plus)%13] += ans[0][k]
				}
			}
		} else {
			s, _ := strconv.Atoi(s)
			plus := (pow10Mod[(N-1-i)%6] * s) % 13
			for k := 0; k < 13; k++ {
				ans[1][(k+plus)%13] += ans[0][k]
			}
		}

		for k := 0; k < 13; k++ {
			ans[0][k] = ans[1][k] % MOD
			ans[1][k] = 0
		}
	}

	fmt.Println(ans[0][5])
}
