package main

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

func intMin(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func intMax(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func main() {
	sc.Split(bufio.ScanWords)

	N, _ := scanInt(), scanInt()
	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = scanInt()
	}

	Amin := make([]int, N)
	Amin[0] = A[0]
	for i := 1; i < N; i++ {
		Amin[i] = intMin(Amin[i-1], A[i])
	}

	max := A[0] - Amin[0]
	for i := 1; i < N; i++ {
		max = intMax(max, A[i]-Amin[i])
	}

	ans := 0
	for i := 0; i < N; i++ {
		if A[i]-Amin[i] == max {
			ans++
		}
	}
	fmt.Println(ans)
}
