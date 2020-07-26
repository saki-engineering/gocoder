package gocoder

import "fmt"

func main() {
	var N int
	fmt.Scan(&N)

	prime := make([]bool, 55556)
	for i := 2; i <= 55555; i++ {
		prime[i] = true
	}

	for i := 2; i <= 55555; i++ {
		if prime[i] {
			j := i * 2
			for j <= 55555 {
				prime[j] = false
				j += i
			}
		}
	}

	cnt := [5]int{}
	for i := 2; i <= 55555; i++ {
		if prime[i] {
			cnt[i%5]++
		}
	}

	for i := 0; i < 5; i++ {
		if cnt[i] >= N {
			cnt := 0
			p := 2
			for cnt < N {
				if prime[p] && p%5 == i {
					fmt.Printf("%d ", p)
					cnt++
				}
				p++
			}
			break
		}
	}
}
