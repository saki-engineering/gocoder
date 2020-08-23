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

func intMax(x, y int) int {
	return int(math.Max(float64(x), float64(y)))
}

func main() {
	sc.Split(bufio.ScanWords)

	N, C := scanInt(), scanInt()
	T := int(math.Pow10(5) * 2)

	table := make([][]int, C)
	for i := 0; i < C; i++ {
		table[i] = make([]int, T+1)
	}

	for i := 0; i < N; i++ {
		s, t, c := scanInt()*2-1, scanInt()*2, scanInt()-1
		for j := s; j <= t; j++ {
			table[c][j] = 1
		}
	}

	ans := 0
	for i := 0; i <= T; i++ {
		tmp := 0
		for j := 0; j < C; j++ {
			tmp += table[j][i]
		}
		ans = intMax(ans, tmp)
	}
	fmt.Println(ans)
}
