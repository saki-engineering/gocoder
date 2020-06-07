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

func intAbs(x int) int {
	return int(math.Abs(float64(x)))
}

func intMin(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

func main() {
	sc.Split(bufio.ScanWords)

	N := scanInt()
	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = scanInt()
	}

	cntPlus, cntZero, cntMinus, absMin := 0, 0, 0, int(math.Pow10(10))
	for i := 0; i < N; i++ {
		if A[i] > 0 {
			cntPlus++
			absMin = intMin(absMin, intAbs(A[i]))
		} else if A[i] == 0 {
			cntZero++
		} else {
			cntMinus++
			absMin = intMin(absMin, intAbs(A[i]))
		}
	}

	ans := 0
	for _, a := range A {
		ans += intAbs(a)
	}
	if cntZero == 0 && cntMinus%2 == 1 {
		ans -= 2 * absMin
	}

	fmt.Println(ans)
}
