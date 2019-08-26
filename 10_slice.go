package main

import "fmt"

// slice 本身不是數組, 它指向底層的數組
//作為變長數組的替代方案, 可以關聯底層數組的局部或全部
//為引用類型
//可以直接創建, 或從底層數組獲取生成
//使用 len() 獲取個數, 使用 cap() 獲取容量
//一般使用 make() 創建
// 如果多個 slice 指向相同底層數組, 其中一個的值改變會影響全部
// make([], len, cap)
// 其中 cap 值若省略, 則和 len 的值相同
// len 表示存數的元素個數, cap 表示容量


//func main(){
//	//數組	創建方式	var a[2]int{0, 0}
//	//數組	創建方式	var a[...]int{1, 2, 3, 4}
//	//slice 創建方式	var s[]int	//空 slice
//	a :=[...]int{1, 2, 3, 4, 5}	//情況, 已有array 共5個值
//	fmt.Println("長度為", len(a), ", 內容為", a)
//	s1 := a[2]		//取第三個值
//	fmt.Println("第三個值為", s1)
//	s2 := a[2:5]	//注意!! 包含起始索引, 不包含終止索引
//	fmt.Println(s2)
//	s3 := a[2:]
//	fmt.Println(s3)
//}

//較正統的 slice 用法為使用 make 函數
//func main(){
//	s1 := make([]int, 3, 10)	//類型, 初始個數, 初始連續內存容量
//								//若最後實際個數超過初始連續內存容量, 系統會重新分配2倍初始連續內存容量
//								//ex: 10-> 20 -> 40
//	fmt.Println(len(s1), cap(s1))
//	fmt.Println(s1)
//}

// reslice 從一個 slice 當中獲得另一個新的 slice
// reslice 索引不可超過被 slice 的切片的容量 cap() 值
// reslice 索引越界會引發錯誤
//func main(){
//	a:=[]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i'}
//	sa := a[2:5]
//	sb := sa[3:4]	//slice 為指向底層的數組, 因此取到了原始數組a[]連續內存塊往後的值
//	fmt.Println(string(sb), "個數", len(sb), ", 容量", cap(sb))	//容量取到了連續內存塊的尾部
//}

// Append 可以在 slice 尾部追加元素
// Append 可以將一個 slice 追加到另一個 slice 尾部
// 如果 append 後的最終長度未超過前 slice 容量, 則返回原始 slice
// 如果 append 後超過前 slice 容量, 則會重新分配數組並拷貝原始數據 10->20->40
//func main(){
//	s1 :=make([]int, 3, 6)
//	fmt.Printf("%p", s1)	// %p 為 Pointer
//	fmt.Println(", ", cap(s1))
//	s1 =append(s1, 9, 8, 7, 6, 5)
//	fmt.Printf("%v %p", s1, s1)	// %v 為值
//	fmt.Println(", ", cap(s1))
//}

// 如果多個 slice 指向相同底層數組, 其中一個的值改變會影響全部
//func main(){
//	a :=[]int{1, 2, 3, 4, 5, 6, 7}
//	fmt.Println(a)
//	s1 := a[2:5]
//	s2 := a[1:3]
//	fmt.Println(s1, s2)
//	s1[0] = 9
//	fmt.Println(a, s1, s2)
//}

// copy 函數
//func main(){
//	s1 :=[]int{1, 2, 3, 4, 5}
//	s2 :=[]int{6, 7, 8}
//	copy(s1, s2)	//s2 複製到 s1 當中
//	fmt.Println(s1, s2)
//	copy(s2, s1)
//	fmt.Println(s2)
//}

// copy 可指定位置
//func main(){
//	s1 :=[]int{1, 2, 3, 4, 5}
//	s2 :=[]int{6, 7, 8}
//	copy(s1[1:3], s2[1:3])	//s2的指定位置 複製到 s1 的指定位置當中
//	fmt.Println(s1, s2)
//}

func main(){
	s1 :=[]int{1, 2, 3, 4, 5}
	s2 := s1
	//s2 := s1[:]
	// s2 :=s1[:5]
	fmt.Println(s2)
}