package main

const MAX_TREE_SIZE  = 100

type PTNode struct {
	data interface{} // 结点数据
	parent int // 双亲位置
}

type PTree struct {
	nodes [MAX_TREE_SIZE]PTNode // 结点数组
	r,n int // 根的位置和结点数
}

func main() {
	
}
