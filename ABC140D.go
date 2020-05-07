package gocoder

import "fmt"

func main() {
	var (
		N, K, ans int
		S         string
	)
	fmt.Scan(&N, &K, &S)

	block := 1
	for i := 1; i < len(S); i++ {
		if S[i-1] != S[i] {
			block++
		}
	}

	if block/2 > K {
		ans = N + K*2 - block
	} else {
		ans = N - 1
	}
	fmt.Println(ans)
}
