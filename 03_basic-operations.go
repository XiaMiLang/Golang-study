package main

import (
	"fmt"
	"strconv"
)

const ff = 'B' //字符, 打印的話會顯示ASCII碼對應的數值. 常量值在函數運算進行前就要能確定, 不能以函數運算結果賦以常量

func main() {
	var a int = 65
	b := string(a) //強轉換成ASCII碼
	fmt.Println(b)
	c := strconv.Itoa(a) //數字轉換成字串
	fmt.Println(c)
	fmt.Println(strconv.Atoi(c)) //字串轉換成數字
	fmt.Println(ff + a)
	fmt.Println(6 & 11)
	fmt.Println(6 | 11)
	fmt.Println(6 ^ 11)
	fmt.Println(6 &^ 11)
}

/*
運算符優先級, 從高到低, 由左至右
^ !			一元運算符
* / % << >> & &^
+ - | ^		二元運算符
== != < <= >= >
<-			專用於channel(併發)
&&			and, 左邊的表達式(&)不成立時, 右邊的(&)就不執行了, 所以一般都是用 &&, 而不使用 &
||			or, 或, 與上式類似
*/

//fmt.Println(1 << 10 <<10 >>10)

/*
 6:	0110
11:	1011
----------------
&	0010	=2	和 運算符
|	1111	=15	或 運算符
^	1101	=13	只有一個是的時候才成立
&^	0100	=4	第二位數是1時, 強制將第一位數改為0
*/
