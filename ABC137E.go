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
	to, weight int
}

type node struct {
	edgeList []edge
	dist     int
}

func main() {
	sc.Split(bufio.ScanWords)
	INF := int(math.Pow10(10))

	N, M, P := scanInt(), scanInt(), scanInt()
	G := make([]node, N)
	for i := 0; i < N; i++ {
		G[i].dist = -INF
	}
	G[0].dist = 0

	for i := 0; i < M; i++ {
		a, b, c := scanInt(), scanInt(), scanInt()
		G[a-1].edgeList = append(G[a-1].edgeList, edge{to: b - 1, weight: c - P})
	}

	// N回きっかりで止めてしまうと、1→Nまでの道のりのどこかで更新されたときを弾いてしまう
	for i := 1; i <= 2*N; i++ {
		updated := false
		for j := 0; j < N; j++ {
			if G[j].dist == -INF {
				continue
			}
			for _, e := range G[j].edgeList {
				if newWeight := G[j].dist + e.weight; G[e.to].dist < newWeight {
					G[e.to].dist = newWeight
					// このif文がないと、after_contestのテストケースが通らない
					// 多分、2*N回程度のループでは頂点Nが更新されないけど、もっと回せば更新されるケースがある
					if i >= N {
						G[e.to].dist = INF
					}
					if e.to == N-1 {
						updated = true
					}
				}
			}
		}
		if updated && i >= N {
			fmt.Println(-1)
			return
		}
	}

	if ans := G[N-1].dist; ans < 0 {
		fmt.Println(0)
	} else {
		fmt.Println(ans)
	}
}
