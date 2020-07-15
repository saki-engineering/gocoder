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

type req struct {
	a, b int
}

type reqList []req

func (a reqList) Len() int      { return len(a) }
func (a reqList) Swap(i, j int) { a[i], a[j] = a[j], a[i] }
func (a reqList) Less(i, j int) bool {
	if a[i].b == a[j].b {
		return a[i].a < a[j].a
	}
	return a[i].b < a[j].b
}

func main() {
	sc.Split(bufio.ScanWords)

	_, M := scanInt(), scanInt()
	L := make(reqList, M)
	for i := 0; i < M; i++ {
		L[i].a, L[i].b = scanInt(), scanInt()
	}
	sort.Sort(L)

	ans, prev := 0, -1
	for _, l := range L {
		if l.a >= prev {
			prev = l.b
			ans++
		}
	}

	fmt.Println(ans)
}
