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

	N, K := scanInt(), scanInt()
	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = scanInt()
	}
	sort.Ints(A)
	A = append(A, 0)

	cntList, cnt := make([]int, 0), 1
	for i := 1; i <= N; i++ {
		if A[i] == A[i-1] {
			cnt++
		} else {
			cntList = append(cntList, cnt)
			cnt = 1
		}
	}

	sort.Ints(cntList)
	ans := 0
	M := len(cntList)
	for i := 0; i < M-K; i++ {
		ans += cntList[i]
	}
	fmt.Println(ans)
}
