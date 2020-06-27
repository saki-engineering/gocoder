package gocoder

import "fmt"

func count(n int, factor *[]int) int {
	ans := 0
	for _, f := range *factor {
		if f >= n {
			ans++
		}
	}
	return ans
}

func main() {
	var N int
	fmt.Scan(&N)

	prime := [...]int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 37, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97}
	factor := make([]int, len(prime))

	for i := 1; i <= N; i++ {
		tmp := i
		p := 0
		for tmp > 1 {
			if tmp%prime[p] == 0 {
				tmp /= prime[p]
				factor[p]++
			} else {
				p++
			}
		}
	}

	ans := 0
	ans += count(74, &factor)
	ans += count(14, &factor) * (count(4, &factor) - 1)
	ans += count(24, &factor) * (count(2, &factor) - 1)
	ans += (count(4, &factor) * (count(4, &factor) - 1) * (count(2, &factor) - 2)) / 2
	fmt.Println(ans)
}
