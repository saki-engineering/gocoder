package gocoder

import (
	"bufio"
	"fmt"
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

func main() {
	sc.Split(bufio.ScanWords)

	N := scanInt()
	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = scanInt()
	}
	sort.Ints(A)

	a, b := 0, 0
	for i := N - 1; i > 0; i-- {
		if A[i] == A[i-1] {
			if a == 0 {
				a = A[i]
			} else {
				b = A[i]
				break
			}
			i--
		}
	}
	fmt.Println(a * b)
}
