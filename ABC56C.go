package gocoder

import "fmt"

func main() {
	var X int
	fmt.Scan(&X)

	sum, t := 0, 1
	for {
		sum += t
		if sum >= X {
			fmt.Println(t)
			return
		}
		t++
	}
}
