package gocoder

import (
	"bufio"
	"fmt"
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

func main() {
	sc.Split(bufio.ScanWords)

	N, M := scanInt(), scanInt()

	G := make([][]bool, N)
	for i := 0; i < N; i++ {
		G[i] = make([]bool, N)
	}
	for i := 0; i < M; i++ {
		a, b := scanInt()-1, scanInt()-1
		G[a][b], G[b][a] = true, true
	}

	L := make([]int, N-1)
	for i := 0; i < N-1; i++ {
		L[i] = i + 1
	}

	ans := 0
	for _, list := range permutation(L) {
		flg := true
		var a, b int
		a = 0
		for _, node := range list {
			b = node
			if !G[a][b] {
				flg = false
				break
			}
			a = b
		}

		if flg {
			ans++
		}
	}
	fmt.Println(ans)
}
