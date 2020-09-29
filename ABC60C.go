package main

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

	N, T := scanInt(), scanInt()
	t := make([]int, N)
	for i := range t {
		t[i] = scanInt()
	}

	ans := T
	for i := 1; i < N; i++ {
		ans += intMin(t[i]-t[i-1], T)
	}
	fmt.Println(ans)
}
