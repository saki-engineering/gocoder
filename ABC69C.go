package gocoder

import (
	"bufio"
	"fmt"
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

func main() {
	sc.Split(bufio.ScanWords)

	N := scanInt()
	c1, c2, c4 := 0, 0, 0
	for i := 0; i < N; i++ {
		a := scanInt()
		if a%4 == 0 {
			c4++
		} else if a%2 == 0 {
			c2++
		} else {
			c1++
		}
	}

	besides := c1
	if c2 > 0 {
		besides++
	}

	if besides > c4+1 {
		fmt.Println("No")
	} else {
		fmt.Println("Yes")
	}
}
