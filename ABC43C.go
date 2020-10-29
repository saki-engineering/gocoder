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
	A, sum := make([]int, N), 0
	for i := range A {
		A[i] = scanInt()
		sum += A[i]
	}

	ave1, ave2 := sum/N, sum/N+1
	cost1, cost2 := 0, 0
	for _, a := range A {
		cost1 += (a - ave1) * (a - ave1)
		cost2 += (a - ave2) * (a - ave2)
	}

	if cost1 < cost2 {
		fmt.Println(cost1)
	} else {
		fmt.Println(cost2)
	}
}
