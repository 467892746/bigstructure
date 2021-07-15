package main

const (
	maxsize = 15
)

type SqList struct {
	r [maxsize  ]int // 用于存储要排序数组，r[0]用做哨兵或临时变量
	length int // 用于记录顺序表的长度
}

func Swap(l *SqList, i int, j int)  {
	l.r[i], l.r[j] = l.r[j], l.r[i]
}
