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

	N := scanInt()
	F := make([][10]int, N)
	for i := 0; i < N; i++ {
		for j := 0; j < 10; j++ {
			F[i][j] = scanInt()
		}
	}
	P := make([][11]int, N)
	for i := 0; i < N; i++ {
		for j := 0; j <= 10; j++ {
			P[i][j] = scanInt()
		}
	}

	ans := -int(math.Pow10(9))
	for i := 1; i < (1 << 10); i++ {
		tmp := 0
		for j := 0; j < N; j++ {
			cnt := 0
			for k := 0; k < 10; k++ {
				if F[j][k] == 1 && i&(1<<k) > 0 {
					cnt++
				}
			}
			tmp += P[j][cnt]
		}
		ans = intMax(ans, tmp)
	}

	fmt.Println(ans)
}
