package gocoder

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

func intMin(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

func main() {
	sc.Split(bufio.ScanWords)

	N := scanInt()
	A, B := make([][]int, N), make([][]int, N)
	for i := 0; i < N; i++ {
		A[i] = make([]int, N)
		B[i] = make([]int, N)
		for j := 0; j < N; j++ {
			A[i][j] = scanInt()
			B[i][j] = A[i][j]
		}
	}

	for k := 0; k < N; k++ {
		for i := 0; i < N; i++ {
			for j := 0; j < N; j++ {
				B[i][j] = intMin(B[i][j], B[i][k]+B[k][j])
			}
		}
	}

	for i := 0; i < N; i++ {
		for j := i; j < N; j++ {
			if B[i][j] < A[i][j] {
				fmt.Println(-1)
				return
			}
		}
	}

	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			for k := 0; k < N; k++ {
				if (i-k)*(j-k) == 0 {
					continue
				}
				if A[i][k]+A[k][j] == A[i][j] {
					B[i][j] = -1
					B[j][i] = -1
					break
				}
			}
		}
	}

	ans := 0
	for i := 0; i < N; i++ {
		for j := i + 1; j < N; j++ {
			if B[i][j] != -1 {
				ans += B[i][j]
			}
		}
	}
	fmt.Println(ans)
}
