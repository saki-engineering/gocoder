package gocoder

import "fmt"

func main() {
	var W, H, x, y float64
	fmt.Scan(&W, &H, &x, &y)

	flg := 0
	if x == W/2 && y == H/2 {
		flg = 1
	}

	fmt.Println((W*H)/2, flg)
}
