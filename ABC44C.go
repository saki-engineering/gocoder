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

	N, A := scanInt(), scanInt()
	X := make([]int, N)
	for i := range X {
		X[i] = scanInt()
	}

	DP, M := make([][][]int, N+1), 50*N
	for i := range DP {
		DP[i] = make([][]int, N+1)
		for j := range DP[i] {
			DP[i][j] = make([]int, M+1)
		}
	}
	DP[0][0][0] = 1

	for i, x := range X {
		for j := 0; j <= N; j++ {
			for k := 0; k <= M; k++ {
				DP[i+1][j][k] += DP[i][j][k]
				if j+1 <= N && k+x <= M {
					DP[i+1][j+1][k+x] += DP[i][j][k]
				}
			}
		}
	}

	ans := 0
	for i := 1; i <= N; i++ {
		ans += DP[N][i][A*i]
	}
	fmt.Println(ans)
}
