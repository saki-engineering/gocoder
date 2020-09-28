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

type edge struct {
	a, b, c int
}

func intMax(x, y int) int {
	return int(math.Max(float64(x), float64(y)))
}

func main() {
	sc.Split(bufio.ScanWords)
	INF := int(math.Pow10(15))

	N, M := scanInt(), scanInt()
	E := make([]edge, M)
	for i := 0; i < M; i++ {
		E[i].a, E[i].b, E[i].c = scanInt()-1, scanInt()-1, scanInt()
	}

	dist := make([]int, N)
	for i := 1; i < N; i++ {
		dist[i] = -INF
	}

	for i := 0; i < N; i++ {
		for _, e := range E {
			if e.a != -INF {
				dist[e.b] = intMax(dist[e.b], dist[e.a]+e.c)
			}
		}
	}

	updated := make([]bool, N)
	for i := 0; i < N; i++ {
		for _, e := range E {
			if e.a == -INF {
				continue
			}
			if dist[e.b] < dist[e.a]+e.c {
				dist[e.b] = dist[e.a] + e.c
				updated[e.b] = true
			}

			if updated[e.a] {
				updated[e.b] = true
			}
		}
	}

	if updated[N-1] {
		fmt.Println("inf")
	} else {
		fmt.Println(dist[N-1])
	}
}
