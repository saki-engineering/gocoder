package gocoder

import "fmt"

func main() {
	var A, B, C, D int
	fmt.Scan(&A, &B, &C, &D)

	ans := f(B, C, D) - f(A-1, C, D)
	fmt.Println(ans)
}

func f(x, C, D int) int {
	return x - (x / C) - (x / D) + (x / lcm(C, D))
}

func gcd(A, B int) int {
	if B == 0 {
		return A
	}
	return gcd(B, A%B)
}

func lcm(A, B int) int {
	return (A * B) / gcd(A, B)
}
