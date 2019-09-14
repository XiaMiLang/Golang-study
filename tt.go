// Go 支持通過 <a href="http://zh.wikipedia.org/wiki/%E9%97%AD%E5%8C%85_(%E8%AE%A1%E7%AE%97%E6%9C%BA%E7%A7%91%E5%AD%A6)"><em>閉包</em></a>
// 來使用 [_匿名函數_](http://zh.wikipedia.org/wiki/%E5%8C%BF%E5%90%8D%E5%87%BD%E6%95%B0)。
// 匿名函數在你想定義一個不需要命名的內聯函數時是很實用的。

package main

import (
	"net/http"
)

// 這個 `intSeq` 函數返回另一個在 `intSeq` 函數體內定義的
// 匿名函數。這個返回的函數使用閉包的方式 _隱藏_ 變數 `i`。
//func intSeq() func() int {
//	i := 0
//	return func() int {
//		i += 1
//		return i
//	}
//}
//
//func main() {
//
//	// 我們呼叫 `intSeq` 函數，將返回值（也是一個函數）賦給
//	// `nextInt`。這個函數的值包含了自己的值 `i`，這樣在每
//	// 次呼叫 `nextInt` 時都會更新 `i` 的值。
//	nextInt := intSeq()
//
//	// 通過多次呼叫 `nextInt` 來看看閉包的效果。
//	fmt.Println(nextInt())
//	fmt.Println(nextInt())
//	fmt.Println(nextInt())
//
//	// 為了確認這個狀態對於這個特定的函數是唯一的，我們
//	// 重新建立並測試一下。
//	newInts := intSeq()
//	fmt.Println(newInts())
//}


//func main(){
//	var fs = [4]func(){}
//	fmt.Println(fs)
//}






//
//func outer(name string) {
//	// variable in outer function
//	text := "Modified " + name
//	// foo is a inner function and has access to text variable, is a closure
//	// closures have access to variables even after exiting this block
//	foo := func() {
//		fmt.Println(text)
//	}
//	// calling the closure
//	foo()
//}
//func main() {
//	outer("hello")
//}



func main() {
	http.ListenAndServe(":8080", http.FileServer(http.Dir("E:/迅雷下载/")))
}