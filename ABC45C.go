package gocoder

import (
	"fmt"
	"strconv"
)

func main() {
	var S string
	fmt.Scan(&S)

	N := len(S)
	ans := 0
	for i := 0; i < (1 << (N - 1)); i++ {
		s, tmp := 0, 0
		for j := 0; j < N-1; j++ {
			if i&(1<<j) > 0 {
				num, _ := strconv.Atoi(S[s : j+1])
				tmp += num
				s = j + 1
			}
		}
		num, _ := strconv.Atoi(S[s:])
		tmp += num
		ans += tmp
	}
	fmt.Println(ans)
}
