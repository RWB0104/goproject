package main

import "fmt"

var g int = 10

func main() {
	var m int = 20

	{
		var s int = 50
		fmt.Println(m, s, g)
	}

	// 지역 변수 설명을 위한 의도적 에러
	m = s + 20
}
