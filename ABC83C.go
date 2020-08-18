package gocoder

import "fmt"

func main() {
	var X, Y int
	fmt.Scan(&X, &Y)

	ans := 0
	for X <= Y {
		X *= 2
		ans++
	}
	fmt.Println(ans)
}
