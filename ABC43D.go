package gocoder

import "fmt"

func main() {
	var S string
	fmt.Scan(&S)

	N := len(S)
	for i := 1; i < N; i++ {
		if S[i-1] == S[i] {
			fmt.Println(i, i+1)
			return
		}
	}
	for i := 2; i < N; i++ {
		if S[i-2] == S[i] {
			fmt.Println(i-1, i+1)
			return
		}
	}
	fmt.Println(-1, -1)
}
