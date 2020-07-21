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

	N, C := scanInt(), scanInt()
	D := make([][]int, C)
	for i := 0; i < C; i++ {
		D[i] = make([]int, C)
		for j := 0; j < C; j++ {
			D[i][j] = scanInt()
		}
	}

	Map := make([][3]int, C)
	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			color := scanInt() - 1
			Map[color][(i+j)%3]++
		}
	}

	ans := int(math.Pow10(10))
	for a := 0; a < C; a++ {
		for b := 0; b < C; b++ {
			for c := 0; c < C; c++ {
				if (a-b)*(b-c)*(c-a) == 0 {
					continue
				}

				color := [3]int{a, b, c}
				tmp := 0
				for i := 0; i < 3; i++ {
					for j := 0; j < C; j++ {
						tmp += D[j][color[i]] * Map[j][i]
					}
				}
				ans = intMin(ans, tmp)
			}
		}
	}
	fmt.Println(ans)
}
