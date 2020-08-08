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

type grid struct {
	x, y int
}

func MP(bfe, aft grid) int {
	ans := 0
	ans += int(math.Abs(float64(bfe.x - aft.x)))
	ans += int(math.Abs(float64(bfe.y - aft.y)))
	return ans
}

func main() {
	sc.Split(bufio.ScanWords)

	H, W, D := scanInt(), scanInt(), scanInt()
	G := make([]grid, H*W+1)
	for i := 1; i <= H; i++ {
		for j := 1; j <= W; j++ {
			a := scanInt()
			G[a].x, G[a].y = i, j
		}
	}

	RS := make([]int, H*W+1)
	for i := D + 1; i <= H*W; i++ {
		RS[i] = RS[i-D] + MP(G[i-D], G[i])
	}

	Q := scanInt()
	for i := 0; i < Q; i++ {
		L, R := scanInt(), scanInt()
		fmt.Println(RS[R] - RS[L])
	}
}
