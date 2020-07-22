package gocoder

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

func main() {
	sc.Split(bufio.ScanWords)
	N := scanInt()
	A := make([][20]int, N)
	for i := 0; i < N; i++ {
		a := scanInt()
		for j := 0; j < 20; j++ {
			if a&(1<<j) > 0 {
				A[i][j] = 1
			}
		}
	}

	tmp, l, r := [20]int{}, 0, 0
	ans := 0
	for r < N {
		// r++できるかどうかの判定
		flg := true
		for i := 0; i < 20; i++ {
			if tmp[i]+A[r][i] >= 2 {
				flg = false
				break
			}
		}

		if flg {
			for i := 0; i < 20; i++ {
				tmp[i] += A[r][i]
			}
			r++
		} else {
			ans += (r - l)
			for i := 0; i < 20; i++ {
				if tmp[i] == 1 {
					tmp[i] -= A[l][i]
				}
			}
			l++
		}
	}
	ans += ((N - l) * (N - l + 1)) / 2
	fmt.Println(ans)
}
