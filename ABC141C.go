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

func abc141c() {
	sc.Split(bufio.ScanWords)

	N, K, Q := scanInt(), scanInt(), scanInt()
	score := make([]int, N)
	for i := 0; i < Q; i++ {
		A := scanInt()
		score[A-1]++
	}

	for _, s := range score {
		if K-Q+s > 0 {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}
	}
}
