package main

import "fmt"

func main() {
	i := 1
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue // kết quả phù hợp sẽ chạy vào đây và ko phù hợp sẽ chạy cái tiếp theo để check

		}
		fmt.Println(n)
	}
}
