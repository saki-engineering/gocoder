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

func main() {
	sc.Split(bufio.ScanWords)

	N := int(math.Pow10(5))
	isPrime := make([]bool, N+1)
	for i := 2; i <= N; i++ {
		isPrime[i] = true
	}
	for i := 2; i <= N; i++ {
		if isPrime[i] {
			j := i * 2
			for j <= N {
				isPrime[j] = false
				j += i
			}
		}
	}

	rsum := make([]int, N+1)
	for i := 1; i <= N; i++ {
		if i%2 == 0 {
			rsum[i] = rsum[i-1]
		} else {
			if isPrime[i] && isPrime[(i+1)/2] {
				rsum[i] = rsum[i-1] + 1
			} else {
				rsum[i] = rsum[i-1]
			}
		}
	}

	Q := scanInt()
	for i := 0; i < Q; i++ {
		l, r := scanInt(), scanInt()
		fmt.Println(rsum[r] - rsum[l-1])
	}
}
