package gocoder

import (
	"fmt"
	"math"
)

func abc144d() {
	var a, b, x float64
	fmt.Scanf("%g %g %g", &a, &b, &x)

	var rad float64
	if a*a*b/2 >= x {
		h := x / (b * a / 2)
		rad = math.Atan(b / h)
	} else {
		b = 2 * (b - x/(a*a))
		rad = math.Atan(b / a)
	}

	ans := rad * 180 / math.Pi
	fmt.Println(ans)
}
