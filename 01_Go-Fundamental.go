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

//關鍵字 const 定義常量
const PI = 3.14
const (
	const1 = "1"
	const2 = 2
	const3 = 3
)

//關鍵字 var 定義全局變量, 在整個 package 當中使用
var name = "John"
var Age int
var (
	name1 = "1"
	name2 = 2
	name3 = 3
)

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
