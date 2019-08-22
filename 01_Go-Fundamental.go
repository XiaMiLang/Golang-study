//main 函數只能出現在 package main 內
//main 包才有可執行程序
package main

//關鍵字 import 導入其它包
//也可寫成 import aa "fmt", 使用 aa 來代表別名, main函數內即可寫成 aa.Println()
//也可寫成 import . "fmt", 使用 . 表示省略調用, main函數內即可寫成 Println()
/*
import(
	"errors"
	"io"
	"os"
)
*/
import "fmt"

//關鍵字 const 定義常量. 常量值在進行函數運算前就要能確定, 不能以函數運算結果當成常量值
//iota是常量的計數器, 從0開始, 組中每定義一個常量, iota計數器值自動增一.
//在相同const組, iota關鍵字出現次數不影響. 從0開始計數.
//在不同const組, iota從0開始計數.
//一般以全大寫字母來表示常量
//如果是只想在自己包內使用, 不被外部包使用, 可以寫成 _MAX_COUNT or cMAX_COUNT
const MAX_COUNT= 'B'  //字符'B', 打印的話會顯示ASCII碼對應的數值.
const PI = 3.14
const (
	const1 = "1"
	const2 = 2
	const3 = 3
	c = const1
	d = PI+const2
	e = iota
)

//關鍵字 var 只能定義全局變量, 在整個 package 當中使用
var name = "John"
var Age int
var (
	name1 = "1"
	name2 = 2
	name3 = 3
)
var a, b, c, d int
var e, f, g, h int = 1, 2, 3, 4

var i, j, k, l int
i, j, k, l = 1, 2, 3, 4


//關鍵字 type 定義類型
type AA int
type bb AA
type (
	newType int
	type1   float32
	type2   string
	type3   byte
)

//關鍵字 type 定義結構(struct)
type student struct {
}

//關鍵字 type 定義接口(interface)
type golang interface {
}

//關鍵字 func 定義函數
func main() {
	fmt.Println()
}

/*可見性規則
Go語言中, 使用 大小寫 來決定該 常量, 變量, 類型, 接口, 結構, 函數 是否可以被外部包所調用
根據約定, 函數名首字母小寫　即為 private, 大寫 即為 public
*/
