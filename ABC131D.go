package main

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

type task struct {
	cost, line int
}
type taskList []task

func (a taskList) Len() int           { return len(a) }
func (a taskList) Less(i, j int) bool { return a[i].line < a[j].line }
func (a taskList) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
	sc.Split(bufio.ScanWords)

	N := scanInt()
	L := make(taskList, N)
	for i := 0; i < N; i++ {
		L[i].cost, L[i].line = scanInt(), scanInt()
	}
	sort.Sort(L)

	T := 0
	for _, l := range L {
		T += l.cost
		if T > l.line {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
