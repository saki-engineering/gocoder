package main

import (
	"bufio"
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

func main() {
	sc.Split(bufio.ScanWords)
	MOD := int(math.Pow10(9)) + 7

	N, M := scanInt(), scanInt()
	X, Y := make([]int, N), make([]int, M)
	for i := 0; i < N; i++ {
		X[i] = scanInt()
	}
	for i := 0; i < M; i++ {
		Y[i] = scanInt()
	}

	x, y := 0, 0
	for i, xe := range X {
		x += (2*i + 1 - N) * xe
	}
	for i, ye := range Y {
		y += (2*i + 1 - M) * ye
	}
	x, y = x%MOD, y%MOD

	ans := (x * y) % MOD
	fmt.Println(ans)
}
