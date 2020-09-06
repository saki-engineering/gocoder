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

	MAX := int(math.Pow10(5))
	N := scanInt()

	A := make([]int, MAX+2)
	for i := 0; i < N; i++ {
		a := scanInt()
		A[a]++
	}

	ans := 0
	for i := 1; i <= MAX; i++ {
		ans = intMax(ans, A[i-1]+A[i]+A[i+1])
	}
	fmt.Println(ans)
}
