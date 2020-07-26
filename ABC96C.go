package gocoder

import (
	"fmt"
)

func main() {
	var H, W int
	fmt.Scan(&H, &W)

	M := make([]string, H)
	for i := 0; i < H; i++ {
		fmt.Scan(&M[i])
	}

	for i := 0; i < H; i++ {
		for j := 0; j < W; j++ {
			if M[i][j] == '.' {
				continue
			}

			flg := false
			if i > 0 && M[i-1][j] == '#' {
				flg = true
			}
			if i < H-1 && M[i+1][j] == '#' {
				flg = true
			}
			if j > 0 && M[i][j-1] == '#' {
				flg = true
			}
			if j < W-1 && M[i][j+1] == '#' {
				flg = true
			}

			if !flg {
				fmt.Println("No")
				return
			}
		}
	}
	fmt.Println("Yes")
}
