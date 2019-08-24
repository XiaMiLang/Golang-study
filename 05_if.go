package main

import "fmt"

func main(){
	var a = 10
	if a, b :=1, 2; a>0 { // 在 if 語句內初始化的變量, 只在 if 語句中生效
		fmt.Println(a)
		fmt.Println(b)
	}
	fmt.Println(a)

}