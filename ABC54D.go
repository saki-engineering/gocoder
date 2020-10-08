package main

import (
	"bufio"
	"fmt"
	"math"
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

type med struct {
	a, b, c int
}

func intMin(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

func main() {
	sc.Split(bufio.ScanWords)
	M, INF := 400, 1<<20

	N, Ma, Mb := scanInt(), scanInt(), scanInt()
	L := make([]med, N)
	for i := 0; i < N; i++ {
		L[i].a, L[i].b, L[i].c = scanInt(), scanInt(), scanInt()
	}

	DP := make([][]int, M+1)
	for i := 0; i <= M; i++ {
		DP[i] = make([]int, M+1)
		for j := 0; j <= M; j++ {
			DP[i][j] = INF
		}
	}
	DP[0][0] = 0

	for _, l := range L {
		for i := M; i >= l.a; i-- {
			for j := M; j >= l.b; j-- {
				if DP[i-l.a][j-l.b] == INF {
					continue
				}
				DP[i][j] = intMin(DP[i][j], DP[i-l.a][j-l.b]+l.c)
			}
		}
	}

	ans, k := INF, 1
	for Ma*k <= M && Mb*k <= M {
		ans = intMin(ans, DP[Ma*k][Mb*k])
		k++
	}

	if ans == INF {
		fmt.Println(-1)
	} else {
		fmt.Println(ans)
	}
}
