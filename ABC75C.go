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
	a, b int
}

type graph [][]bool

func deleteEdge(G *graph, e edge) {
	(*G)[e.a][e.b], (*G)[e.b][e.a] = false, false
}

func addEdge(G *graph, e edge) {
	(*G)[e.a][e.b], (*G)[e.b][e.a] = true, true
}

func DFS(G *graph, n int, d *[]bool) {
	(*d)[n] = true
	for i, isLinked := range (*G)[n] {
		if isLinked && !(*d)[i] {
			DFS(G, i, d)
		}
	}
}

func isBridge(G *graph, e edge) bool {
	deleteEdge(G, e)

	N := len((*G)[0])
	dist := make([]bool, N)

	DFS(G, 0, &dist)

	addEdge(G, e)

	for _, isVisited := range dist {
		if !isVisited {
			return true
		}
	}
	return false
}

func main() {
	sc.Split(bufio.ScanWords)

	N, M := scanInt(), scanInt()
	G := make(graph, N)
	for i := 0; i < N; i++ {
		G[i] = make([]bool, N)
	}

	L := make([]edge, M)
	for i := 0; i < M; i++ {
		a, b := scanInt()-1, scanInt()-1
		L[i].a, L[i].b = a, b
		addEdge(&G, L[i])
	}

	ans := 0
	for _, ed := range L {
		if isBridge(&G, ed) {
			ans++
		}
	}
	fmt.Println(ans)
}
