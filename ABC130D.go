package gocoder

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func scanInt() int {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return i
}

func main() {
	sc.Split(bufio.ScanWords)

	N, K := scanInt(), scanInt()
	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = scanInt()
	}

	l, r, ans, S := 0, 1, 0, A[0]
	for {
		if S >= K {
			ans += (N - r + 1)
			S -= A[l]
			l++
		} else {
			if r == N {
				break
			}
			S += A[r]
			r++
		}
	}
	fmt.Println(ans)
}
