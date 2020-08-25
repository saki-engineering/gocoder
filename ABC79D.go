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

func main() {
	sc.Split(bufio.ScanWords)

	H, W := scanInt(), scanInt()

	C := [10][10]int{}
	for i := 0; i < 10; i++ {
		for j := 0; j < 10; j++ {
			C[i][j] = scanInt()
		}
	}

	Map := [10][100]int{}
	for i := 0; i < 10; i++ {
		Map[i][1] = C[i][1]
	}

	INF := int(math.Pow10(10))
	for j := 2; j < 100; j++ {
		for i := 0; i < 10; i++ {
			Map[i][j] = INF
			for k := 0; k < 10; k++ {
				Map[i][j] = intMin(Map[i][j], C[i][k]+Map[k][j-1])
			}
		}
	}

	A := [10]int{}
	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			a := scanInt()
			if a != -1 {
				A[a]++
			}
		}
	}

	ans := 0
	for i := 0; i < 10; i++ {
		ans += A[i] * Map[i][99]
	}
	fmt.Println(ans)
}
