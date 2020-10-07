package main

import "fmt"

func main() {
	var N, M, ans int
	fmt.Scan(&N, &M)

	if 2*N > M {
		ans = M / 2
	} else {
		ans = N + (M-2*N)/4
	}
	fmt.Println(ans)
}
