package main

import "fmt"

// 函數(function) 不支持 嵌套, 重載, 及 默認參數
// 函數(function) 無需聲明原型, 不定長度變參, 多返回值, 命名返回值參數, 匿名函數, 閉包
// 函數也可作為一種類型使用
//
//func main(){
//	A(1, 2, 3, 4, 5,6,7)
//}

//func A(a int, b string) (int, string, int){
//// func A() int{
//// func A(a, b, c int) {
//// func A() (a, b, c int) {
//// func A() (int, int, int){
//	fmt.Println("My name is", b, ". And I am", a, "years old.")
//}

//func A()(int, int, int){	//如果沒有命名返回值, 那麼 return 後面就必須跟上 a, b, c
//	a, b, c :=1, 2, 3
//	return a, b, c
//}

//func A()(a, b, c int){		//如果已有命名返回值, 那麼 return 後面就不須跟上
//	a, b, c =1, 2, 3
//	return
//}


//func main(){
//	A(1, 2, 3, 4, 5,6,7)
//}
//func A(b, string, a ...int){	// 不定長變參(a ...int), 在使用時, 規定只能作為參數列表的最後一個參數.
//	fmt.Println(a)
//}

//func main(){			//拷貝傳遞
//	a, b := 1, 2
//	A(a, b)
//	fmt.Println(a, b)
//}
//func A(a ...int){		//須注意, 傳進來的是一個 slice, slice 是一個引用類型
//	a[0] = 3
//	a[1] = 4
//	fmt.Println(a)
//}

//func main(){			//修改原資料
//	s1 :=[]int{1, 2, 3, 4}
//	A(s1)
//	fmt.Println(s1)
//}
//func A(a []int){
//	a[0] = 5
//	a[1] = 6
//	a[2] = 7
//	a[3] = 8
//	fmt.Println(a)
//}

//func main(){
//	a := 1
//	A(a)				//為值的拷貝
//	fmt.Println(a)
//}
//
//func A(a int){
//	a = 2
//	fmt.Println(a)
//}

//func main(){
//	a := 1
//	A(&a)				//用指向位址方式
//	fmt.Println(a)
//}
//
//func A(a *int){			//設為指針
//	*a = 2				//設為指針指針
//	fmt.Println(*a)
//}

//func main(){
//	a := A
//	a()
//}
//func A(){
//	fmt.Println("Func A()")
//}

//func main(){		//匿名函數, 也就是沒有名稱的
//	a := func(){
//		fmt.Println("Func A()")
//	}
//	a()
//}


//閉包的作用: 返回一個匿名函數
func main(){
	f := closure(10)
	fmt.Println(f(1))
	fmt.Println(f(2))
}
func closure(x int) func(int) int{	//作用, 返回一個匿名函數
	fmt.Printf("%p\n", &x)	//打印x的位址, 證明形成了一個閉包
	return func (y int) int{
		fmt.Printf("%p\n", &x)
		return x+y
	}
}










