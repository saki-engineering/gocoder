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

func gcd(x, y int) int {
	if x < y {
		x, y = y, x
	}

	if y == 0 {
		return x
	}
	return gcd(y, x%y)
}

func lcm(x, y int) int {
	return x * (y / gcd(x, y))
}

func main() {
	sc.Split(bufio.ScanWords)

	N := scanInt()
	T := make([]int, N)
	for i := 0; i < N; i++ {
		T[i] = scanInt()
	}

	ans := T[0]
	for _, t := range T {
		ans = lcm(ans, t)
	}
	fmt.Println(ans)
}
