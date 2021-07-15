package main

import (
	"errors"
	"fmt"
)

type StaticLinkList struct {
	data interface{}
	cur int
}

func MallocSll(list []StaticLinkList) int {
	if list[0].cur == len(list ) - 1 {
		return 0
	}
	i := list[0].cur
	if list[0].cur != 0 {
		list[0].cur = list[list[0].cur].cur
	}
	return i
}

func ListDataLen(list []StaticLinkList) int {
	var j = 0
	i := list[len(list) - 1].cur
	for i != 0 {
		i = list[i].cur
		j++
	}
	return j
}

func ListInsert(list []StaticLinkList, i int, e interface{}) error {
	var (
		j,k,l int
	)
	k = len(list) - 1
	if i < 1 || i > ListDataLen(list) + 1{
		return errors.New("i错误")
	}
	j = MallocSll(list)
	if j != 0 {
		list[j].data = e
		for l = 1; l <= i -1; l++  {
			k = list[k].cur
		}
		list[j].cur = list[k].cur
		list[k].cur = j
		return nil
	}
	return errors.New("没有可用分配空间了")
}

func Free_SSL(list []StaticLinkList, k int)  {
	list[k].cur = list[0].cur
	list[0].cur = k
}

func ListDelete(list []StaticLinkList, i int) error {
	if i < 1 || i > ListDataLen(list) {
		return errors.New("没有此数据")
	}
	var j,k int
	k = len(list) - 1
	for j = 1 ;j <= i - 1 ; j ++ {
		k = list[k].cur
	}
	j = list[k].cur
	list[k].cur = list[j].cur
	Free_SSL(list, j)
	return nil
}

func main() {
	size := 4
	var staticLinkLists = make([]StaticLinkList, size)
	i := 0
	for i = 0; i < size-1;i++  {
		staticLinkLists[i].cur = i + 1
	}
	staticLinkLists[size-1].cur = 0
	err := ListInsert(staticLinkLists, 1 , "ddddd")
	if err != nil {
		panic(err)
	}
	err = ListInsert(staticLinkLists, 1 , "dd")
	if err != nil {
		panic(err)
	}
	err = ListDelete(staticLinkLists, 1)
	if err != nil {
		panic(err)
	}
	fmt.Println(staticLinkLists)
}
