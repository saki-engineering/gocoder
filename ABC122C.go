package gocoder

import (
	"fmt"
)

func main() {
	var (
		N, Q, l, r int
		S          string
	)
	fmt.Scan(&N, &Q, &S)

	cnt := make([]int, N+1)
	for i := 0; i < N; i++ {
		if i+1 < N && string(S[i]) == "A" && string(S[i+1]) == "C" {
			cnt[i+1] = cnt[i] + 1
		} else {
			cnt[i+1] = cnt[i]
		}
	}

	for i := 0; i < Q; i++ {
		fmt.Scan(&l, &r)
		fmt.Println(cnt[r-1] - cnt[l-1])
	}
}
