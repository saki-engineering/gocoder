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

type station struct {
	C, S, F int
}

func main() {
	sc.Split(bufio.ScanWords)

	N := scanInt()
	L := make([]station, N-1)
	for i := 0; i < N-1; i++ {
		L[i].C, L[i].S, L[i].F = scanInt(), scanInt(), scanInt()
	}

	for i := 0; i < N; i++ {
		ans := 0
		for _, l := range L[i:] {
			if ans < l.S {
				ans = l.S + l.C
			} else {
				q := ans / l.F
				if r := ans % l.F; r > 0 {
					q++
				}
				ans = q*l.F + l.C
			}
		}
		fmt.Println(ans)
	}
}
