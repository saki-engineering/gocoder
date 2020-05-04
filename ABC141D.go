package gocoder

import (
	"bufio"
	"container/heap"
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

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] > h[j] } // maxにするためにあえて変更
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]

	return x
}

func main() {
	sc.Split(bufio.ScanWords)

	N, M := scanInt(), scanInt()
	A := make(IntHeap, N)
	for i := 0; i < N; i++ {
		A[i] = scanInt()
	}

	heap.Init(&A)
	for i := 0; i < M; i++ {
		max := heap.Pop(&A).(int)
		heap.Push(&A, max/2)
	}

	ans := 0
	for _, a := range A {
		ans += a
	}
	fmt.Println(ans)
}
