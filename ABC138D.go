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
	no      int
	cnt     int
	visited bool
	next    []int
}

type graph []node

func dfs(G *graph, x int) {
	(*G)[x].visited = true
	for _, i := range (*G)[x].next {
		if !(*G)[i].visited {
			(*G)[i].cnt += (*G)[x].cnt
			dfs(G, i)
		}
	}

}

func main() {
	sc.Split(bufio.ScanWords)
	N, Q := scanInt(), scanInt()

	G := make(graph, N)
	for i := 0; i < N; i++ {
		G[i].no = i
		G[i].next = make([]int, 0)
	}

	for i := 0; i < N-1; i++ {
		a, b := scanInt(), scanInt()
		G[a-1].next = append(G[a-1].next, b-1)
		G[b-1].next = append(G[b-1].next, a-1)
	}

	for i := 0; i < Q; i++ {
		p, x := scanInt(), scanInt()
		G[p-1].cnt += x
	}
	dfs(&G, 0)

	for _, node := range G {
		fmt.Printf("%d ", node.cnt)
	}
	fmt.Printf("\n")
}
