package main

import "fmt"

//數組 Array
//定義數組的格式 var <varName> [n]<type>, n>=0
//				var a [5]int
//數組長度也為類型的一部分, 不同長度數組為不同類型
//注意區分指向數組的指針和指針數組
//數組在 Golang 當中為值類型//數組在 Golang 當中為值類型, 當函數引用時, 值是複製過去的.
//數組之間可以用 == 或 != 進行比較, 但不可使用 < 或 >
//可用 new 來創建數組, 返回一個指向數組的指針
//Golang 支持多維數組

//func main(){
//	var a [2]int
//	var b [1]int
//	b ==a  //(mismatched types [1]int and [2]int)
//	fmt.Println(b)
//}

func main(){
	var a [2]int
	var b [2]int
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(a == b)
}


//func main(){
//	a:=[2]int{5, 6}
//	fmt.Println(a)
//}

//func main(){
//	a :=[20]int{19:1} //索引第20個元素賦值為1
//	fmt.Println(a)
//}

//func main(){
//	a := [...]int{1,2,3,4,5} //...表不直接指定元素數量, 讓 golang 對個數進行計算
//	fmt.Println(a)
//	fmt.Println(len(a))
//}

//func main(){
//	a := [...]int{19: 1} //...表不直接指定元素數量, 讓 golang 對個數進行計算
//	fmt.Println(a)
//	fmt.Println(len(a))
//}

//func main(){
//	a := [...]int{19: 1}
//	var p *[20]int = &a
//	fmt.Println(a)
//	fmt.Println(len(a))
//	fmt.Println(p)
//}

//指針數組
//func main(){
//	x, y :=1, 2
//	a := [...]*int{&x, &y}
//	b := [...]int{x, y}
//	fmt.Println(a)
//	fmt.Println(len(a))
//	fmt.Println(b)
//	fmt.Println(len(b))
//}

//數組在 Golang 當中為值類型, 當函數引用時, 值是複製過去的.