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

func nextInt() int {
	sc.Scan()
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func abc144e() {
	sc.Split(bufio.ScanWords)

	N, K := nextInt(), nextInt()

	A, F := make([]int, N), make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = nextInt()
	}
	sort.Ints(A)

	for i := 0; i < N; i++ {
		F[i] = nextInt()
	}
	sort.Sort(sort.Reverse(sort.IntSlice(F)))

	l, r := 0, int(math.Pow(10, 12))
	for l < r {
		mid := l + (r-l)/2
		trning := 0

		for i := 0; i < N; i++ {
			if A[i]*F[i] > mid {
				trning += A[i] - (mid / F[i])
			}
		}

		if K < trning {
			l = mid + 1
		} else {
			r = mid
		}
	}

	fmt.Println(l)
}
