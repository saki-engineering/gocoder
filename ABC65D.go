package gocoder

import (
	"bufio"
	"container/heap"
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

type point struct {
	no, val int
}

type pointList []point

func (a pointList) Len() int           { return len(a) }
func (a pointList) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a pointList) Less(i, j int) bool { return a[i].val < a[j].val }

type edge struct {
	to, dist int
}

type edgeList []edge

type node struct {
	edgeL     edgeList
	isVisited bool
}

type graph []node

func (a edgeList) Len() int           { return len(a) }
func (a edgeList) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a edgeList) Less(i, j int) bool { return a[i].dist < a[j].dist }
func (a *edgeList) Push(x interface{}) {
	*a = append(*a, x.(edge))
}
func (a *edgeList) Pop() interface{} {
	old := *a
	n := len(old)
	x := old[n-1]
	*a = old[0 : n-1]

	return x
}

func main() {
	sc.Split(bufio.ScanWords)

	N := scanInt()
	X, Y := make(pointList, N), make(pointList, N)
	for i := 0; i < N; i++ {
		X[i].no, X[i].val = i, scanInt()
		Y[i].no, Y[i].val = i, scanInt()
	}
	sort.Sort(X)
	sort.Sort(Y)

	G := make(graph, N)
	for i := 0; i < N-1; i++ {
		xa, xb, xwei := X[i].no, X[i+1].no, X[i+1].val-X[i].val
		ya, yb, ywei := Y[i].no, Y[i+1].no, Y[i+1].val-Y[i].val

		G[xa].edgeL = append(G[xa].edgeL, edge{to: xb, dist: xwei})
		G[xb].edgeL = append(G[xb].edgeL, edge{to: xa, dist: xwei})
		G[ya].edgeL = append(G[ya].edgeL, edge{to: yb, dist: ywei})
		G[yb].edgeL = append(G[yb].edgeL, edge{to: ya, dist: ywei})
	}

	E := make(edgeList, 0)
	heap.Init(&E)
	G[0].isVisited = true
	for _, ed := range G[0].edgeL {
		heap.Push(&E, ed)
	}

	ans := 0
	for i := 1; i < N; i++ {
		haveAdded := false
		for !haveAdded {
			edkouho := heap.Pop(&E).(edge)
			if !G[edkouho.to].isVisited {
				ans += edkouho.dist
				G[edkouho.to].isVisited = true
				haveAdded = true
				for _, ed := range G[edkouho.to].edgeL {
					heap.Push(&E, ed)
				}
			}
		}
	}

	fmt.Println(ans)
}
