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

func main() {
	sc.Split(bufio.ScanWords)

	N := scanInt()
	p := make([]int, N+1)
	for i := 1; i <= N; i++ {
		p[i] = scanInt()
	}

	ans := 0
	for i := 1; i < N; i++ {
		if p[i] == i {
			ans++
			p[i], p[i+1] = p[i+1], p[i]
		}
	}
	if p[N] == N {
		ans++
	}

	fmt.Println(ans)
}
