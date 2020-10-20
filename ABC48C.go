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

	N, x := scanInt(), scanInt()
	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = scanInt()
	}

	ans := 0
	for i := 1; i < N; i++ {
		if diff := A[i-1] + A[i] - x; diff > 0 {
			ans += diff
			A[i] = intMax(A[i]-diff, 0)
		}
	}
	fmt.Println(ans)
}
