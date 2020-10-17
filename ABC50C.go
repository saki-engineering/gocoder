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

func main() {
	sc.Split(bufio.ScanWords)
	MOD := int(math.Pow10(9)) + 7

	N := scanInt()
	A := make([]int, N)
	for i := range A {
		A[i] = scanInt()
	}
	sort.Ints(A)

	ans, index, num := 1, 0, 0
	if N%2 == 0 {
		num = 1
		if A[index] == num && A[index+1] == num {
			ans *= 2
			index += 2
			num += 2
		} else {
			ans = 0
			fmt.Println(ans)
			return
		}
	} else {
		if A[index] == num {
			index++
			num += 2
		} else {
			ans = 0
			fmt.Println(ans)
			return
		}
	}
	for index < N {
		if A[index] == num && A[index+1] == num {
			ans = (ans * 2) % MOD
			index += 2
			num += 2
		} else {
			ans = 0
			fmt.Println(ans)
			return
		}
	}
	fmt.Println(ans)
}
