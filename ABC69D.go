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

func main() {
	sc.Split(bufio.ScanWords)

	H, W, N := scanInt(), scanInt(), scanInt()
	MAP := make([][]int, H)
	for i := 0; i < H; i++ {
		MAP[i] = make([]int, W)
	}

	cnt := 0
	for i := 1; i <= N; i++ {
		a := scanInt()
		for j := 0; j < a; j++ {
			h, w := cnt/W, cnt%W
			MAP[h][w] = i
			cnt++
		}
	}

	for i := 0; i < H; i++ {
		if i%2 == 0 {
			for j := 0; j < W; j++ {
				fmt.Printf("%d ", MAP[i][j])
			}
		} else {
			for j := W - 1; j >= 0; j-- {
				fmt.Printf("%d ", MAP[i][j])
			}
		}
		fmt.Printf("\n")
	}
}
