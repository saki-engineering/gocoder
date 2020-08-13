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
	to, dist int
}

type node struct {
	no        int
	edgeList  []edge
	weight    int
	isVisited bool
}

type graph []node

func main() {
	sc.Split(bufio.ScanWords)

	N, M := scanInt(), scanInt()
	G := make(graph, N)
	for i := 0; i < N; i++ {
		G[i].no = i
	}
	for i := 0; i < M; i++ {
		l, r, d := scanInt(), scanInt(), scanInt()
		G[l-1].edgeList = append(G[l-1].edgeList, edge{to: r - 1, dist: d})
		G[r-1].edgeList = append(G[r-1].edgeList, edge{to: l - 1, dist: -d})
	}

	queue := make([]*node, 0)
	for i := 0; i < N; i++ {
		if G[i].isVisited {
			continue
		}

		queue = append(queue, &G[i])
		G[i].isVisited = true

		for len(queue) > 0 {
			grayNode := queue[0]
			// ここを queue := とtypoすると、for文中でのみ使えるqueueの定義となり無限ループになる
			queue = queue[1:]
			for _, ed := range grayNode.edgeList {
				if G[ed.to].isVisited {
					if G[ed.to].weight != G[grayNode.no].weight+ed.dist {
						fmt.Println("No")
						return
					}
				} else {
					queue = append(queue, &G[ed.to])
					G[ed.to].weight = G[grayNode.no].weight + ed.dist
					G[ed.to].isVisited = true
				}
			}
		}
	}

	fmt.Println("Yes")
}
