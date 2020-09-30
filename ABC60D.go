package main

import (
	"bufio"
	"fmt"
	"math"
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

func intMax(x, y int) int {
	return int(math.Max(float64(x), float64(y)))
}

func intMin(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

func main() {
	sc.Split(bufio.ScanWords)

	N, W, w0 := scanInt(), scanInt(), scanInt()

	L := [4][]int{}
	for i := 0; i < 4; i++ {
		L[i] = make([]int, 0)
	}
	L[0] = append(L[0], scanInt())

	for i := 1; i < N; i++ {
		w, v := scanInt(), scanInt()
		L[w-w0] = append(L[w-w0], v)
	}

	num := [...]int{len(L[0]), len(L[1]), len(L[2]), len(L[3])}
	for i := 0; i < 4; i++ {
		sort.Sort(sort.Reverse(sort.IntSlice(L[i])))
		for j := 1; j < num[i]; j++ {
			L[i][j] += L[i][j-1]
		}
		L[i] = append([]int{0}, L[i]...)
	}

	ans := 0
	for a := 0; a <= num[0]; a++ {
		for b := 0; b <= num[1]; b++ {
			for c := 0; c <= num[2]; c++ {
				tmpW := w0*a + (w0+1)*b + (w0+2)*c
				if tmpW > W {
					continue
				}

				d := intMin((W-tmpW)/(w0+3), num[3])
				tmpV := L[0][a] + L[1][b] + L[2][c] + L[3][d]
				ans = intMax(ans, tmpV)
			}
		}
	}
	fmt.Println(ans)
}
