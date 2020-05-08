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
	H := make([]int, N)
	for i := 0; i < N; i++ {
		H[i] = scanInt()
	}

	ans := 0
	tmp := 0
	for i := 0; i < N-1; i++ {
		if H[i] >= H[i+1] {
			tmp++
		} else {
			ans = intMax(ans, tmp)
			tmp = 0
		}
	}
	ans = intMax(ans, tmp)

	fmt.Println(ans)
}
