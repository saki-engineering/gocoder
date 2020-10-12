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
	for i := range A {
		A[i] = scanInt()
	}
	sort.Ints(A)

	cnt := 1
	for i := 1; i < N; i++ {
		if A[i-1] != A[i] {
			cnt++
		}
	}

	if cnt%2 == 0 {
		cnt--
	}
	fmt.Println(cnt)
}
