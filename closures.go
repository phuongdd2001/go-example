package main

import "fmt"

func intSeq() func() int {
	i := 0
	// khởi tạo một hàm ẩn danh
	return func() int {
		i++
		return i
	}
}

func main() {
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}
