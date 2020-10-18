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

type city struct {
	link  [2][]int
	group [2]int
}

func DP(G *[]city, from, n int) {
	fromGroupNum := (*G)[from].group[n]
	for _, to := range (*G)[from].link[n] {
		if toGroupNum := (*G)[to].group[n]; toGroupNum >= 0 {
			continue
		}
		(*G)[to].group[n] = fromGroupNum
		DP(G, to, n)
	}
}

func main() {
	sc.Split(bufio.ScanWords)

	N, K, L := scanInt(), scanInt(), scanInt()

	G := make([]city, N)
	for i := 0; i < N; i++ {
		G[i].group = [2]int{-1, -1}
	}
	for i := 0; i < K; i++ {
		p, q := scanInt()-1, scanInt()-1
		G[p].link[0] = append(G[p].link[0], q)
		G[q].link[0] = append(G[q].link[0], p)
	}
	for i := 0; i < L; i++ {
		r, s := scanInt()-1, scanInt()-1
		G[r].link[1] = append(G[r].link[1], s)
		G[s].link[1] = append(G[s].link[1], r)
	}

	groupNum := [2]int{}
	for i := 0; i < 2; i++ {
		for j, ct := range G {
			if ct.group[i] >= 0 {
				continue
			}
			G[j].group[i] = groupNum[i]
			DP(&G, j, i)
			groupNum[i]++
		}
	}

	cntMap := make(map[int]int)
	for _, ct := range G {
		cntMap[ct.group[0]*N+ct.group[1]]++
	}

	for _, ct := range G {
		fmt.Printf("%d ", cntMap[ct.group[0]*N+ct.group[1]])
	}
	fmt.Printf("\n")
}
