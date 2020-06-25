package gocoder

import "fmt"

type burger struct {
	P, B, Sum int
}

func count(N, X, tmp int, L []burger) int {
	if N == 0 {
		return tmp + 1
	}

	if X == 1 {
		return tmp
	} else if X < (L[N].Sum/2 + 1) {
		return count(N-1, X-1, tmp, L)
	} else if X == (L[N].Sum/2 + 1) {
		return tmp + L[N-1].P + 1
	} else if X < L[N].Sum {
		return count(N-1, X-L[N-1].Sum-2, tmp+L[N-1].P+1, L)
	}
	return tmp + L[N].P

}

func main() {
	var N, X int
	fmt.Scan(&N, &X)

	L := make([]burger, N+1)
	L[0].P, L[0].B, L[0].Sum = 1, 0, 1
	for i := 1; i <= N; i++ {
		L[i].P, L[i].B, L[i].Sum = 2*L[i-1].P+1, 2*L[i-1].B+2, 2*L[i-1].Sum+3
	}

	ans := count(N, X, 0, L)
	fmt.Println(ans)
}
