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
	A := make([]int, N+1)
	for i := 1; i <= N; i++ {
		A[i] = scanInt()
		A[i] = (A[i] + A[i-1]) % M
	}
	sort.Ints(A)

	ans, cnt := 0, 1
	for i := 1; i <= N; i++ {
		if A[i] == A[i-1] {
			cnt++
		} else {
			ans += ((cnt) * (cnt - 1)) / 2
			cnt = 1
		}
	}
	ans += ((cnt) * (cnt - 1)) / 2
	fmt.Println(ans)
}
