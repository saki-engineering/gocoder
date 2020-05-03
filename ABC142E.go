package gocoder

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

type key struct {
	price int
	box   int
}

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

func abc142e() {
	INF := int(math.Pow10(9))
	sc.Split(bufio.ScanWords)

	N, M := uint(scanInt()), scanInt()
	K := make([]key, M)
	for i := 0; i < M; i++ {
		a, b := scanInt(), scanInt()

		K[i].price = a
		K[i].box = 0
		for j := 0; j < b; j++ {
			K[i].box += 1 << (uint(scanInt() - 1))
		}
	}

	// DP[i]...ビットマスクiの箱の鍵が開けられる最安値
	DP := make([]int, 1<<N)
	DP[0] = 0
	for i := 1; i < (1 << N); i++ {
		DP[i] = INF
	}

	for _, key := range K {
		for i := 0; i < (1 << N); i++ {
			DP[i|key.box] = intMin(DP[i|key.box], DP[i]+key.price)
		}
	}

	if ans := DP[(1<<N)-1]; ans == INF {
		fmt.Println(-1)
	} else {
		fmt.Println(ans)
	}
}
