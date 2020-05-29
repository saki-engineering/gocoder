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
	MOD := int(math.Pow10(9) + 7)

	N, M := scanInt(), scanInt()
	A := make([]int, N+1)
	A[0], A[1] = 1, 1
	for i := 0; i < M; i++ {
		m := scanInt()
		A[m] = -1
	}

	for i := 2; i <= N; i++ {
		if A[i] != -1 {
			A[i] = intMax(A[i-2], 0) + intMax(A[i-1], 0)
			A[i] %= MOD
		}
	}
	fmt.Println(A[N])
}
