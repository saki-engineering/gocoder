package gocoder

import (
	"fmt"
	"strconv"
)

func main() {
	var N int
	fmt.Scan(&N)

	ans := ""
	for N != 0 {
		ans = strconv.Itoa(N&1) + ans
		N = -(N >> 1)
	}
	if len(ans) == 0 {
		fmt.Println(0)
	} else {
		fmt.Println(ans)
	}
}
