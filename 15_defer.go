package main

import "fmt"

// defer
// 在函數體執行結束後, 按照調用順序的相反順序執行. 先定義的後執行, 最後定義的最先執行
// 即使函數發生嚴重錯誤也會執行
// 支持匿名函數的調用
// 常用於資源清理, 文件關閉, 解鎖, 以及紀錄時間等操作
// 通過與匿名函數的配合, 可在 return 之後修改函數計算結果
// 如果函數體內某個變量作為 defer 時匿名函數的參數,
// 則在定義 defer 時即已經獲得了拷貝, 否則則引用某個變量的地址
// go 沒有異常機制, 但有 panic/recover 模式來處理錯誤
// panic 可以在任何地方引發, 但 recover 只有在 defer 調用的函數中有效
//func main(){
//	fmt.Println("a")
//	defer fmt.Println("b")
//	defer fmt.Println("c")
//}

//func main(){
//	for i :=0; i<3; i++{
//		defer fmt.Println(i)
//	}
//}

// 最後兩個括號的意思是調用這個函數
//func main(){
//	for i :=0; i<3; i++{
//		defer func() {
//			fmt.Println(i)
//		}()
//	}
//}

// panic / recover
//func main(){
//	A()
//	B()
//	C()
//}
//func A(){
//	fmt.Println("Func A")
//}
//func B(){
//	defer func() {
//		if err := recover(); err != nil{
//			fmt.Println("Recover in B")
//		}
//	}()
//	panic("Panic in B")
//}
//func C(){
//	fmt.Println("Func C")
//}

//課堂作業
func main(){
	var fs = [4]func(){}
	for i :=0; i<4; i++{
		defer fmt.Println("defer i = ", i)
		defer func(){
			fmt.Println("defer_closure i =", i)
		}()
		fs[i] = func (){
			fmt.Println("closure i = ", i)
		}
	}
	for _, f := range fs{
		f()
	}
}