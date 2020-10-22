package gocoder

import (
	"bufio"
	"fmt"
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
	if x > y {
		return x
	}
	return y
}

func divOver(p, q int) int {
	ans := p / q
	if p%q > 0 {
		ans++
	}
	return ans
}

func main() {
	sc.Split(bufio.ScanWords)

	N := scanInt()

	X, Y := 1, 1
	for i := 0; i < N; i++ {
		T, A := scanInt(), scanInt()

		n := intMax(divOver(X, T), divOver(Y, A))
		X, Y = n*T, n*A
	}
	fmt.Println(X + Y)
}
