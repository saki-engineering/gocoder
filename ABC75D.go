package gocoder

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
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
	x, y int
}

type pointList []point

type xList struct{ pointList }
type yList struct{ pointList }

func (a pointList) Len() int      { return len(a) }
func (a pointList) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a xList) Less(i, j int) bool { return a.pointList[i].x < a.pointList[j].x }
func (a yList) Less(i, j int) bool { return a.pointList[i].y < a.pointList[j].y }

func main() {
	sc.Split(bufio.ScanWords)

	N, K := scanInt(), scanInt()
	L := make(pointList, N)
	for i := 0; i < N; i++ {
		L[i].x, L[i].y = scanInt(), scanInt()
	}
	sort.Sort(xList{L})

	ans := int(math.Pow10(18)) * 4
	for i := 0; i < N-1; i++ {
		for j := i + 1; j < N; j++ {
			M := j - i + 1
			if M < K {
				continue
			}

			L2 := make(pointList, M)
			for k := 0; k < M; k++ {
				L2[k].x, L2[k].y = L[i+k].x, L[i+k].y
			}
			sort.Sort(yList{L2})

			xrange := L[j].x - L[i].x
			for k := 0; k < M-K+1; k++ {
				yrange := L2[k+K-1].y - L2[k].y
				if new := xrange * yrange; new < ans {
					ans = new
				}
			}
		}
	}
	fmt.Println(ans)
}
