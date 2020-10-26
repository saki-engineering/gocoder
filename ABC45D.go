package gocoder

import (
	"bufio"
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

type panel struct {
	a, b int
}

type panelList []panel

func (a panelList) Len() int      { return len(a) }
func (a panelList) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a panelList) Less(i, j int) bool {
	if a[i].a == a[j].a {
		return a[i].b < a[j].b
	}
	return a[i].a < a[j].a
}

func main() {
	sc.Split(bufio.ScanWords)

	H, W, N := scanInt(), scanInt(), scanInt()

	L, cnt := make(panelList, 9*N), 0
	for i := 0; i < N; i++ {
		a, b := scanInt(), scanInt()
		for j := -1; j <= 1; j++ {
			if a+j <= 1 || H <= a+j {
				continue
			}
			for k := -1; k <= 1; k++ {
				if b+k <= 1 || W <= b+k {
					continue
				}
				L[cnt].a, L[cnt].b = a+j, b+k
				cnt++
			}
		}
	}
	L = L[:cnt]
	sort.Sort(L)

	ans, c := [10]int{}, 1
	for i := 1; i < cnt; i++ {
		if L[i].a == L[i-1].a && L[i].b == L[i-1].b {
			c++
		} else {
			ans[c]++
			c = 1
		}
	}
	if cnt > 0 {
		ans[c]++
	}

	sum := 0
	for i := 1; i < 10; i++ {
		sum += ans[i]
	}
	ans[0] = (H-2)*(W-2) - sum

	for _, a := range ans {
		fmt.Println(a)
	}
}
