package main

import "fmt"

// Go 的循環語句只有 for, 沒有 while
// Go 的循環語句支持三種型式


//func main(){
//	a:=1
//	for {
//		a++
//		if a >3 {
//			break
//		}
//	}
//	fmt.Println(a)
//}


//func main(){
//	a :=1
//	for a <=3{
//		a++
//	}
//	fmt.Println(a)
//}

//func main(){
//	 a :=1
//	 for i :=0; i<3; i++{
//	 	a++
//	 }
//	 fmt.Println(a)
//}

//func main(){
//	a := 1
//	for {
//		a++
//		if a >3{
//			break
//		}
//	fmt.Println(a)
//	}
//	fmt.Println("Over")
//	fmt.Println(a)
//}

//func main(){
//	a :=1
//	for a <=3{
//		a++
//		fmt.Println(a)
//	}
//	fmt.Println("Over")
//}

func main(){
	a:=1
	for i:=0; i<3; i++{
		a++
		fmt.Println(a)
	}
	fmt.Println("Over")
}





