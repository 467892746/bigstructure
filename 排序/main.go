package main

import "fmt"

func main() {
	s := SqList{
		r:      [maxsize]int{1,6,4,3,6,8,9,11,2,5,10,100,120,0,90},
		length: maxsize ,
	}
	MergeSort(&s)
	fmt.Println(s.r)
}
