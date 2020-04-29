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
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

func abc143e() {
	sc.Split(bufio.ScanWords)

	INF := int(math.Pow10(10))
	N, M, L := scanInt(), scanInt(), scanInt()

	d := make([][]int, N)
	for i := 0; i < N; i++ {
		d[i] = make([]int, N)
		for j := 0; j < N; j++ {
			if i == j {
				d[i][j] = 0
			} else {
				d[i][j] = INF
			}
		}
	}

	for i := 0; i < M; i++ {
		a, b, c := scanInt(), scanInt(), scanInt()
		if c <= L {
			d[a-1][b-1] = c
			d[b-1][a-1] = c
		}
	}

	for k := 0; k < N; k++ {
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				d[i][j] = int(math.Min(float64(d[i][j]), float64(d[i][k]+d[k][j])))
			}
		}
	}

	for i := 0; i < N; i++ {
		for j := 0; j < N; j++ {
			if d[i][j] == 0 {
				d[i][j] = 0
			} else if d[i][j] <= L {
				d[i][j] = 1
			} else {
				d[i][j] = INF
			}
		}
	}

	for k := 0; k < N; k++ {
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				d[i][j] = int(math.Min(float64(d[i][j]), float64(d[i][k]+d[k][j])))
			}
		}
	}

	Q := scanInt()
	ans := make([]int, Q)
	for i := 0; i < Q; i++ {
		s, t := scanInt(), scanInt()
		if dist := d[s-1][t-1]; dist >= INF {
			ans[i] = 0
		} else {
			ans[i] = dist
		}
	}

	for _, a := range ans {
		fmt.Println(a - 1)
	}
}
