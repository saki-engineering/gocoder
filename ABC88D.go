package gocoder

import (
	"fmt"
	"math"
)

type node struct {
	dist               int
	isVisited, isWhite bool
}

func intMin(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

func main() {
	var H, W int
	fmt.Scan(&H, &W)
	INF := int(math.Pow10(5))

	Map := make([]string, H)
	for i := 0; i < H; i++ {
		fmt.Scan(&Map[i])
	}

	cntWhite := 0
	Dist := make([][]node, H)
	for i := 0; i < H; i++ {
		Dist[i] = make([]node, W)
		for j := 0; j < W; j++ {
			Dist[i][j].dist = INF
			if Map[i][j] == '.' {
				Dist[i][j].isWhite = true
				cntWhite++
			}
		}
	}

	Dist[0][0].dist, Dist[0][0].isVisited = 1, true
	queue := make([]int, 0)
	queue = append(queue, 0)
	for len(queue) > 0 {
		pno := queue[0]
		queue = queue[1:]
		x, y := pno/W, pno%W
		p := &Dist[x][y]
		if p.isWhite {
			if x > 0 {
				p.dist = intMin(p.dist, Dist[x-1][y].dist+1)
				if !Dist[x-1][y].isVisited && Dist[x-1][y].isWhite {
					queue = append(queue, (x-1)*W+y)
					Dist[x-1][y].isVisited = true
				}
			}
			if x < H-1 {
				p.dist = intMin(p.dist, Dist[x+1][y].dist+1)
				if !Dist[x+1][y].isVisited && Dist[x+1][y].isWhite {
					queue = append(queue, (x+1)*W+y)
					Dist[x+1][y].isVisited = true
				}
			}
			if y > 0 {
				p.dist = intMin(p.dist, Dist[x][y-1].dist+1)
				if !Dist[x][y-1].isVisited && Dist[x][y-1].isWhite {
					queue = append(queue, x*W+(y-1))
					Dist[x][y-1].isVisited = true
				}
			}
			if y < W-1 {
				p.dist = intMin(p.dist, Dist[x][y+1].dist+1)
				if !Dist[x][y+1].isVisited && Dist[x][y+1].isWhite {
					queue = append(queue, x*W+(y+1))
					Dist[x][y+1].isVisited = true
				}
			}
		}
	}

	if !Dist[H-1][W-1].isVisited {
		fmt.Println(-1)
		return
	}
	ans := cntWhite - Dist[H-1][W-1].dist
	fmt.Println(ans)
}
