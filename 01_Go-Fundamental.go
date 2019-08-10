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

//關鍵字 var 定義全局變量, 在整個 package 當中使用
var name = "John"
var Age int

//關鍵字 type 定義類型
type AA int
type bb AA

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
