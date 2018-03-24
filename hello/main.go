package main

import (
	"fmt"
)

func main() {

	e := mult(1, 2, 3)
	fmt.Println(e)
}

func mult(a int, b int, c int) int {
	var d int
	d = a * b * c
	return d
}
