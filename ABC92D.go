package gocoder

import (
	"fmt"
	"strings"
)

func main() {
	var A, B int
	fmt.Scan(&A, &B)

	N := 50
	h, w := 2*N, 2*N

	Map := make([]string, h)
	for i := 0; i < N; i++ {
		Map[i] = strings.Repeat("#%", N)
		Map[h-1-i] = strings.Repeat(".%", N)
	}

	cnt := 0
	for A-1 > 0 {
		if A-1 >= N {
			Map[cnt*2] = strings.Replace(Map[cnt*2], "%", ".", N)
			A -= N
		} else {
			Map[cnt*2] = strings.Replace(Map[cnt*2], "%", ".", A-1)
			A -= (A - 1)
		}
		cnt++
	}
	for i := 0; i < N; i++ {
		Map[i] = strings.ReplaceAll(Map[i], "%", "#")
	}

	cnt = 0
	for B-1 > 0 {
		if B-1 >= N {
			Map[h-1-cnt*2] = strings.Replace(Map[h-1-cnt*2], "%", "#", N)
			B -= N
		} else {
			Map[h-1-cnt*2] = strings.Replace(Map[h-1-cnt*2], "%", "#", B-1)
			B -= (B - 1)
		}
		cnt++
	}
	for i := 0; i < N; i++ {
		Map[h-1-i] = strings.ReplaceAll(Map[h-1-i], "%", ".")
	}

	fmt.Println(h, w)
	for i := 0; i < h; i++ {
		fmt.Println(Map[i])
	}
}
