package gocoder

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type dpFact struct {
	digit, from, num int
}

func main() {
	match := [...]int{0, 2, 5, 5, 4, 5, 6, 3, 7, 6}
	var N, M int
	fmt.Scan(&N, &M)

	A := make([]int, M)
	for i := 0; i < M; i++ {
		fmt.Scan(&A[i])
	}
	sort.Ints(A)

	DP := make([]dpFact, N+1)
	for i := 1; i <= N; i++ {
		DP[i].digit, DP[i].from, DP[i].num = -1, -1, -1
	}

	for i := 0; i <= N; i++ {
		for _, a := range A {
			if index := i - match[a]; index >= 0 && DP[index].digit >= 0 {
				if DP[index].digit+1 >= DP[i].digit {
					DP[i].digit = DP[index].digit + 1
					DP[i].from = index
					DP[i].num = a
				}
			}
		}
	}

	ansList := [10]int{}
	t := N
	for t > 0 {
		ansList[DP[t].num]++
		t = DP[t].from
	}

	ans := ""
	for i := 9; i > 0; i-- {
		if ansList[i] > 0 {
			ans = ans + strings.Repeat(strconv.Itoa(i), ansList[i])
		}
	}
	fmt.Println(ans)
}
