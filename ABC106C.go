package gocoder

import "fmt"

func main() {
	var (
		S string
		K int
	)
	fmt.Scan(&S, &K)

	for i, s := range S {
		if rune(s) != '1' {
			fmt.Println(string(s))
			return
		}
		if i == K-1 {
			fmt.Println(string(s))
			return
		}
	}
}
