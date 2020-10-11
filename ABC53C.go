package gocoder

import "fmt"

func main() {
	var x int
	fmt.Scan(&x)

	q, r := x/11, x%11
	ans := 0
	if r == 0 {
		ans = 2 * q
	} else if r <= 6 {
		ans = 2*q + 1
	} else {
		ans = 2*q + 2
	}
	fmt.Println(ans)
}
