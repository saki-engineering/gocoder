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

type edge struct {
	to, weight int
}

type node struct {
	dist     int
	visited  bool
	edgeList []edge
}

type graph []node

func DFS(G *graph, K int) {
	(*G)[K].visited = true
	for _, ed := range (*G)[K].edgeList {
		if !(*G)[ed.to].visited {
			(*G)[ed.to].dist = (*G)[K].dist + ed.weight
			DFS(G, ed.to)
		}
	}
}

func main() {
	sc.Split(bufio.ScanWords)

	N := scanInt()
	G := make(graph, N)
	for i := 0; i < N; i++ {
		G[i].edgeList = make([]edge, 0)
	}

	for i := 0; i < N-1; i++ {
		a, b, c := scanInt()-1, scanInt()-1, scanInt()
		G[a].edgeList = append(G[a].edgeList, edge{to: b, weight: c})
		G[b].edgeList = append(G[b].edgeList, edge{to: a, weight: c})
	}

	Q, K := scanInt(), scanInt()-1
	DFS(&G, K)

	for i := 0; i < Q; i++ {
		x, y := scanInt()-1, scanInt()-1
		ans := G[x].dist + G[y].dist
		fmt.Println(ans)
	}
}
