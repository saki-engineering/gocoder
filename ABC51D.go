package gocoder

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

func intMin(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

type edge struct {
	a, b, c int
}

func main() {
	sc.Split(bufio.ScanWords)
	INF := 1 << 20

	N, M := scanInt(), scanInt()
	G := make([][]int, N)
	for i := 0; i < N; i++ {
		G[i] = make([]int, N)
		for j := 0; j < N; j++ {
			G[i][j] = INF
		}
		G[i][i] = 0
	}

	L := make([]edge, M)
	for i := range L {
		a, b, c := scanInt()-1, scanInt()-1, scanInt()

		L[i].a, L[i].b, L[i].c = a, b, c
		G[a][b], G[b][a] = c, c
	}

	for k := 0; k < N; k++ {
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				G[i][j] = intMin(G[i][j], G[i][k]+G[k][j])
			}
		}
	}

	ans := 0
	for _, ed := range L {
		if G[ed.a][ed.b] < ed.c {
			ans++
		}
	}
	fmt.Println(ans)
}
