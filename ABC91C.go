package gocoder

import (
	"bufio"
	"fmt"
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
	X, Y       int
	IsSelected bool
}

type pointList []point

type redList struct{ pointList }
type blueList struct{ pointList }

func (a pointList) Len() int      { return len(a) }
func (a pointList) Swap(i, j int) { a[i], a[j] = a[j], a[i] }

func (a redList) Less(i, j int) bool  { return a.pointList[i].Y > a.pointList[j].Y }
func (a blueList) Less(i, j int) bool { return a.pointList[i].X < a.pointList[j].X }

func isClose(red, blue point) bool {
	if red.X < blue.X && red.Y < blue.Y {
		return true
	}
	return false
}

func main() {
	sc.Split(bufio.ScanWords)

	N := scanInt()
	R, B := redList{make(pointList, N)}, blueList{make(pointList, N)}
	for i := 0; i < N; i++ {
		R.pointList[i].X, R.pointList[i].Y = scanInt(), scanInt()
	}
	for i := 0; i < N; i++ {
		B.pointList[i].X, B.pointList[i].Y = scanInt(), scanInt()
	}
	sort.Sort(R)
	sort.Sort(B)

	ans := 0
	for _, blue := range B.pointList {
		for i, red := range R.pointList {
			if isClose(red, blue) && !red.IsSelected {
				ans++
				R.pointList[i].IsSelected = true
				break
			}
		}
	}

	fmt.Println(ans)
}
