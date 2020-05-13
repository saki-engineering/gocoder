package gocoder

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
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

	S := make([]string, N)
	for i := 0; i < N; i++ {
		sc.Scan()
		S[i] = sc.Text()

		sp := strings.Split(S[i], "")
		sort.Strings(sp)
		S[i] = strings.Join(sp, "")
	}

	sort.Strings(S)
	ans := 0
	cnt := 1
	for i := 1; i < N; i++ {
		if S[i] == S[i-1] {
			cnt++
		} else {
			ans += ((cnt) * (cnt - 1)) / 2
			cnt = 1
		}
	}
	ans += ((cnt) * (cnt - 1)) / 2

	fmt.Println(ans)
}
