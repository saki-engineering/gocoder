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

	max := 0
	max2 := 0
	for _, a := range A {
		if a >= max {
			max2 = max
			max = a
		} else if a >= max2 {
			max2 = a
		}
	}

	for _, a := range A {
		if a == max {
			fmt.Println(max2)
		} else {
			fmt.Println(max)
		}
	}
}
