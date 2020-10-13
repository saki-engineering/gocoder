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

func main() {
	sc.Split(bufio.ScanWords)

	N, A, B := scanInt(), scanInt(), scanInt()
	X := make([]int, N)
	for i := range X {
		X[i] = scanInt()
	}

	ans := 0
	for i := 1; i < N; i++ {
		ans += intMin((X[i]-X[i-1])*A, B)
	}
	fmt.Println(ans)
}
