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

	N := scanInt()
	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = scanInt()
	}

	ans := make([]int, N)
	for i := 0; i < N; i++ {
		if i%2 == 0 {
			ans[0] += A[i]
		} else {
			ans[0] -= A[i]
		}
	}
	for i := 1; i < N; i++ {
		ans[i] = 2 * (A[i-1] - ans[i-1]/2)
	}

	for i := 0; i < N; i++ {
		fmt.Printf("%d ", ans[i])
	}
	fmt.Printf("\n")
}
