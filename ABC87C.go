package gocoder

import (
	"bufio"
	"fmt"
	"math"
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

func intMax(x, y int) int {
	return int(math.Max(float64(x), float64(y)))
}

func main() {
	sc.Split(bufio.ScanWords)

	N := scanInt()
	A := [2][]int{}
	A[0], A[1] = make([]int, N), make([]int, N)
	for i := 0; i < 2; i++ {
		for j := 0; j < N; j++ {
			A[i][j] = scanInt()
		}
	}
	for i := 1; i < N; i++ {
		A[0][i] += A[0][i-1]
	}
	for i := N - 2; i >= 0; i-- {
		A[1][i] += A[1][i+1]
	}

	ans := 0
	for i := 0; i < N; i++ {
		ans = intMax(ans, A[0][i]+A[1][i])
	}

	fmt.Println(ans)
}
