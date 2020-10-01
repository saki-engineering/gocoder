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

	N := scanInt()
	A := make([]int, N)
	for i := range A {
		A[i] = scanInt()
	}

	ans := int(math.Pow10(14))
	for _, t := range []int{1, -1} {
		r, cnt, T := 0, 0, t
		for _, a := range A {
			r += a
			if T > 0 && r <= 0 {
				cnt += 1 - r
				r += 1 - r
			}
			if T < 0 && r >= 0 {
				cnt += r + 1
				r -= (r + 1)
			}
			T *= -1
		}
		ans = intMin(ans, cnt)
	}
	fmt.Println(ans)
}
