package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
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

func intMax(x, y int) int {
	return int(math.Max(float64(x), float64(y)))
}

type sushi struct {
	t, d int
}

type sushiList []sushi

func (a sushiList) Len() int           { return len(a) }
func (a sushiList) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a sushiList) Less(i, j int) bool { return a[i].d > a[j].d } // 実態はMax Heap
func (a *sushiList) Push(x interface{}) {
	*a = append(*a, x.(sushi))
}
func (a *sushiList) Pop() interface{} {
	old := *a
	n := len(*a)
	x := old[n-1]
	*a = old[0 : n-1]

	return x
}

func main() {
	sc.Split(bufio.ScanWords)

	N, K := scanInt(), scanInt()
	L := make(sushiList, N)
	for i := 0; i < N; i++ {
		L[i].t, L[i].d = scanInt()-1, scanInt()
	}
	sort.Sort(L)

	var selected, menu sushiList
	tMap := make([]bool, N)

	heap.Init(&selected)
	for i := 0; i < K; i++ {
		if tMap[L[i].t] {
			heap.Push(&selected, sushi{t: L[i].t, d: -L[i].d})
		} else {
			tMap[L[i].t] = true
		}
	}

	heap.Init(&menu)
	for i := K; i < N; i++ {
		if !tMap[L[i].t] {
			heap.Push(&menu, L[i])
		}
	}

	tSum := 0
	for _, flg := range tMap {
		if flg {
			tSum++
		}
	}
	dSum := 0
	for i := 0; i < K; i++ {
		dSum += L[i].d
	}

	ans := (tSum * tSum) + dSum

LOOP:
	for {
		if selected.Len() == 0 {
			break LOOP
		}
		removed := heap.Pop(&selected).(sushi)
		dSum += removed.d

		var added sushi
		for added.t == 0 {
			if menu.Len() == 0 {
				break LOOP
			}

			kouho := heap.Pop(&menu).(sushi)

			if !tMap[kouho.t] {
				added = kouho
				tMap[added.t] = true
				tSum++
				dSum += added.d
			}
		}

		ans = intMax(ans, (tSum*tSum)+dSum)
	}

	fmt.Println(ans)
}
