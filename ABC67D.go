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

type node struct {
	edge    []int
	dist    [2]int
	Visited [2]bool
}

type graph []node

func DFS(G *graph, s, i int) {
	(*G)[s].Visited[i] = true
	for _, t := range (*G)[s].edge {
		if !(*G)[t].Visited[i] {
			(*G)[t].dist[i] = (*G)[s].dist[i] + 1
			DFS(G, t, i)
		}
	}
}

func main() {
	sc.Split(bufio.ScanWords)

	N := scanInt()
	G := make(graph, N)
	for i := 0; i < N; i++ {
		G[i].edge = make([]int, 0)
	}

	for i := 0; i < N-1; i++ {
		a, b := scanInt()-1, scanInt()-1
		G[a].edge = append(G[a].edge, b)
		G[b].edge = append(G[b].edge, a)
	}
	DFS(&G, 0, 0)
	DFS(&G, N-1, 1)

	cnt := [2]int{0, 0}
	for i := 0; i < N; i++ {
		if G[i].dist[0] <= G[i].dist[1] {
			cnt[0]++
		} else {
			cnt[1]++
		}
	}

	if cnt[0] > cnt[1] {
		fmt.Println("Fennec")
	} else {
		fmt.Println("Snuke")
	}
}
