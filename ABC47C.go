package main

import "fmt"

func main() {
	var S string
	fmt.Scan(&S)

	N, ans := len(S), 0
	for i := 1; i < N; i++ {
		if S[i-1] != S[i] {
			ans++
		}
	}
	fmt.Println(ans)
}
