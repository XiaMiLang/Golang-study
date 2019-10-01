package main

import "fmt"

// 結構struct
// Go 沒有 class, 有 struct 承擔class任務, 有 method 概念
// 使用 type<Name> struct{} 定義結構, 名稱遵循可見性規則
// 支持指向自身的指針類型成員
// 支持匿名結構, 可用作成員,或定義成員變量
// 匿名結構也可用於 map 的值
// 可使用字面值對結構進行初始化
// 允許直接通過指針讀寫結構成員
// 同類型成員可直接拷貝賦值
// 支持 == 與 != 比較運算符, 但不支持 > 或 <
// 支持匿名字段, 本質上是定義了以某個類型名為名稱的字段
// 嵌入結構作為匿名字段看起來像繼承, 但非繼承
// 可以使用匿名字段指針

//type person struct {
//	Name string
//	Age int
//}
//func main(){
//	//a := person{"John", 13}
//	a := person{}
//	a.Name = "Tom"
//	a.Age = 24
//	fmt.Println(a)
//	A(&a)
//	fmt.Println(a)
//	B(&a)
//	fmt.Println(a)
//}
//func A(p *person){
//	p.Age = 13
//	fmt.Println("A", p)
//}
//func B(p *person){
//	p.Age = 30
//	fmt.Println("B", p)
//}

//type person struct {
//	Name string
//	Age int
//}
//func main(){
//	//a := person{"John", 13}
//	a := &person{}
//	a.Name = "Tom"
//	a.Age = 24
//	fmt.Println(a)
//	A(a)
//	fmt.Println(a)
//	B(a)
//	fmt.Println(a)
//}
//func A(p *person){
//	p.Age = 13
//	fmt.Println("A", p)
//}
//func B(p *person){
//	p.Age = 30
//	fmt.Println("B", p)
//}
//
//type person struct {
//	Name string
//	Age int
//}
//func main(){
//	//a := person{"John", 13}
//	a := &person{}
//	a.Name = "Tom"
//	a.Age = 24
//	fmt.Println(a)
//	A(a)
//	fmt.Println(a)
//	B(a)
//	fmt.Println(a)
//	a.Name = "Ximo"
//	fmt.Println(a)
//}
//func A(p *person){
//	p.Age = 13
//	fmt.Println("A", p)
//}
//func B(p *person){
//	p.Age = 30
//	fmt.Println("B", p)
//}


//func main(){	//匿名結構
//	//a := person{"John", 13}
//	a := &struct{
//		Name string
//		Age int
//	}{
//		"John",
//		55,
//	}
//	fmt.Println(a)
//	a.Name ="Ximo"
//	a.Age = 23
//	fmt.Println(a)
//}

//type person struct {
//	Name string
//	Age int
//	Contact struct{		//匿名結構
//		Phone, City string
//	}
//}
//func main(){
//	a := person{}
//	fmt.Println(a)
//	a.Contact.City = "Tainan"	//匿名結構
//}


//type person struct {
//	string		//匿名字段
//	int			//匿名字段
//	Contact struct{		//匿名結構
//		Phone, City string
//	}
//}
//func main(){
//	a := person{
//		string: "John",
//		int:    33,
//		Contact: struct {
//    Phone, City string
//	}{"0123", "Taipei"},
//	}
//	var b person
//	b =a
//	fmt.Println(a)
//	fmt.Println(b)
//	a.Contact.City = "Tainan"
//}

//
//type person struct {
//	string		//匿名字段
//	int			//匿名字段
//	Contact struct{		//匿名結構
//		Phone, City string
//	}
//}
//type person1 struct {
//	string		//匿名字段
//	int			//匿名字段
//	Contact struct{		//匿名結構
//		Phone, City string
//	}
//}
//func main(){
//	a := person{
//		string: "John",
//		int:    33,
//		Contact: struct {
//			Phone, City string
//		}{"0123", "Taipei"},
//	}
//	b := person{
//		string: "John",
//		int:    33,
//		Contact: struct {
//			Phone, City string
//		}{"0123", "Taipei"},
//	}
//	fmt.Println(a == b)		//比較只能在相同類型之間做比較
//}


//嵌入結構, 在go語言當中稱為組合(go語言不存在繼承)
//type human struct{
//	Sex int
//}
//type teacher struct{
//	human
//	Name string
//	Age int
//}
//type students struct {
//	human			//嵌入結構
//	Name string
//	Age int
//}
//func main(){
//	a := teacher{
//		human: human{1},	//需要前面打入human做初始化
//		Name: "",
//		Age:  0,
//	}
//	b:= students{
//		human: human{0},	//需要前面打入human做初始化
//		Name:  "",
//		Age:   0,
//	}
//	a.Age = 33
//	a.Name = "John"
//	a.Sex = 1
//	a.human.Sex = 100
//	fmt.Println(a, b)
//}


//課堂作業
//如果匿名字段和外層結構有同名字段, 應該如何操作?
type A struct {
	B
	C
	Name string
}
type B struct {
	Name string
}
type C struct {
	Name string
}
func main(){
	a:=A{
		B:    B{"John"},
		Name: "A",
		C:C{Name:"Jason"},
	}
	fmt.Println(a.Name, a.B.Name, a.C.Name)
}





