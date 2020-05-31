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

	N, M := scanInt(), scanInt()

	Map := make([][]int, N)
	for i := 0; i < N; i++ {
		Map[i] = make([]int, M)
	}

	for j := 0; j < M; j++ {
		k := scanInt()
		for i := 0; i < k; i++ {
			s := scanInt() - 1
			Map[s][j] = 1
		}
	}

	p := make([]int, M)
	for i := 0; i < M; i++ {
		p[i] = scanInt()
	}

	ans := 0
	for i := 0; i < (1 << uint(N)); i++ {
		cnt := make([]int, M)
		for k := 0; k < N; k++ {
			if i&(1<<uint(k)) > 0 {
				for j := 0; j < M; j++ {
					cnt[j] += Map[k][j]
				}
			}
		}
		flg := 1
		for j := 0; j < M; j++ {
			if cnt[j]%2 != p[j] {
				flg = 0
				break
			}
		}
		if flg == 1 {
			ans++
		}
	}
	fmt.Println(ans)
}
