package gocoder

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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

func intAbs(x int) int {
	return int(math.Abs(float64(x)))
}

func main() {
	sc.Split(bufio.ScanWords)

	N := scanInt()
	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = scanInt() - (i + 1)
	}
	sort.Ints(A)

	var b int
	if N%2 == 1 {
		b = A[(N-1)/2]
	} else {
		b = (A[N/2] + A[N/2-1]) / 2
	}
	ans := []int{0, 0}
	for _, a := range A {
		ans[0] += intAbs(a - b)
		ans[1] += intAbs(a - b - 1)
	}
	sort.Ints(ans)

	fmt.Println(ans[0])
}
