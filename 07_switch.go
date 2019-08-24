package main

import "fmt"

//可以使用任何類型或表達式作為條件問句"
//不需 break, 一旦條件符合自動終止
//如需繼續執行下去, 使用 fallthrough 語句
//支持初始化表達式

//func main(){
//	a :=1
//	switch a {
//	case 0:
//		fmt.Println("a=0")
//	case 1:
//		fmt.Println("a=1")
//	case 2:
//		fmt.Println("a=2")
//	default:
//		fmt.Println("no found")
//
//	}
//}

//func main(){
//	a :=1
//	switch {
//	case a>=0:
//		fmt.Println("a=0")
//		fallthrough
//	case a>=1:
//		fmt.Println("a=1")
//	case a>=2:
//		fmt.Println("a=2")
//	default:
//		fmt.Println("no found")
//
//	}
//}

//初始化語句放在 switch 後面
func main(){
	switch a :=1; {
	case a>=0:
		fmt.Println("a=0")
		fallthrough
	case a>=1:
		fmt.Println("a=1")
	case a>=2:
		fmt.Println("a=2")
	default:
		fmt.Println("no found")

	}
}