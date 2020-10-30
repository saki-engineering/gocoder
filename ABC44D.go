package gocoder

import "fmt"

func f(b, n int) int {
	ans, tmp := 0, n
	for tmp > 0 {
		r := tmp % b
		ans += r
		tmp = (tmp - r) / b
	}
	return ans
}

func main() {
	var N, S int
	fmt.Scan(&N, &S)

	if N == S {
		fmt.Println(N + 1)
		return
	}

	Nsqrt := 0
	for b := 2; b*b <= N; b++ {
		if f(b, N) == S {
			fmt.Println(b)
			return
		}
		Nsqrt = b
	}

	ans := -1
	for p := 1; p*p < N; p++ {
		b := (N-S)/p + 1
		if Nsqrt < b && f(b, N) == S {
			ans = b
		}
	}

	fmt.Println(ans)
}
