package gocoder

import (
	"fmt"
	"math"
	"sort"
)

func intMin(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

func main() {
	var H, W int
	fmt.Scan(&H, &W)

	if (H*W)%3 == 0 {
		fmt.Println(0)
		return
	}

	ans := H * W
	for i := 1; i < H; i++ {
		tmp := []int{i * W, (H - i) * (W / 2), (H - i) * (W - W/2)}
		sort.Ints(tmp)
		ans = intMin(ans, tmp[2]-tmp[0])
	}
	for i := 1; i < W; i++ {
		tmp := []int{i * H, (W - i) * (H / 2), (W - i) * (H - H/2)}
		sort.Ints(tmp)
		ans = intMin(ans, tmp[2]-tmp[0])
	}
	ans = intMin(ans, H)
	ans = intMin(ans, W)
	fmt.Println(ans)
}
