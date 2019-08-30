package main

import "fmt"

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