package gocoder

import (
	"fmt"
	"strings"
)

func main() {
	var S string
	fmt.Scan(&S)

	N := len(S)
	ans := N/2 - strings.Count(S, "p")
	fmt.Println(ans)
}
