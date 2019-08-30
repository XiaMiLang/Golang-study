package main

import "fmt"

//map 以 key-value 形式存儲數據
//key 必須是支持 == 或 != 比較運算的類型, 不可為 函數, map or slice (這三種都不支持比較運算)
//map 查找比線性搜索快很多, 但比使用索引訪問數據的類型慢100倍 (盡量使用 slice or array)
//map 使用 make() 創建, 支持使用 := 簡潔方式
//		make(map[keyType]valueType, cap), cap 表示容量, 可省略
// m :=	make(map[int]string)
//超出容量時會自動擴容, 但請盡量提供一個合理的初始值
//使用 len() 獲取元素個數
//鍵值對不存在時自動添加, 使用 delete() 刪除某鍵值對
//使用 for range 對 map 和 slice 進行迭代操作

//func main(){
//	//var m map[int]string
//	m := make(map[int]string)
//	fmt.Println(m)
//	m[1] = "ok"
//	a := m[1]
//	fmt.Println(a)
//	fmt.Println(m)
//	delete(m, 1)	//刪除內容的方式
//	fmt.Println(m)
//}

// map 的初始化
// var m map[int]string
// m = map[int]string{} or m=make(map[int]string)

// map 的初始化, 省略版
// var m map[int]string = make(map[int]string)

// map 的初始化, 省略版
// m := make(map[int]string)


//func main(){
//	m := make(map[int]map[int]string)
//	m[1] = make(map[int]string)
//	m[1][1] = "ok"
//	a:= m[1][1]
//	fmt.Println(a)
//}

//func main(){
//	var m map[int]map[int]string
//	m = make(map[int]map[int]string)
//	a, ok :=m[2][1]
//	if !ok{
//		m[2]=make(map[int]string)
//	}
//	m[2][1]="good"
//	a, ok = m[2][1]
//	fmt.Println(a, ok)
//}

//func main(){
//	for i, v := range slice{	// i 為索引(計數器),
//								// v 為slice當中所存儲的值, 為slice當中值的拷貝
//								// 對 v 的任何修改不會影響到 slice 本身
//		slice[i]				// 但可通過索引 slice[i] 對 slice 本身進行操作
//
//	}
//}

//func main(){
//	for k, v := range map{		// k 為 key
//								// k 為 map 當中的 key, 為 key 值的拷貝
//								// 對 k 的任何修改不會影響到 key 本身
//		map[k]					// 但可通過 map[k] 對 map 本身進行操作
//
//	}
//}

//func main(){
//	sm := make([]map[int]string, 5)
//	for _, v := range sm{
//		v = make(map[int]string, 2)
//		v[1]="ok"
//		v[2]="oo"
//		fmt.Println(v)
//	}
//	fmt.Println(sm)
//}

//func main(){
//	sm := make([]map[int]string, 5)
//	for i := range sm{
//		sm[i] = make(map[int]string)
//		sm[i][1]="ok"
//		fmt.Println(sm[i])
//	}
//	fmt.Println(sm)
//}

//對 slice 進行間接排序
//func main(){
//	m :=map[int]string{6:"a", 7:"b", 3:"c", 4:"d", 5:"e"}
//	s :=make([]int, len(m))
//	i :=0
//	for k, _ := range m{
//		s[i] = k
//		i++
//	}
//	sort.Ints(s)
//	fmt.Println(s)
//	for i,_ :=range s{
//		fmt.Println(m[s[i]])
//	}
//}

// 課堂作業. 將 map[int]string 鍵值交換, 變成 map[string]int
func main(){
	m1 :=map[int]string{1:"a", 2:"b", 3:"c", 4:"d", 5:"e"}
	m2 :=make(map[string]int)
	fmt.Println(m1)
	for k,v := range m1{
		m2[v] = k
	}
	fmt.Println(m2)
}
















