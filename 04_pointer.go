package main

import "fmt"

/*
Go 不支持指針運算，以及 "->" 運算符, 採用 "." 選擇符來操作指針目標對象的成員
"*" 定義為指針類型,  指向"&"變數的位址
*/

func main() {
	a := 1
	var p *int = &a
	fmt.Println(&a)
	fmt.Println(p)
	fmt.Println(*p) //加上星號取值

}
