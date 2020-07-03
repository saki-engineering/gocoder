package gocoder

import "fmt"

func main() {
	var S, T string
	fmt.Scan(&S, &T)

	N := len(S)
	M1 := make(map[uint8]uint8, 26)
	M2 := make(map[uint8]uint8, 26)

	for i := 0; i < N; i++ {
		if t, ok := M1[S[i]]; ok {
			if T[i] != t {
				fmt.Println("No")
				return
			}
		} else {
			M1[S[i]] = T[i]
			if s, ok := M2[T[i]]; ok {
				if S[i] != s {
					fmt.Println("No")
					return
				}
			} else {
				M2[T[i]] = S[i]
			}
		}
	}
	fmt.Println("Yes")
}
