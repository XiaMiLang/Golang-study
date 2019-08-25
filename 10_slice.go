package main

import "fmt"

func main(){
LABEL1:
	for {	//第一層循環
		for i:=0; i<10; i++ {	//第二層循環
			if i>3{
				fmt.Println("if i>3", i)
				//break	// 跳出此層循環
				break LABEL1	// 跳出到 LABEL1 繼續執行
			}
		fmt.Println(i)
		}
	}
fmt.Println("測試有沒有繼續執行")
}
