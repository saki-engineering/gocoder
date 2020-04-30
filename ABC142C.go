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
	i, e := strconv.Atoi(sc.Text())
	if e != nil {
		panic(e)
	}
	return i
}

type student struct {
	no, arv int
}
type stdList []student

func (a stdList) Len() int           { return len(a) }
func (a stdList) Less(i, j int) bool { return a[i].arv < a[j].arv }
func (a stdList) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func abc142c() {
	sc.Split(bufio.ScanWords)

	N := scanInt()
	A := make(stdList, N)
	for i := 0; i < N; i++ {
		A[i].no = i + 1
		A[i].arv = scanInt()
	}
	sort.Sort(A)

	for i := 0; i < N; i++ {
		fmt.Printf("%d ", A[i].no)
	}
	fmt.Printf("\n")
}
