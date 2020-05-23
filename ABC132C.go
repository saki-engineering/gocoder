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

func main() {
	sc.Split(bufio.ScanWords)

	N := scanInt()
	d := make([]int, N)
	for i := 0; i < N; i++ {
		d[i] = scanInt()
	}
	sort.Ints(d)

	fmt.Println(d[N/2] - d[N/2-1])
}
