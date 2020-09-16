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

func intMin(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

func intAbs(x int) int {
	return int(math.Abs(float64(x)))
}

func main() {
	sc.Split(bufio.ScanWords)

	N := scanInt()
	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = scanInt()
	}

	S := 0
	for _, a := range A {
		S += a
	}

	for i := 1; i < N; i++ {
		A[i] += A[i-1]
	}

	ans := intAbs(S - 2*A[0])
	for i := 1; i < N-1; i++ {
		ans = intMin(ans, intAbs(S-2*A[i]))
	}
	fmt.Println(ans)
}
