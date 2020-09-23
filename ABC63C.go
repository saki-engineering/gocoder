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
	s := make([]int, N)
	SUM := 0
	for i := 0; i < N; i++ {
		s[i] = scanInt()
		SUM += s[i]
	}

	ans := SUM
	if SUM%10 == 0 {
		ans = 0
	}
	for _, point := range s {
		if tmp := (SUM - point); tmp%10 != 0 {
			ans = intMax(ans, tmp)
		}
	}

	fmt.Println(ans)
}
