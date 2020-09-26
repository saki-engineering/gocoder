package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"math"
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

type intList []int

func (a intList) Len() int           { return len(a) }
func (a intList) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a intList) Less(i, j int) bool { return a[i] < a[j] }
func (a *intList) Push(x interface{}) {
	*a = append(*a, x.(int))
}
func (a *intList) Pop() interface{} {
	old := *a
	n := len(*a)
	x := old[n-1]
	*a = old[:n-1]

	return x
}

func makeFormer(A *intList) (intList, int) {
	M := len(*A)
	N := M / 3

	B, sum := make(intList, N), 0
	copy(B, (*A)[:N])
	for _, b := range B {
		sum += b
	}

	return B, sum
}

func makeLatter(A *intList) (intList, int) {
	M := len(*A)
	N := M / 3

	C := make(intList, N)
	copy(C, (*A)[2*N:])

	sum := 0
	for i, c := range C {
		sum += c
		C[i] = -c
	}

	return C, sum
}

func intMax(x, y int) int {
	return int(math.Max(float64(x), float64(y)))
}

func main() {
	sc.Split(bufio.ScanWords)

	N := scanInt()
	A := make(intList, 3*N)
	for i := 0; i < 3*N; i++ {
		A[i] = scanInt()
	}

	Former, Fsum := makeFormer(&A)
	heap.Init(&Former)
	FsumList := make(intList, N+1)
	FsumList[0] = Fsum
	for i := 1; i <= N; i++ {
		heap.Push(&Former, A[N-1+i])
		poped := heap.Pop(&Former).(int)

		Fsum = Fsum + A[N-1+i] - poped
		FsumList[i] = Fsum
	}

	Latter, Lsum := makeLatter(&A)
	heap.Init(&Latter)
	LsumList := make(intList, N+1)
	LsumList[N] = Lsum
	for i := N - 1; i >= 0; i-- {
		heap.Push(&Latter, -A[N+i])
		poped := -heap.Pop(&Latter).(int)

		Lsum = Lsum + A[N+i] - poped
		LsumList[i] = Lsum
	}

	ans := FsumList[0] - LsumList[0]
	for i := 0; i <= N; i++ {
		ans = intMax(ans, FsumList[i]-LsumList[i])
	}

	fmt.Println(ans)
}
