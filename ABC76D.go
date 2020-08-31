package main

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

type section struct {
	start, end, maxspeed int
}

func main() {
	sc.Split(bufio.ScanWords)

	N := scanInt()
	L := make([]section, N+2)

	for i := 1; i <= N; i++ {
		L[i].start = L[i-1].end
		L[i].end = L[i-1].end + scanInt()*2
	}
	T := L[N].end
	L[N+1].start, L[N+1].end = T, T

	for i := 1; i <= N; i++ {
		L[i].maxspeed = scanInt()
	}

	v := make([]float64, T+1)
	for t := 0; t <= T; t++ {
		nowSpeed := 10000.0
		for _, l := range L {
			var tmpSpeed float64
			if t < l.start {
				tmpSpeed = float64(l.maxspeed) + float64(l.start-t)/2
			} else if l.end < t {
				tmpSpeed = float64(l.maxspeed) + float64(t-l.end)/2
			} else {
				tmpSpeed = float64(l.maxspeed)
			}
			nowSpeed = math.Min(nowSpeed, tmpSpeed)
		}
		v[t] = nowSpeed
	}

	ans := 0.0
	for t := 1; t <= T; t++ {
		ans += (v[t] + v[t-1]) / 4
	}

	fmt.Println(ans)
}
