package gocoder

import "fmt"

func main() {
	var N, M int
	fmt.Scan(&N, &M)

	ans := (M*1900 + (N-M)*100) * (1 << M)
	fmt.Println(ans)
}
