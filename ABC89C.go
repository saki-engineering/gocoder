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

func scanString() string {
	sc.Scan()
	return sc.Text()
}

func main() {
	sc.Split(bufio.ScanWords)

	N := scanInt()
	S := make([]string, N)
	for i := 0; i < N; i++ {
		S[i] = scanString()
	}

	T := "MARCH"
	cnt := make([]int, len(T))
	for _, s := range S {
		for j := 0; j < len(T); j++ {
			if s[0] == T[j] {
				cnt[j]++
			}
		}
	}

	ans := 0
	for i := 0; i < len(T); i++ {
		for j := i + 1; j < len(T); j++ {
			for k := j + 1; k < len(T); k++ {
				ans += cnt[i] * cnt[j] * cnt[k]
			}
		}
	}
	fmt.Println(ans)
}
