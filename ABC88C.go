package gocoder

import "fmt"

func main() {
	C := make([][]int, 3)
	for i := 0; i < 3; i++ {
		C[i] = make([]int, 3)
	}

	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Scan(&C[i][j])
		}
	}

	for i := 0; i < 3; i++ {
		p := C[0][1] - C[0][0]
		q := C[0][2] - C[0][1]
		if C[i][1]-C[i][0] != p {
			fmt.Println("No")
			return
		}
		if C[i][2]-C[i][1] != q {
			fmt.Println("No")
			return
		}
	}
	for i := 0; i < 3; i++ {
		p := C[1][0] - C[0][0]
		q := C[2][0] - C[1][0]
		if C[1][i]-C[0][i] != p {
			fmt.Println("No")
			return
		}
		if C[2][i]-C[1][i] != q {
			fmt.Println("No")
			return
		}
	}
	fmt.Println("Yes")
}
