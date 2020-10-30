package gocoder

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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

func main() {
	sc.Split(bufio.ScanWords)

	N, K := scanInt(), scanInt()
	D := make([]string, K)
	for i := range D {
		D[i] = strconv.Itoa(scanInt())
	}

	for {
		Nstr := strconv.Itoa(N)
		flg := true
		for _, d := range D {
			if strings.Contains(Nstr, d) {
				flg = false
				break
			}
		}
		if flg {
			fmt.Println(N)
			return
		}
		N++
	}
}
