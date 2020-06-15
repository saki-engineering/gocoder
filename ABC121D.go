package gocoder

import "fmt"

func main() {
	var A, B int
	fmt.Scan(&A, &B)
	fmt.Println(XOR(A-1) ^ XOR(B))
}

func XOR(A int) int {
	cnt1 := ((A + 1) / 2) % 2
	if A%2 == 0 {
		return cnt1 ^ A
	}
	return cnt1
}
