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
	A = append(A, 0)

	ans := 0
	cnt := 1
	for i := 1; i <= N; i++ {
		if A[i] == A[i-1] {
			cnt++
		} else {
			if A[i-1] > cnt {
				ans += cnt
			} else {
				ans += (cnt - A[i-1])
			}
			cnt = 1
		}
	}

	fmt.Println(ans)
}
