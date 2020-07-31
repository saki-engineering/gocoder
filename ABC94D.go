package gocoder

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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
	sort.Ints(A)

	index, abs := 0, intAbs((A[N-1]/2)-A[0])
	if A[N-1]%2 == 1 {
		abs = intMin(intAbs((A[N-1]/2+1)-A[0]), abs)
	}
	for i, a := range A {
		if b := intAbs((A[N-1] / 2) - a); b < abs {
			abs = b
			index = i
		}
		if A[N-1]%2 == 1 {
			if b := intAbs((A[N-1] / 2) + 1 - a); b < abs {
				abs = b
				index = i
			}
		}
	}
	fmt.Println(A[N-1], A[index])
}
