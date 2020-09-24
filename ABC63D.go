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

func isEnough(T, A, B int, h *[]int) bool {
	cnt := 0
	for _, point := range *h {
		if remain := point - T*B; remain > 0 {
			cnt += remain / (A - B)
			if remain%(A-B) != 0 {
				cnt++
			}
		}
	}

	if cnt > T {
		return false
	}
	return true
}

func main() {
	sc.Split(bufio.ScanWords)

	N, A, B := scanInt(), scanInt(), scanInt()
	h := make([]int, N)
	for i := 0; i < N; i++ {
		h[i] = scanInt()
	}
	sort.Ints(h)

	left, right := 0, h[N-1]
	for left < right {
		middle := ((right - left) / 2) + left
		if isEnough(middle, A, B, &h) {
			right = middle
		} else {
			left = middle + 1
		}
	}

	fmt.Println(left)
}
