package main

import (
	"fmt"
	"strconv"
)

const(
	B float64 = 1 << (iota*10)
	KB
	MB
)

func main(){

	a := int(MB)
	c := strconv.Itoa(a)

	fmt.Println(B)
	fmt.Println(KB)
	fmt.Println(c)
}