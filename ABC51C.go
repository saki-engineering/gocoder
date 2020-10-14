package gocoder

import (
	"fmt"
	"strings"
)

func main() {
	var sx, sy, tx, ty int
	fmt.Scan(&sx, &sy, &tx, &ty)

	X, Y := tx-sx, ty-sy
	ans := ""
	ans += strings.Repeat("U", Y)
	ans += strings.Repeat("R", X)
	ans += strings.Repeat("D", Y)
	ans += strings.Repeat("L", X+1)
	ans += strings.Repeat("U", Y+1)
	ans += strings.Repeat("R", X+1)
	ans += "DR"
	ans += strings.Repeat("D", Y+1)
	ans += strings.Repeat("L", X+1)
	ans += "U"

	fmt.Println(ans)
}
