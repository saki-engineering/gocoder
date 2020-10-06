package gocoder

import "fmt"

var (
	N int
	S string
)

// n=00→SS, 01→SW, 10→WS, 11→WW
func getAnimalArray(n int) (string, bool) {
	animal := map[bool]byte{true: 'S', false: 'W'}

	bArray := make([]bool, N)
	for i := 0; i < 2; i++ {
		if n&(1<<i) == 0 {
			bArray[i] = true
		} else {
			bArray[i] = false
		}
	}

	for i := 2; i < N; i++ {
		if S[i-1] == 'o' {
			if bArray[i-1] {
				bArray[i] = bArray[i-2]
			} else {
				bArray[i] = !bArray[i-2]
			}
		} else if S[i-1] == 'x' {
			if !bArray[i-1] {
				bArray[i] = bArray[i-2]
			} else {
				bArray[i] = !bArray[i-2]
			}
		}
	}

	isEqualBetween0 := (bArray[1] == bArray[N-1])
	if S[0] == 'o' {
		if bArray[0] && !isEqualBetween0 {
			return "", false
		} else if !bArray[0] && isEqualBetween0 {
			return "", false
		}
	} else if S[0] == 'x' {
		if bArray[0] && isEqualBetween0 {
			return "", false
		} else if !bArray[0] && !isEqualBetween0 {
			return "", false
		}
	}

	isEqualBetweenN := (bArray[0] == bArray[N-2])
	if S[N-1] == 'o' {
		if bArray[N-1] && !isEqualBetweenN {
			return "", false
		} else if !bArray[N-1] && isEqualBetweenN {
			return "", false
		}
	} else if S[N-1] == 'x' {
		if bArray[N-1] && isEqualBetweenN {
			return "", false
		} else if !bArray[N-1] && !isEqualBetweenN {
			return "", false
		}
	}

	array := make([]byte, N)
	for i := range bArray {
		array[i] = animal[bArray[i]]
	}

	return string(array), true
}

func main() {
	fmt.Scan(&N, &S)

	for i := 0; i < 4; i++ {
		ans, flg := getAnimalArray(i)
		if flg {
			fmt.Println(ans)
			return
		}
	}
	fmt.Println(-1)
}
