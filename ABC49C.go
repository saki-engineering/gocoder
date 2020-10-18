package main

import (
	"fmt"
	"strings"
)

func main() {
	var S string
	fmt.Scan(&S)
	N := len(S)

	A := [4]string{"dream", "erase", "eraser", "dreamer"}
	DP := make([]bool, N+1)
	DP[0] = true

	for i := 0; i < N; i++ {
		for _, pre := range A {
			if strings.HasPrefix(S[i:], pre) && DP[i] {
				DP[i+len(pre)] = true
			}
		}
	}

	if DP[N] {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}
}
