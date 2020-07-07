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

func intMax(x, y int) int {
	return int(math.Max(float64(x), float64(y)))
}

func intAbs(x int) int {
	return int(math.Abs(float64(x)))
}

func main() {
	sc.Split(bufio.ScanWords)

	N, K := scanInt(), scanInt()
	X := make([]int, N)
	for i := 0; i < N; i++ {
		X[i] = scanInt()
	}

	ans := int(math.Pow10(9))
	for i := K - 1; i < N; i++ {
		if X[i]*X[i+1-K] >= 0 {
			ans = intMin(ans, intMax(intAbs(X[i]), intAbs(X[i+1-K])))
		} else {
			ans = intMin(ans, X[i]*2-X[i+1-K])
			ans = intMin(ans, X[i+1-K]*(-2)+X[i])
		}
	}
	fmt.Println(ans)
}
