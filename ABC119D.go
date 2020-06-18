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

func scanInt() int {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return i
}

func intMin(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

func intAbs(x int) int {
	return int(math.Abs(float64(x)))
}

func main() {
	sc.Split(bufio.ScanWords)

	INF := int(math.Pow10(11))
	A, B, Q := scanInt(), scanInt(), scanInt()
	L, X := make([][]int, 2), make([]int, Q)

	L[0] = make([]int, A+2)
	L[0][0], L[0][A+1] = -INF, INF
	for i := 1; i <= A; i++ {
		L[0][i] = scanInt()
	}
	L[1] = make([]int, B+2)
	L[1][0], L[1][B+1] = -INF, INF
	for i := 1; i <= B; i++ {
		L[1][i] = scanInt()
	}
	for i := 0; i < Q; i++ {
		X[i] = scanInt()
	}

	for _, x := range X {
		index := make([]int, 2)
		index[0] = sort.Search(len(L[0]), func(i int) bool { return L[0][i] >= x })
		index[1] = sort.Search(len(L[1]), func(i int) bool { return L[1][i] >= x })

		ans := 4 * INF
		for i := 0; i < 2; i++ {
			for j := 0; j < 2; j++ {
				for k := 0; k < 2; k++ {
					ans = intMin(ans, intAbs(x-L[k][index[k]-i])+intAbs(L[k][index[k]-i]-L[1-k][index[1-k]-j]))
				}
			}
		}

		fmt.Println(ans)
	}
}
