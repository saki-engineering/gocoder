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

func removeElement(L []int, i int) []int {
	ans := make([]int, i)
	copy(ans, L[:i])
	ans = append(ans, L[i+1:]...)
	return ans
}

func permutation(L []int) (ans [][]int) {
	if N := len(L); N == 1 {
		ans = append(ans, L)
		return
	}

	for i, l := range L {
		for _, R := range permutation(removeElement(L, i)) {
			elm := []int{l}
			elm = append(elm, R...)
			ans = append(ans, elm)
		}
	}
	return
}

func intMin(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

func main() {
	sc.Split(bufio.ScanWords)

	INF := int(math.Pow10(10))
	N, M, R := scanInt(), scanInt(), scanInt()

	r := make([]int, R)
	for i := 0; i < R; i++ {
		r[i] = scanInt() - 1
	}

	G := make([][]int, N)
	for i := 0; i < N; i++ {
		G[i] = make([]int, N)
		for j := 0; j < N; j++ {
			G[i][j] = INF
		}
		G[i][i] = 0
	}

	for i := 0; i < M; i++ {
		A, B, C := scanInt()-1, scanInt()-1, scanInt()
		G[A][B], G[B][A] = C, C
	}

	for k := 0; k < N; k++ {
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				G[i][j] = intMin(G[i][j], G[i][k]+G[k][j])
			}
		}
	}

	ans := INF
	for _, order := range permutation(r) {
		tmp := 0
		for i := 0; i < R-1; i++ {
			tmp += G[order[i]][order[i+1]]
		}
		ans = intMin(ans, tmp)
	}
	fmt.Println(ans)
}
