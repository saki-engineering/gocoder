package gocoder

import (
	"bufio"
	"fmt"
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

func gcd(a, b int) int {
	for b > 0 {
		a, b = b, a%b
	}
	return a
}

func main() {
	sc.Split(bufio.ScanWords)

	N := scanInt()
	X := make([]int, N+1)
	for i := 0; i <= N; i++ {
		X[i] = scanInt()
	}
	sort.Ints(X)

	ans := X[1] - X[0]
	for i := 1; i < N; i++ {
		ans = gcd(ans, X[i+1]-X[i])
	}
	fmt.Println(ans)
}
