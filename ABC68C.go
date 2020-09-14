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

func main() {
	sc.Split(bufio.ScanWords)

	N, M := scanInt(), scanInt()
	A, B := make([]int, 0), make([]int, 0)
	for i := 0; i < M; i++ {
		a, b := scanInt(), scanInt()
		if a == 1 {
			B = append(B, b)
		}
		if b == N {
			A = append(A, a)
		}
	}

	sort.Ints(A)
	sort.Ints(B)
	i, j := 0, 0
	for i < len(A) && j < len(B) {
		if A[i] == B[j] {
			fmt.Println("POSSIBLE")
			return
		} else if A[i] < B[j] {
			i++
		} else {
			j++
		}
	}
	fmt.Println("IMPOSSIBLE")
}
