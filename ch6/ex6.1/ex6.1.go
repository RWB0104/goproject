package main

import "fmt"

func main() {
	const C int = 10

	var b int = C * 20
	// 의도적 에러
	C = 20
	fmt.Println(&C)
}
