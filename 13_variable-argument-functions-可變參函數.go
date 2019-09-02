// [_可變參數函數_](http://zh.wikipedia.org/wiki/可變參數函數)。可以用任意
// 數量的參數呼叫。例如，`fmt.Println` 是一個常見的變參函數。

package main

import "fmt"

// 這個函數使用任意數目的 `int` 作為參數。
func sum(nums ...int) {
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)
}

func main() {

	// 變參函數使用常規的呼叫方式，除了參數比較特殊。
	sum(1, 2)
	sum(1, 2, 3)

	// 如果你的 slice 已經有了多個值，想把它們作為變參
	// 使用，你要這樣呼叫 `func(slice...)`。
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}