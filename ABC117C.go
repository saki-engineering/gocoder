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

	N, M := scanInt(), scanInt()
	X := make([]int, M)
	for i := 0; i < M; i++ {
		X[i] = scanInt()
	}
	sort.Ints(X)

	Y := make([]int, M-1)
	for i := 0; i < M-1; i++ {
		Y[i] = X[i+1] - X[i]
	}
	sort.Ints(Y)

	ans := 0
	for i := 0; i < M-N; i++ {
		ans += Y[i]
	}

	fmt.Println(ans)
}
