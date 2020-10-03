package gocoder

import (
	"fmt"
	"math"
	"strings"
)

func intMin(x, y int) int {
	return int(math.Min(float64(x), float64(y)))
}

func main() {
	var N int
	fmt.Scanf("%d", &N)

	var S string
	countLetter := make([]int, 26)
	for i := 0; i < 26; i++ {
		countLetter[i] = 50
	}

	for i := 0; i < N; i++ {
		fmt.Scanf("%s", &S)
		for j := 0; j < 26; j++ {
			countLetter[j] = intMin(countLetter[j], strings.Count(S, string('a'+j)))
		}
	}

	for i, cnt := range countLetter {
		fmt.Printf("%s", strings.Repeat(string('a'+i), cnt))
	}
	fmt.Printf("\n")
}
