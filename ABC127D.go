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

type set struct {
	B, C int
}

type setlist []set

func (a setlist) Len() int           { return len(a) }
func (a setlist) Less(i, j int) bool { return a[i].C > a[j].C } // 降順にした
func (a setlist) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
	sc.Split(bufio.ScanWords)

	N, M := scanInt(), scanInt()
	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = scanInt()
	}
	sort.Ints(A)

	L := make(setlist, M)
	for i := 0; i < M; i++ {
		L[i].B, L[i].C = scanInt(), scanInt()
	}
	sort.Sort(L)

	a, l := 0, 0
	for a < N && l < M {
		if A[a] < L[l].C {
			A[a] = L[l].C
			L[l].B--
			if L[l].B == 0 {
				l++
			}
		}
		a++
	}

	ans := 0
	for _, v := range A {
		ans += v
	}

	fmt.Println(ans)
}
