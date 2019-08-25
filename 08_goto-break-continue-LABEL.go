package main

import "fmt"

//三個語法都可配合標籤使用
//標籤區分大小寫, 標籤放了不用會造成編譯錯誤
//標籤盡量放在 goto 之後以避免死循環
//func main(){
//	LABEL1:
//		for {
//			for i:=0; i<10; i++{
//				if i >3{
//					fmt.Println(i)
//					break LABEL1
//				}
//			}
//		}
//		fmt.Println("ok")
//}

//goto :調整程序執行的位置
//func main() {
//LABEL1:
//	fmt.Println("start")
//	for a:=1;a<10;a++{
//		for i := 0; i < 10; i++ {
//			if i > 3 {
//				fmt.Println(i)
//				goto LABEL2
//			}
//		}
//	}
//
//	fmt.Println("ok")
//
//LABEL2:
//	fmt.Println("LABEL2")
//	goto LABEL1
//
//}

//continue 跳過該次循環的所有代碼, 繼續下次循環
//func main(){
//	for i:=1; i<=10; i++{
//		if i==5{
//			continue
//		}
//		fmt.Println(i)
//	}
//}

//func main(){
//LABEL1:
//	for i :=0; i<10;i++{
//		for{
//			fmt.Println(i)
//			continue LABEL1
//			fmt.Println("會執行嗎?",i)
//		}
//	}
//	fmt.Println("ok")
//}

//func main(){
//LABEL1:
//	for i :=0; i<10;i++{
//		for{
//			fmt.Println(i)
//			continue LABEL1  //作業, 如果 continue 換成 goto, 結果會一樣嗎?
//			fmt.Println("會執行嗎?",i)
//		}
//	}
//	fmt.Println("ok")
//}

func main(){
LABEL1:
	for {	//第一層循環
		for i:=0; i<10; i++ {	//第二層循環
			if i>3{
				fmt.Println("if i>3", i)
				//break	// 跳出此層循環
				break LABEL1	// 跳出到 LABEL1 繼續執行
			}
			fmt.Println(i)
		}
	}
	fmt.Println("測試有沒有繼續執行")
}