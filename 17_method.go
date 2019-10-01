package main

// 方法 method
// Go 沒有 class, 有 struct 承擔class任務, 有 method 概念

//type A struct{
//	Name string
//}
//type B struct{
//	Name string
//}
//func main(){
//	a:=A{}				//聲明一個結構A{}
//	a.Name = "John"
//	a.Print()
//	fmt.Println(a.Name)
//
//	b := B{}
//	b.Print()
//}
//func (a *A)Print(){
//	a.Name = "AA"
//	fmt.Println("A", a.Name)
//}
//
//func (b B)Print(){
//	fmt.Println("B")
//}


//type TZ int
//func main(){
//	var a TZ
//	a.Print()
//	var b TZ
//	b.Print()
//}
//func (tz *TZ) Print(){
//	fmt.Println("TZ")
//}

//type TZ int
//func main(){
//	var a TZ
//	a.Print()			//method value
//	(*TZ).Print(&a)		//method expression
//	var b TZ
//	b.Print()
//}
//func (tz *TZ) Print(){
//	fmt.Println("TZ")
//}


//演示方法訪問權限的問題
//type A struct{
//	name string			// <- 以package為級別, 開頭小寫為私有自段
//}
//func main(){
//	a := A{}
//	a.Print()
//	fmt.Println(a.name)
//}
//func (a *A) Print(){
//	a.name = "123"
//	fmt.Println(a.name)
//}

// 課堂作業
// 根據為結構增加方法的知識, 嘗試聲明一個底層類型為 int 的類型,
// 並實現調用某個方法就遞增100
// 如 a:=0, 調用 a.Increase()之後, a 從 0 變成 100

func main(){

}





