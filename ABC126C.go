package gocoder

import "fmt"

func main() {
	var (
		N, K int
		ans  float64
	)
	fmt.Scan(&N, &K)

	for i := 1; i <= N; i++ {
		rate := 1 / float64(N)
		point := i
		for point < K {
			point *= 2
			rate /= 2
		}
		ans += rate
	}
	fmt.Println(ans)
}
