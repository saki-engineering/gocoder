package gocoder

import (
	"bufio"
	"fmt"
	"math"
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

type point struct {
	x, y, h int
}

func height(c point, x, y int) int {
	h := float64(c.h) - math.Abs(float64(c.x-x)) - math.Abs(float64(c.y-y))
	return int(math.Max(h, 0))
}

func main() {
	sc.Split(bufio.ScanWords)

	N := scanInt()
	L := make([]point, N)
	for i := 0; i < N; i++ {
		L[i].x, L[i].y, L[i].h = scanInt(), scanInt(), scanInt()
	}

	for Cx := 0; Cx <= 100; Cx++ {
		for Cy := 0; Cy <= 100; Cy++ {
			H := 0
			for _, l := range L {
				if l.h > 0 {
					H = int(float64(l.h) + math.Abs(float64(l.x-Cx)) + math.Abs(float64(l.y-Cy)))
					break
				}
			}

			flg := true
			for _, l := range L {
				if l.h != height(point{Cx, Cy, H}, l.x, l.y) {
					flg = false
					break
				}
			}

			if flg {
				fmt.Printf("%d %d %d\n", Cx, Cy, H)
				return
			}
		}
	}

}
