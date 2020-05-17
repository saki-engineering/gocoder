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
	A, B := make([]int, N+1), make([]int, N)
	for i := 0; i <= N; i++ {
		A[i] = scanInt()
	}
	for i := 0; i < N; i++ {
		B[i] = scanInt()
	}

	ans := 0
	for i := 0; i < N; i++ {
		if A[i]+A[i+1] >= B[i] {
			ans += B[i]
			A[i+1] -= intMax(B[i]-A[i], 0)
		} else {
			ans += (A[i] + A[i+1])
			A[i+1] = 0
		}
	}
	fmt.Println(ans)
}
